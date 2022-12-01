```shell
$ flex --prefix=yy --header-file=n.lex.h -o n.lex.c n.l
$ goyacc -o n.y.go -p "n" n.y
```