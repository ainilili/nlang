package main

import (
	"fmt"
	"nlang/src/parser"
)

func main() {
	//str := "connect abc\na = `select * from user`"
	//str := "a = \"ab\t \n\"\"c\""
	str := `func hello(a int, b string){
	a = "abc"
}`
	parser.Parse(parser.NewNLex([]byte(str)))

	for _, fn := range parser.Ast.Functions {
		fmt.Println(fn.Name)
		for _, arg := range fn.Args {
			fmt.Println(arg.Name, arg.Type)
		}
		for _, block := range fn.Blocks {
			fmt.Println(block.Value, block.Type)
		}
	}
}
