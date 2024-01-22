package ast

import (
	"fmt"
	"sysy/ast/types"
)

type Expr interface {
	GetType() types.Type
}

type ConstExpression struct {
	Expr
	Value any
}

type DoubleElement struct {
	Type  types.Type
	Left  Expr
	Right Expr
}

type SingleElement struct {
	Type types.Type
	Exp  Expr
}

type Plus SingleElement

type Neg SingleElement

type Not SingleElement

type Add DoubleElement

type Sub DoubleElement

type Mul DoubleElement

type Div DoubleElement

type Mod DoubleElement

type Symbol struct {
	Type       types.Type
	Identifier string
}

type IntegerConst struct {
	Value int
}

type FloatConst struct {
	Value float32
}

type Member struct {
	Type  types.Type
	LVal  Expr
	Index Expr
}

type FuncCall struct {
	Type       types.Type
	Identifier string
	ParamsR    []Expr
}

type ArrayExp struct {
	Type   types.Type
	Member []Expr
}

type LOrExp DoubleElement

type LAndExp DoubleElement

type EqExp DoubleElement

type NeqExp DoubleElement

type LessExp DoubleElement

type GreaterExp DoubleElement

type LessEqExp DoubleElement

type GreaterEqExp DoubleElement

func (t *Plus) GetType() types.Type {
	return t.Type
}

func NewPlus(exp Expr) (*Plus, error) {
	check, ok := exp.GetType().(*types.Base)
	if !ok || check.Type != types.Int && check.Type != types.Float {
		return nil, fmt.Errorf("plus can only be int or float")
	}

	return &Plus{
		Type: check,
		Exp:  exp,
	}, nil
}

func (n *Neg) GetType() types.Type {
	return n.Type
}

func NewNeg(exp Expr) (*Neg, error) {
	check, ok := exp.GetType().(*types.Base)
	if !ok || check.Type != types.Int && check.Type != types.Float {
		return nil, fmt.Errorf("neg can only be int or float")
	}

	return &Neg{
		Type: check,
		Exp:  exp,
	}, nil
}

func (n *Not) GetType() types.Type {
	return &types.Base{Type: types.Int}
}

func NewNot(exp Expr) (*Not, error) {
	check, ok := exp.GetType().(*types.Base)
	if !ok || check.Type != types.Int {
		return nil, fmt.Errorf("not can only be int")
	}

	return &Not{
		Type: check,
		Exp:  exp,
	}, nil
}

func (t *Add) GetType() types.Type {
	return t.Type
}

func NewAdd(left, right Expr) (*Add, error) {
	newType, err := commonCalNewType(left, right)
	if err != nil {
		return nil, err
	}

	return &Add{
		Left:  left,
		Right: right,
		Type:  newType,
	}, nil
}

func (t *Sub) GetType() types.Type {
	return t.Type
}

func NewSub(left, right Expr) (*Sub, error) {
	newType, err := commonCalNewType(left, right)
	if err != nil {
		return nil, err
	}

	return &Sub{
		Left:  left,
		Right: right,
		Type:  newType,
	}, nil
}

func (t *Mul) GetType() types.Type {
	return t.Type
}

func NewMul(left, right Expr) (*Mul, error) {
	newType, err := commonCalNewType(left, right)
	if err != nil {
		return nil, err
	}

	return &Mul{
		Left:  left,
		Right: right,
		Type:  newType,
	}, nil
}

func (t *Div) GetType() types.Type {
	return t.Type
}

func NewDiv(left, right Expr) (*Div, error) {
	newType, err := commonCalNewType(left, right)
	if err != nil {
		return nil, err
	}

	return &Div{
		Left:  left,
		Right: right,
		Type:  newType,
	}, nil

}

func (t *Mod) GetType() types.Type {
	return t.Type
}

func NewMod(left, right Expr) (*Mod, error) {
	if left.GetType().Equal(right.GetType()) && left.GetType().Equal(&types.Base{Type: types.Int}) {
		return &Mod{
			Left:  left,
			Right: right,
			Type:  &types.Base{Type: types.Int},
		}, nil
	}

	return nil, fmt.Errorf("mod can only be int")
}

func (t *Symbol) GetType() types.Type {
	return t.Type
}

