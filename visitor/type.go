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
	return nil, Invalid(pctx.GetStart(), "Invalid Func")
}

func (c *Context) Type(pctx antlr.TerminalNode) (types.Type, error) {
	switch pctx.GetText() {
	case "int":
		return &types.Base{Type: types.Int}, nil
	case "float":
		return &types.Base{Type: types.Float}, nil
	default:
		return nil, Invalid(pctx.GetSymbol(), "Invalid Type")
	}
}

func (c *Context) FuncFParams(pctx parser.IFuncFParamsContext) ([]*ast.FuncParam, error) {
	params := pctx.AllFuncFParam()
	var astParams []*ast.FuncParam
	for _, param := range params {
		astParam, err := c.FuncFParam(param)
		if err != nil {
			return nil, err
		}
		astParams = append(astParams, astParam)
	}
	return astParams, nil
}

func (c *Context) FuncFParam(pctx parser.IFuncFParamContext) (*ast.FuncParam, error) {
	BaseType, err := c.Type(pctx.Type())
	if err != nil {
		return nil, err
	}

	ident := pctx.Identifier().GetText()

	constExps := pctx.AllConstExp()
	if len(constExps) == 0 {
		return &ast.FuncParam{
			Identifier: ident,
			Type:       BaseType,
		}, nil
	}

	var dims []int
	for _, exp := range constExps {
		dim, err := c.ConstExp(exp)
		if err != nil {
			return nil, err
		}

		dimIndex, ok := dim.Value.(int)
		if !ok {
			return nil, Invalid(exp.GetStart(), "Invalid array index type")
		}
		dims = append(dims, dimIndex)
	}

	return &ast.FuncParam{
		Identifier: ident,
		Type: &types.Array{
			ElementType: BaseType,
			Shape:       dims,
		},
	}, nil
}
