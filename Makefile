generate:
	flex --prefix=yy --header-file=./src/parser/n.lex.h -o ./src/parser/n.lex.c ./src/parser/n.l
	goyacc -o ./src/parser/n.y.go -p "n" ./src/parser/n.y

run:
	make generate
	go run main.go