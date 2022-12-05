package ast

type ValueType int

const (
	_ ValueType = iota
	String
	Int
	Float
	Bool
	SQL
)

type NLang struct {
	Functions []Function
}

type Function struct {
	Name    string
	Args    []Argument
	Results interface{}
}

type Argument struct {
	Type ValueType
	Name string
}

type Assignment struct {
	Variable   string
	Expression Expression
}

type Value struct {
	Type  ValueType
	Value interface{}
}
