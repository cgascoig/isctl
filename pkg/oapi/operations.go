package oapi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/icza/dyno"
	log "github.com/sirupsen/logrus"
)

func getOperation(path, method string, operation map[string]any) *Operation {
	operationId, err := dyno.GetString(operation, "operationId")
	if err != nil {
		log.Error("operation with no operationId")
		return nil
	}

	returnType, _ := dyno.GetString(operation, "responses", "200", "content", "application/json", "schema", "$ref")
	summary, _ := dyno.GetString(operation, "summary")
	baseName, _ := dyno.GetString(operation, "tags", 0)

	newOp := Operation{
		Path:           path,
		OperationID:    operationId,
		HTTPMethod:     strings.ToUpper(method),
		ReturnType:     returnType,
		ReturnBaseType: returnType,
		Summary:        summary,
		BaseName:       ToCamelCase(baseName),
		Params:         GetParams(operation),
	}

	log.WithFields(log.Fields{
		"path":        path,
		"method":      method,
		"operationId": operationId,
	}).Debug("processing operation")

	return &newOp
}

func ToCamelCase(in string) string {
	in = strings.TrimSpace(in)

	regex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	words := regex.Split(in, -1)

	ret := ""

	for _, word := range words {
		ret = ret + strings.ToUpper(string(word[0])) + word[1:]
	}

	return ret
}

func SchemaRefToType(sr *openapi3.SchemaRef) string {
	if sr == nil {
		return ""
	}

	if sr.Ref == "" && sr.Value != nil {
		return sr.Value.Type
	}

	return SchemaNameToType(sr.Ref)
}

func SchemaNameToType(sn string) string {
	regex := regexp.MustCompile(`^#/components/schemas/([^/]+)$`)
	m := regex.FindStringSubmatch(sn)
	if m == nil || len(m) != 2 {
		fmt.Printf("WARN: SchemaNameToType returning unknown for '%s'", sn)
		return ""
	}

	return ToCamelCase(m[1])
}

func getParam(paramName string) *Param {
	r := regexp.MustCompile(`^#/components/parameters/(.+)$`)
	m := r.FindStringSubmatch(paramName)
	if m != nil {
		paramName = m[1]
	}

	p, err := dyno.GetMapS(lazyLoadSpec(), "components", "parameters", paramName)
	if err != nil {
		log.WithField("param_name", paramName).Debug("parameter not found")
		return nil
	}

	name, _ := dyno.GetString(p, "name")
	description, _ := dyno.GetString(p, "description")
	typ, _ := dyno.GetString(p, "schema", "type")
	req, _ := dyno.GetBoolean(p, "required")

	newParam := &Param{
		ParamName:   name,
		Description: description,
		DataType:    typ,
		Required:    req,
	}

	in, _ := dyno.GetString(p, "in")
	switch in {
	case "path":
		newParam.IsPathParam = true
	case "query":
		newParam.IsQueryParam = true
	case "body":
		newParam.IsBodyParam = true
	}

	return newParam
}

func GetParams(op map[string]any) []Param {
	ret := []Param{}
	params, err := dyno.GetSlice(op, "parameters")
	if err != nil {
		return ret
	}
	for _, p := range params {
		if ref, err := dyno.GetString(p, "$ref"); err == nil {
			p := getParam(ref)
			if p == nil {
				log.Error("param was nil")
				continue
			}
			ret = append(ret, *p)
		}
	}

	ref, err := dyno.GetString(op, "requestBody", "content", "application/json", "schema", "$ref")
	if err == nil {
		desc, _ := dyno.GetString(op, "requestBody", "description")
		req, _ := dyno.GetBoolean(op, "requestBody", "required")
		newParam := Param{
			ParamName:   SchemaNameToType(ref),
			Description: desc,
			DataType:    ref,
			Required:    req,
			IsBodyParam: true,
		}
		ret = append(ret, newParam)
	}

	return ret
}

