/* Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved. */
/* Use of this source code is governed by a Apache */
/* license that can be found in the LICENSE file. */

/* simplest version of calculator */

%{
package parser

import (
	"nlang/src/ast"
)
var Ast ast.NLang
%}

%union {
	id    string
	value ast.Value
	value_type ast.ValueType
	assignment ast.Assignment
	argument ast.Argument
	function ast.Function
}

%type <value> value
%type <assignment> assignment
%type <value_type> value_type
%type <function> function
%type <argument> argument

%token <id> ID STRING_LITERAL SQL_LITERAL
%token ASSIGN FUNC EOL STRING INT FLOAT BOOL '(' ')'


%%
nlang
	: function {
		Ast = ast.NLang{Functions: []ast.Function{$1}}
	}
	;

function
	: FUNC ID '(' ')'{
		$$ = ast.Function{Name: $2}
	}
	| FUNC ID '(' argument ')'{
		$$ = ast.Function{Name: $2, Args: []ast.Argument{$4}}
	}
	;

argument
	: ID value_type {
		$$ = ast.Argument{Name: $1, Type: $2}
	}
	;

assignment
	: ID ASSIGN value{
		$$ = ast.Assignment{Variable: $1, Value: $3}
	}
	;

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
	;
%%
