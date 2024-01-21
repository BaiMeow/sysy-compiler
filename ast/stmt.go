package ast

type Block struct {
	Children []any `draw:"unfold"`
}

type If struct {
	Condition any
	Then      any
}

type IfElse struct {
	Condition any
	Then      any
	Else      any
}

type While struct {
	Condition any
	Body      any
}

type Break struct{}

type Continue struct{}

type Return struct {
	Value any
}

type Assign struct {
	LVal any
}