func getOperations() []Operation {
	operations := []Operation{}
	paths, err := dyno.GetMapS(lazyLoadSpec(), "paths")
	if err != nil {
		log.Error("error getting paths from spec")
	}

	for path, pathSpec := range paths {
		if pathSpec, ok := pathSpec.(map[string]any); ok {
			for method, operation := range pathSpec {
				if operation, ok := operation.(map[string]any); ok {
					newOp := getOperation(path, method, operation)
					if newOp != nil {
						operations = append(operations, *newOp)
					}
				}
			}
		}
	}

	return operations
}

func FindOperation(method, classID string) *Operation {
	classID = SchemaNameToClassId(classID)
	var opType string
	var suffix = ""
	if strings.EqualFold(method, "get") {
		opType = "get"
		suffix = "List"
	} else if strings.EqualFold(method, "post") {
		opType = "create"
	} else if strings.EqualFold(method, "patch") {
		opType = "update"
	} else if strings.EqualFold(method, "delete") {
		opType = "delete"
	} else {
		return nil
	}

	operationID := ToCamelCase(fmt.Sprintf("%s %s %s", opType, classID, suffix))
	log.Tracef("FindOperation searching for operationID %s", operationID)

	operations := getOperations()

	for _, operation := range operations {
		if operationID == operation.OperationID {
			return &operation
		}
	}

	return nil
}

// CliItem is a node in the CLI tree
type CliItem struct {
	Token                 string
	Help                  string
	Children              map[string]*CliItem
	Parameter             string //the name of the field in the request struct, nil if this isn't a parameter
	Operation             *Operation
	BodyParamType         string
	BodyParamVars         []*Var // this includes all inherited parameters of the BodyParamType
	RequiredBodyParamVars []*Var // this is just required parameters in the concrete BodyParamType
	Alias                 bool   // this indicates that this CliItem is really an alias for another operation - we don't need to generate and Operation in operations.go for these
}

// Operation represents the YAML of one operation
type Operation struct {
	Path           string  `yaml:"path"`
	OperationID    string  `yaml:"operationId"`
	HTTPMethod     string  `yaml:"httpMethod"`
	ReturnType     string  `yaml:"returnType"`
	ReturnBaseType string  `yaml:"returnBaseType"`
	Summary        string  `yaml:"summary"`
	BaseName       string  `yaml:"baseName"`
	Params         []Param `yaml:"params"`
}

func (o *Operation) IsListOperation() bool {
	r := regexp.MustCompile(`List$`)
	return r.MatchString(o.OperationID)
}

func (o *Operation) IsUpdateOperation() bool {
	r := regexp.MustCompile(`^Update`)
	return r.MatchString(o.OperationID)
}

func (o *Operation) IsCreateOperation() bool {
	r := regexp.MustCompile(`^Create`)
	return r.MatchString(o.OperationID)
}

func (o *Operation) IsDeleteOperation() bool {
	r := regexp.MustCompile(`^Delete`)
	return r.MatchString(o.OperationID)
}

func (o *Operation) DeleteOperationDataType() string {
	re := fmt.Sprintf(`^Delete%s(\w+)$`, o.BaseName)
	r := regexp.MustCompile(re)
	m := r.FindStringSubmatch(o.OperationID)
	if len(m) != 2 {
		return ""
	}
	return fmt.Sprintf("%s.%s", strings.ToLower(o.BaseName), m[1])
}

