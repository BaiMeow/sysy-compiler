package visitor

import (
	"sysy/ast"
	"sysy/parser"
)

func (c *Context) Comp(pctx parser.ICompContext) (*ast.Program, error) {
	program := &ast.Program{}
	for _, v := range pctx.AllCompUnit() {
		u, err := c.CompUnit(v)
		if err != nil {
			return nil, err
		}
		program.Children = append(program.Children, u)
	}
	return program, nil
}

func (c *Context) CompUnit(pctx parser.ICompUnitContext) (any, error) {
	if decl := pctx.Decl(); decl != nil {
		return c.Decl(decl)
	}
	if funcdel := pctx.FuncDef(); funcdel != nil {
		return c.FuncDef(funcdel)
	}
	return nil, nil
}
