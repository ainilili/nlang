package ast

type ValueType int
type BlockType int

const (
	_ ValueType = iota
	String
	Int
	Float
	Bool
	SQL

	_ BlockType = iota
	AssignmentBlock
)

type NLang struct {
	Functions []Function
}

type Function struct {
	Name   string
	Args   []Argument
	Type   ValueType
	Blocks []Block
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

type Block struct {
	Type  BlockType
	Value interface{}
}
