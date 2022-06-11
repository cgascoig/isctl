package main

import (
	"fmt"
	"os"
	"path/filepath"

	openapi "github.com/cgascoig/intersight-go-sdk/intersight"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"
)

var (
	applyFilenames []string
	applyDelete    bool
)

type rawMO map[string]interface{}

type applyConfig struct {
	client *openapi.APIClient
}

func newCmdApply(client *openapi.APIClient) *cobra.Command {
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

	return cmd
}

func init() {
	auxCommandsGenerators = append(auxCommandsGenerators, newCmdApply)
}

func (config *applyConfig) runCmdApply(cmd *cobra.Command, args []string) {
	config.client.GetConfig().Debug = verbose

	rawMOs, err := loadRawMOs(applyFilenames)
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

func destroyMOs(client *openapi.APIClient, rawMOs []rawMO) error {
	// when destroying we process in reverse order
	for i := len(rawMOs) - 1; i >= 0; i-- {
		mo := rawMOs[i]

		classID, err := mo.getString("ClassId")
		if err != nil {
			return err
		}

		name, err := mo.getString("Name")
		if err != nil {
			return err
		}

		getOperation := getGetOperationForClassID(classID)
		res, _, err := getOperation.Execute(client, nil, map[string]string{"filter": fmt.Sprintf("Name eq '%s'", name)})
		if err != nil {
			return fmt.Errorf("Error checking if MO already exists: %v", err)
		}

		moid, ok := getMoid(res)
		if ok {
			log.Printf("Performing delete operation on existing MO (Name: %s, Moid: %s, ClassId: %s)", name, moid, classID)

			delOperation := getDeleteOperationForClassID(classID)

			_, _, err = delOperation.Execute(client, []string{moid}, nil)
			if err != nil {
				return fmt.Errorf("Error executing operation: %v", err)
			}
		} else {
			log.Printf("Skipping non-existent MO (Name: %s, ClassId: %s)", name, classID)
		}

	}

	return nil
}

func applyMOs(client *openapi.APIClient, rawMOs []rawMO) error {
	for _, mo := range rawMOs {
		classID, err := mo.getString("ClassId")
		if err != nil {
			return err
		}

		name, err := mo.getString("Name")
		if err != nil {
			return err
		}

		getOperation := getGetOperationForClassID(classID)
		res, _, err := getOperation.Execute(client, nil, map[string]string{"filter": fmt.Sprintf("Name eq '%s'", name)})
		if err != nil {
			return fmt.Errorf("Error checking if MO already exists: %v", err)
		}

		var args []string
		var op Operation

		moid, ok := getMoid(res)
		if ok {
			log.Printf("Performing update operation on existing MO (Name: %s, Moid: %s, ClassId: %s)", name, moid, classID)
			op = getUpdateOperationForClassID(classID)
			args = []string{moid}
		} else {
			log.Printf("Performing create operation on new MO (Name: %s, ClassId: %s)", name, classID)
			op = getCreateOperationForClassID(classID)
			args = []string{}
		}

		err = op.SetBodyParams(client, mo)
		if err != nil {
			return fmt.Errorf("Error setting up operation body: %v", err)
		}

		_, _, err = op.Execute(client, args, nil)
		if err != nil {
			return fmt.Errorf("Error executing operation: %v", err)
		}
	}
	return nil
}

func loadRawMOs(applyFilenames []string) ([]rawMO, error) {
	rawMOs := []rawMO{}

	for _, filePath := range applyFilenames {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return nil, fmt.Errorf("Unable to get file info: %v", err)
		}

		switch mode := fileInfo.Mode(); {
		case mode.IsDir():
			filenames1, err1 := filepath.Glob(filepath.Join(filePath, "*.yaml"))
			filenames2, err2 := filepath.Glob(filepath.Join(filePath, "*.yml"))
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("Error finding yaml/yml files: %v, %v", err1, err2)
			}
			filenames := append(filenames1, filenames2...)
			for _, filename := range filenames {
				mos, err := loadFile(filename)
				if err != nil {
					return nil, fmt.Errorf("Error reading file: %v", err)
				}

				rawMOs = append(rawMOs, mos...)
			}
		case mode.IsRegular():
			mos, err := loadFile(filePath)
			if err != nil {
				return nil, fmt.Errorf("Error reading file: %v", err)
			}

			rawMOs = append(rawMOs, mos...)
		default:
			return nil, fmt.Errorf("Invalid file type")
		}
	}

	return rawMOs, nil
}

func loadFile(filename string) ([]rawMO, error) {
	in, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	ret := []rawMO{}

	dec := yaml.NewDecoder(in)

	for {
		var mo rawMO
		if dec.Decode(&mo) != nil {
			return ret, nil // no more documents in YAML file
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
		classID, err := mo.getString("ClassId")
		if err != nil {
			return nil, err
		}

		mosForClassID[classID] = append(mosForClassID[classID], mo)

		finalised[classID] = false
		processing[classID] = false

		op := getUpdateOperationForClassID(classID)
		if op == nil {
			return nil, fmt.Errorf("Unable to get update operation for ClassID: %v", classID)
		}

		deps, err := op.GetReferencedClasses(mo)
		if err != nil {
			return nil, fmt.Errorf("Unable to get referenced classes: %v", err)
		}
		for _, dep := range deps {
			if dependencies[classID] == nil {
				dependencies[classID] = map[string]bool{}
			}
			dependencies[classID][dep] = true
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
		return fmt.Errorf("Cyclic reference detected")
	}

	(*processing)[classID] = true

	for dep, _ := range (*dependencies)[classID] {
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

func (mo rawMO) getString(attr string) (string, error) {
	attrIntf, ok := mo[attr]
	if !ok {
		return "", fmt.Errorf("Invalid MO - no %s attribute", attr)
	}
	attrVal, ok := attrIntf.(string)
	if !ok {
		return "", fmt.Errorf("Invalid MO - %s attribute is not string", attr)
	}

	return attrVal, nil
}
