package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/icza/dyno"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"

	"strings"

	"github.com/cgascoig/isctl/pkg/gen"
	"github.com/cgascoig/isctl/pkg/oapi"
	"github.com/cgascoig/isctl/pkg/util"
)

var (
	applyFilenames []string
	applyDelete    bool
)

type rawMO = map[string]any

type applyConfig struct {
	client   *util.IsctlClient
	varFlags []string
	varFile  string
}

func newCmdApply(client *util.IsctlClient) *cobra.Command {
	log.Trace("Running apply cmd generator")
	config := applyConfig{
		client: client,
	}

	cmd := &cobra.Command{
		Use:   "apply",
		Run:   config.runCmdApply,
		Short: "Apply a configuration (or set of configurations) to Intersight",
	}

	cmd.Flags().StringSliceVarP(&applyFilenames, "filename", "f", []string{}, "Filename(s) that contains the configuration to apply (comma-separated list)")
	cmd.Flags().BoolVarP(&applyDelete, "delete", "d", false, "Destroy the configuration instead of creating it")
	cmd.Flags().StringArrayVar(&config.varFlags, "var", []string{}, "Set variable values (key=value)")
	cmd.Flags().StringVar(&config.varFile, "var-file", "", "Load variable values from a YAML file")

	return cmd
}

func init() {
	auxCommandsGenerators = append(auxCommandsGenerators, newCmdApply)
}

func (config *applyConfig) runCmdApply(cmd *cobra.Command, args []string) {
	// config.client.GetConfig().Debug = verbose

	vars, err := config.getVariables(applyFilenames)
	if err != nil {
		log.Fatalf("Error loading variables: %v", err)
	}

	rawMOs, err := loadRawMOs(applyFilenames, vars)
	if err != nil {
		log.Fatalf("Unable to load MOs: %v", err)
	}

	rawMOs, err = getOrderedMOs(rawMOs)
	if err != nil {
		log.Fatalf("Unable to determine order to apply MOs: %v", err)
	}

	if !applyDelete {
		// Normal apply
		err = applyMOs(config.client, rawMOs)
		if err != nil {
			log.Fatalf("Error while applying MOs: %v", err)
		}

		fmt.Println("Apply completed successfully")
	} else {
		// Destroy
		err = destroyMOs(config.client, rawMOs)
		if err != nil {
			log.Fatalf("Error while destroying MOs: %v", err)
		}

		fmt.Println("Destroy completed successfully")
	}
}

func destroyMOs(client *util.IsctlClient, rawMOs []rawMO) error {
	// when destroying we process in reverse order
	for i := len(rawMOs) - 1; i >= 0; i-- {
		mo := rawMOs[i]

		classID, err := getString(mo, "ClassId")
		if err != nil {
			return err
		}

		name, _ := getString(mo, "Name")

		meta, err := oapi.GetMeta()
		if err != nil {
			return err
		}

		getOperation := gen.GetGetOperationForClassID(classID)

		filter, err := buildIdentityFilter(client, mo, meta)
		if err != nil {
			return err
		}

		res, err := getOperation.Execute(client, nil, map[string]string{"filter": filter})
		if err != nil {
			return fmt.Errorf("error checking if MO already exists: %v", err)
		}

		moid, ok := util.GetMoid(res)
		if ok {
			log.Printf("Performing delete operation on existing MO (Name: %s, Moid: %s, ClassId: %s)", name, moid, classID)

			delOperation := gen.GetDeleteOperationForClassID(classID)

			_, err = delOperation.Execute(client, []string{moid}, nil)
			if err != nil {
				return fmt.Errorf("error executing operation: %v", err)
			}
		} else {
			log.Printf("Skipping non-existent MO (Name: %s, ClassId: %s)", name, classID)
		}

	}

	return nil
}

