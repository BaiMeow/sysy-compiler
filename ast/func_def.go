package ast

import (
	"sysy/ast/types"
)

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
