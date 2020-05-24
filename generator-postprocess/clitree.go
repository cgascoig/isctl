package main

import "strings"

// CliItem is a node in the CLI tree
type CliItem struct {
	Token     string
	Children  map[string]*CliItem
	Parameter string //the name of the field in the request struct, nil if this isn't a parameter
	Operation *Operation
}

func getTokenListForOperation(op *Operation) []string {
	return []string{
		op.HttpMethod,
		op.BaseName,
		op.OperationId,
	}
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

func generateCliTree(opData *OperationsFile) *CliItem {
	cliTree := CliItem{
		Children: map[string]*CliItem{},
	}

	for i := range opData.Operations {
		var op *Operation = &opData.Operations[i]
		tokenList := getTokenListForOperation(op)
		cliItem := &cliTree

		for _, token := range tokenList {
			cliItem = getOrCreateChildCliItem(cliItem, token, false)
		}

		for _, param := range op.Params {
			if param.IsPathParam {
				cliItem = getOrCreateChildCliItem(cliItem, param.ParamName, true)
			}
		}

		cliItem.Operation = op
	}

	return &cliTree
}
