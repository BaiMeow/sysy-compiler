package visitor

import (
	"sysy/ast"
	"sysy/parser"
)

func (c *Context) Cond(pctx parser.ICondContext) (ast.Expr, error) {
	lOrExpNode := pctx.LOrExp()
	if lOrExpNode == nil {
		return nil, Invalid(pctx.GetStart(), "invalid cond")
	}
	return c.lOrExp(lOrExpNode)
}

func (c *Context) lOrExp(pctx parser.ILOrExpContext) (ast.Expr, error) {
	lAndExpNode := pctx.LAndExp()
	if lAndExpNode == nil {
		return nil, Invalid(pctx.GetStart(), "invalid lOrExp")
	}

	lExpr, err := c.lAndExp(lAndExpNode)
	if err != nil {
		return nil, err
	}

	lOrExpNode := pctx.LOrExp()
	if lOrExpNode == nil {
		return lExpr, nil
	}

	rExpr, err := c.lOrExp(lOrExpNode)
	if err != nil {
		return nil, err
	}

	return &ast.LOrExp{
		Left:  lExpr,
		Right: rExpr,
	}, nil
}

func (c *Context) lAndExp(pctx parser.ILAndExpContext) (ast.Expr, error) {
	eqExpNode := pctx.EqExp()
	if eqExpNode == nil {
		return nil, Invalid(pctx.GetStart(), "Invalid LAndExp")
	}
	lExpr, err := c.EqExp(eqExpNode)
	if err != nil {
		return nil, err
	}

	lAndNode := pctx.LAndExp()
	if lAndNode == nil {
		return lExpr, nil
	}
	rExpr, err := c.lAndExp(lAndNode)
	if err != nil {
		return nil, err
	}

	return &ast.LAndExp{
		Left:  lExpr,
		Right: rExpr,
	}, nil
}

func (c *Context) EqExp(pctx parser.IEqExpContext) (ast.Expr, error) {
	relExpNode := pctx.RelExp()
	if relExpNode == nil {
		return nil, Invalid(pctx.GetStart(), "Invalid EqExp")
	}
	relExpr, err := c.RelExp(relExpNode)
	if err != nil {
		return nil, err
	}

	eqExpNode := pctx.EqExp()
	if eqExpNode == nil {
		return relExpr, nil
	}

	eqExpr, err := c.EqExp(eqExpNode)
	if err != nil {
		return nil, err
	}

	return &ast.EqExp{
		Left:  relExpr,
		Right: eqExpr,
	}, nil
}

func (c *Context) RelExp(pctx parser.IRelExpContext) (ast.Expr, error) {
	addExpNode := pctx.AddExp()
	if addExpNode == nil {
		return nil, Invalid(pctx.GetStart(), "Invalid RelExp")
	}
	addExpr, err := c.AddExp(addExpNode)
	if err != nil {
		return nil, err
	}

	relExpNode := pctx.RelExp()
	if relExpNode == nil {
		return addExpr, nil
	}

	relExpr, err := c.RelExp(relExpNode)
	if err != nil {
		return nil, err
	}

	switch {
	case pctx.Less() != nil:
		return &ast.LessExp{
			Left:  addExpr,
			Right: relExpr,
		}, nil
	case pctx.Greater() != nil:
		return &ast.GreaterExp{
			Left:  addExpr,
			Right: relExpr,
		}, nil
	case pctx.LessEqual() != nil:
		return &ast.LessEqExp{
			Left:  addExpr,
			Right: relExpr,
		}, nil
	case pctx.GreaterEqual() != nil:
		return &ast.GreaterEqExp{
			Left:  addExpr,
			Right: relExpr,
		}, nil
	default:
		return nil, Invalid(pctx.GetStart(), "Invalid RelExp")
	}
}
