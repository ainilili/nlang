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
	number int
	value ast.Value
	value_list []ast.Value
	value_type ast.ValueType
	assignment ast.Assignment
	argument_list []ast.Argument
	argument ast.Argument
	function ast.Function
	function_list []ast.Function
	block ast.Block
	block_list []ast.Block
	caller ast.Caller
	statement ast.Statement
	statement_list []ast.Statement
}

%type <value> value
%type <assignment> assignment
%type <value_type> value_type
%type <function> function prefix_function postfix_function
%type <argument> argument
%type <argument_list> argument_list
%type <block> block
%type <block_list> block_list
%type <caller> caller
%type <value_list> value_list
%type <statement> statement
%type <statement_list> statement_list

%token <id> ID STRING_LITERAL SQL_LITERAL
%token <number> INT_LITERAL

%token ASSIGN FUNC EOL STRING INT FLOAT BOOL SQL


%%
nlang
	: statement_list {
		Ast = ast.NLang{Statements: $1}
	}
	;

statement
	: function {
		$$ = ast.Statement{Type: ast.FunctionStatement, Value: $1}
	}
	| assignment {
		$$ = ast.Statement{Type: ast.AssignmentStatement, Value: $1}
	}
	;

statement_list
	: statement {
		$$ = []ast.Statement{$1}
	}
	| statement_list EOL statement {
		$$ = append($1, $3)
	}
	;

function
	: postfix_function '{' '}'{
		$$ = $1
	}
	| postfix_function '{' EOL block_list EOL '}'{
		$1.Blocks = $4
		$$ = $1
	}

postfix_function
	: prefix_function {
		$$ = $1
	}
	| prefix_function value_type {
		$1.Type = $2
		$$ = $1
	}
	;

prefix_function
	: FUNC ID '(' ')'{
		$$ = ast.Function{Name: $2}
	}
	| FUNC ID '(' argument_list ')'{
		$$ = ast.Function{Name: $2, Args: $4}
	}
	;

block_list
	: block {
		$$ = []ast.Block{$1}
	}
	| block_list EOL block {
		$$ = append($1, $3)
	}
	;

block
	: assignment{
		$$ = ast.Block{Type: ast.AssignmentBlock, Value: $1}
	}
	| caller {
		$$ = ast.Block{Type: ast.CallerBlock, Value: $1}
	}
	;

argument_list
	: argument {
		$$ = []ast.Argument{$1}
	}
	| argument_list ',' argument{
		$$ = append($1, $3)
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
		$$ = ast.Value{Type: ast.StringType, Value:$1}
	}
	| SQL_LITERAL {
		$$ = ast.Value{Type: ast.SQLType, Value:$1}
	}
	| INT_LITERAL {
		$$ = ast.Value{Type: ast.IntType, Value:$1}
	}
	| caller {
		$$ = ast.Value{Type: ast.CallerType, Value: $1}
	}
	;

value_list
	: value {
		$$ = []ast.Value{$1}
	}
	| value_list ',' value {
		$$ = append($1, $3)
	}
	;

value_type
	: STRING { $$ = ast.StringType }
	| INT { $$ = ast.IntType }
	| FLOAT { $$ = ast.FloatType }
	| BOOL { $$ = ast.BoolType }
	| SQL { $$ = ast.SQLType }
	;

caller
	: ID '(' ')'{
		$$ = ast.Caller{Name: $1}
	}
	| ID '(' value_list ')'{
		$$ = ast.Caller{Name: $1, Values: $3}
        }
	;
%%
