package visitor

import (
	"sysy/ast"
	"sysy/ast/types"
	"sysy/parser"
)

func (c *Context) Decl(pctx parser.IDeclContext) (*ast.Declare, error) {
	if decl := pctx.ConstDecl(); decl != nil {
		return c.ConstDecl(decl)
	}

	if decl := pctx.VarDecl(); decl != nil {
		return c.VarDecl(decl)
	}

	return nil, Invalid(pctx.GetStart(), "Invalid Decl")
}

func (c *Context) ConstDecl(pctx parser.IConstDeclContext) (*ast.Declare, error) {
	decl := &ast.Declare{}
	Type := types.ParseBase(pctx.Type(), true)
	if Type == nil {
		return nil, Invalid(pctx.Type().GetSymbol(), "Invalid Type")
	}

	for _, def := range pctx.AllConstDef() {
		def, err := c.ConstDef(def, Type)
		if err != nil {
			return nil, err
		}
		decl.Definitions = append(decl.Definitions, def)
	}

	return decl, nil
}

func (c *Context) VarDecl(pctx parser.IVarDeclContext) (*ast.Declare, error) {
	decl := &ast.Declare{}
	Type := types.ParseBase(pctx.Type(), false)
	if Type == nil {
		return nil, Invalid(pctx.Type().GetSymbol(), "Invalid Type")
	}
	for _, def := range pctx.AllVarDef() {
		def, err := c.VarDef(def, Type)
		if err != nil {
			return nil, err
		}
		decl.Definitions = append(decl.Definitions, def)
	}
	return decl, nil
}