func (o *Operation) ReturnClassID() string {
	regExpClassID := regexp.MustCompile(`\.`)
	if regExpClassID.MatchString(o.ReturnBaseType) {
		// ReturnBaseType is not empty and contains a "." - so it must be a classID
		return o.ReturnBaseType
	}

	if regexp.MustCompile(`^Delete`).MatchString(o.OperationID) {
		return o.DeleteOperationDataType()
	}

	re := fmt.Sprintf(`^%s(\w+)$`, o.BaseName)
	r := regexp.MustCompile(re)
	m := r.FindStringSubmatch(o.ReturnBaseType)
	if len(m) == 2 {
		r = regexp.MustCompile(`^(\w+)Response$`)
		m2 := r.FindStringSubmatch(m[1])
		if len(m2) == 2 {
			return fmt.Sprintf("%s.%s.Response", strings.ToLower(o.BaseName), m2[1])
		} else {
			return fmt.Sprintf("%s.%s", strings.ToLower(o.BaseName), m[1])
		}
	} else {
		return ""
	}

}

func (o *Operation) OperationClassID() string {
	if o.ReturnBaseType != "" {
		r := regexp.MustCompile(`^(\w+\.\w+)\.Response$`)
		m := r.FindStringSubmatch(o.ReturnBaseType)
		if m == nil || len(m) != 2 {
			return o.ReturnBaseType
		}
		return m[1]
	}
	re := fmt.Sprintf(`^\w+%s(\w+)$`, o.BaseName)
	r := regexp.MustCompile(re)
	m := r.FindStringSubmatch(o.OperationID)
	if len(m) != 2 {
		return ""
	}
	return fmt.Sprintf("%s.%s", strings.ToLower(o.BaseName), m[1])
}

// Param represents the YAML of one parameter
type Param struct {
	ParamName    string `yaml:"paramName"`
	Description  string `yaml:"description"`
	DataType     string `yaml:"dataType"`
	Required     bool   `yaml:"required"`
	IsBodyParam  bool   `yaml:"isBodyParam"`
	IsPathParam  bool   `yaml:"isPathParam"`
	IsQueryParam bool   `yaml:"isQueryParam"`
}

// Var represents the YAML of one member var of the OpenAPI model
type Var struct {
	Name     string `yaml:"name"`
	DataType string `yaml:"dataType"`
	Required bool   `yaml:"required"`
	ReadOnly bool   `yaml:"readOnly"`
	Nullable bool   `yaml:"nullable"`
	Default  string `yaml:"default"`
	Ignore   bool   // Used in cli.gotmpl to skip vars
}

func (v *Var) IsList() bool {
	validTypeRegExp := regexp.MustCompile(`^\[\][a-zA-Z0-9]+$`)
	return validTypeRegExp.MatchString(v.DataType) && v.DataType != "[]int64" && v.DataType != "[]float32" && v.DataType != "[]int32"
}

func (v *Var) ListElementType() string {
	validTypeRegExp := regexp.MustCompile(`^\[\]([a-zA-Z0-9]+)$`)
	m := validTypeRegExp.FindStringSubmatch(v.DataType)
	return m[1]
}

func (v *Var) ValidGenericType() bool {
	validTypeRegExp := regexp.MustCompile(`^(\[\])?([a-zA-Z0-9]+)?([a-zA-Z0-9]+)$`)
	return validTypeRegExp.MatchString(v.DataType)
}

func (v *Var) NormalisedType() string {
	var t string
	isList := false
	if v.IsList() {
		isList = true
		t = v.ListElementType()
	} else {
		t = v.DataType
	}

	// if type is already in pkg.type, do not change, otherwise change type to openapi.type
	qualifiedTypeRegExp := regexp.MustCompile(`^([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)$`)
	if !qualifiedTypeRegExp.MatchString(t) {
		t = fmt.Sprintf("openapi.%s", t)
	}

	if isList {
		return fmt.Sprintf("[]%s", t)
	} else {
		return t
	}
}

func (ci *CliItem) HasBaseMo() bool {
	hasClassID := false
	hasObjectType := false
	for _, v := range ci.BodyParamVars {
		if v.Name == "ClassId" {
			hasClassID = true
		}
		if v.Name == "ObjectType" {
			hasObjectType = true
		}
	}

	return hasClassID && hasObjectType
}

type fixupFunc func(tokenlist []string) []string

