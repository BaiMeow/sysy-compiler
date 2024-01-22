package visitor

import (
	"sysy/ast"
	"sysy/ast/types"
	"sysy/parser"
	"sysy/utils"
)

func (c *Context) FuncDef(pctx parser.IFuncDefContext) (*ast.FuncDef, error) {
	funcTypeNode := pctx.FuncType()
	if funcTypeNode == nil {
		return nil, Invalid(pctx.GetStart(), "Missing Func")
	}

	recType, err := c.FuncType(funcTypeNode)
	if err != nil {
		return nil, err
	}

	var funcParams []*ast.FuncParam
	funcParamsNode := pctx.FuncFParams()
	if funcParamsNode != nil {
		funcParams, err = c.FuncFParams(funcParamsNode)
		if err != nil {
			return nil, err
		}
	}

	fType := &types.Func{
		ParamsList: utils.Map(funcParams, func(param *ast.FuncParam) types.Type {
			return param.Type
		}),
		Return: recType,
	}

	id := pctx.Identifier().GetText()

	if !c.DeclSymbol(id, fType) {
		return nil, Invalid(pctx.Identifier().GetSymbol(), "Duplicate Identifier")
	}

	c.PushSymbolTable()
	for _, param := range funcParams {
		if !c.DeclSymbol(param.Identifier, param.Type) {
			return nil, Invalid(pctx.Identifier().GetSymbol(), "Duplicate Identifier")
		}
	}

	fblock, err := c.Block(pctx.Block())
	if err != nil {
		return nil, err
	}

	return &ast.FuncDef{
		Identifier: id,
		Params:     funcParams,
		Return:     recType,
		Body:       fblock,
	}, nil
}

func (c *Context) VarDef(pctx parser.IVarDefContext, BaseType types.Type) (*ast.Define, error) {
	id := pctx.Identifier().GetText()

	dims := pctx.AllConstExp()

	var Type types.Type
	allLen := 1

	mustInt := func(exp parser.IConstExpContext) (int, error) {
		node, err := c.ConstExp(exp)
		if err != nil {
			return 0, err
		}

		l, ok := node.Value.(int)
		if !ok {
			return 0, Invalid(pctx.GetStart(), "Invalid Array Length")
		}
		return l, nil
	}

	if len(dims) == 0 {
		Type = BaseType
	} else if len(dims) == 1 {
		l, err := mustInt(dims[0])
		if err != nil {
			return nil, err
		}
		allLen *= l
		Type = &types.Array{
			ElementType: BaseType,
			Shape:       nil,
		}
	} else {
		shape := make([]int, len(dims)-1)
		for i := 0; i < len(dims)-1; i++ {
			l, err := mustInt(dims[i])
			if err != nil {
				return nil, err
			}
			shape[i] = l
			allLen *= l
		}

		shape = shape[1:]

		Type = &types.Array{
			ElementType: BaseType,
			Shape:       shape,
		}
	}

	if !c.DeclSymbol(id, Type) {
		return nil, Invalid(pctx.Identifier().GetSymbol(), "Duplicate Identifier")
	}

	initExps, err := c.InitVal(pctx.InitVal())
	if err != nil {
		return nil, err
	}

	if allLen < len(initExps) {
		return nil, Invalid(pctx.InitVal().GetStart(), "too many init val")
	}

	if len(initExps) == 1 {
		if !initExps[0].GetType().Equal(Type) {
			return nil, Invalid(pctx.InitVal().GetStart(), "Invalid InitVal Type")
		}
		return &ast.Define{
			Identifier:   id,
			Type:         Type,
			InitialValue: initExps[0],
		}, nil
	}

	initExpNode := &ast.ArrayExp{
		Type:   BaseType,
		Member: make([]ast.Expr, allLen),
	}

	for i := 0; i < allLen; i++ {
		if i < len(initExps) {
			initExpNode.Member[i] = initExps[i]
			if !initExpNode.Member[i].GetType().Assign(BaseType) {
				return nil, Invalid(pctx.InitVal().GetStart(), "Invalid InitVal Type")
			}
		} else {
			initExpNode.Member[i], err = ast.BaseTypeDefault(BaseType)
			if err != nil {
				return nil, Invalid(pctx.InitVal().GetStart(), err.Error())
			}
		}
	}

	return &ast.Define{
		Identifier:   id,
		Type:         BaseType,
		InitialValue: initExpNode,
	}, nil
}

func (c *Context) ConstDef(pctx parser.IConstDefContext, BaseType types.Type) (def *ast.Define, err error) {
	defer func() {
		if err != nil {
			return
		}
		if !c.DeclSymbol(def.Identifier, def.Type) {
			def = nil
			err = Invalid(pctx.Identifier().GetSymbol(), "Duplicate Identifier")
		}
	}()

	ident := pctx.Identifier().GetText()
	vals, err := c.ConstInitVal(pctx.ConstinitVal())
	if err != nil {
		return nil, err
	}

	dims := pctx.AllConstExp()
	if len(dims) == 0 {
		if len(vals) == 0 {
			def, err := ast.BaseTypeDefault(BaseType)
			if err != nil {
				return nil, Invalid(pctx.GetStart(), err.Error())
			}
			return &ast.Define{
				Type:         BaseType,
				Identifier:   ident,
				InitialValue: def,
			}, nil
		} else if len(vals) == 1 {
			return &ast.Define{
				Type:         BaseType,
				Identifier:   ident,
				InitialValue: vals[0],
			}, nil
		} else {
			return nil, Invalid(pctx.ConstinitVal().GetStart(), "init val not match definition")
		}
	}

	Type := &types.Array{
		ElementType: BaseType,
		Shape:       nil,
	}

	for _, constExp := range pctx.AllConstExp() {
		exp, err := c.ConstExp(constExp)
		if err != nil {
			return nil, err
		}

		length, ok := exp.Value.(int)
		if !ok {
			return nil, Invalid(constExp.GetStart(), "Invalid array index type")
		}

		Type.Shape = append(Type.Shape, length)
	}

	return &ast.Define{
		Identifier: ident,
		Type:       Type,
		InitialValue: &ast.ArrayExp{
			Type:   BaseType,
			Member: vals,
		},
	}, nil
}
