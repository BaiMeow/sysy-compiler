package visitor

import (
	"sysy/ast"
	"sysy/parser"
)

func (c *Context) Cond(pctx parser.ICondContext) (ast.Expr, error) {
	lOrExp := pctx.LOrExp()
	if lOrExp == nil {
		return nil, Invalid(pctx.GetStart(), "Invalid Cond")
	}
	return c.lOrExp(lOrExp)
}

func (c *Context) lOrExp(pctx parser.ILOrExpContext) (ast.Expr, error) {
	lAndExp := pctx.LAndExp()
	if lAndExp == nil {
		return nil, Invalid(pctx.GetStart(), "Invalid LOrExp")
	}

	lexpr, err := c.lAndExp(lAndExp)
	if err != nil {
		return nil, err
	}

	lorexp := pctx.LOrExp()
	if lorexp == nil {
		return lexpr, nil
	}

	expr, err := c.lOrExp(lorexp)
	if err != nil {
		return nil, err
	}

	return &ast.LOrExp{
		Left:  lexpr,
		Right: expr,
	}, nil
}

func (c *Context) lAndExp(pctx parser.ILAndExpContext) (ast.Expr, error) {
	eqExp := pctx.EqExp()
	if eqExp == nil {
		return nil, Invalid(pctx.GetStart(), "Invalid LAndExp")
	}
	left, err := c.EqExp(eqExp)
	if err != nil {
		return nil, err
	}

	landexp := pctx.LAndExp()
	if landexp == nil {
		return left, nil
	}
	right, err := c.lAndExp(landexp)
	if err != nil {
		return nil, err
	}

	return &ast.LAndExp{
		Left:  left,
		Right: right,
	}, nil
}

func (c *Context) EqExp(pctx parser.IEqExpContext) (ast.Expr, error) {
	relExp := pctx.RelExp()
	if relExp == nil {
		return nil, Invalid(pctx.GetStart(), "Invalid EqExp")
	}
	relNode, err := c.RelExp(relExp)
	if err != nil {
		return nil, err
	}

	eqexp := pctx.EqExp()
	if eqexp == nil {
		return relNode, nil
	}

	expr, err := c.EqExp(eqexp)
	if err != nil {
		return nil, err
	}

	return &ast.EqExp{
		Left:  relNode,
		Right: expr,
	}, nil
}

func (c *Context) RelExp(pctx parser.IRelExpContext) (ast.Expr, error) {
	addExp := pctx.AddExp()
	if addExp == nil {
		return nil, Invalid(pctx.GetStart(), "Invalid RelExp")
	}
	addNode, err := c.AddExp(addExp)
	if err != nil {
		return nil, err
	}

	relexp := pctx.RelExp()
	if relexp == nil {
		return addNode, nil
	}

	expr, err := c.RelExp(relexp)
	if err != nil {
		return nil, err
	}

	switch {
	case pctx.Less() != nil:
		return &ast.LessExp{
			Left:  addNode,
			Right: expr,
		}, nil
	case pctx.Greater() != nil:
		return &ast.GreaterExp{
			Left:  addNode,
			Right: expr,
		}, nil
	case pctx.LessEqual() != nil:
		return &ast.LessEqExp{
			Left:  addNode,
			Right: expr,
		}, nil
	case pctx.GreaterEqual() != nil:
		return &ast.GreaterEqExp{
			Left:  addNode,
			Right: expr,
		}, nil
	default:
		return nil, Invalid(pctx.GetStart(), "Invalid RelExp")
	}
}
