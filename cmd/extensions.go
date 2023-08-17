package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"

	"github.com/cgascoig/isctl/pkg/extension"
	"github.com/cgascoig/isctl/pkg/gen"
	"github.com/cgascoig/isctl/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//go:embed extensions
var extensionFS embed.FS

func loadExtensions(client *util.IsctlClient) []*cobra.Command {
	em := extension.NewExtensionManager()

	eoFn := getExecuteOperationFunction(client)

	subFs, err := fs.Sub(extensionFS, "extensions")
	if err == nil {
		em.AddFS(subFs, eoFn, structuredOutputHandler)
	}

	em.AddFS(os.DirFS("./scrap/test-extensions"), eoFn, structuredOutputHandler)

	return em.GetCommands()
}

func getExecuteOperationFunction(client *util.IsctlClient) extension.ExecuteOperationFunction {
	return func(method, classID string, args []string, body map[string]interface{}, queryParams map[string]string, paginationBatchSize int) (interface{}, error) {
		op := gen.GetOperationForClassID(method, classID)
		if op == nil {
			log.Errorf("ExecuteOperation: operation for method %s on classID %s not found", method, classID)
			return nil, fmt.Errorf("error executing operation")
		}

		if body != nil {
			op.SetBodyParams(client, body)
		}

		var res any
		var err error

		if paginationBatchSize == 0 {
			log.Trace("Batch size is 0, not using pagination")
			res, err = op.Execute(client, args, queryParams)
		} else {
			log.Tracef("Using pagination with batch size %d", paginationBatchSize)
			res, err = gen.ExecuteWithPagination(op, paginationBatchSize, client, args, queryParams)
		}

		if err != nil {
			log.Errorf("ExecuteOperation: error executing operation, res: %v, err: %v", res, err)
			return nil, fmt.Errorf("error executing operation")
		}

		return res, nil
	}
}
