/* Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved. */
/* Use of this source code is governed by a Apache */
/* license that can be found in the LICENSE file. */

%option noyywrap

%{
#include "const.h"
%}

%%

[c][o][n][n][e][c][t]    { return CONNECT; }

`  {return BACK_QUOTE;}
= 	{return ASSIGN;}
[_a-zA-Z]+ { return ID; }
[0-9]+     { return NUMBER; }

\n     { return EOL; }
[ \t]  { /* ignore whitespace */ }

%%

