package ast

type ValueType int
type BlockType int
type StatementType int

const (
	_ ValueType = iota
	StringType
	IntType
	FloatType
	BoolType
	SQLType
	CallerType
	VariableType

	_ BlockType = iota
	AssignmentBlock
	CallerBlock

	_ StatementType = iota
	FunctionStatement
	AssignmentStatement
)

type NLang struct {
	Statements []Statement
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

type Caller struct {
	Name   string
	Values []Value
}

type Statement struct {
	Type  StatementType
	Value interface{}
}
