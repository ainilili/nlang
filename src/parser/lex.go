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

// The parser calls this method to get each new token. This
// implementation returns operators and NUM.
func (p *nLex) Lex(yylval *nSymType) int {
	p.lastErr = nil

	var tok = C.yylex()

	p.yylineno = int(C.yylineno)
	p.yytext = C.GoString(C.yytext)

	switch tok {
	case C.ID:
		yylval.id = p.yytext
		return ID

	case C.NUMBER:
		yylval.value, _ = strconv.Atoi(p.yytext)
		return NUMBER

	case C.SUB:
		return SUB
	case C.MUL:
		return MUL
	case C.DIV:
		return DIV
	case C.ABS:
		return ABS

	case C.LPAREN:
		return LPAREN
	case C.RPAREN:
		return RPAREN
	case C.ASSIGN:
		return ASSIGN

	case C.EOL:
		return EOL
	case C.CONNECT:
		return CONNECT
	}

	if tok == C.ILLEGAL {
		log.Printf("lex: ILLEGAL token, yytext = %q, yylineno = %d", p.yytext, p.yylineno)
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
