package ast

import "sysy/ast/types"

type Declare struct {
	Definitions []*Define
}

type Define struct {
	Type         types.Type
	Identifier   string
	InitialValue Expr
}

type FuncDef struct {
	Identifier string
	Params     []*FuncParam
	Return     types.Type
	Body       any
	Type       types.Func `draw:"ignore"`
}

type FuncParam struct {
	Identifier string
	Type       types.Type
}
