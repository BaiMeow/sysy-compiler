package visitor

import (
	"sysy/parser"
)

func (c *Context) CompUnit(pctx parser.ICompUnitContext) (any, error) {
	if decl := pctx.Decl(); decl != nil {
		return c.Decl(decl)
	}
	if funcdel := pctx.FuncDef(); funcdel != nil {
		return c.FuncDef(funcdel)
	}
	return nil, nil
}