func simpleRegexpFixup(r *regexp.Regexp, sub string) fixupFunc {
	return func(tokens []string) []string {
		for i := range tokens {
			tokens[i] = r.ReplaceAllString(tokens[i], sub)
		}

		return tokens
	}
}

func removeRedundantPrefixes() fixupFunc {
	return func(tokens []string) []string {
		prefix := ""
		for i := range tokens {
			t := tokens[i]
			tokens[i] = regexp.MustCompile(`(?i)^`+prefix).ReplaceAllString(t, "")
			prefix += t
		}

		return tokens
	}
}

var fixups = []fixupFunc{
	//remove redundant prefixes (e.g. with get ntp getntppolicylist remove 'getntp')
	// this is done both before and after the post->create prefix substitutions to catch all permutations
	removeRedundantPrefixes(),

	//replace post with create
	simpleRegexpFixup(regexp.MustCompile(`(?i)^post$`), "create"),
	//replace patch with update
	simpleRegexpFixup(regexp.MustCompile(`(?i)^patch$`), "update"),

	//remove redundant prefixes (e.g. with get ntp getntppolicylist remove 'getntp')
	// this is done both before and after the post->create prefix substitutions to catch all permutations
	removeRedundantPrefixes(),

	simpleRegexpFixup(regexp.MustCompile(`By[A-Z]\w*$`), ""),

	simpleRegexpFixup(regexp.MustCompile(`List$`), ""),
}

func fixupTokenList(tokenList []string) []string {
	for _, fixup := range fixups {
		tokenList = fixup(tokenList)
	}
	return tokenList
}

func getTokenListForOperation(op *Operation) ([]string, []string) {
	methodVerbs := map[string]string{
		"get":    "Get or list",
		"patch":  "Update",
		"delete": "Delete",
		"post":   "Create",
	}

	tokenList := []string{
		op.HTTPMethod,
		op.BaseName,
		op.OperationID,
	}

	methodVerb := methodVerbs[strings.ToLower(op.HTTPMethod)]

	helpList := []string{
		methodVerb + " resouce(s)",
		fmt.Sprintf("%s %s resource(s)", methodVerb, op.BaseName),
		op.Summary,
	}

	return fixupTokenList(tokenList), helpList
}

func getOrCreateChildCliItem(cliItem *CliItem, token string, parameter bool) *CliItem {
	if token == "" {
		return cliItem
	}

	token = strings.ToLower(token)
	_, ok := cliItem.Children[token]
	if !ok {
		newItem := &CliItem{
			Token:    token,
			Children: map[string]*CliItem{},
		}

		if parameter {
			newItem.Parameter = token
		}
		cliItem.Children[token] = newItem
	}

	return cliItem.Children[token]
}

func schemaToVars(s map[string]any) []*Var {
	ret := []*Var{}

	props, err := dyno.GetMapS(s, "properties")
	if err == nil {
		for propName, prop := range props {
			dt, _ := dyno.GetString(prop, "type")
			if dt == "array" {
				if arrayType, err := dyno.GetString(prop, "items", "type"); err == nil {
					dt = fmt.Sprintf("[]%s", arrayType)
				} else if arrayTypeRef, err := dyno.GetString(prop, "items", "$ref"); err == nil {
					dt = fmt.Sprintf("[]%s", SchemaNameToClassId(arrayTypeRef))
				} else {
					log.WithField("prop", prop).Info("malformed array type")
				}
			}
			nullable, _ := dyno.GetBoolean(prop, "nullable")
			readonly, _ := dyno.GetBoolean(prop, "readonly")
			def, _ := dyno.Get(prop, "default")

			newVar := Var{
				Name:     propName,
				DataType: dt,
				ReadOnly: readonly,
				Nullable: nullable,
				Default:  fmt.Sprintf("%v", def),
			}

			req, err := dyno.GetSlice(s, "required")
			if err == nil {

				for _, req := range req {
					if req, ok := req.(string); ok {
						if propName == req {
							newVar.Required = true
						}
					}
				}
			}

			ret = append(ret, &newVar)
		}
	}

	return ret
}

