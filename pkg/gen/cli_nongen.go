package gen

import (
	"encoding/json"
	"fmt"
	"regexp"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/cgascoig/isctl/pkg/oapi"
	"github.com/cgascoig/isctl/pkg/util"
)

const (
	autoPaginateBatchSizeDefault = 1000
)

func getQueryParams(flagset *pflag.FlagSet) map[string]string {
	queryParams := map[string]string{}

	if filter, err := flagset.GetString("filter"); err == nil && filter != "" {
		queryParams["filter"] = filter
	}

	if sel, err := flagset.GetString("select"); err == nil && sel != "" {
		queryParams["select"] = sel
	}

	if expand, err := flagset.GetString("expand"); err == nil && expand != "" {
		queryParams["expand"] = expand
	}

	if orderby, err := flagset.GetString("orderby"); err == nil && orderby != "" {
		queryParams["orderby"] = orderby
	}

	if top, err := flagset.GetString("top"); err == nil && top != "" {
		queryParams["top"] = top
	}

	if skip, err := flagset.GetString("skip"); err == nil && skip != "" {
		queryParams["skip"] = skip
	}

	if name, err := flagset.GetString("name"); err == nil && name != "" {
		queryParams["filter"] = fmt.Sprintf("Name eq '%s'", name)
	}

	return queryParams
}

func flagSetToMap(flagset *pflag.FlagSet, ret *map[string]interface{}) {
	flagset.Visit(func(f *pflag.Flag) {

		// Filter out isctl flags that aren't body parameters (all body params are capitalised)
		paramRe := regexp.MustCompile(`^[A-Z]`)
		if !paramRe.MatchString(f.Name) {
			return
		}

		switch f.Value.Type() {
		case "string":
			(*ret)[f.Name], _ = flagset.GetString(f.Name)
		case "bool":
			(*ret)[f.Name], _ = flagset.GetBool(f.Name)
		case "stringSlice":
			(*ret)[f.Name], _ = flagset.GetStringSlice(f.Name)
		case "int64":
			(*ret)[f.Name], _ = flagset.GetInt64(f.Name)
		case "[]int64":
			(*ret)[f.Name], _ = flagset.GetInt64Slice(f.Name)
		case "int32":
			(*ret)[f.Name], _ = flagset.GetInt(f.Name)
		case "float64":
			(*ret)[f.Name], _ = flagset.GetFloat64(f.Name)
		case "float32":
			(*ret)[f.Name], _ = flagset.GetFloat32(f.Name)
		case "json":
			var paramVal any

			err := json.Unmarshal([]byte(f.Value.String()), &paramVal)
			if err == nil {
				(*ret)[f.Name] = paramVal
			} else {
				(*ret)[f.Name] = f.Value.String()
			}
		default:
			log.Warnf("Unknown type for %s flag", f.Name)
		}
	})
}

