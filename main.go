package main

import "nlang/src/parser"

func main() {
	//str := "connect abc\na = `select * from user`"
	str := "\"abc\""
	parser.Parse(parser.NewNLex([]byte(str)))
}
