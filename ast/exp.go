package ast

import (
	"fmt"
	"log"
	"sysy/ast/types"
)

type Expression interface {
	GetType() types.Type
}

type ConstExpression struct {
	Expression
	Value any
}

type Add struct {
	Type  types.Type
	Left  Expression
	Right Expression
}

type Sub struct {
	Type        types.Type
	Left, Right Expression
}

type Mul struct {
	Type        types.Type
	Left, Right Expression
}

type Div struct {
	Type        types.Type
	Left, Right Expression
}

type Mod struct {
	Type        types.Type
	Left, Right Expression
}

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
	LVal  Expression
	Index Expression
}

type FuncCall struct {
	Identifier string
	ParamsR    []Expression
	types.Type
}

type ArrayExp struct {
	Type   types.Type
	Member []Expression
}

func (t *Add) GetType() types.Type {
	return t.Type
}

func NewAdd(left, right Expression) (*Add, error) {
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

func NewSub(left, right Expression) (*Sub, error) {
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

func NewMul(left, right Expression) (*Mul, error) {
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

func NewDiv(left, right Expression) (*Div, error) {
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

func NewMod(left, right Expression) (*Mod, error) {
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

func NewMember(lval Expression, index Expression) (*Member, error) {
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

func commonCalNewType(left, right Expression) (types.Type, error) {
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

func CalConst(exp Expression) any {
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
		log.Println(fmt.Sprintf("Unknown Expression Type: %T", exp))
		return false
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
