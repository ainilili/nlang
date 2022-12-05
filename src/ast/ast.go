package ast

type ValueType int
type ExpressionType int

const (
	_ ValueType = iota
	String
	Int
	Float
	Bool
	SQL

	_ ExpressionType = iota
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
	Variable string
	Value    Value
}

type Value struct {
	Type  ValueType
	Value interface{}
}
