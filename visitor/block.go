package visitor

import (
	"sysy/ast"
	"sysy/parser"
)

func (c *Context) Block(pctx parser.IBlockContext) (*ast.Block, error) {
	block := &ast.Block{}

	for _, item := range pctx.AllBlockItem() {
		if decl := item.Decl(); decl != nil {
			decl, err := c.Decl(decl)
			if err != nil {
				return nil, err
			}
			block.Children = append(block.Children, decl)
			continue
		}

		if stmt := item.Stmt(); stmt != nil {
			stmt, err := c.Stmt(stmt)
			if err != nil {
				return nil, err
			}
			block.Children = append(block.Children, stmt)
			continue
		}

		return nil, Invalid(pctx.GetStart(), "Invalid BlockItem")
	}

	return block, nil
}
