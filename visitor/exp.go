package visitor

import (
	"fmt"
	"strconv"
	"strings"
	"sysy/ast"
	"sysy/ast/types"
	"sysy/parser"
)

func (c *Context) Exp(pc parser.IExpContext) (ast.Expr, error) {
	add := pc.AddExp()
	if add == nil {
		return nil, Invalid(pc.GetStart(), "Missing AddExp")
	}
	return c.AddExp(add)
}

func (c *Context) ConstExp(pc parser.IConstExpContext) (*ast.ConstExpression, error) {
	add := pc.AddExp()
	if add == nil {
		return nil, Invalid(pc.GetStart(), "Missing AddExp")
	}

	rec, err := c.AddExp(add)
	if err != nil {
		return nil, err
	}

	Value := ast.CalConst(rec)
	if Value == nil {
		return nil, Invalid(pc.GetStart(), "Not a const")
	}

	return &ast.ConstExpression{
		Expr:  rec,
		Value: Value,
	}, nil
}

func (c *Context) AddExp(pc parser.IAddExpContext) (ast.Expr, error) {
	mul := pc.MulExp()
	if mul == nil {
		return nil, Invalid(pc.GetStart(), "Missing MulExp")
	}

	mulNode, err := c.MulExp(mul)
	if err != nil {
		return nil, err
	}

	addExp := pc.AddExp()

	if addExp == nil {
		return mulNode, nil
	}

	addNode, err := c.AddExp(addExp)
	if err != nil {
		return nil, err
	}

	switch {
	case pc.Add() != nil:

		addNode, err := ast.NewAdd(mulNode, addNode)
		if err != nil {
			return nil, Invalid(pc.Add().GetSymbol(), "Invalid AddExp: "+err.Error())
		}
		return addNode, nil
	case pc.Sub() != nil:
		subNode, err := ast.NewSub(mulNode, addNode)
		if err != nil {
			return nil, Invalid(pc.Sub().GetSymbol(), "Invalid AddExp: "+err.Error())
		}
		return subNode, nil
	default:
		return nil, Invalid(pc.MulExp().GetStop(), "Missing operator")
	}
}

func (c *Context) MulExp(pc parser.IMulExpContext) (ast.Expr, error) {
	unary := pc.UnaryExp()
	if unary == nil {
		return nil, Invalid(pc.GetStart(), "Missing UnaryExp")
	}

	unaryExp, err := c.UnaryExp(unary)
	if err != nil {
		return nil, err
	}

	mulExp := pc.MulExp()

	if mulExp == nil {
		return unaryExp, nil
	}

	mulNode, err := c.MulExp(mulExp)
	if err != nil {
		return nil, err
	}

	switch {
	case pc.Mul() != nil:
		mulNode, err := ast.NewMul(unaryExp, mulNode)
		if err != nil {
			return nil, Invalid(pc.Mul().GetSymbol(), "Invalid AddExp: "+err.Error())
		}
		return mulNode, nil
	case pc.Div() != nil:
		divNode, err := ast.NewDiv(unaryExp, mulNode)
		if err != nil {
			return nil, Invalid(pc.Div().GetSymbol(), "Invalid AddExp: "+err.Error())
		}
		return divNode, nil
	case pc.Mod() != nil:
		modNode, err := ast.NewMod(unaryExp, mulNode)
		if err != nil {
			return nil, Invalid(pc.Mod().GetSymbol(), "Invalid AddExp: "+err.Error())
		}
		return modNode, nil
	default:
		return nil, Invalid(pc.UnaryExp().GetStop(), "Missing operator")
	}
}