func getBodyParamVars(dataType string) []*Var {
	ret := []*Var{}

	s := getSchema(dataType)
	if s == nil {
		log.WithField("schema_name", dataType).Error("getBodyParamVars: schema not found")
	}

	// is the schema an "object" itself
	if typ, err := dyno.GetString(s, "type"); err == nil && typ == "object" {
		ret = append(ret, schemaToVars(s)...)
	}

	// go through the allOfs
	allOf, err := dyno.GetSlice(s, "allOf")
	if err == nil {
		for _, v := range allOf {

			// this item in allOf is a literal object
			typ, err := dyno.GetString(v, "type")
			if v, ok := v.(map[string]any); err == nil && typ == "object" && ok {
				ret = append(ret, schemaToVars(v)...)
			}

			// this item in allOf is a ref
			ref, err := dyno.GetString(v, "$ref")
			if err == nil {
				ret = append(ret, getBodyParamVars(ref)...)
			}
		}
	}

	return ret
}

func removeDuplicateBodyParamVars(vars []*Var) []*Var {
	seen := map[string]bool{}
	ret := []*Var{}

	for _, v := range vars {
		if _, ok := seen[v.Name]; !ok {
			ret = append(ret, v)
			seen[v.Name] = true
		}
	}
	return ret
}

var ignoredOperations map[string]bool = map[string]bool{
	"QueryTelemetryTimeSeries":         true,
	"QueryTelemetryDatasourceMetadata": true,
	"QueryTelemetryGroupBy":            true,
	"QueryTelemetryScan":               true,
	"QueryTelemetrySearch":             true,
	"QueryTelemetrySegmentMetadata":    true,
	"QueryTelemetryTimeBoundary":       true,
	"QueryTelemetryTopN":               true,
}

func GenerateCliTree() *CliItem {
	cliTree := CliItem{
		Children: map[string]*CliItem{},
	}

	operations := getOperations()

	for i := range operations {
		var op *Operation = &operations[i]

		if _, ok := ignoredOperations[op.OperationID]; ok {
			continue
		}

		tokenList, tokenHelp := getTokenListForOperation(op)
		cliItem := &cliTree

		for i, token := range tokenList {
			cliItem = getOrCreateChildCliItem(cliItem, token, false)
			cliItem.Help = tokenHelp[i]
		}

		lastNonParamCliItem := cliItem
		for _, param := range op.Params {
			if param.IsPathParam {
				cliItem = getOrCreateChildCliItem(cliItem, param.ParamName, true)
			}
		}

		var vars []*Var

		// Check if there is a body param
		for _, param := range op.Params {
			if param.IsBodyParam {
				vars = removeDuplicateBodyParamVars(getBodyParamVars(param.DataType))
				cliItem.BodyParamType = param.DataType
				cliItem.BodyParamVars = vars
				cliItem.RequiredBodyParamVars = vars
			}
		}

		cliItem.Operation = op

		// if the CLI Item is an "moid" parameter, we generate a sibling CLI Item for "name"
		//  (so you can do "isctl get ntp policy name XYZ" instead of having to lookup the Moid first)
		if cliItem.Parameter == "moid" {
			cliItem = getOrCreateChildCliItem(lastNonParamCliItem, "name", true)
			// Check if there is a body param
			for _, param := range op.Params {
				if param.IsBodyParam {
					cliItem.BodyParamType = param.DataType
					cliItem.BodyParamVars = vars
					cliItem.RequiredBodyParamVars = vars
				}
			}

			cliItem.Operation = op
			cliItem.Alias = true // this is an alias for the operation with "moid" parameter - don't generate a duplicate operation in operations.go
		}
	}

	return &cliTree
}
