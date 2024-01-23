package visitor

import (
	"fmt"
	"strconv"
	"strings"
	"sysy/ast"
	"sysy/ast/types"
	"sysy/parser"
)

func (c *Context) Exp(pctx parser.IExpContext) (ast.Expr, error) {
	addExpNode := pctx.AddExp()
	if addExpNode == nil {
		return nil, Invalid(pctx.GetStart(), "missing addExp")
	}
	return c.AddExp(addExpNode)
}

func (c *Context) ConstExp(pctx parser.IConstExpContext) (*ast.ConstExpression, error) {
	addExpNode := pctx.AddExp()
	if addExpNode == nil {
		return nil, Invalid(pctx.GetStart(), "missing addExp")
	}

	addExpr, err := c.AddExp(addExpNode)
	if err != nil {
		return nil, err
	}

	value := ast.CalConst(addExpr)
	if value == nil {
		return nil, Invalid(pctx.GetStart(), "not const")
	}

	return &ast.ConstExpression{
		Expr:  addExpr,
		Value: value,
	}, nil
}

func (c *Context) AddExp(pc parser.IAddExpContext) (ast.Expr, error) {
	mulExpNode := pc.MulExp()
	if mulExpNode == nil {
		return nil, Invalid(pc.GetStart(), "missing mulExp")
	}

	lExpr, err := c.MulExp(mulExpNode)
	if err != nil {
		return nil, err
	}

	addExpNode := pc.AddExp()

	if addExpNode == nil {
		return lExpr, nil
	}

	rExpr, err := c.AddExp(addExpNode)
	if err != nil {
		return nil, err
	}

	switch {
	case pc.Add() != nil:
		addExpr, err := ast.NewAdd(lExpr, rExpr)
		if err != nil {
			return nil, Invalid(pc.Add().GetSymbol(), "invalid addExp: "+err.Error())
		}
		return addExpr, nil
	case pc.Sub() != nil:
		subExpr, err := ast.NewSub(lExpr, rExpr)
		if err != nil {
			return nil, Invalid(pc.Sub().GetSymbol(), "invalid subExp: "+err.Error())
		}
		return subExpr, nil
	default:
		return nil, Invalid(pc.MulExp().GetStop(), "missing operator")
	}
}

func (c *Context) MulExp(pctx parser.IMulExpContext) (ast.Expr, error) {
	unaryExpNode := pctx.UnaryExp()
	if unaryExpNode == nil {
		return nil, Invalid(pctx.GetStart(), "missing unaryExp")
	}

	lExpr, err := c.UnaryExp(unaryExpNode)
	if err != nil {
		return nil, err
	}

	mulExpNode := pctx.MulExp()

	if mulExpNode == nil {
		return lExpr, nil
	}

	rExpr, err := c.MulExp(mulExpNode)
	if err != nil {
		return nil, err
	}

	switch {
	case pctx.Mul() != nil:
		mulExpr, err := ast.NewMul(lExpr, rExpr)
		if err != nil {
			return nil, Invalid(pctx.Mul().GetSymbol(), "invalid addExp: "+err.Error())
		}
		return mulExpr, nil
	case pctx.Div() != nil:
		divExpr, err := ast.NewDiv(lExpr, rExpr)
		if err != nil {
			return nil, Invalid(pctx.Div().GetSymbol(), "invalid divExp: "+err.Error())
		}
		return divExpr, nil
	case pctx.Mod() != nil:
		modExpr, err := ast.NewMod(lExpr, rExpr)
		if err != nil {
			return nil, Invalid(pctx.Mod().GetSymbol(), "invalid modExp: "+err.Error())
		}
		return modExpr, nil
	default:
		return nil, Invalid(pctx.UnaryExp().GetStop(), "missing operator")
	}
}

