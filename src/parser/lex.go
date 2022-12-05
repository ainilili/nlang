// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package parser

//#include "tok.h"
//#include "n.lex.h"
import "C"

import (
	"errors"
	"log"
	"strconv"
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

func (p *nLex) Lex(yylval *nSymType) int {
	p.lastErr = nil

	var tok = C.yylex()

	p.yylineno = int(C.yylineno)
	p.yytext = C.GoString(C.yytext)

	switch tok {
	case C.ASSIGN:
		return ASSIGN
	case C.ID:
		yylval.id = p.yytext
		return ID
	case C.STRING_LITERAL:
		yylval.id = p.yytext
		return STRING_LITERAL
	case C.SQL_LITERAL:
		yylval.id = p.yytext
		return SQL_LITERAL
	case C.INT_LITERAL:
		num, _ := strconv.Atoi(p.yytext)
		yylval.number = num
		return INT_LITERAL
	case C.FUNC:
		return FUNC
	case C.EOL:
		return EOL
	case C.STRING:
		return STRING
	case C.INT:
		return INT
	case C.FLOAT:
		return FLOAT
	case C.BOOL:
		return BOOL
	}
	return int(tok)
}

// The parser calls this method on a parse error.
func (p *nLex) Error(s string) {
	p.lastErr = errors.New("yacc: " + s)
	if err := p.lastErr; err != nil {
		log.Println(err)
	}
}
