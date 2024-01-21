package ast

import "sysy/ast/types"

type Declare struct {
	Definitions []*Define
}

type Define struct {
	Type         types.Type
	Identifier   string
	InitialValue Expression
}
