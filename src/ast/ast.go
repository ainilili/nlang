package ast

type NLang struct {
	Connection *Connection
	Functions  []Function
}

type Connection struct {
	URL string
}

type Function struct {
}
