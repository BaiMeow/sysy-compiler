package visitor

import (
	"sysy/ast"
	"sysy/parser"
)

func (c *Context) InitVal(pctx parser.IInitValContext) ([]ast.Expr, error) {
	var exps []ast.Expr
	allExpNode := pctx.AllExp()
	for _, exp := range allExpNode {
		expNode, err := c.Exp(exp)
		if err != nil {
			return nil, err
		}
		exps = append(exps, expNode)
	}
	return exps, nil
}

func (c *Context) ConstInitVal(pctx parser.IConstinitValContext) ([]ast.Expr, error) {
	var exps []ast.Expr
	allExpNode := pctx.AllConstExp()
	for _, exp := range allExpNode {
		expNode, err := c.ConstExp(exp)
		if err != nil {
			return nil, err
		}
		exps = append(exps, expNode)
	}
	return exps, nil
}
