package visitor

import (
	"github.com/antlr4-go/antlr/v4"
	"sysy/ast"
	"sysy/ast/types"
	"sysy/parser"
)

func (c *Context) FuncType(pctx parser.IFuncTypeContext) (types.Type, error) {
	if Type := pctx.Type(); Type != nil {
		return c.Type(Type)
	}
	if Void := pctx.Void(); Void != nil {
		return &types.Base{Type: types.Void}, nil
	}
	return nil, Invalid(pctx.GetStart(), "invalid func")
}

func (c *Context) Type(pctx antlr.TerminalNode) (types.Type, error) {
	switch pctx.GetText() {
	case "int":
		return &types.Base{Type: types.Int}, nil
	case "float":
		return &types.Base{Type: types.Float}, nil
	default:
		return nil, Invalid(pctx.GetSymbol(), "invalid type")
	}
}

func (c *Context) FuncFParams(pctx parser.IFuncFParamsContext) ([]*ast.FuncParam, error) {
	paramNodes := pctx.AllFuncFParam()
	var params []*ast.FuncParam
	for _, param := range paramNodes {
		astParam, err := c.FuncFParam(param)
		if err != nil {
			return nil, err
		}
		params = append(params, astParam)
	}
	return params, nil
}

func (c *Context) FuncFParam(pctx parser.IFuncFParamContext) (*ast.FuncParam, error) {
	baseType, err := c.Type(pctx.Type())
	if err != nil {
		return nil, err
	}

	ident := pctx.Identifier().GetText()

	constExpNodes := pctx.AllConstExp()
	if len(constExpNodes) == 0 {
		return &ast.FuncParam{
			Identifier: ident,
			Type:       baseType,
		}, nil
	}

	var dims []int
	for _, constExpNode := range constExpNodes {
		dim, err := c.ConstExp(constExpNode)
		if err != nil {
			return nil, err
		}

		dimIndex, ok := dim.Value.(int)
		if !ok {
			return nil, Invalid(constExpNode.GetStart(), "invalid array index type")
		}
		dims = append(dims, dimIndex)
	}

	return &ast.FuncParam{
		Identifier: ident,
		Type: &types.Array{
			ElementType: baseType,
			Shape:       dims,
		},
	}, nil
}
