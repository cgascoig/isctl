package extension

import (
	"encoding/json"
	"fmt"

	"github.com/go-python/gpython/py"
	_ "github.com/go-python/gpython/stdlib"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ExecuteOperationFunction func(method, classID string, args []string, body map[string]interface{}, queryParams map[string]string, paginationBatchSize int) (interface{}, error)
type OutputFunction func(res any, multiPartResults bool)

type Extension struct {
	ExecuteOperation ExecuteOperationFunction
	Output           OutputFunction

	module *py.Module
}

func NewExtension() *Extension {
	e := &Extension{}

	return e
}

func (e *Extension) SetCode(code []byte) error {
	ctx := py.NewContext(py.DefaultContextOpts())
	compiledCode, err := py.Compile(string(code), "extension", py.ExecMode, 0, false)
	if err != nil {
		return fmt.Errorf("Error compiling extension code: %v", err)
	}

	_, err = ctx.ModuleInit(&py.ModuleImpl{
		Info: py.ModuleInfo{
			Name: "isctl",
		},
		Methods: e.getMethods(),
	})
	if err != nil {
		return fmt.Errorf("Error setting up Python interpreter with utility functions: %v", err)
	}

	m, err := py.RunCode(ctx, compiledCode, "extension", nil)
	if err != nil {
		return fmt.Errorf("Error loading extension code: %v", err)
	}
	e.module = m

	return nil
}

func (e *Extension) GetCommand() (*cobra.Command, error) {
	res, err := e.module.Call("cmd", py.Tuple{}, nil)
	if err != nil {
		log.Tracef("Error running extension entrypoint: %v", err)
		return nil, fmt.Errorf("Error running extension entrypoint: %v", err)
	}

	resGo := pyToGo(res)

	if resMap, ok := resGo.(map[string]interface{}); ok && err == nil {
		return objectToCommand(e, resMap)
	}

	return nil, fmt.Errorf("Returned extension configuration invalid (not map[string]interface{})")
}

func pyToGo(o py.Object) interface{} {

	switch v := o.(type) {
	case py.StringDict:
		ret := map[string]interface{}{}
		for key, value := range v {
			ret[key] = pyToGo(value)
		}
		return ret
	case py.Tuple:
		ret := []interface{}{}
		for _, value := range v {
			ret = append(ret, pyToGo(value))
		}
		return ret
	case *py.List:
		ret := []interface{}{}
		for _, value := range v.Items {
			ret = append(ret, pyToGo(value))
		}
		return ret
	case py.Bool:
		if v {
			return true
		} else {
			return false
		}
	case py.Int:
		return int64(v)
	case py.Float:
		return float64(v)
	case py.String:
		return string(v)
	case py.NoneType:
		return nil
	default:
		log.Tracef("Unsupported object type from Python %T", o)
	}

	return nil
}

func goToPy(o interface{}) py.Object {
	switch v := o.(type) {
	case map[string]interface{}:
		ret := py.StringDict{}
		for key, value := range v {
			ret[key] = goToPy(value)
		}
		return ret
	case []interface{}:
		ret := &py.List{}
		for _, value := range v {
			ret.Items = append(ret.Items, goToPy(value))
		}
		return ret
	case []string:
		ret := &py.List{}
		for _, value := range v {
			ret.Items = append(ret.Items, goToPy(value))
		}
		return ret
	case bool:
		return py.Bool(v)
	case int:
		return py.Int(v)
	case string:
		return py.String(v)
	case nil:
		return py.None
	case float64:
		return py.Float(v)
	case float32:
		return py.Float(v)
	default:
		log.Tracef("Unsupported object type for conversion to Python %T", o)
	}

	return nil
}

func tupleToFormat(args py.Tuple) (string, []interface{}) {
	values := pyToGo(args).([]interface{})
	format := values[0].(string)
	values = values[1:]
	return format, values
}

func getOutputFn(of OutputFunction) func(_ py.Object, args py.Tuple) (py.Object, error) {
	return func(_ py.Object, args py.Tuple) (py.Object, error) {
		argsFromPyIntf := pyToGo(args)
		argsFromPy, ok := argsFromPyIntf.([]interface{})
		if !ok {
			return nil, fmt.Errorf("output: unable to unpack arguments from extension")
		}

		data, ok := argsFromPy[0].(any)
		if !ok {
			return nil, fmt.Errorf("output: unable to unpack arguments from extension - data should be interface{}")
		}

		var multiPartResult = false
		if len(argsFromPy) >= 2 {
			if mpr, ok := argsFromPy[1].(bool); ok {
				multiPartResult = mpr
			}
		}

		of(data, multiPartResult)
		return nil, nil
	}
}

func (e *Extension) getMethods() []*py.Method {
	ret := []*py.Method{
		//isctl.debug function
		py.MustNewMethod("debug", func(_ py.Object, args py.Tuple) (py.Object, error) {
			f, a := tupleToFormat(args)
			log.Debugf(f, a...)
			return nil, nil
		}, 0, ""),

		//isctl.info function
		py.MustNewMethod("info", func(_ py.Object, args py.Tuple) (py.Object, error) {
			f, a := tupleToFormat(args)
			log.Infof(f, a...)
			return nil, nil
		}, 0, ""),

		//isctl.warn function
		py.MustNewMethod("warn", func(_ py.Object, args py.Tuple) (py.Object, error) {
			f, a := tupleToFormat(args)
			log.Warnf(f, a...)
			return nil, nil
		}, 0, ""),

		//isctl.error function
		py.MustNewMethod("error", func(_ py.Object, args py.Tuple) (py.Object, error) {
			f, a := tupleToFormat(args)
			log.Errorf(f, a...)
			return nil, nil
		}, 0, ""),

		//isctl.fatal function
		py.MustNewMethod("fatal", func(_ py.Object, args py.Tuple) (py.Object, error) {
			f, a := tupleToFormat(args)
			log.Fatalf(f, a...)
			return nil, nil
		}, 0, ""),

		//isctl.output function
		py.MustNewMethod("output", getOutputFn(e.Output), 0, ""),
	}

	if e.ExecuteOperation != nil {
		method := py.MustNewMethod("executeOperation", func(_ py.Object, args py.Tuple) (py.Object, error) {
			if len(args) < 5 {
				log.Tracef("Incorrect number of arguments from extension")
				return nil, fmt.Errorf("incorrect number of arguments from extension")
			}

			argsFromPyIntf := pyToGo(args)
			log.Tracef("argsFromPy: %v", argsFromPyIntf)

			argsFromPy, ok := argsFromPyIntf.([]interface{})
			if !ok {
				return nil, fmt.Errorf("unable to unpack arguments from extension")
			}

			method, ok := argsFromPy[0].(string)
			if !ok {
				return nil, fmt.Errorf("unable to unpack arguments from extension - method should be string")
			}

			classID, ok := argsFromPy[1].(string)
			if !ok {
				return nil, fmt.Errorf("unable to unpack arguments from extension - classID should be string")
			}

			operationArgs, ok := argsFromPy[2].([]interface{})
			if !ok && argsFromPy[2] != nil {
				return nil, fmt.Errorf("unable to unpack arguments from extension - args should be list")
			}

			bodyParams, ok := argsFromPy[3].(map[string]interface{})
			if !ok && argsFromPy[3] != nil {
				return nil, fmt.Errorf("unable to unpack arguments from extension - bodyParams should be dict")
			}

			queryParams, ok := argsFromPy[4].(map[string]interface{})
			if !ok && argsFromPy[4] != nil {
				return nil, fmt.Errorf("unable to unpack arguments from extension - queryParams should be dict")
			}

			var paginationBatchSize = 0
			if len(argsFromPy) >= 6 {
				if pbs, ok := argsFromPy[5].(int64); ok {
					paginationBatchSize = int(pbs)
				}
			}

			var operationArgsString []string = make([]string, len(operationArgs))
			for i, v := range operationArgs {
				operationArgsString[i] = fmt.Sprintf("%v", v)
			}

			queryParamsString := map[string]string{}
			for k, v := range queryParams {
				queryParamsString[k] = fmt.Sprintf("%v", v)
			}

			res, err := e.ExecuteOperation(method, classID, operationArgsString, bodyParams, queryParamsString, paginationBatchSize)
			if err != nil {
				log.Tracef("error calling executeOperation: %v", err)
				return nil, fmt.Errorf("error calling executeOperation")
			}

			buf, err := json.Marshal(res)
			if err != nil {
				return nil, fmt.Errorf("JSON marshalling error: %v", err)
			}

			var js interface{}
			err = json.Unmarshal(buf, &js)
			if err != nil {
				return nil, fmt.Errorf("JSON unmarshalling error: %v", err)
			}

			return goToPy(js), nil
		}, 0, "")
		ret = append(ret, method)
	}

	return ret
}

func getString(o map[string]interface{}, k string) string {
	str, ok := o[k]
	if !ok {
		return ""
	}

	if str, ok := str.(string); ok {
		return str
	} else {
		return ""
	}
}

func flagSetToMap(flagset *pflag.FlagSet) map[string]interface{} {
	ret := map[string]interface{}{}

	flagset.Visit(func(f *pflag.Flag) {
		switch f.Value.Type() {
		case "string":
			ret[f.Name], _ = flagset.GetString(f.Name)
		case "bool":
			ret[f.Name], _ = flagset.GetBool(f.Name)
		case "stringSlice":
			ret[f.Name], _ = flagset.GetStringSlice(f.Name)
		case "int64":
			ret[f.Name], _ = flagset.GetInt64(f.Name)
		case "[]int64":
			ret[f.Name], _ = flagset.GetInt64Slice(f.Name)
		case "int32":
			ret[f.Name], _ = flagset.GetInt(f.Name)
		case "float64":
			ret[f.Name], _ = flagset.GetFloat64(f.Name)
		case "float32":
			ret[f.Name], _ = flagset.GetFloat32(f.Name)
		case "json":
			ret[f.Name] = f.Value
		default:
			log.Tracef("Unknown type for %s flag", f.Name)
		}
	})

	return ret
}

func getRunFunction(e *Extension, f string) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		_, err := e.module.Call(f, py.Tuple{goToPy(args), goToPy(flagSetToMap(cmd.Flags()))}, nil)
		if err != nil {
			log.Tracef("Error executing extension function \"%v\": %v", f, err)
		}
	}
}

