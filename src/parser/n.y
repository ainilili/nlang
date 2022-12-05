/* Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved. */
/* Use of this source code is governed by a Apache */
/* license that can be found in the LICENSE file. */

/* simplest version of calculator */

%{
package parser

import (
	"fmt"
	"nlang/src/ast"
)

%}

%union {
	id    string
	value ast.Value
	valueType ast.ValueType
	assignment ast.Assignment
	argument ast.Argument
}

%type <value> value
%type <assignment> assignment
%token <id> ID STRING_LITERAL SQL_LITERAL
%token ASSIGN FUNC EOL STRING INT FLOAT BOOL


%%
nlang
	: assignment {
		fmt.Println($1)
	}
	;

function
	: func ID '(' argument ')' {

	}



assignment
	: ID ASSIGN expression{
		$$ = ast.Assignment{Parameter: $1, Expression: $3}
	}
	;

argument
	: ID value_type

value
	: STRING_LITERAL {
		$$ = ast.Value{Type: ast.String, Value:$1}
	}
	| SQL_LITERAL {
		$$ = ast.Value{Type: ast.SQL, Value:$1}
	}
	;

value_type
	: STRING { $$ = ast.String }
	| INT { $$ = ast.Int }
	| FLOAT { $$ = ast.Float }
	| BOOL { $$ = ast.Bool }
%%
