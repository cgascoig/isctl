package main

import (
	"fmt"
	"regexp"
	"strings"
)

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
	Alias bool // this indicates that this CliItem is really an alias for another operation - we don't need to generate and Operation in operations.go for these
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

func getRequiredBodyParamVars(opData *OperationsFile, dataType string) []*Var {
	ret := []*Var{}
	validTypeRegExp := regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	if model, ok := opData.Models[dataType]; ok {
		for i := range model.Vars {
			dt := model.Vars[i].DataType
			if !validTypeRegExp.MatchString(dt) || dt == "int64" || dt == "float32" || dt == "int32" {
				model.Vars[i].Ignore = true
			}

			ret = append(ret, &model.Vars[i])
		}
	}

	return ret
}

func getBodyParamVars(opData *OperationsFile, dataType string) []*Var {
	ret := []*Var{}
	validTypeRegExp := regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	if model, ok := opData.Models[dataType]; ok {
		for _, p := range model.Parents {
			ret = append(ret, getBodyParamVars(opData, p)...)
		}
		for i := range model.Vars {
			dt := model.Vars[i].DataType
			if !validTypeRegExp.MatchString(dt) || dt == "int64" || dt == "float32" || dt == "int32" {
				model.Vars[i].Ignore = true
			}

			ret = append(ret, &model.Vars[i])
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

func generateCliTree(opData *OperationsFile) *CliItem {
	cliTree := CliItem{
		Children: map[string]*CliItem{},
	}

	for i := range opData.Operations {
		var op *Operation = &opData.Operations[i]

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

		// Check if there is a body param
		for _, param := range op.Params {
			if param.IsBodyParam {
				cliItem.BodyParamType = param.DataType
				cliItem.BodyParamVars = removeDuplicateBodyParamVars(getBodyParamVars(opData, param.DataType))
				cliItem.RequiredBodyParamVars = getRequiredBodyParamVars(opData, param.DataType)
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
					cliItem.BodyParamVars = removeDuplicateBodyParamVars(getBodyParamVars(opData, param.DataType))
					cliItem.RequiredBodyParamVars = getRequiredBodyParamVars(opData, param.DataType)
				}
			}

			cliItem.Operation = op
			cliItem.Alias = true // this is an alias for the operation with "moid" parameter - don't generate a duplicate operation in operations.go
		}
	}

	return &cliTree
}
