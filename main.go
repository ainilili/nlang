package main

import (
	"fmt"
	"nlang/src/ast"
	"nlang/src/parser"
)

func main() {
	//str := "connect abc\na = `select * from user`"
	//str := "a = \"ab\t \n\"\"c\""
	str := `func hello(a int, b string){
	a = "abc"
	b = "23"
	c = 100
}`
	parser.Parse(parser.NewNLex([]byte(str)))

	for _, fn := range parser.Ast.Functions {
		fmt.Print("func ", fn.Name, "(")
		for _, arg := range fn.Args {
			fmt.Print(arg.Name, arg.Type)
		}
		fmt.Println("){")
		for _, block := range fn.Blocks {

			switch block.Type {
			case ast.AssignmentBlock:
				ass := block.Value.(ast.Assignment)
				fmt.Printf("    %v = %v\n", ass.Variable, ass.Value.Value)
			}
		}
		fmt.Println("}")
	}
}
