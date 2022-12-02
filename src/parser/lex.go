// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package parser

//#include "const.h"
//#include "n.lex.h"
import "C"

import (
	"errors"
	"log"
)

type nLex struct {
	yylineno int
	yytext   string
	lastErr  error
}

var _ nLexer = (*nLex)(nil)

func NewNLex(data []byte) *nLex {
	p := new(nLex)

	C.yy_scan_bytes(
		(*C.char)(C.CBytes(data)),
		C.yy_size_t(len(data)),
	)

	return p
}

func Parse(nlex nLexer) int {
	return nParse(nlex)
}

// The parser calls this method to get each new token. This
// implementation returns operators and NUM.
func (p *nLex) Lex(yylval *nSymType) int {
	p.lastErr = nil

	var tok = C.yylex()

	p.yylineno = int(C.yylineno)
	p.yytext = C.GoString(C.yytext)

	switch tok {
	case C.STR:
		yylval.id = p.yytext
		return STR
	}

	return 0 // eof
}

// The parser calls this method on a parse error.
func (p *nLex) Error(s string) {
	p.lastErr = errors.New("yacc: " + s)
	if err := p.lastErr; err != nil {
		log.Println(err)
	}
}