func applyMOs(client *util.IsctlClient, rawMOs []rawMO) error {
	for _, mo := range rawMOs {
		classID, err := getString(mo, "ClassId")
		if err != nil {
			return err
		}

		var args []string
		var op *gen.Operation

		getOperation := gen.GetGetOperationForClassID(classID)
		meta, err := oapi.GetMeta()
		if err != nil {
			return err
		}

		// Check if we can identify the object
		constraints := meta.GetIdentityConstraints(classID)
		canIdentify := false
		if len(constraints) > 0 {
			// If there are identity constraints, we assume we can identify the object
			// Missing fields will simply be omitted from the filter
			canIdentify = true
		} else {
			_, canIdentify = mo["Name"]
		}

		if getOperation == nil || !canIdentify {
			log.Printf("Performing create operation on new MO (ClassId: %s)", classID)
			op = gen.GetCreateOperationForClassID(classID)
			args = []string{}
		} else {
			filter, err := buildIdentityFilter(client, mo, meta)
			if err != nil {
				return err
			}

			name, _ := getString(mo, "Name")
			log.Tracef("applyMOs: Checking for existing object with filter %s", filter)

			res, err := getOperation.Execute(client, nil, map[string]string{"filter": filter})
			if err != nil {
				return fmt.Errorf("error checking if MO already exists: %v", err)
			}

			moid, ok := util.GetMoid(res)
			if ok {
				log.Printf("Performing update operation on existing MO (Name: %s, Moid: %s, ClassId: %s)", name, moid, classID)
				op = gen.GetUpdateOperationForClassID(classID)
				args = []string{moid}
			} else {
				log.Printf("Performing create operation on new MO (Name: %s, ClassId: %s)", name, classID)
				op = gen.GetCreateOperationForClassID(classID)
				args = []string{}
			}
		}

		if op == nil {
			return fmt.Errorf("unable to determine operation (this is definitely a bug - please submit an issue in GitHub)")
		}
		err = op.SetBodyParams(client, mo)
		if err != nil {
			return fmt.Errorf("error setting up operation body: %v", err)
		}

		_, err = op.Execute(client, args, nil)
		if err != nil {
			return fmt.Errorf("error executing operation: %v", err)
		}
	}
	return nil
}

func loadRawMOs(applyFilenames []string, vars map[string]interface{}) ([]rawMO, error) {
	rawMOs := []rawMO{}

	for _, filePath := range applyFilenames {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return nil, fmt.Errorf("unable to get file info: %v", err)
		}

		switch mode := fileInfo.Mode(); {
		case mode.IsDir():
			filenames1, err1 := filepath.Glob(filepath.Join(filePath, "*.yaml"))
			filenames2, err2 := filepath.Glob(filepath.Join(filePath, "*.yml"))
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("error finding yaml/yml files: %v, %v", err1, err2)
			}
			filenames := append(filenames1, filenames2...)
			for _, filename := range filenames {
				if filepath.Base(filename) == "isctl.vars.yaml" || filepath.Base(filename) == "isctl.vars.yml" {
					continue
				}
				mos, err := loadFile(filename, vars)
				if err != nil {
					return nil, fmt.Errorf("error reading file: %v", err)
				}

				rawMOs = append(rawMOs, mos...)
			}
		case mode.IsRegular():
			mos, err := loadFile(filePath, vars)
			if err != nil {
				return nil, fmt.Errorf("error reading file: %v", err)
			}

			rawMOs = append(rawMOs, mos...)
		default:
			return nil, fmt.Errorf("invalid file type")
		}
	}

	return rawMOs, nil
}

func loadFile(filename string, vars map[string]interface{}) ([]rawMO, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	processedContent, err := processTemplate(content, vars)
	if err != nil {
		return nil, fmt.Errorf("error processing template in %s: %w", filename, err)
	}

	log.Debugf("Processed template for %s:\n%s", filename, string(processedContent))

	ret := []rawMO{}

	dec := yaml.NewDecoder(bytes.NewReader(processedContent))

	for {
		var mo rawMO
		if dec.Decode(&mo) != nil {
			return ret, nil // no more documents in YAML file
		}

		// skip empty documents
		if len(mo) == 0 {
			continue
		}
		ret = append(ret, mo)
	}
}

