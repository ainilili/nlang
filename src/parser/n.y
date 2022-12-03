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

%token <id> ID STRING
%token ASSIGN


%%
nlang:
	assign
	;

assign:
	ID STRING{
		fmt.Printf("%v = %v\n", $1, $2)
	}
	;

%%
