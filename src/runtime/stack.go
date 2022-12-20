package runtime

import "nlang/src/ast"

type Stack struct {
	variables map[string]*ast.Value
}
