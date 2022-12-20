package runtime

import (
	"fmt"
	"nlang/src/ast"
)

type Runtime struct {
	ast       ast.NLang
	root      Stack
	functions map[string]*ast.Function
}

func New(nl ast.NLang) Runtime {
	return Runtime{
		ast:       nl,
		root:      Stack{variables: map[string]*ast.Value{}},
		functions: map[string]*ast.Function{},
	}
}

func (r *Runtime) Main() error {
	root := r.root
	for _, stat := range r.ast.Statements {
		switch stat.Type {
		case ast.AssignmentStatement:
			ass := stat.Value.(ast.Assignment)
			root.variables[ass.Variable] = &ass.Value
		case ast.FunctionStatement:
			fn := stat.Value.(ast.Function)
			r.functions[fn.Name] = &fn
		}
	}
	main, err := r.getFunction("main")
	if err != nil {
		return err
	}
	_, err = r.Call(main)
	return err
}

func (r *Runtime) Call(fn *ast.Function, args ...*ast.Value) (*ast.Value, error) {
	if len(fn.Args) != len(args) {
		return nil, fmt.Errorf("func %s, not matched arguments, expected %d, actual: %d", fn.Name, len(fn.Args), len(args))
	}
	tempVariables := r.copyRootVariables()
	for i, arg := range fn.Args {
		tempVariables[arg.Name] = args[i]
	}
	for _, block := range fn.Blocks {
		switch block.Type {
		case ast.AssignmentBlock:
			ass := block.Value.(ast.Assignment)
			val, err := r.parseValue(tempVariables, &ass.Value)
			if err != nil {
				return nil, err
			}
			tempVariables[ass.Variable] = val
		case ast.CallerBlock:
			_, err := r.parseCaller(tempVariables, block.Value.(ast.Caller))
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

func (r *Runtime) copyRootVariables() map[string]*ast.Value {
	tempVariables := map[string]*ast.Value{}
	for key, val := range r.root.variables {
		tempVariables[key] = val
	}
	return tempVariables
}

func (r *Runtime) getFunction(name string) (*ast.Function, error) {
	fn, ok := r.functions[name]
	if !ok {
		return fn, fmt.Errorf("not defined func %s. ", name)
	}
	return fn, nil
}

func (r *Runtime) parseValue(tempVariables map[string]*ast.Value, val *ast.Value) (*ast.Value, error) {
	switch val.Type {
	case ast.StringType, ast.IntType, ast.FloatType,
		ast.BoolType, ast.SQLType:
		return val, nil
	case ast.CallerType:
		return r.parseCaller(tempVariables, val.Value.(ast.Caller))
	case ast.VariableType:
		arg, ok := tempVariables[val.Value.(string)]
		if !ok {
			return nil, fmt.Errorf("required arg %v is missing! ", val.Value)
		}
		return arg, nil
	default:
		return val, fmt.Errorf("unknown val type: %v. ", val.Type)
	}
}

func (r *Runtime) parseCaller(tempVariables map[string]*ast.Value, caller ast.Caller) (*ast.Value, error) {
	if caller.Name == "print" {
		for i := range caller.Values {
			val, err := r.parseValue(tempVariables, &caller.Values[i])
			if err != nil {
				return nil, err
			}
			fmt.Println(val.Value)
		}
		return nil, nil
	}
	fn, err := r.getFunction(caller.Name)
	if err != nil {
		return nil, err
	}
	args := make([]*ast.Value, 0)
	for i := range caller.Values {
		val, err := r.parseValue(tempVariables, &caller.Values[i])
		if err != nil {
			return nil, err
		}
		args = append(args, val)
	}
	return r.Call(fn, args...)
}
