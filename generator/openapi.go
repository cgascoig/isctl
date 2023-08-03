package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func ReadOpenAPIFile(filename string) (*OperationsFile, error) {
	doc, err := openapi3.NewLoader().LoadFromFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error loading spec file: %v", err)
	}
	ret := &OperationsFile{}

	requiredTypes := &TypeList{}

	// Get Operations
	for path, pathItem := range doc.Paths {
		for _, method := range []string{"GET", "PUT", "POST", "PATCH", "DELETE"} {
			op := pathItem.GetOperation(method)

			if op != nil {
				var returnType string
				if response, ok := op.Responses["200"]; ok && response != nil {
					if response.Value != nil {
						if mediaType := response.Value.Content.Get("application/json"); mediaType != nil {
							if mediaType.Schema != nil {
								returnType = SchemaRefToType(mediaType.Schema)
							}
						}
					}
				}

				fmt.Printf("Adding operation: %v\n", op.OperationID)

				newOp := Operation{
					Path:           path,
					OperationID:    op.OperationID,
					HTTPMethod:     method,
					ReturnType:     returnType,
					ReturnBaseType: returnType,
					Summary:        op.Summary,
					BaseName:       ToCamelCase(op.Tags[0]),
					Params:         GetParams(op, requiredTypes),
					//TODO: finish
				}
				ret.Operations = append(ret.Operations, newOp)

				fmt.Printf("Added operation: %v\n", newOp.OperationID)
			}

		}
	}

	// Get Models
	ret.Models = map[string]Model{}
	for schemaName, schemaRef := range doc.Components.Schemas {
		fmt.Printf("Processing schemaName: %s\n", schemaName)
		name, model := GetModelFromSchemaRef(schemaName, schemaRef)
		if name != "" && model != nil {
			ret.Models[name] = *model
		}
	}

	// fmt.Printf("\n\n%v", requiredTypes.Types())

	// out, err := yaml.Marshal(ret)
	// if err != nil {
	// 	fmt.Printf("ERROR marshalling yaml")
	// } else {
	// 	fmt.Printf("---\n")
	// 	fmt.Print(string(out))
	// }

	return ret, nil
}

func GetParams(op *openapi3.Operation, requiredTypes *TypeList) []Param {
	ret := []Param{}
	for _, p := range op.Parameters {
		if p.Value != nil {
			newParam := Param{
				ParamName:   p.Value.Name,
				Description: p.Value.Description,
				DataType:    p.Value.Schema.Value.Type,
				Required:    p.Value.Required,
			}

			switch p.Value.In {
			case "path":
				newParam.IsPathParam = true
			case "query":
				newParam.IsQueryParam = true
			case "body":
				newParam.IsBodyParam = true
			}

			ret = append(ret, newParam)
			requiredTypes.AddType(newParam.DataType)
		} else {
			panic("Unimplemented")
		}
	}

	if op.RequestBody != nil && op.RequestBody.Value.Content.Get("application/json") != nil {
		p := op.RequestBody
		newParam := Param{
			ParamName:   SchemaRefToType(p.Value.Content.Get("application/json").Schema),
			Description: p.Value.Description,
			DataType:    SchemaRefToType(p.Value.Content.Get("application/json").Schema),
			Required:    p.Value.Required,
		}

		newParam.IsBodyParam = true

		ret = append(ret, newParam)
		requiredTypes.AddType(newParam.DataType)
	}

	return ret
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

func GetModelFromSchemaRef(schemaName string, schema *openapi3.SchemaRef) (string, *Model) {
	if schema == nil || schema.Value == nil {
		return "", nil
	}
	v := schema.Value

	name := ToCamelCase(schemaName)
	model := Model{}

	if v.OneOf != nil {
		fmt.Printf("WARN: oneOf type detected, using empty model\n")
	}

	if v.AnyOf != nil {
		fmt.Printf("WARN: anyOf type detected, using empty model\n")
	}

	if v.AllOf != nil {
		for _, varSchema := range v.AllOf {
			if varSchema.Value == nil || varSchema.Ref != "" {
				model.Parents = append(model.Parents, SchemaNameToType(varSchema.Ref))
			} else {
				model.Vars = append(model.Vars, SchemaToVars(varSchema.Value)...)
			}
		}
	}

	if v.Type == "object" {
		model.Vars = append(model.Vars, SchemaToVars(v)...)
	}

	return name, &model

}

func SchemaRefToDataType(sr *openapi3.SchemaRef) string {
	if sr.Ref == "" && sr.Value == nil {
		fmt.Printf("ERROR: SchemaRefToDataType - both Ref and Value are empty")
		return ""
	}

	// if sr.Ref != "" && sr.Value != nil {
	// 	fmt.Printf("ERROR: SchemaRefToDataType - both Ref and Value are non-empty")
	// 	return ""
	// }

	if sr.Ref != "" {
		return SchemaNameToType(sr.Ref)
	}

	if sr.Value != nil {
		switch sr.Value.Type {
		case "array":
			return fmt.Sprintf("[]%s", SchemaRefToDataType(sr.Value.Items))
		case "integer":
			return "int64"
		case "boolean":
			return "bool"
		case "number":
			return "float32"
		case "string":
			return "string"
		case "object":
			if sr.Ref != "" {
				return SchemaNameToType(sr.Ref)
			} else {
				fmt.Printf("ERROR: object type without Ref")
				return ""
			}
		case "":
			fmt.Printf("WARN: empty type")
			return ""
		}
	}

	panic(fmt.Sprintf("Shouldn't get here: %v", sr.Value.Type))
	return ""
}

func SchemaToVars(schema *openapi3.Schema) []Var {
	if schema.Type != "object" {
		fmt.Printf("FATAL: unhandled schema type. Schema: %v\n", schema)
		panic("Unimplemented schema type for vars")
	}

	vars := []Var{}
	for k, v := range schema.Properties {
		if v.Value.AdditionalProperties != nil {
			continue
		}
		newVar := Var{
			Name:     k,
			DataType: SchemaRefToDataType(v), //v.Value.Type,
			ReadOnly: v.Value.ReadOnly,
			Nullable: v.Value.Nullable,
			Default:  fmt.Sprintf("%v", v.Value.Default),
		}

		for _, req := range schema.Required {
			if k == req {
				newVar.Required = true
			}
		}
		vars = append(vars, newVar)
	}

	return vars
}

type TypeList map[string]bool

func (tl *TypeList) AddType(newType string) {
	(*tl)[newType] = true
}

func (tl *TypeList) Types() []string {
	ret := []string{}
	for k := range *tl {
		ret = append(ret, k)
	}
	return ret
}
