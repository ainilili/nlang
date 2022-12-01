/* Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved. */
/* Use of this source code is governed by a Apache */
/* license that can be found in the LICENSE file. */

/* simplest version of calculator */

%{
package parser

import (
	"fmt"
)

var idValueMap = map[string]interface{}{}
%}

%union {
	value int
	id    string
}

%token <value> NUMBER
%token <id>    ID

%token ADD SUB MUL DIV ABS CONNECT
%token LPAREN RPAREN ASSIGN BACK_QUOTE
%token EOL


%%
nlang:
	CONNECT ID{
		fmt.Printf("connect %v\n", $2)
	}
	| ID ASSIGN BACK_QUOTE ID BACK_QUOTE{
		fmt.Printf("%v = %v", $1, $4)
	}
	;

%%
