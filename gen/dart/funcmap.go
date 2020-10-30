package dart

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/webrpc/webrpc/schema"
)

var fieldTypeMap = map[schema.DataType]string{
	schema.T_Uint:      "int",
	schema.T_Uint8:     "int",
	schema.T_Uint16:    "int",
	schema.T_Uint32:    "int",
	schema.T_Uint64:    "int",
	schema.T_Int:       "int",
	schema.T_Int8:      "int",
	schema.T_Int16:     "int",
	schema.T_Int32:     "int",
	schema.T_Int64:     "int",
	schema.T_Float32:   "double",
	schema.T_Float64:   "double",
	schema.T_String:    "String",
	schema.T_Timestamp: "DateTime",
	schema.T_Null:      "null",
	schema.T_Any:       "dynamic",
	schema.T_Byte:      "String",
	schema.T_Bool:      "bool",
}

func fieldType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_Map:
		typK, ok := fieldTypeMap[in.Map.Key]
		if !ok {
			return "", fmt.Errorf("unknown type mapping %v", in.Map.Key)
		}
		typV, err := fieldType(in.Map.Value)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Map<%s, %s>", typK, typV), nil

	case schema.T_List:
		z, err := fieldType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("List<%s>", z), nil

	case schema.T_Struct:
		return in.Struct.Name, nil

	default:
		if fieldTypeMap[in.Type] != "" {
			return fieldTypeMap[in.Type], nil
		}
	}
	return "", fmt.Errorf("could not represent type: %#v", in)
}

func downcaseName(v interface{}) (string, error) {
	downFn := func(s string) string {
		if s == "" {
			return ""
		}
		return strings.ToLower(s[0:1]) + s[1:]
	}
	switch t := v.(type) {
	case schema.VarName:
		return downFn(string(t)), nil
	case string:
		return downFn(t), nil
	default:
		return "", errors.New("downcaseFieldName, unknown arg type")
	}
}

