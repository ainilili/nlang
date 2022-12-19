package runtime

import (
	"errors"
	"fmt"
	"nlang/src/ast"
)

type Runtime struct {
	ast       ast.NLang
	root      Stack
	functions map[string]ast.Function
}

func New(nl ast.NLang) Runtime {
	return Runtime{ast: nl, root: Stack{variables: map[string]ast.Value{}}}
}

func (r *Runtime) Run() error {
	root := r.root
	for _, stat := range r.ast.Statements {
		switch stat.Type {
		case ast.AssignmentStatement:
			ass := stat.Value.(ast.Assignment)
			root.variables[ass.Variable] = ass.Value
		case ast.FunctionStatement:
			fn := stat.Value.(ast.Function)
			r.functions[fn.Name] = fn
		}
	}
	main, ok := r.functions["main"]
	if !ok {
		return errors.New("not defined main. ")
	}
	_, err := r.Call(main)
	return err
}

func (r *Runtime) Call(fn ast.Function, args ...ast.Value) (*ast.Value, error) {
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
			tempVariables[ass.Variable] = ass.Value
		case ast.CallerBlock:
			//caller := block.Value.(ast.Caller)
			//for _, a := range caller.Values {
			//
			//}
		}
	}
	return nil, nil
}

func (r *Runtime) copyRootVariables() map[string]ast.Value {
	tempVariables := map[string]ast.Value{}
	for key, val := range r.root.variables {
		tempVariables[key] = val
	}
	return tempVariables
}
