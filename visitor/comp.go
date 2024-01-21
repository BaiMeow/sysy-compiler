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
