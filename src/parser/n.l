/* Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved. */
/* Use of this source code is governed by a Apache */
/* license that can be found in the LICENSE file. */

%option noyywrap

%{
#include "const.h"
%}
%%

L?\"(\\.|[^\\"\n])*\"	{
	printf("%s\n", yytext);
	return STR;
}

%%

