package main

import "nlang/src/parser"

func main() {
	str := "connect abc\na = `select * from user`"
	parser.Parse(parser.NewNLex([]byte(str)))
}