func makeLowerCase(in *schema.MessageField) (string, error) {
	return strings.ToLower(string(in.Name)), nil
}
func methodName(in interface{}) string {
	v, _ := downcaseName(in)
	return v
}
func isEnum(t schema.MessageType) bool {
	return t == "enum"
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

func lastField(msg schema.Message) schema.MessageField {
	return *msg.Fields[len(msg.Fields)-1]
}

func exportableField(in schema.MessageField) bool {
	for _, meta := range in.Meta {
		for k := range meta {
			if k == "json" {
				if meta[k] == "-" {
					return false
				}
			}
		}
	}
	return true
}

func serviceClassName(in schema.VarName) (string, error) {
	s := string(in)
	return s, nil
}

func exportedJSONField(in schema.MessageField) (string, error) {
	for _, meta := range in.Meta {
		for k := range meta {
			if k == "json" {
				s := strings.Split(fmt.Sprintf("%v", meta[k]), ",")
				return s[0], nil
			}
		}
	}
	return string(in.Name), nil
}

func jsonKey(in schema.MessageField) (string, error) {
	field, err := exportedJSONField(in)
	if field != string(in.Name) {
		return fmt.Sprintf("@JsonKey(name: '%s')", field), err
	}
	return "", nil
}

func constPathPrefix(in schema.VarName) (string, error) {
	return string(in) + "PathPrefix", nil
}

func methodInputName(in *schema.MethodArgument) string {
	name := string(in.Name)
	if name != "" {
		return name
	}
	if in.Type != nil {
		return in.Type.String()
	}
	return ""
}

func methodInputType(in *schema.MethodArgument) string {
	z, _ := fieldType(in.Type)
	return z
}

func methodArgumentInputClassName(in *schema.Method) string {
	if len(in.Service.Schema.Services) == 1 {
		return fmt.Sprintf("%s%s", in.Name, "Args")
	} else {
		return fmt.Sprintf("%s%s%s", in.Service.Name, in.Name, "Args")
	}
}

func methodNameUsed(in *schema.Method) bool {
	for _, srvc := range in.Service.Schema.Services {
		if srvc.Name != in.Service.Name {
			for _, meth := range srvc.Methods {
				if meth.Name == in.Name {
					return true
				}
			}
		}
	}
	return false
}

func isFirstService(in *schema.Method) bool {
	for i, srvc := range in.Service.Schema.Services {
		if i == 0 && srvc.Name == in.Service.Name {
			return true
		}
	}
	return false
}

func methodArgumentOutputClassName(in *schema.Method) string {
	if !methodNameUsed(in) || isFirstService(in) {
		return fmt.Sprintf("%s%s", in.Name, "Result")
	} else {
		return fmt.Sprintf("%s%s%s", in.Service.Name, in.Name, "Result")
	}
}

func clientMethodInputs(in *schema.Method) (string, error) {
	inputs := []string{}
	if len(in.Inputs) > 0 {
		inputs = append(inputs, fmt.Sprintf("@required %s args", methodArgumentInputClassName(in)))
		inputs = append(inputs, "Map<String, String> headers,")
		return fmt.Sprintf("{%s}", strings.Join(inputs, ", ")), nil
	}
	return fmt.Sprintf("{%s}", "Map<String, String> headers"), nil
}

func serverMethodInputs(in *schema.Method) (string, error) {
	inputs := []string{}
	if len(in.Inputs) > 0 {
		inputs = append(inputs, fmt.Sprintf("@required %s params", methodArgumentInputClassName(in)))
		return fmt.Sprintf("{%s}", strings.Join(inputs, ", ")), nil
	}
	return "", nil
}

func methodOutputs(in *schema.Method) (string, error) {
	if len(in.Outputs) == 0 {
		return fmt.Sprintf("FutureOr<%s>", "void"), nil
	}
	return fmt.Sprintf("FutureOr<%s>", methodArgumentOutputClassName(in)), nil
}

func methodOutputsClient(in *schema.Method) (string, error) {
	if len(in.Outputs) == 0 {
		return fmt.Sprintf("Stream<RpcState<%s>>", "void"), nil
	}
	return fmt.Sprintf("Stream<RpcState<%s>>", methodArgumentOutputClassName(in)), nil
}

func methodOutputsClientBloc(in *schema.Method) (string, error) {
	return fmt.Sprintf("Stream<RpcState<%s>>", methodArgumentOutputClassName(in)), nil
}
func blocStateFactory(in *schema.Method) (string, error) {
	return downcaseName(methodArgumentOutputClassName(in))
}

func serviceImplName(in schema.VarName) (string, error) {
	return downcaseName(string(in))
}

func useFlutter(in string) bool {
	return in == "flutter"
}

func useFlutterBloc(in string) bool {
	return in == "flutter_bloc"
}

func useAngularBloc(in string) bool {
	return in == "angular_bloc"
}

func useBloc(in string) bool {
	return in == "angular_bloc" || in == "flutter_bloc"
}

func templateFuncMap(proto *schema.WebRPCSchema) map[string]interface{} {
	return map[string]interface{}{
		"fieldType":                     fieldType,
		"methodName":                    methodName,
		"isEnum":                        isEnum,
		"exportedJSONField":             exportedJSONField,
		"exportableField":               exportableField,
		"isStruct":                      isStruct,
		"lastField":                     lastField,
		"jsonKey":                       jsonKey,
		"serviceClassName":              serviceClassName,
		"methodInputName":               methodInputName,
		"methodInputType":               methodInputType,
		"methodArgumentInputClassName":  methodArgumentInputClassName,
		"methodArgumentOutputClassName": methodArgumentOutputClassName,
		"constPathPrefix":               constPathPrefix,
		"clientMethodInputs":            clientMethodInputs,
		"serverMethodInputs":            serverMethodInputs,
		"methodOutputs":                 methodOutputs,
		"downcaseName":                  downcaseName,
		"makeLowerCase":                 makeLowerCase,
		"serviceImplName":               serviceImplName,
		"useFlutter":                    useFlutter,
		"methodOutputsClient":           methodOutputsClient,
		"useFlutterBloc":                useFlutterBloc,
		"useAngularBloc":                useAngularBloc,
		"useBloc":                       useBloc,
		"methodOutputsClientBloc":       methodOutputsClientBloc,
		"blocStateFactory":              blocStateFactory,
	}
}