func safeGetStringFromMapInterface(m map[string]interface{}, k string) string {
	if valIntf, ok := m[k]; ok {
		if val, ok := valIntf.(string); ok {
			return val
		}
	}
	return ""
}

func objectToCommand(e *Extension, o map[string]interface{}) (*cobra.Command, error) {
	cmd := &cobra.Command{}

	cmd.Use = getString(o, "use")
	cmd.Short = getString(o, "short")

	if cmd.Use == "" {
		return nil, fmt.Errorf("Command without \"use\" set")
	}

	run := getString(o, "run")
	if run != "" {
		cmd.Run = getRunFunction(e, run)
	}

	if args, ok := o["args"]; ok {
		if argsMap, ok := args.(map[string]interface{}); ok {
			if exact, ok := argsMap["exact"]; ok {
				if exactInt, ok := exact.(int64); ok {
					cmd.Args = cobra.ExactArgs(int(exactInt))
				}
			}
		}
	}

	if flagsIntf, ok := o["flags"]; ok {
		switch flags := flagsIntf.(type) {
		case map[string]interface{}:
			for flag, flagDefIntf := range flags {
				if flagDef, ok := flagDefIntf.(map[string]interface{}); ok {
					flagType := safeGetStringFromMapInterface(flagDef, "type")
					flagShort := safeGetStringFromMapInterface(flagDef, "short")
					flagUsage := safeGetStringFromMapInterface(flagDef, "usage")

					switch flagType {
					case "bool":
						if flagDefaultIntf, ok := flagDef["default"]; ok {
							if flagDefault, ok := flagDefaultIntf.(bool); ok {
								cmd.Flags().BoolP(flag, flagShort, flagDefault, flagUsage)
							}
						}
					default:
						log.Tracef("Unknown flag type %s", flagType)
					}
				}
			}
		}
	}

	childrenInterface, ok := o["children"]
	if ok {
		if childrenListInterface, ok := childrenInterface.([]interface{}); ok {
			for _, childInterface := range childrenListInterface {
				if child, ok := childInterface.(map[string]interface{}); ok {
					subCmd, err := objectToCommand(e, child)
					if err != nil {
						return nil, err
					}

					cmd.AddCommand(subCmd)
				}
			}
		}
	}

	return cmd, nil
}
