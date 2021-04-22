package golang

import (
	"fmt"
	"go/types"

	"github.com/pkg/errors"
	"github.com/webrpc/webrpc/schema"
)

func (p *parser) parsePkgInterfaces(scope *types.Scope) error {
	for _, name := range scope.Names() {
		iface, ok := scope.Lookup(name).Type().Underlying().(*types.Interface)
		if !ok {
			continue
		}

		service := &schema.Service{
			Name:   schema.VarName(name),
			Schema: p.schema, // denormalize/back-reference
		}

		fmt.Printf("interface %v {\n", name)
		defer fmt.Printf("}\n")

		// Loop over the interface's methods.
		for i := 0; i < iface.NumMethods(); i++ {
			method := iface.Method(i)
			if !method.Exported() {
				continue
			}

			methodName := method.Id()
			fmt.Printf("  %v\n", methodName)

			methodSignature, ok := method.Type().(*types.Signature)
			if !ok {
				return errors.Errorf("interface %v method %v(): failed to get method signature", name, methodName)
			}

			methodParams := methodSignature.Params()
			inputs, err := p.getMethodArguments(methodParams, true)
			if err != nil {
				return errors.Wrapf(err, "interface %v method %v(): failed to get inputs (method arguments)", name, methodName)
			}

			// First method argument must be of type context.Context.
			if methodParams.Len() == 0 {
				return errors.Errorf("interface %v method %v(): first method argument must be context.Context: no arguments defined", name, methodName)
			}
			if err := ensureContextType(methodParams.At(0).Type()); err != nil {
				return errors.Wrapf(err, "interface %v method %v(): first method argument must be context.Context", name, methodName)
			}
			inputs = inputs[1:] // Cut it off. The gen/golang adds context.Context as first method argument automatically.

			methodResults := methodSignature.Results()
			outputs, err := p.getMethodArguments(methodResults, false)
			if err != nil {
				return errors.Wrapf(err, "interface %v method %v(): failed to get outputs (method results)", name, methodName)
			}

			// Last method return value must be of type error.
			if methodResults.Len() == 0 {
				return errors.Errorf("interface %v method %v(): last return value must be context.Context: no return values defined", name, methodName)
			}
			if err := ensureErrorType(methodResults.At(methodResults.Len() - 1).Type()); err != nil {
				return errors.Wrapf(err, "interface %v method %v(): first method argument must be context.Context", name, methodName)
			}
			outputs = outputs[:len(outputs)-1] // Cut it off. The gen/golang adds error as a last return value automatically.

			service.Methods = append(service.Methods, &schema.Method{
				Name:    schema.VarName(methodName),
				Inputs:  inputs,
				Outputs: outputs,
				Service: service, // denormalize/back-reference
			})
		}

		p.schema.Services = append(p.schema.Services, service)
	}

	return nil
}

func (p *parser) getMethodArguments(params *types.Tuple, isInput bool) ([]*schema.MethodArgument, error) {
	var args []*schema.MethodArgument

	for i := 0; i < params.Len(); i++ {
		param := params.At(i)
		typ := param.Type()

		name := param.Name()
		if name == "" {
			if isInput {
				name = fmt.Sprintf("arg%v", i) // 0 is `ctx context.Context`
			} else {
				name = fmt.Sprintf("ret%v", i+1)
			}
		}

		varType, err := p.parseType(typ) // Type name will be resolved deeper down the stack.
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse argument %v", name)
		}

		arg := &schema.MethodArgument{
			Name:      schema.VarName(name),
			Type:      varType,
			InputArg:  isInput,  // denormalize/back-reference
			OutputArg: !isInput, // denormalize/back-reference
		}

		args = append(args, arg)
	}

	return args, nil
}

func ensureContextType(typ types.Type) (err error) {
	namedType, ok := typ.(*types.Named)
	if !ok {
		return errors.Errorf("expected named type: found type %T (%+v)", typ, typ)
	}

	if _, ok := namedType.Underlying().(*types.Interface); !ok {
		return errors.Errorf("expected underlying interface: found type %T (%+v)", typ, typ)
	}

	pkgName := namedType.Obj().Pkg().Name()
	typeName := namedType.Obj().Name()

	if pkgName != "context" && typeName != "Context" {
		return errors.Errorf("expected context.Context: found %v.%v", pkgName, typeName)
	}

	return nil
}

func ensureErrorType(typ types.Type) (err error) {
	namedType, ok := typ.(*types.Named)
	if !ok {
		return errors.Errorf("expected named type: found type %T (%+v)", typ, typ)
	}

	if _, ok := namedType.Underlying().(*types.Interface); !ok {
		return errors.Errorf("expected underlying interface: found type %T (%+v)", typ, typ)
	}

	pkgName := namedType.Obj().Pkg()
	typeName := namedType.Obj().Name()

	if pkgName != nil && typeName != "error" {
		return errors.Errorf("expected error: found %v.%v", pkgName, typeName)
	}

	return nil
}
