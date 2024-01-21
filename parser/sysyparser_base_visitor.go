// Code generated from sysyParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // sysyParser
import "github.com/antlr4-go/antlr/v4"

type BasesysyParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasesysyParserVisitor) VisitComp(ctx *CompContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitCompUnit(ctx *CompUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitDecl(ctx *DeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitConstDecl(ctx *ConstDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitConstDef(ctx *ConstDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitConstinitVal(ctx *ConstinitValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitVarDef(ctx *VarDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitInitVal(ctx *InitValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitFuncDef(ctx *FuncDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitFuncType(ctx *FuncTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitFuncFParams(ctx *FuncFParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitFuncFParam(ctx *FuncFParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitBlockItem(ctx *BlockItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitExp(ctx *ExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitCond(ctx *CondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitLVal(ctx *LValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitPrimaryExp(ctx *PrimaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitUnaryExp(ctx *UnaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitUnaryOp(ctx *UnaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitFuncRParams(ctx *FuncRParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitMulExp(ctx *MulExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitAddExp(ctx *AddExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitRelExp(ctx *RelExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitEqExp(ctx *EqExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitLAndExp(ctx *LAndExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitLOrExp(ctx *LOrExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesysyParserVisitor) VisitConstExp(ctx *ConstExpContext) interface{} {
	return v.VisitChildren(ctx)
}