func isFlagPassed(flagset *pflag.FlagSet, name string) bool {
	found := false
	flagset.Visit(func(f *pflag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// func runCmd(cmd *cobra.Command, args []string) {
// 	log.Printf("Running command %s with args %v", cmd.Use, args)
// }

// ResultHandler is the function signature to handle API results
type ResultHandler = func(result interface{}, err error, opts ...util.ResultOpt)

func getCommand(cliItem *oapi.CliItem, client *util.IsctlClient, resultHandler ResultHandler) *cobra.Command {
	cmd := &cobra.Command{
		Use:   cliItem.Token,
		Short: cliItem.Help,
	}

	if cliItem.Operation != nil {
		cmd.Run = getRunCmd(cliItem, client, resultHandler)
	}

	if cliItem.Parameter != "" {
		cmd.Args = cobra.ExactArgs(1)
	}

	if cliItem.BodyParamType != "" {
		cmd.Long = fmt.Sprintf("%s\nProvide resource body as JSON on standard input", cliItem.Help)
		cmd.Flags().String("bodyformat", "", "Format of request body passed on stdin (e.g. \"json\"). Default: don't read from stdin")
	}

	if cliItem.Children != nil {
		for _, child := range cliItem.Children {
			cmd.AddCommand(getCommand(child, client, resultHandler))
		}
	}

	for _, bodyParam := range cliItem.BodyParamVars {
		if bodyParam.ReadOnly {
			continue
		}
		switch bodyParam.DataType {
		case "string":
			cmd.Flags().String(bodyParam.Name, "", bodyParam.Name)
		case "boolean":
			cmd.Flags().Bool(bodyParam.Name, false, bodyParam.Name)
		case "[]string":
			cmd.Flags().StringSlice(bodyParam.Name, []string{}, bodyParam.Name)
		case "int64":
			cmd.Flags().Int64(bodyParam.Name, 0, bodyParam.Name)
		case "[]int64":
			cmd.Flags().Int64Slice(bodyParam.Name, []int64{}, bodyParam.Name)
		case "int32":
			cmd.Flags().Int(bodyParam.Name, 0, bodyParam.Name)
		case "integer":
			cmd.Flags().Int(bodyParam.Name, 0, bodyParam.Name)
		case "float64":
			cmd.Flags().Float64(bodyParam.Name, 0, bodyParam.Name)
		case "float32":
			cmd.Flags().Float32(bodyParam.Name, 0, bodyParam.Name)
		case "number":
			cmd.Flags().Float32(bodyParam.Name, 0, bodyParam.Name)
		case "[]float32":
			cmd.Flags().Float32Slice(bodyParam.Name, []float32{}, bodyParam.Name)
		case "":
			cmd.Flags().Var(&util.ComplexValue{}, bodyParam.Name, fmt.Sprintf("%s as JSON", bodyParam.Name))
		default:
			// if bobodyParam.ValidGenericType() {
			// 	if false /*isMoRef({{ .NormalisedType }}{})*/ { // TODO - need to fix isMoRef
			// 		cmd.Flags().String("{{ .Name }}", "", "{{ .Name }}")
			// 	} else {
			// 		cmd.Flags().Var(&util.ComplexValue{}, "{{ .Name }}", "{{ .Name }} as JSON")
			// 	}
			// }
			// Skipping unhandled param '{{ .Name }}' with type '{{ .DataType }}'
			//
			log.Tracef("Falling back to JSON for param %s with type %s", bodyParam.Name, bodyParam.DataType)
			cmd.Flags().Var(&util.ComplexValue{}, bodyParam.Name, fmt.Sprintf("%s as JSON", bodyParam.Name))

		}
	}

	if cliItem.Operation != nil && cliItem.Operation.HTTPMethod == "GET" && cliItem.Parameter == "" {
		cmd.Flags().String("filter", "", "Filter query (e.g. \"Name eq 'Bob'\" - note the inner quotes must be single-quotes)")
		cmd.Flags().String("select", "", "Select which attributes the API should return")
		cmd.Flags().String("expand", "", "Expand query https://intersight.com/apidocs/introduction/query/#expand-query-option-returning-related-resources")
		cmd.Flags().String("orderby", "", "Sort returned results (e.g. \"CreateTime desc\")")
		cmd.Flags().String("top", "", "Maximum number of results to return. Use with --skip for pagination of results")
		cmd.Flags().String("skip", "", "Number of objects to skip and not include in results. Use with --top for pagination of results")
		cmd.Flags().String("name", "", "Filter by exact name (note: if both --filter and --name are supplied, --name takes precedence")
		cmd.Flags().Bool("auto-paginate", false, "Automatically execute batch requests to retrieve all results")
		cmd.Flags().Int("auto-paginate-batch-size", autoPaginateBatchSizeDefault, "Number of results to request per batch when using auto-paginate")
	}

	return cmd
}

// GetCommands returns the cobra command tree for the API
func GetCommands(client *util.IsctlClient, resultHandler ResultHandler) *cobra.Command {
	cliTree := oapi.GenerateCliTree()

	return getCommand(cliTree, client, resultHandler)
}

func getRunCmd(cliItem *oapi.CliItem, client *util.IsctlClient, resultHandler ResultHandler) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var res any
		var err error

		// singleResult is true if the query should return a single result - used to decide whether to collapse the result into an object or single entry list for display/jsonpath processing
		var singleResult = false

		if cliItem.Parameter == "name" {

			singleResult = true
			// Lookup moid from name
			getOp := GetGetOperationForClassID(cliItem.Operation.ReturnClassID())
			if getOp == nil {
				resultHandler(nil, fmt.Errorf("no operation to lookup Moid from '%s'", cliItem.Operation.ReturnClassID()))
				return
			}
			res, err = getOp.Execute(client, nil, map[string]string{"filter": fmt.Sprintf("Name eq '%s'", args[0])})
			if err != nil {
				resultHandler(nil, fmt.Errorf("error executing query to lookup Moid from Name: %v", err))
				return
			}

			moid, ok := util.GetMoid(res)
			if !ok {
				resultHandler(nil, fmt.Errorf("error executing query to lookup Moid from Name: number of results must exactly equal 1"))
				return
			}

			args = []string{moid}
		}

		bodyParamMap := map[string]interface{}{}

		bodyFormat, err := cmd.Flags().GetString("bodyformat")
		if err == nil && bodyFormat != "" {
			err = util.ReadBody(bodyFormat, &bodyParamMap)
			if err != nil {
				resultHandler(nil, fmt.Errorf("error reading body from standard input: %v", err))
				return
			}
		}

		operation := newOperationFromOperation(cliItem.Operation) //newOperation(cliItem.Operation.HTTPMethod, cliItem.Operation.OperationClassID())

		flagSetToMap(cmd.Flags(), &bodyParamMap)
		if err = operation.SetBodyParams(client, bodyParamMap); err != nil {
			resultHandler(nil, fmt.Errorf("error generating API body: %v", err))
			return
		}

		queryParams := map[string]string{}
		if cliItem.Operation.HTTPMethod == "GET" && cliItem.Parameter == "" {
			queryParams = getQueryParams(cmd.Flags())
			if isFlagPassed(cmd.Flags(), "name") {
				singleResult = true
			}
		}

		if autoPaginate, err2 := cmd.Flags().GetBool("auto-paginate"); err2 == nil && autoPaginate {
			var batchSize = autoPaginateBatchSizeDefault
			if _batchSize, err3 := cmd.Flags().GetInt("auto-paginate-batch-size"); err3 == nil {
				batchSize = _batchSize
			}
			res, err = ExecuteWithPagination(operation, batchSize, client, args, queryParams)
		} else {
			res, err = operation.Execute(client, args, queryParams)
		}
		resultHandler(res, err, util.ResultOpt{SingleResult: &singleResult})
	}
}
