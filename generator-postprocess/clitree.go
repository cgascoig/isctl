package main

import (
	"regexp"
	"strings"
)

// CliItem is a node in the CLI tree
type CliItem struct {
	Token     string
	Children  map[string]*CliItem
	Parameter string //the name of the field in the request struct, nil if this isn't a parameter
	Operation *Operation
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
	removeRedundantPrefixes(),

	simpleRegexpFixup(regexp.MustCompile(`By[A-Z]\w*$`), ""),

	simpleRegexpFixup(regexp.MustCompile(`List$`), ""),

	//replace post with create
	simpleRegexpFixup(regexp.MustCompile(`(?i)^post$`), "create"),
	//replace patch with update
	simpleRegexpFixup(regexp.MustCompile(`(?i)^patch$`), "update"),
}

func fixupTokenList(tokenList []string) []string {
	for _, fixup := range fixups {
		tokenList = fixup(tokenList)
	}
	return tokenList
}

func getTokenListForOperation(op *Operation) []string {
	tokenList := []string{
		op.HttpMethod,
		op.BaseName,
		op.OperationId,
	}

	return fixupTokenList(tokenList)
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