// Take a list if MOs to apply and return a new list ordered based on dependencies
// this is doing a topological sort using the Depth-First search algorithm
func getOrderedMOs(mos []rawMO) ([]rawMO, error) {
	finalised := map[string]bool{}
	processing := map[string]bool{}

	dependencies := map[string]map[string]bool{}
	orderedClasses := []string{}
	mosForClassID := map[string][]rawMO{}

	// setup the finalised, processing, dependencies and mosForClassID structures from the input mos
	for _, mo := range mos {
		classID, err := getString(mo, "ClassId")
		if err != nil {
			return nil, err
		}

		mosForClassID[classID] = append(mosForClassID[classID], mo)

		finalised[classID] = false
		processing[classID] = false

		var deps = []string{}
		mo := map[string]any(mo)
		oapi.CanonicaliseMoRefs(&mo, classID)
		deps = gen.GetReferencedClasses(mo)

		for _, dep := range deps {
			if dependencies[classID] == nil {
				dependencies[classID] = map[string]bool{}
			}
			dependencies[classID][dep] = true
		}
	}

	// Resolve abstract class dependencies to concrete classes present in the MO set.
	// When a MoRef uses an abstract relationship type (e.g. resource.AbstractResourceQualificationPolicy),
	// the extracted dependency won't match any classID in the MO set. We replace such abstract deps
	// with their concrete implementations that are actually in the set.
	meta, err := oapi.GetMeta()
	if err != nil {
		log.Warnf("Failed to load metadata for abstract class resolution: %v", err)
	} else {
		for classID, deps := range dependencies {
			var toAdd []string
			var toRemove []string
			for dep := range deps {
				if _, inSet := mosForClassID[dep]; inSet {
					continue
				}
				if !meta.IsConcreteClass(dep) {
					impls := meta.GetConcreteImplementations(dep)
					for _, impl := range impls {
						if _, inSet := mosForClassID[impl]; inSet {
							toAdd = append(toAdd, impl)
						}
					}
					toRemove = append(toRemove, dep)
				}
			}
			for _, dep := range toRemove {
				delete(deps, dep)
			}
			for _, dep := range toAdd {
				deps[dep] = true
			}
			dependencies[classID] = deps
		}
	}

	// while there are still unfinalised classIds, pick one and perform the recursive search for dependencies
	for classID := getOrderedMOsGetUnfinalisedClassID(&finalised); classID != ""; classID = getOrderedMOsGetUnfinalisedClassID(&finalised) {
		err := getOrderedMOsVisit(classID, &finalised, &processing, &dependencies, &orderedClasses)
		if err != nil {
			return nil, err
		}
	}

	ret := []rawMO{}
	for _, classID := range orderedClasses {
		ret = append(ret, mosForClassID[classID]...)
	}

	return ret, nil
}

// find an unfinalised classId or return "" if there are none
func getOrderedMOsGetUnfinalisedClassID(finalised *map[string]bool) string {
	for classID, state := range *finalised {
		if !state {
			return classID
		}
	}
	return ""
}

// recursive function to visit a classId (node) and follow its dependencies
// return an error if we encounter another node being processed as this indicates a cyclic reference
func getOrderedMOsVisit(classID string, finalised, processing *map[string]bool, dependencies *map[string]map[string]bool, result *[]string) error {
	if (*finalised)[classID] {
		return nil
	}
	if (*processing)[classID] {
		return fmt.Errorf("cyclic reference detected")
	}

	(*processing)[classID] = true

	for dep := range (*dependencies)[classID] {
		err := getOrderedMOsVisit(dep, finalised, processing, dependencies, result)
		if err != nil {
			return err
		}
	}

	(*processing)[classID] = false
	(*finalised)[classID] = true

	*result = append(*result, classID)

	return nil
}

