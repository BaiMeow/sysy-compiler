package types

import "github.com/antlr4-go/antlr/v4"

func ParseBase(node antlr.TerminalNode, Const bool) *Base {
	switch node.GetText() {
	case "int":
		return &Base{Type: Int, Const: Const}
	case "float":
		return &Base{Type: Float, Const: Const}
	case "void":
		return &Base{Type: Void, Const: Const}
	}
	return nil
}
