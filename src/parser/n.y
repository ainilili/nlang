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
	id    string
}

%token <id> STR


%%
nlang:
	STR{
		fmt.Printf("%v1\n", $1)
	}
	;
%%