func getString(mo rawMO, attr string) (string, error) {
	attrIntf, ok := mo[attr]
	if !ok {
		return "", fmt.Errorf("invalid MO - no %s attribute", attr)
	}
	attrVal, ok := attrIntf.(string)
	if !ok {
		return "", fmt.Errorf("invalid MO - %s attribute is not string", attr)
	}

	return attrVal, nil
}

func buildIdentityFilter(client *util.IsctlClient, mo rawMO, meta *oapi.Meta) (string, error) {
	classID, err := getString(mo, "ClassId")
	if err != nil {
		return "", err
	}

	constraints := meta.GetIdentityConstraints(classID)
	if len(constraints) == 0 {
		// Fallback logic
		name, err := getString(mo, "Name")
		if err != nil {
			return "", fmt.Errorf("cannot identify object: no identity constraints and no Name attribute")
		}

		if oapi.ClassIdHasProperty(classID, "Organization") {
			var cMoRef *oapi.MoRef
			orgAttr, err := dyno.Get(mo, "Organization")
			if err != nil {
				cMoRef = oapi.CanonicaliseMoRef("default", "organization.Organization.Relationship")
			} else {
				switch orgAttr := orgAttr.(type) {
				case string:
					cMoRef = oapi.CanonicaliseMoRef(orgAttr, "organization.Organization.Relationship")
				case *oapi.MoRef:
					cMoRef = orgAttr
				default:
					return "", fmt.Errorf("error: unable to determine Organization reference")
				}
			}

			orgMoRef, err := gen.GetMoMoRef(client, cMoRef)
			if err != nil {
				return "", fmt.Errorf("error finding organization: %v", err)
			}

			orgMoid, err := dyno.GetString(orgMoRef, "Moid")
			if err != nil {
				return "", fmt.Errorf("error finding organization: %v", err)
			}
			return fmt.Sprintf("Name eq '%s' and Organization/Moid eq '%s'", name, orgMoid), nil
		}
		return fmt.Sprintf("Name eq '%s'", name), nil
	}

	// Constraints logic
	parts := []string{}
	for _, field := range constraints {
		refType, isRef := meta.GetRefType(classID, field)
		if isRef {
			var cMoRef *oapi.MoRef
			attr, err := dyno.Get(mo, field)
			if err != nil {
				if field == "Account" {
					continue
				}
				if field == "Organization" {
					cMoRef = oapi.CanonicaliseMoRef("default", refType)
				} else {
					continue
				}
			} else {
				switch attr := attr.(type) {
				case string:
					cMoRef = oapi.CanonicaliseMoRef(attr, refType)
				case *oapi.MoRef:
					cMoRef = attr
				default:
					return "", fmt.Errorf("error: unable to determine reference for field %s", field)
				}
			}

			if cMoRef == nil {
				return "", fmt.Errorf("error: unable to canonicalise reference for field %s", field)
			}

			resolvedMoRef, err := gen.GetMoMoRef(client, cMoRef)
			if err != nil {
				return "", fmt.Errorf("error finding reference for %s: %v", field, err)
			}

			moid, err := dyno.GetString(resolvedMoRef, "Moid")
			if err != nil {
				return "", fmt.Errorf("error finding moid for %s: %v", field, err)
			}
			parts = append(parts, fmt.Sprintf("%s/Moid eq '%s'", field, moid))

		} else {
			val, err := dyno.Get(mo, field)
			if err != nil {
				continue
			}
			switch val.(type) {
			case string:
				parts = append(parts, fmt.Sprintf("%s eq '%v'", field, val))
			default:
				parts = append(parts, fmt.Sprintf("%s eq %v", field, val))
			}
		}
	}

	return strings.Join(parts, " and "), nil
}
