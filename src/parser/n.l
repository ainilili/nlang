/* Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved. */
/* Use of this source code is governed by a Apache */
/* license that can be found in the LICENSE file. */

%option noyywrap

%{
#include "tok.h"
%}
%%

[_a-zA-Z]+ 			{ return ID; }
= 				{ return ASSIGN; }
L?\"(\\.|\t|\n|\"|[^\\"\n])*\"	{ return STRING_LITERAL; }
L?\`(\\.|\t|\n|\"|[^\\"\n])*\`	{ return SQL_LITERAL; }
\n				{ return EOL; }
"("				{ return '('; }
")"				{ return ')'; }


"func" 				{ return FUNC;}
"int" 				{ return INT; }
"float" 			{ return FLOAT; }
"string" 			{ return STRING; }
"bool" 				{ return BOOL; }

%%