func (c *Context) UnaryExp(pctx parser.IUnaryExpContext) (ast.Expr, error) {
	primary := pctx.PrimaryExp()
	if primary != nil {
		return c.Primary(primary)
	}

	// func call
	if funcRParams := pctx.FuncRParams(); funcRParams != nil {
		expList, err := c.FuncRParams(funcRParams)
		if err != nil {
			return nil, err
		}

		ident := pctx.Identifier().GetText()
		symbolType := c.GetSymbol(ident)
		if symbolType == nil {
			return nil, Invalid(pctx.Identifier().GetSymbol(), "Undeclared Identifier")
		}

		fType, ok := symbolType.(*types.Func)
		if !ok {
			return nil, Invalid(pctx.Identifier().GetSymbol(), "Not a function")
		}

		if len(expList) != len(fType.ParamsList) {
			return nil, Invalid(pctx.Identifier().GetSymbol(), "Wrong number of parameters")
		}

		for i, v := range fType.ParamsList {
			if !v.Assign(expList[i].GetType()) {
				return nil, Invalid(pctx.Identifier().GetSymbol(), fmt.Sprintf("%s cannot be assigned to %s", expList[i].GetType().String(), v.String()))
			}
		}

		return &ast.FuncCall{
			Identifier: ident,
			ParamsR:    expList,
			Type:       fType.Return,
		}, nil
	}

	if unaryNode := pctx.UnaryExp(); unaryNode != nil {
		unary, err := c.UnaryExp(unaryNode)
		if err != nil {
			return nil, err
		}

		op := pctx.UnaryOp()
		switch {
		case op.Add() != nil:
			return ast.NewPlus(unary)
		case op.Sub() != nil:
			return ast.NewNeg(unary)
		case op.Not() != nil:
			return ast.NewNot(unary)
		default:
			return nil, Invalid(pctx.UnaryExp().GetStop(), "Missing operator")
		}
	}

	return nil, Invalid(pctx.GetStart(), "Invalid UnaryExp")
}

func (c *Context) FuncRParams(pctx parser.IFuncRParamsContext) ([]ast.Expr, error) {
	var exprs []ast.Expr
	for _, v := range pctx.AllExp() {
		exp, err := c.Exp(v)
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, exp)
	}
	return exprs, nil
}

func (c *Context) Primary(pctx parser.IPrimaryExpContext) (ast.Expr, error) {
	if exp := pctx.Exp(); exp != nil {
		return c.Exp(exp)
	}

	if lval := pctx.LVal(); lval != nil {
		return c.Lval(lval)
	}

	if number := pctx.Number(); number != nil {
		return c.Number(number)
	}

	return nil, Invalid(pctx.GetStart(), "Invalid PrimaryExp")
}

func (c *Context) Lval(pctx parser.ILValContext) (ast.Expr, error) {
	id := pctx.Identifier().GetText()
	symbolType := c.GetSymbol(id)
	if symbolType == nil {
		return nil, Invalid(pctx.Identifier().GetSymbol(), "Undeclared Identifier")
	}

	var node ast.Expr = &ast.Symbol{
		Type:       symbolType,
		Identifier: id,
	}

	exps := pctx.AllExp()
	if len(exps) == 0 {
		// todo: check lval
		return node, nil
	}

	for _, exp := range exps {
		expNode, err := c.Exp(exp)
		if err != nil {
			return nil, Invalid(exp.GetStart(), "Cannot be member")
		}
		node, err = ast.NewMember(node, expNode)
		if err != nil {
			return nil, Invalid(exp.GetStart(), "Cannot be member "+err.Error())
		}
	}

	return node, nil
}

func (c *Context) Number(pctx parser.INumberContext) (ast.Expr, error) {
	strval := pctx.GetText()
	if strings.Contains(strval, ".") {
		f64, err := strconv.ParseFloat(strval, 32)
		if err != nil {
			return nil, Invalid(pctx.GetStart(), "Invalid FloatConst: "+err.Error())
		}
		return &ast.FloatConst{
			Value: float32(f64),
		}, nil
	}

	if (strings.HasPrefix(strval, "0x") || strings.HasPrefix(strval, "0X")) && len(strval) > 2 {
		i64, err := strconv.ParseInt(strval[2:], 16, 32)
		if err != nil {
			return nil, Invalid(pctx.GetStart(), "Invalid IntegerConst: "+err.Error())
		}
		return &ast.IntegerConst{
			Value: int(i64),
		}, nil
	}

	if strings.HasPrefix(strval, "0") && len(strval) > 1 {
		i64, err := strconv.ParseInt(strval[1:], 8, 32)
		if err != nil {
			return nil, Invalid(pctx.GetStart(), "Invalid IntegerConst: "+err.Error())
		}
		return &ast.IntegerConst{
			Value: int(i64),
		}, nil
	}

	i64, err := strconv.ParseInt(strval, 10, 32)
	if err != nil {
		return nil, Invalid(pctx.GetStart(), "Invalid IntegerConst: "+err.Error())
	}
	return &ast.IntegerConst{
		Value: int(i64),
	}, nil
}
