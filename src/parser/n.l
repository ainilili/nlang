/* Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved. */
/* Use of this source code is governed by a Apache */
/* license that can be found in the LICENSE file. */

%option noyywrap

%{
#include "tok.h"
%}
%%


= 				{ return ASSIGN; }
L?\"(\\.|\"|[^\\"\n])*\"	{ return STRING_LITERAL; }
L?\`(\\.|\"|[^\\"\n])*\`	{ return SQL_LITERAL; }
\n				{ return EOL; }
"("				{ return '('; }
")"				{ return ')'; }
","				{ return ','; }
"{"				{ return '{'; }
"}"				{ return '}'; }


"func" 				{ return FUNC;}
"int" 				{ return INT; }
"float" 			{ return FLOAT; }
"string" 			{ return STRING; }
"bool" 				{ return BOOL; }

[_a-zA-Z]+ 			{ return ID; }

%%