func (c *Context) UnaryExp(pctx parser.IUnaryExpContext) (ast.Expr, error) {
	primaryExpNode := pctx.PrimaryExp()
	if primaryExpNode != nil {
		return c.Primary(primaryExpNode)
	}

	// func call
	if funcRParamsNode := pctx.FuncRParams(); funcRParamsNode != nil {
		paramList, err := c.FuncRParams(funcRParamsNode)
		if err != nil {
			return nil, err
		}

		ident := pctx.Identifier().GetText()
		symbolType := c.GetSymbol(ident)
		if symbolType == nil {
			return nil, Invalid(pctx.Identifier().GetSymbol(), "undeclared identifier")
		}

		funcType, ok := symbolType.(*types.Func)
		if !ok {
			return nil, Invalid(pctx.Identifier().GetSymbol(), "not a function")
		}

		if len(paramList) != len(funcType.ParamsList) {
			return nil, Invalid(pctx.Identifier().GetSymbol(), "wrong number of parameters")
		}

		for i, v := range funcType.ParamsList {
			if !v.Assign(paramList[i].GetType()) {
				return nil, Invalid(pctx.Identifier().GetSymbol(), fmt.Sprintf("%s cannot be assigned to %s", paramList[i].GetType().String(), v.String()))
			}
		}

		return &ast.FuncCall{
			Identifier: ident,
			ParamsR:    paramList,
			Type:       funcType.Return,
		}, nil
	}

	if unaryNode := pctx.UnaryExp(); unaryNode != nil {
		unaryExpr, err := c.UnaryExp(unaryNode)
		if err != nil {
			return nil, err
		}

		op := pctx.UnaryOp()
		switch {
		case op.Add() != nil:
			return ast.NewPlus(unaryExpr)
		case op.Sub() != nil:
			return ast.NewNeg(unaryExpr)
		case op.Not() != nil:
			return ast.NewNot(unaryExpr)
		default:
			return nil, Invalid(pctx.UnaryExp().GetStop(), "missing operator")
		}
	}

	return nil, Invalid(pctx.GetStart(), "invalid unaryExp")
}

func (c *Context) FuncRParams(pctx parser.IFuncRParamsContext) ([]ast.Expr, error) {
	var exprs []ast.Expr
	for _, expNode := range pctx.AllExp() {
		expr, err := c.Exp(expNode)
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, expr)
	}
	return exprs, nil
}

func (c *Context) Primary(pctx parser.IPrimaryExpContext) (ast.Expr, error) {
	if expNode := pctx.Exp(); expNode != nil {
		return c.Exp(expNode)
	}

	if lValNode := pctx.LVal(); lValNode != nil {
		return c.LVal(lValNode)
	}

	if numberNode := pctx.Number(); numberNode != nil {
		return c.Number(numberNode)
	}

	return nil, Invalid(pctx.GetStart(), "invalid primaryExp")
}

func (c *Context) LVal(pctx parser.ILValContext) (ast.Expr, error) {
	ident := pctx.Identifier().GetText()
	symbolType := c.GetSymbol(ident)
	if symbolType == nil {
		return nil, Invalid(pctx.Identifier().GetSymbol(), "undeclared identifier")
	}

	var lValExpr ast.Expr = &ast.Symbol{
		Type:       symbolType,
		Identifier: ident,
	}

	expNodes := pctx.AllExp()
	if len(expNodes) == 0 {
		// todo: check lval
		return lValExpr, nil
	}

	for _, expNode := range expNodes {
		expr, err := c.Exp(expNode)
		if err != nil {
			return nil, Invalid(expNode.GetStart(), "cannot be member")
		}
		lValExpr, err = ast.NewMember(lValExpr, expr)
		if err != nil {
			return nil, Invalid(expNode.GetStart(), "cannot be member "+err.Error())
		}
	}

	return lValExpr, nil
}

func (c *Context) Number(pctx parser.INumberContext) (ast.Expr, error) {
	literal := pctx.GetText()
	if strings.Contains(literal, ".") {
		f64, err := strconv.ParseFloat(literal, 32)
		if err != nil {
			return nil, Invalid(pctx.GetStart(), "invalid floatConst: "+err.Error())
		}
		return &ast.FloatConst{
			Value: float32(f64),
		}, nil
	}

	if (strings.HasPrefix(literal, "0x") || strings.HasPrefix(literal, "0X")) && len(literal) > 2 {
		i64, err := strconv.ParseInt(literal[2:], 16, 32)
		if err != nil {
			return nil, Invalid(pctx.GetStart(), "invalid integerConst: "+err.Error())
		}
		return &ast.IntegerConst{
			Value: int(i64),
		}, nil
	}

	if strings.HasPrefix(literal, "0") && len(literal) > 1 {
		i64, err := strconv.ParseInt(literal[1:], 8, 32)
		if err != nil {
			return nil, Invalid(pctx.GetStart(), "invalid integerConst: "+err.Error())
		}
		return &ast.IntegerConst{
			Value: int(i64),
		}, nil
	}

	i64, err := strconv.ParseInt(literal, 10, 32)
	if err != nil {
		return nil, Invalid(pctx.GetStart(), "invalid integerConst: "+err.Error())
	}
	return &ast.IntegerConst{
		Value: int(i64),
	}, nil
}