func (t *IntegerConst) GetType() types.Type {
	return &types.Base{
		Type: types.Int,
	}
}

func (t *FloatConst) GetType() types.Type {
	return &types.Base{
		Type: types.Float,
	}
}

func (t *Member) GetType() types.Type {
	return t.Type
}

func NewMember(lval Expr, index Expr) (*Member, error) {
	checkArray, ok := lval.GetType().(*types.Array)
	if !ok {
		return nil, fmt.Errorf("member can only modify array")
	}

	if !index.GetType().Equal(&types.Base{Type: types.Int}) {
		return nil, fmt.Errorf("index can only be int")
	}

	return &Member{
		Type:  checkArray.ElementType,
		LVal:  lval,
		Index: index,
	}, nil
}

func (t *FuncCall) GetType() types.Type {
	return t.Type
}

func (t *ArrayExp) GetType() types.Type {
	return t.Type
}

func (t *LOrExp) GetType() types.Type {
	return &types.Base{Type: types.Int}
}

func (t *LAndExp) GetType() types.Type {
	return &types.Base{Type: types.Int}
}

func (t *EqExp) GetType() types.Type {
	return &types.Base{Type: types.Int}
}

func (t *NeqExp) GetType() types.Type {
	return &types.Base{Type: types.Int}
}

func (t *LessExp) GetType() types.Type {
	return &types.Base{Type: types.Int}
}

func (t *GreaterExp) GetType() types.Type {
	return &types.Base{Type: types.Int}
}

func (t *LessEqExp) GetType() types.Type {
	return &types.Base{Type: types.Int}
}

func (t *GreaterEqExp) GetType() types.Type {
	return &types.Base{Type: types.Int}
}

func commonCalNewType(left, right Expr) (types.Type, error) {
	lcheck, ok := left.GetType().(*types.Base)
	if !ok {
		return nil, fmt.Errorf("left type can only be int or float")
	}

	rcheck, ok := right.GetType().(*types.Base)
	if !ok {
		return nil, fmt.Errorf("right type can only be int or float")
	}

	if lcheck.Type == types.Int && rcheck.Type == types.Int {
		return &types.Base{Type: types.Int}, nil
	} else {
		return &types.Base{Type: types.Float}, nil
	}
}

func CalConst(exp Expr) any {
	switch exp := exp.(type) {
	case *IntegerConst:
		return exp.Value
	case *FloatConst:
		return exp.Value
	case *Symbol, *FuncCall, *ArrayExp:
		return nil
	case *Add:
		return CalculateWithType(CalConst(exp.Left), CalConst(exp.Right), actionAdd[float32], actionAdd[int])
	case *Sub:
		return CalculateWithType(CalConst(exp.Left), CalConst(exp.Right), actionSub[float32], actionSub[int])
	case *Mul:
		return CalculateWithType(CalConst(exp.Left), CalConst(exp.Right), actionMul[float32], actionMul[int])
	case *Div:
		return CalculateWithType(CalConst(exp.Left), CalConst(exp.Right), actionDiv[float32], actionDiv[int])
	case *Mod:
		return CalculateWithType(CalConst(exp.Left), CalConst(exp.Right), actionNil, actionMod)
	default:
		return nil
	}
}

func CalculateWithType(l, r any, actionFloat func(float32, float32) any, actionInt func(int, int) any) any {
	switch l := l.(type) {
	case int:
		switch r := r.(type) {
		case int:
			return actionInt(l, r)
		case float32:
			return actionFloat(float32(l), r)
		default:
			return nil
		}
	case float32:
		switch r := r.(type) {
		case int:
			return actionFloat(l, float32(r))
		case float32:
			return actionFloat(l, r)
		default:
			return nil
		}
	default:
		return nil
	}
}

func actionAdd[T int | float32](a, b T) any {
	return a + b
}

func actionSub[T int | float32](a, b T) any {
	return a - b
}

func actionMul[T int | float32](a, b T) any {
	return a * b
}

func actionDiv[T int | float32](a, b T) any {
	if b == 0 {
		return nil
	}
	return a / b
}

func actionMod(a, b int) any {
	return a % b
}

func actionNil(a, b float32) any {
	return nil
}
