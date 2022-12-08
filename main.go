package main

import (
	_ "embed"
	"fmt"
	"nlang/src/ast"
	"nlang/src/parser"
)

var (
	//go:embed test.n
	code []byte
)

func main() {
	//str := "connect abc\na = `select * from user`"
	//str := "a = \"ab\t \n\"\"c\""
	parser.Parse(parser.NewNLex(code))
	for _, fn := range parser.Ast.Functions {
		fmt.Print("func ", fn.Name, "(")
		for _, arg := range fn.Args {
			fmt.Print(arg.Name, " ", arg.Type, ",")
		}
		fmt.Println("){")
		for _, block := range fn.Blocks {
			switch block.Type {
			case ast.AssignmentBlock:
				ass := block.Value.(ast.Assignment)
				fmt.Printf("    %v = %v\n", ass.Variable, ass.Value.Value)
			case ast.CallerBlock:
				caller := block.Value.(ast.Caller)
				fmt.Printf("    %v(%v)\n", caller.Name, caller.Values)
			}
		}
		fmt.Println("}")
	}
}
