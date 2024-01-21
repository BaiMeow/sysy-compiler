// Code generated from sysyParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // sysyParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by sysyParser.
type sysyParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by sysyParser#comp.
	VisitComp(ctx *CompContext) interface{}

	// Visit a parse tree produced by sysyParser#compUnit.
	VisitCompUnit(ctx *CompUnitContext) interface{}

	// Visit a parse tree produced by sysyParser#decl.
	VisitDecl(ctx *DeclContext) interface{}

	// Visit a parse tree produced by sysyParser#constDecl.
	VisitConstDecl(ctx *ConstDeclContext) interface{}

	// Visit a parse tree produced by sysyParser#constDef.
	VisitConstDef(ctx *ConstDefContext) interface{}

	// Visit a parse tree produced by sysyParser#constinitVal.
	VisitConstinitVal(ctx *ConstinitValContext) interface{}

	// Visit a parse tree produced by sysyParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by sysyParser#varDef.
	VisitVarDef(ctx *VarDefContext) interface{}

	// Visit a parse tree produced by sysyParser#initVal.
	VisitInitVal(ctx *InitValContext) interface{}

	// Visit a parse tree produced by sysyParser#funcDef.
	VisitFuncDef(ctx *FuncDefContext) interface{}

	// Visit a parse tree produced by sysyParser#funcType.
	VisitFuncType(ctx *FuncTypeContext) interface{}

	// Visit a parse tree produced by sysyParser#funcFParams.
	VisitFuncFParams(ctx *FuncFParamsContext) interface{}

	// Visit a parse tree produced by sysyParser#funcFParam.
	VisitFuncFParam(ctx *FuncFParamContext) interface{}

	// Visit a parse tree produced by sysyParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by sysyParser#blockItem.
	VisitBlockItem(ctx *BlockItemContext) interface{}

	// Visit a parse tree produced by sysyParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by sysyParser#exp.
	VisitExp(ctx *ExpContext) interface{}

	// Visit a parse tree produced by sysyParser#cond.
	VisitCond(ctx *CondContext) interface{}

	// Visit a parse tree produced by sysyParser#lVal.
	VisitLVal(ctx *LValContext) interface{}

	// Visit a parse tree produced by sysyParser#primaryExp.
	VisitPrimaryExp(ctx *PrimaryExpContext) interface{}

	// Visit a parse tree produced by sysyParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by sysyParser#unaryExp.
	VisitUnaryExp(ctx *UnaryExpContext) interface{}

	// Visit a parse tree produced by sysyParser#unaryOp.
	VisitUnaryOp(ctx *UnaryOpContext) interface{}

	// Visit a parse tree produced by sysyParser#funcRParams.
	VisitFuncRParams(ctx *FuncRParamsContext) interface{}

	// Visit a parse tree produced by sysyParser#mulExp.
	VisitMulExp(ctx *MulExpContext) interface{}

	// Visit a parse tree produced by sysyParser#addExp.
	VisitAddExp(ctx *AddExpContext) interface{}

	// Visit a parse tree produced by sysyParser#relExp.
	VisitRelExp(ctx *RelExpContext) interface{}

	// Visit a parse tree produced by sysyParser#eqExp.
	VisitEqExp(ctx *EqExpContext) interface{}

	// Visit a parse tree produced by sysyParser#lAndExp.
	VisitLAndExp(ctx *LAndExpContext) interface{}

	// Visit a parse tree produced by sysyParser#lOrExp.
	VisitLOrExp(ctx *LOrExpContext) interface{}

	// Visit a parse tree produced by sysyParser#constExp.
	VisitConstExp(ctx *ConstExpContext) interface{}
}
