package visitor

import (
	"sysy/ast"
	"sysy/parser"
)

func (c *Context) Stmt(pctx parser.IStmtContext) (any, error) {
	if lval := pctx.LVal(); lval != nil {
		exp := pctx.Exp()
		if exp == nil {
			return nil, Invalid(pctx.GetStart(), "assignment missing exp")
		}

		lvalNode, err := c.Lval(lval)
		if err != nil {
			return nil, err
		}

		expNode, err := c.Exp(exp)
		if err != nil {
			return nil, err
		}

		return &ast.Assign{
			LVal: lvalNode,
			Exp:  expNode,
		}, nil
	}

	if hasReturn := pctx.Return(); hasReturn != nil {
		returnNode := &ast.Return{}
		exp := pctx.Exp()
		if exp == nil {
			return returnNode, nil
		}

		expNode, err := c.Exp(exp)
		if err != nil {
			return nil, err
		}
		returnNode.Value = expNode

		return returnNode, nil
	}

	if hasBreak := pctx.Break(); hasBreak != nil {
		return &ast.Break{}, nil
	}

	if hasContinue := pctx.Continue(); hasContinue != nil {
		return &ast.Continue{}, nil
	}

	if blockNode := pctx.Block(); blockNode != nil {
		return c.Block(blockNode)
	}

	if expNode := pctx.Exp(); expNode != nil {
		return c.Exp(expNode)
	}

	if hasIf := pctx.If(); hasIf != nil {
		cond, err := c.Cond(pctx.Cond())
		if err != nil {
			return nil, err
		}

		stmtNode := pctx.AllStmt()
		stmt, err := c.Stmt(stmtNode[0])
		if err != nil {
			return nil, err
		}

		if len(stmtNode) == 1 {
			return &ast.If{
				Condition: cond,
				Then:      stmt,
			}, nil
		}

		stmt2, err := c.Stmt(stmtNode[1])
		if err != nil {
			return nil, err
		}
		return &ast.IfElse{
			Condition: cond,
			Then:      stmt,
			Else:      stmt2,
		}, nil
	}

	if hasWhile := pctx.While(); hasWhile != nil {
		cond, err := c.Cond(pctx.Cond())
		if err != nil {
			return nil, err
		}

		stmt, err := c.Stmt(pctx.Stmt(0))
		if err != nil {
			return nil, err
		}

		return &ast.While{
			Condition: cond,
			Body:      stmt,
		}, nil
	}

	return nil, Invalid(pctx.GetStart(), "Invalid Stmt")
}
