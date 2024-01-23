package visitor

import (
	"sysy/ast"
	"sysy/parser"
)

func (c *Context) Stmt(pctx parser.IStmtContext) (any, error) {
	if lValNode := pctx.LVal(); lValNode != nil {
		exp := pctx.Exp()
		if exp == nil {
			return nil, Invalid(pctx.GetStart(), "assignment missing exp")
		}

		lValExpr, err := c.LVal(lValNode)
		if err != nil {
			return nil, err
		}

		expr, err := c.Exp(exp)
		if err != nil {
			return nil, err
		}

		return &ast.Assign{
			LVal: lValExpr,
			Exp:  expr,
		}, nil
	}

	if hasReturn := pctx.Return(); hasReturn != nil {
		returnStmt := &ast.Return{}
		expNode := pctx.Exp()
		if expNode == nil {
			return returnStmt, nil
		}

		expr, err := c.Exp(expNode)
		if err != nil {
			return nil, err
		}
		returnStmt.Value = expr

		return returnStmt, nil
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

		stmtNodes := pctx.AllStmt()
		stmt, err := c.Stmt(stmtNodes[0])
		if err != nil {
			return nil, err
		}

		if len(stmtNodes) == 1 {
			return &ast.If{
				Condition: cond,
				Then:      stmt,
			}, nil
		}

		stmt2, err := c.Stmt(stmtNodes[1])
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

	return nil, Invalid(pctx.GetStart(), "invalid stmt")
}
