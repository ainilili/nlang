/* Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved. */
/* Use of this source code is governed by a Apache */
/* license that can be found in the LICENSE file. */

/* simplest version of calculator */

%{
package parser

import (
	"fmt"
)

%}

%union {
	id    string
}

%type <id> expression
%token <id> ID STRING SQL
%token ASSIGN


%%
nlang:
	assign
	;

assign
	: ID ASSIGN expression{
		fmt.Printf("%v = %v\n", $1, $3)
	}
	;

expression
	: STRING
	| SQL
	;

%%
