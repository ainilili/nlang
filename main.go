package main

import (
	_ "embed"
	"log"
	"nlang/src/parser"
	"nlang/src/runtime"
)

var (
	//go:embed test.n
	code []byte
)

func main() {
	//str := "connect abc\na = `select * from user`"
	//str := "a = \"ab\t \n\"\"c\""
	parser.Parse(parser.NewNLex(code))

	r := runtime.New(parser.Ast)
	log.Fatalln(r.Main())
}
