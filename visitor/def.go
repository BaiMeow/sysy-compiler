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
		return nil, Invalid(pctx.GetStart(), "missing func")
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

	funcType := &types.Func{
		ParamsList: utils.Map(funcParams, func(param *ast.FuncParam) types.Type {
			return param.Type
		}),
		Return: recType,
	}

	ident := pctx.Identifier().GetText()

	if !c.DeclSymbol(ident, funcType) {
		return nil, Invalid(pctx.Identifier().GetSymbol(), "duplicate identifier")
	}

	c.PushSymbolTable()
	for _, param := range funcParams {
		if !c.DeclSymbol(param.Identifier, param.Type) {
			return nil, Invalid(pctx.Identifier().GetSymbol(), "duplicate identifier")
		}
	}

	block, err := c.Block(pctx.Block())
	if err != nil {
		return nil, err
	}

	return &ast.FuncDef{
		Identifier: ident,
		Params:     funcParams,
		Return:     recType,
		Body:       block,
	}, nil
}

func (c *Context) VarDef(pctx parser.IVarDefContext, baseType types.Type) (*ast.Define, error) {
	ident := pctx.Identifier().GetText()

	dims := pctx.AllConstExp()

	var varType types.Type
	size := 1

	mustInt := func(exp parser.IConstExpContext) (int, error) {
		node, err := c.ConstExp(exp)
		if err != nil {
			return 0, err
		}

		i, ok := node.Value.(int)
		if !ok {
			return 0, Invalid(pctx.GetStart(), "Invalid Array Length")
		}
		return i, nil
	}

	if len(dims) == 0 {
		varType = baseType
	} else if len(dims) == 1 {
		l, err := mustInt(dims[0])
		if err != nil {
			return nil, err
		}
		size *= l
		varType = &types.Array{
			ElementType: baseType,
			Shape:       nil,
		}
	} else {
		arrShape := make([]int, len(dims)-1)
		for i := 0; i < len(dims)-1; i++ {
			l, err := mustInt(dims[i])
			if err != nil {
				return nil, err
			}
			arrShape[i] = l
			size *= l
		}

		arrShape = arrShape[1:]

		varType = &types.Array{
			ElementType: baseType,
			Shape:       arrShape,
		}
	}

	if !c.DeclSymbol(ident, varType) {
		return nil, Invalid(pctx.Identifier().GetSymbol(), "duplicate identifier")
	}

	initExprs, err := c.InitVal(pctx.InitVal())
	if err != nil {
		return nil, err
	}

	if size < len(initExprs) {
		return nil, Invalid(pctx.InitVal().GetStart(), "too many init val")
	}

	if len(initExprs) == 1 {
		if !initExprs[0].GetType().Equal(varType) {
			return nil, Invalid(pctx.InitVal().GetStart(), "invalid initVal type")
		}
		return &ast.Define{
			Identifier:   ident,
			Type:         varType,
			InitialValue: initExprs[0],
		}, nil
	}

	initMemory := &ast.ArrayExp{
		Type:   baseType,
		Member: make([]ast.Expr, size),
	}

	for i := 0; i < size; i++ {
		if i < len(initExprs) {
			initMemory.Member[i] = initExprs[i]
			if !initMemory.Member[i].GetType().Assign(baseType) {
				return nil, Invalid(pctx.InitVal().GetStart(), "Invalid InitVal Type")
			}
		} else {
			initMemory.Member[i], err = ast.BaseTypeDefault(baseType)
			if err != nil {
				return nil, Invalid(pctx.InitVal().GetStart(), err.Error())
			}
		}
	}

	return &ast.Define{
		Identifier:   ident,
		Type:         baseType,
		InitialValue: initMemory,
	}, nil
}

func (c *Context) ConstDef(pctx parser.IConstDefContext, baseType types.Type) (def *ast.Define, err error) {
	ident := pctx.Identifier().GetText()

	dims := pctx.AllConstExp()

	var varType types.Type
	size := 1

	mustInt := func(exp parser.IConstExpContext) (int, error) {
		node, err := c.ConstExp(exp)
		if err != nil {
			return 0, err
		}

		i, ok := node.Value.(int)
		if !ok {
			return 0, Invalid(pctx.GetStart(), "Invalid Array Length")
		}
		return i, nil
	}

	if len(dims) == 0 {
		varType = baseType
	} else if len(dims) == 1 {
		l, err := mustInt(dims[0])
		if err != nil {
			return nil, err
		}
		size *= l
		varType = &types.Array{
			ElementType: baseType,
			Shape:       nil,
		}
	} else {
		arrShape := make([]int, len(dims)-1)
		for i := 0; i < len(dims)-1; i++ {
			l, err := mustInt(dims[i])
			if err != nil {
				return nil, err
			}
			arrShape[i] = l
			size *= l
		}

		arrShape = arrShape[1:]

		varType = &types.Array{
			ElementType: baseType,
			Shape:       arrShape,
		}
	}

	if !c.DeclSymbol(ident, varType) {
		return nil, Invalid(pctx.Identifier().GetSymbol(), "duplicate identifier")
	}

	initExprs, err := c.ConstInitVal(pctx.ConstinitVal())
	if err != nil {
		return nil, err
	}

	if size < len(initExprs) {
		return nil, Invalid(pctx.ConstinitVal().GetStart(), "too many init val")
	}

	if len(initExprs) == 1 {
		if !initExprs[0].GetType().Equal(varType) {
			return nil, Invalid(pctx.ConstinitVal().GetStart(), "invalid initVal type")
		}
		return &ast.Define{
			Identifier:   ident,
			Type:         varType,
			InitialValue: initExprs[0],
		}, nil
	}

	initMemory := &ast.ArrayExp{
		Type:   baseType,
		Member: make([]ast.Expr, size),
	}

	for i := 0; i < size; i++ {
		if i < len(initExprs) {
			initMemory.Member[i] = initExprs[i]
			if !initMemory.Member[i].GetType().Assign(baseType) {
				return nil, Invalid(pctx.ConstinitVal().GetStart(), "Invalid InitVal Type")
			}
		} else {
			initMemory.Member[i], err = ast.BaseTypeDefault(baseType)
			if err != nil {
				return nil, Invalid(pctx.ConstinitVal().GetStart(), err.Error())
			}
		}
	}

	return &ast.Define{
		Identifier:   ident,
		Type:         baseType,
		InitialValue: initMemory,
	}, nil
}
