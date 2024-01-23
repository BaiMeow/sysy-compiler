package visitor

import (
	"sysy/ast"
	"sysy/ast/types"
	"sysy/parser"
)

func (c *Context) Decl(pctx parser.IDeclContext) (*ast.Declare, error) {
	if declNode := pctx.ConstDecl(); declNode != nil {
		return c.ConstDecl(declNode)
	}

	if varDeclNode := pctx.VarDecl(); varDeclNode != nil {
		return c.VarDecl(varDeclNode)
	}

	return nil, Invalid(pctx.GetStart(), "Invalid Decl")
}

func (c *Context) ConstDecl(pctx parser.IConstDeclContext) (*ast.Declare, error) {
	decl := &ast.Declare{}
	Type := types.ParseBase(pctx.Type(), true)
	if Type == nil {
		return nil, Invalid(pctx.Type().GetSymbol(), "invalid type")
	}

	for _, constDefNode := range pctx.AllConstDef() {
		constDef, err := c.ConstDef(constDefNode, Type)
		if err != nil {
			return nil, err
		}
		decl.Definitions = append(decl.Definitions, constDef)
	}

	return decl, nil
}

func (c *Context) VarDecl(pctx parser.IVarDeclContext) (*ast.Declare, error) {
	decl := &ast.Declare{}
	Type := types.ParseBase(pctx.Type(), false)
	if Type == nil {
		return nil, Invalid(pctx.Type().GetSymbol(), "invalid type")
	}
	for _, varDefNode := range pctx.AllVarDef() {
		varDef, err := c.VarDef(varDefNode, Type)
		if err != nil {
			return nil, err
		}
		decl.Definitions = append(decl.Definitions, varDef)
	}
	return decl, nil
}
