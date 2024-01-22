package ast

type Block struct {
	Children []any `draw:"unfold"`
}

type If struct {
	Condition Expr
	Then      any
}

type IfElse struct {
	Condition Expr
	Then      any
	Else      any
}

type While struct {
	Condition Expr
	Body      any
}

type Break struct{}

type Continue struct{}

type Return struct {
	Value Expr
}

type Assign struct {
	LVal Expr
	Exp  Expr
}
