// Code generated from sysyParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // sysyParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type sysyParser struct {
	*antlr.BaseParser
}

var SysyParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sysyparserParserInit() {
	staticData := &SysyParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'const'", "'void'", "", "'if'", "'else'", "'while'", "'break'",
		"'continue'", "'return'", "'('", "')'", "'['", "']'", "'{'", "'}'",
		"','", "';'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'<'", "'>'",
		"'=='", "'!='", "'<='", "'>='", "'&&'", "'||'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "Const", "Void", "Type", "If", "Else", "While", "Break", "Continue",
		"Return", "LeftParen", "RightParen", "LeftBracket", "RightBracket",
		"LeftBrace", "RightBrace", "Comma", "Semicolon", "Assign", "Add", "Sub",
		"Mul", "Div", "Mod", "Less", "Greater", "Equal", "NotEqual", "LessEqual",
		"GreaterEqual", "And", "Or", "Not", "Identifier", "IntegerConstant",
		"FloatConstant", "WS", "LineComment", "BlockComment", "Preprocessor",
	}
	staticData.RuleNames = []string{
		"comp", "compUnit", "decl", "constDecl", "constDef", "constinitVal",
		"varDecl", "varDef", "initVal", "funcDef", "funcType", "funcFParams",
		"funcFParam", "block", "blockItem", "stmt", "exp", "cond", "lVal", "primaryExp",
		"number", "unaryExp", "unaryOp", "funcRParams", "mulExp", "addExp",
		"relExp", "eqExp", "lAndExp", "lOrExp", "constExp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 319, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 5,
		0, 64, 8, 0, 10, 0, 12, 0, 67, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 73,
		8, 1, 1, 2, 1, 2, 3, 2, 77, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 84,
		8, 3, 10, 3, 12, 3, 87, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		5, 4, 96, 8, 4, 10, 4, 12, 4, 99, 9, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 5, 5, 109, 8, 5, 10, 5, 12, 5, 112, 9, 5, 3, 5, 114, 8,
		5, 1, 5, 3, 5, 117, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 123, 8, 6, 10,
		6, 12, 6, 126, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 135,
		8, 7, 10, 7, 12, 7, 138, 9, 7, 1, 7, 1, 7, 3, 7, 142, 8, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 5, 8, 149, 8, 8, 10, 8, 12, 8, 152, 9, 8, 3, 8, 154,
		8, 8, 1, 8, 3, 8, 157, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 163, 8, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 173, 8, 11, 10,
		11, 12, 11, 176, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 5, 12, 186, 8, 12, 10, 12, 12, 12, 189, 9, 12, 3, 12, 191, 8, 12,
		1, 13, 1, 13, 5, 13, 195, 8, 13, 10, 13, 12, 13, 198, 9, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 3, 14, 204, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 3, 15, 212, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 3, 15, 223, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 237, 8, 15, 1,
		15, 3, 15, 240, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 5, 18, 251, 8, 18, 10, 18, 12, 18, 254, 9, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 262, 8, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 275, 8,
		21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 5, 23, 282, 8, 23, 10, 23, 12, 23,
		285, 9, 23, 1, 24, 1, 24, 1, 24, 3, 24, 290, 8, 24, 1, 25, 1, 25, 1, 25,
		3, 25, 295, 8, 25, 1, 26, 1, 26, 1, 26, 3, 26, 300, 8, 26, 1, 27, 1, 27,
		1, 27, 3, 27, 305, 8, 27, 1, 28, 1, 28, 1, 28, 3, 28, 310, 8, 28, 1, 29,
		1, 29, 1, 29, 3, 29, 315, 8, 29, 1, 30, 1, 30, 1, 30, 0, 0, 31, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 7, 1, 0, 2, 3, 1, 0, 34, 35, 2,
		0, 19, 20, 32, 32, 1, 0, 21, 23, 1, 0, 19, 20, 2, 0, 24, 25, 28, 29, 1,
		0, 26, 27, 329, 0, 65, 1, 0, 0, 0, 2, 72, 1, 0, 0, 0, 4, 76, 1, 0, 0, 0,
		6, 78, 1, 0, 0, 0, 8, 90, 1, 0, 0, 0, 10, 116, 1, 0, 0, 0, 12, 118, 1,
		0, 0, 0, 14, 129, 1, 0, 0, 0, 16, 156, 1, 0, 0, 0, 18, 158, 1, 0, 0, 0,
		20, 167, 1, 0, 0, 0, 22, 169, 1, 0, 0, 0, 24, 177, 1, 0, 0, 0, 26, 192,
		1, 0, 0, 0, 28, 203, 1, 0, 0, 0, 30, 239, 1, 0, 0, 0, 32, 241, 1, 0, 0,
		0, 34, 243, 1, 0, 0, 0, 36, 245, 1, 0, 0, 0, 38, 261, 1, 0, 0, 0, 40, 263,
		1, 0, 0, 0, 42, 274, 1, 0, 0, 0, 44, 276, 1, 0, 0, 0, 46, 278, 1, 0, 0,
		0, 48, 286, 1, 0, 0, 0, 50, 291, 1, 0, 0, 0, 52, 296, 1, 0, 0, 0, 54, 301,
		1, 0, 0, 0, 56, 306, 1, 0, 0, 0, 58, 311, 1, 0, 0, 0, 60, 316, 1, 0, 0,
		0, 62, 64, 3, 2, 1, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63,
		1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0,
		68, 69, 5, 0, 0, 1, 69, 1, 1, 0, 0, 0, 70, 73, 3, 4, 2, 0, 71, 73, 3, 18,
		9, 0, 72, 70, 1, 0, 0, 0, 72, 71, 1, 0, 0, 0, 73, 3, 1, 0, 0, 0, 74, 77,
		3, 6, 3, 0, 75, 77, 3, 12, 6, 0, 76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0,
		77, 5, 1, 0, 0, 0, 78, 79, 5, 1, 0, 0, 79, 80, 5, 3, 0, 0, 80, 85, 3, 8,
		4, 0, 81, 82, 5, 16, 0, 0, 82, 84, 3, 8, 4, 0, 83, 81, 1, 0, 0, 0, 84,
		87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 1, 0, 0,
		0, 87, 85, 1, 0, 0, 0, 88, 89, 5, 17, 0, 0, 89, 7, 1, 0, 0, 0, 90, 97,
		5, 33, 0, 0, 91, 92, 5, 12, 0, 0, 92, 93, 3, 60, 30, 0, 93, 94, 5, 13,
		0, 0, 94, 96, 1, 0, 0, 0, 95, 91, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95,
		1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 100, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0,
		100, 101, 5, 18, 0, 0, 101, 102, 3, 10, 5, 0, 102, 9, 1, 0, 0, 0, 103,
		117, 3, 60, 30, 0, 104, 113, 5, 14, 0, 0, 105, 110, 3, 60, 30, 0, 106,
		107, 5, 16, 0, 0, 107, 109, 3, 60, 30, 0, 108, 106, 1, 0, 0, 0, 109, 112,
		1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 114, 1, 0,
		0, 0, 112, 110, 1, 0, 0, 0, 113, 105, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0,
		114, 115, 1, 0, 0, 0, 115, 117, 5, 15, 0, 0, 116, 103, 1, 0, 0, 0, 116,
		104, 1, 0, 0, 0, 117, 11, 1, 0, 0, 0, 118, 119, 5, 3, 0, 0, 119, 124, 3,
		14, 7, 0, 120, 121, 5, 16, 0, 0, 121, 123, 3, 14, 7, 0, 122, 120, 1, 0,
		0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		125, 127, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 128, 5, 17, 0, 0, 128,
		13, 1, 0, 0, 0, 129, 136, 5, 33, 0, 0, 130, 131, 5, 12, 0, 0, 131, 132,
		3, 60, 30, 0, 132, 133, 5, 13, 0, 0, 133, 135, 1, 0, 0, 0, 134, 130, 1,
		0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0,
		0, 137, 141, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 140, 5, 18, 0, 0, 140,
		142, 3, 16, 8, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 15,
		1, 0, 0, 0, 143, 157, 3, 32, 16, 0, 144, 153, 5, 14, 0, 0, 145, 150, 3,
		32, 16, 0, 146, 147, 5, 16, 0, 0, 147, 149, 3, 32, 16, 0, 148, 146, 1,
		0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0,
		0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 145, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 157, 5, 15, 0, 0, 156, 143,
		1, 0, 0, 0, 156, 144, 1, 0, 0, 0, 157, 17, 1, 0, 0, 0, 158, 159, 3, 20,
		10, 0, 159, 160, 5, 33, 0, 0, 160, 162, 5, 10, 0, 0, 161, 163, 3, 22, 11,
		0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		165, 5, 11, 0, 0, 165, 166, 3, 26, 13, 0, 166, 19, 1, 0, 0, 0, 167, 168,
		7, 0, 0, 0, 168, 21, 1, 0, 0, 0, 169, 174, 3, 24, 12, 0, 170, 171, 5, 16,
		0, 0, 171, 173, 3, 24, 12, 0, 172, 170, 1, 0, 0, 0, 173, 176, 1, 0, 0,
		0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 23, 1, 0, 0, 0, 176,
		174, 1, 0, 0, 0, 177, 178, 5, 3, 0, 0, 178, 190, 5, 33, 0, 0, 179, 180,
		5, 12, 0, 0, 180, 187, 5, 13, 0, 0, 181, 182, 5, 12, 0, 0, 182, 183, 3,
		60, 30, 0, 183, 184, 5, 13, 0, 0, 184, 186, 1, 0, 0, 0, 185, 181, 1, 0,
		0, 0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0,
		188, 191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 179, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 25, 1, 0, 0, 0, 192, 196, 5, 14, 0, 0, 193, 195,
		3, 28, 14, 0, 194, 193, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1,
		0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 1, 0, 0, 0, 198, 196, 1, 0, 0,
		0, 199, 200, 5, 15, 0, 0, 200, 27, 1, 0, 0, 0, 201, 204, 3, 4, 2, 0, 202,
		204, 3, 30, 15, 0, 203, 201, 1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 29,
		1, 0, 0, 0, 205, 206, 3, 36, 18, 0, 206, 207, 5, 18, 0, 0, 207, 208, 3,
		32, 16, 0, 208, 209, 5, 17, 0, 0, 209, 240, 1, 0, 0, 0, 210, 212, 3, 32,
		16, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0,
		213, 240, 5, 17, 0, 0, 214, 240, 3, 26, 13, 0, 215, 216, 5, 4, 0, 0, 216,
		217, 5, 10, 0, 0, 217, 218, 3, 34, 17, 0, 218, 219, 5, 11, 0, 0, 219, 222,
		3, 30, 15, 0, 220, 221, 5, 5, 0, 0, 221, 223, 3, 30, 15, 0, 222, 220, 1,
		0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 240, 1, 0, 0, 0, 224, 225, 5, 6, 0,
		0, 225, 226, 5, 10, 0, 0, 226, 227, 3, 34, 17, 0, 227, 228, 5, 11, 0, 0,
		228, 229, 3, 30, 15, 0, 229, 240, 1, 0, 0, 0, 230, 231, 5, 7, 0, 0, 231,
		240, 5, 17, 0, 0, 232, 233, 5, 8, 0, 0, 233, 240, 5, 17, 0, 0, 234, 236,
		5, 9, 0, 0, 235, 237, 3, 32, 16, 0, 236, 235, 1, 0, 0, 0, 236, 237, 1,
		0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 240, 5, 17, 0, 0, 239, 205, 1, 0, 0,
		0, 239, 211, 1, 0, 0, 0, 239, 214, 1, 0, 0, 0, 239, 215, 1, 0, 0, 0, 239,
		224, 1, 0, 0, 0, 239, 230, 1, 0, 0, 0, 239, 232, 1, 0, 0, 0, 239, 234,
		1, 0, 0, 0, 240, 31, 1, 0, 0, 0, 241, 242, 3, 50, 25, 0, 242, 33, 1, 0,
		0, 0, 243, 244, 3, 58, 29, 0, 244, 35, 1, 0, 0, 0, 245, 252, 5, 33, 0,
		0, 246, 247, 5, 12, 0, 0, 247, 248, 3, 32, 16, 0, 248, 249, 5, 13, 0, 0,
		249, 251, 1, 0, 0, 0, 250, 246, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252,
		250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 37, 1, 0, 0, 0, 254, 252, 1,
		0, 0, 0, 255, 256, 5, 10, 0, 0, 256, 257, 3, 32, 16, 0, 257, 258, 5, 11,
		0, 0, 258, 262, 1, 0, 0, 0, 259, 262, 3, 36, 18, 0, 260, 262, 3, 40, 20,
		0, 261, 255, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 260, 1, 0, 0, 0, 262,
		39, 1, 0, 0, 0, 263, 264, 7, 1, 0, 0, 264, 41, 1, 0, 0, 0, 265, 275, 3,
		38, 19, 0, 266, 267, 5, 33, 0, 0, 267, 268, 5, 10, 0, 0, 268, 269, 3, 46,
		23, 0, 269, 270, 5, 11, 0, 0, 270, 275, 1, 0, 0, 0, 271, 272, 3, 44, 22,
		0, 272, 273, 3, 42, 21, 0, 273, 275, 1, 0, 0, 0, 274, 265, 1, 0, 0, 0,
		274, 266, 1, 0, 0, 0, 274, 271, 1, 0, 0, 0, 275, 43, 1, 0, 0, 0, 276, 277,
		7, 2, 0, 0, 277, 45, 1, 0, 0, 0, 278, 283, 3, 32, 16, 0, 279, 280, 5, 16,
		0, 0, 280, 282, 3, 32, 16, 0, 281, 279, 1, 0, 0, 0, 282, 285, 1, 0, 0,
		0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 47, 1, 0, 0, 0, 285,
		283, 1, 0, 0, 0, 286, 289, 3, 42, 21, 0, 287, 288, 7, 3, 0, 0, 288, 290,
		3, 48, 24, 0, 289, 287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 49, 1, 0,
		0, 0, 291, 294, 3, 48, 24, 0, 292, 293, 7, 4, 0, 0, 293, 295, 3, 50, 25,
		0, 294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 51, 1, 0, 0, 0, 296,
		299, 3, 50, 25, 0, 297, 298, 7, 5, 0, 0, 298, 300, 3, 52, 26, 0, 299, 297,
		1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 53, 1, 0, 0, 0, 301, 304, 3, 52,
		26, 0, 302, 303, 7, 6, 0, 0, 303, 305, 3, 54, 27, 0, 304, 302, 1, 0, 0,
		0, 304, 305, 1, 0, 0, 0, 305, 55, 1, 0, 0, 0, 306, 309, 3, 54, 27, 0, 307,
		308, 5, 30, 0, 0, 308, 310, 3, 56, 28, 0, 309, 307, 1, 0, 0, 0, 309, 310,
		1, 0, 0, 0, 310, 57, 1, 0, 0, 0, 311, 314, 3, 56, 28, 0, 312, 313, 5, 31,
		0, 0, 313, 315, 3, 58, 29, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0,
		0, 315, 59, 1, 0, 0, 0, 316, 317, 3, 50, 25, 0, 317, 61, 1, 0, 0, 0, 34,
		65, 72, 76, 85, 97, 110, 113, 116, 124, 136, 141, 150, 153, 156, 162, 174,
		187, 190, 196, 203, 211, 222, 236, 239, 252, 261, 274, 283, 289, 294, 299,
		304, 309, 314,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// sysyParserInit initializes any static state used to implement sysyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewsysyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SysyParserInit() {
	staticData := &SysyParserParserStaticData
	staticData.once.Do(sysyparserParserInit)
}

// NewsysyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewsysyParser(input antlr.TokenStream) *sysyParser {
	SysyParserInit()
	this := new(sysyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SysyParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "sysyParser.g4"

	return this
}

// sysyParser tokens.
const (
	sysyParserEOF             = antlr.TokenEOF
	sysyParserConst           = 1
	sysyParserVoid            = 2
	sysyParserType            = 3
	sysyParserIf              = 4
	sysyParserElse            = 5
	sysyParserWhile           = 6
	sysyParserBreak           = 7
	sysyParserContinue        = 8
	sysyParserReturn          = 9
	sysyParserLeftParen       = 10
	sysyParserRightParen      = 11
	sysyParserLeftBracket     = 12
	sysyParserRightBracket    = 13
	sysyParserLeftBrace       = 14
	sysyParserRightBrace      = 15
	sysyParserComma           = 16
	sysyParserSemicolon       = 17
	sysyParserAssign          = 18
	sysyParserAdd             = 19
	sysyParserSub             = 20
	sysyParserMul             = 21
	sysyParserDiv             = 22
	sysyParserMod             = 23
	sysyParserLess            = 24
	sysyParserGreater         = 25
	sysyParserEqual           = 26
	sysyParserNotEqual        = 27
	sysyParserLessEqual       = 28
	sysyParserGreaterEqual    = 29
	sysyParserAnd             = 30
	sysyParserOr              = 31
	sysyParserNot             = 32
	sysyParserIdentifier      = 33
	sysyParserIntegerConstant = 34
	sysyParserFloatConstant   = 35
	sysyParserWS              = 36
	sysyParserLineComment     = 37
	sysyParserBlockComment    = 38
	sysyParserPreprocessor    = 39
)

// sysyParser rules.
const (
	sysyParserRULE_comp         = 0
	sysyParserRULE_compUnit     = 1
	sysyParserRULE_decl         = 2
	sysyParserRULE_constDecl    = 3
	sysyParserRULE_constDef     = 4
	sysyParserRULE_constinitVal = 5
	sysyParserRULE_varDecl      = 6
	sysyParserRULE_varDef       = 7
	sysyParserRULE_initVal      = 8
	sysyParserRULE_funcDef      = 9
	sysyParserRULE_funcType     = 10
	sysyParserRULE_funcFParams  = 11
	sysyParserRULE_funcFParam   = 12
	sysyParserRULE_block        = 13
	sysyParserRULE_blockItem    = 14
	sysyParserRULE_stmt         = 15
	sysyParserRULE_exp          = 16
	sysyParserRULE_cond         = 17
	sysyParserRULE_lVal         = 18
	sysyParserRULE_primaryExp   = 19
	sysyParserRULE_number       = 20
	sysyParserRULE_unaryExp     = 21
	sysyParserRULE_unaryOp      = 22
	sysyParserRULE_funcRParams  = 23
	sysyParserRULE_mulExp       = 24
	sysyParserRULE_addExp       = 25
	sysyParserRULE_relExp       = 26
	sysyParserRULE_eqExp        = 27
	sysyParserRULE_lAndExp      = 28
	sysyParserRULE_lOrExp       = 29
	sysyParserRULE_constExp     = 30
)

// ICompContext is an interface to support dynamic dispatch.
type ICompContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllCompUnit() []ICompUnitContext
	CompUnit(i int) ICompUnitContext

	// IsCompContext differentiates from other interfaces.
	IsCompContext()
}

type CompContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompContext() *CompContext {
	var p = new(CompContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_comp
	return p
}

func InitEmptyCompContext(p *CompContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_comp
}

func (*CompContext) IsCompContext() {}

func NewCompContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompContext {
	var p = new(CompContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_comp

	return p
}

func (s *CompContext) GetParser() antlr.Parser { return s.parser }

func (s *CompContext) EOF() antlr.TerminalNode {
	return s.GetToken(sysyParserEOF, 0)
}

func (s *CompContext) AllCompUnit() []ICompUnitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICompUnitContext); ok {
			len++
		}
	}

	tst := make([]ICompUnitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICompUnitContext); ok {
			tst[i] = t.(ICompUnitContext)
			i++
		}
	}

	return tst
}

func (s *CompContext) CompUnit(i int) ICompUnitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompUnitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompUnitContext)
}

func (s *CompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitComp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) Comp() (localctx ICompContext) {
	localctx = NewCompContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, sysyParserRULE_comp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0 {
		{
			p.SetState(62)
			p.CompUnit()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(sysyParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompUnitContext is an interface to support dynamic dispatch.
type ICompUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl() IDeclContext
	FuncDef() IFuncDefContext

	// IsCompUnitContext differentiates from other interfaces.
	IsCompUnitContext()
}

type CompUnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompUnitContext() *CompUnitContext {
	var p = new(CompUnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_compUnit
	return p
}

func InitEmptyCompUnitContext(p *CompUnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_compUnit
}

func (*CompUnitContext) IsCompUnitContext() {}

func NewCompUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompUnitContext {
	var p = new(CompUnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_compUnit

	return p
}

func (s *CompUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompUnitContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *CompUnitContext) FuncDef() IFuncDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *CompUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitCompUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) CompUnit() (localctx ICompUnitContext) {
	localctx = NewCompUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, sysyParserRULE_compUnit)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Decl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.FuncDef()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConstDecl() IConstDeclContext
	VarDecl() IVarDeclContext

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_decl
	return p
}

func InitEmptyDeclContext(p *DeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_decl
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) ConstDecl() IConstDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

func (s *DeclContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, sysyParserRULE_decl)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case sysyParserConst:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.ConstDecl()
		}

	case sysyParserType:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.VarDecl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstDeclContext is an interface to support dynamic dispatch.
type IConstDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const() antlr.TerminalNode
	Type() antlr.TerminalNode
	AllConstDef() []IConstDefContext
	ConstDef(i int) IConstDefContext
	Semicolon() antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsConstDeclContext differentiates from other interfaces.
	IsConstDeclContext()
}

type ConstDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDeclContext() *ConstDeclContext {
	var p = new(ConstDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_constDecl
	return p
}

func InitEmptyConstDeclContext(p *ConstDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_constDecl
}

func (*ConstDeclContext) IsConstDeclContext() {}

func NewConstDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDeclContext {
	var p = new(ConstDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_constDecl

	return p
}

func (s *ConstDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDeclContext) Const() antlr.TerminalNode {
	return s.GetToken(sysyParserConst, 0)
}

func (s *ConstDeclContext) Type() antlr.TerminalNode {
	return s.GetToken(sysyParserType, 0)
}

func (s *ConstDeclContext) AllConstDef() []IConstDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstDefContext); ok {
			len++
		}
	}

	tst := make([]IConstDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstDefContext); ok {
			tst[i] = t.(IConstDefContext)
			i++
		}
	}

	return tst
}

func (s *ConstDeclContext) ConstDef(i int) IConstDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDefContext)
}

func (s *ConstDeclContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(sysyParserSemicolon, 0)
}

func (s *ConstDeclContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(sysyParserComma)
}

func (s *ConstDeclContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserComma, i)
}

func (s *ConstDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitConstDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) ConstDecl() (localctx IConstDeclContext) {
	localctx = NewConstDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, sysyParserRULE_constDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(sysyParserConst)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(sysyParserType)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.ConstDef()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == sysyParserComma {
		{
			p.SetState(81)
			p.Match(sysyParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.ConstDef()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(88)
		p.Match(sysyParserSemicolon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstDefContext is an interface to support dynamic dispatch.
type IConstDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Assign() antlr.TerminalNode
	ConstinitVal() IConstinitValContext
	AllLeftBracket() []antlr.TerminalNode
	LeftBracket(i int) antlr.TerminalNode
	AllConstExp() []IConstExpContext
	ConstExp(i int) IConstExpContext
	AllRightBracket() []antlr.TerminalNode
	RightBracket(i int) antlr.TerminalNode

	// IsConstDefContext differentiates from other interfaces.
	IsConstDefContext()
}

type ConstDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDefContext() *ConstDefContext {
	var p = new(ConstDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_constDef
	return p
}

func InitEmptyConstDefContext(p *ConstDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_constDef
}

func (*ConstDefContext) IsConstDefContext() {}

func NewConstDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDefContext {
	var p = new(ConstDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_constDef

	return p
}

func (s *ConstDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDefContext) Identifier() antlr.TerminalNode {
	return s.GetToken(sysyParserIdentifier, 0)
}

func (s *ConstDefContext) Assign() antlr.TerminalNode {
	return s.GetToken(sysyParserAssign, 0)
}

func (s *ConstDefContext) ConstinitVal() IConstinitValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstinitValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstinitValContext)
}

func (s *ConstDefContext) AllLeftBracket() []antlr.TerminalNode {
	return s.GetTokens(sysyParserLeftBracket)
}

func (s *ConstDefContext) LeftBracket(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserLeftBracket, i)
}

func (s *ConstDefContext) AllConstExp() []IConstExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstExpContext); ok {
			len++
		}
	}

	tst := make([]IConstExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstExpContext); ok {
			tst[i] = t.(IConstExpContext)
			i++
		}
	}

	return tst
}

func (s *ConstDefContext) ConstExp(i int) IConstExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstExpContext)
}

func (s *ConstDefContext) AllRightBracket() []antlr.TerminalNode {
	return s.GetTokens(sysyParserRightBracket)
}

func (s *ConstDefContext) RightBracket(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserRightBracket, i)
}

func (s *ConstDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitConstDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) ConstDef() (localctx IConstDefContext) {
	localctx = NewConstDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, sysyParserRULE_constDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(sysyParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == sysyParserLeftBracket {
		{
			p.SetState(91)
			p.Match(sysyParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.ConstExp()
		}
		{
			p.SetState(93)
			p.Match(sysyParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(100)
		p.Match(sysyParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.ConstinitVal()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstinitValContext is an interface to support dynamic dispatch.
type IConstinitValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConstExp() []IConstExpContext
	ConstExp(i int) IConstExpContext
	LeftBrace() antlr.TerminalNode
	RightBrace() antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsConstinitValContext differentiates from other interfaces.
	IsConstinitValContext()
}

type ConstinitValContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstinitValContext() *ConstinitValContext {
	var p = new(ConstinitValContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_constinitVal
	return p
}

func InitEmptyConstinitValContext(p *ConstinitValContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_constinitVal
}

func (*ConstinitValContext) IsConstinitValContext() {}

func NewConstinitValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstinitValContext {
	var p = new(ConstinitValContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_constinitVal

	return p
}

func (s *ConstinitValContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstinitValContext) AllConstExp() []IConstExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstExpContext); ok {
			len++
		}
	}

	tst := make([]IConstExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstExpContext); ok {
			tst[i] = t.(IConstExpContext)
			i++
		}
	}

	return tst
}

func (s *ConstinitValContext) ConstExp(i int) IConstExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstExpContext)
}

func (s *ConstinitValContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(sysyParserLeftBrace, 0)
}

func (s *ConstinitValContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(sysyParserRightBrace, 0)
}

func (s *ConstinitValContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(sysyParserComma)
}

func (s *ConstinitValContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserComma, i)
}

func (s *ConstinitValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstinitValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstinitValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitConstinitVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) ConstinitVal() (localctx IConstinitValContext) {
	localctx = NewConstinitValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, sysyParserRULE_constinitVal)
	var _la int

	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case sysyParserLeftParen, sysyParserAdd, sysyParserSub, sysyParserNot, sysyParserIdentifier, sysyParserIntegerConstant, sysyParserFloatConstant:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.ConstExp()
		}

	case sysyParserLeftBrace:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Match(sysyParserLeftBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&64426083328) != 0 {
			{
				p.SetState(105)
				p.ConstExp()
			}
			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == sysyParserComma {
				{
					p.SetState(106)
					p.Match(sysyParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(107)
					p.ConstExp()
				}

				p.SetState(112)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(115)
			p.Match(sysyParserRightBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type() antlr.TerminalNode
	AllVarDef() []IVarDefContext
	VarDef(i int) IVarDefContext
	Semicolon() antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) Type() antlr.TerminalNode {
	return s.GetToken(sysyParserType, 0)
}

func (s *VarDeclContext) AllVarDef() []IVarDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDefContext); ok {
			len++
		}
	}

	tst := make([]IVarDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDefContext); ok {
			tst[i] = t.(IVarDefContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclContext) VarDef(i int) IVarDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDefContext)
}

func (s *VarDeclContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(sysyParserSemicolon, 0)
}

func (s *VarDeclContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(sysyParserComma)
}

func (s *VarDeclContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserComma, i)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, sysyParserRULE_varDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(sysyParserType)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.VarDef()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == sysyParserComma {
		{
			p.SetState(120)
			p.Match(sysyParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.VarDef()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(127)
		p.Match(sysyParserSemicolon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDefContext is an interface to support dynamic dispatch.
type IVarDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	AllLeftBracket() []antlr.TerminalNode
	LeftBracket(i int) antlr.TerminalNode
	AllConstExp() []IConstExpContext
	ConstExp(i int) IConstExpContext
	AllRightBracket() []antlr.TerminalNode
	RightBracket(i int) antlr.TerminalNode
	Assign() antlr.TerminalNode
	InitVal() IInitValContext

	// IsVarDefContext differentiates from other interfaces.
	IsVarDefContext()
}

type VarDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDefContext() *VarDefContext {
	var p = new(VarDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_varDef
	return p
}

func InitEmptyVarDefContext(p *VarDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_varDef
}

func (*VarDefContext) IsVarDefContext() {}

func NewVarDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDefContext {
	var p = new(VarDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_varDef

	return p
}

func (s *VarDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDefContext) Identifier() antlr.TerminalNode {
	return s.GetToken(sysyParserIdentifier, 0)
}

func (s *VarDefContext) AllLeftBracket() []antlr.TerminalNode {
	return s.GetTokens(sysyParserLeftBracket)
}

func (s *VarDefContext) LeftBracket(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserLeftBracket, i)
}

func (s *VarDefContext) AllConstExp() []IConstExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstExpContext); ok {
			len++
		}
	}

	tst := make([]IConstExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstExpContext); ok {
			tst[i] = t.(IConstExpContext)
			i++
		}
	}

	return tst
}

func (s *VarDefContext) ConstExp(i int) IConstExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstExpContext)
}

func (s *VarDefContext) AllRightBracket() []antlr.TerminalNode {
	return s.GetTokens(sysyParserRightBracket)
}

func (s *VarDefContext) RightBracket(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserRightBracket, i)
}

func (s *VarDefContext) Assign() antlr.TerminalNode {
	return s.GetToken(sysyParserAssign, 0)
}

func (s *VarDefContext) InitVal() IInitValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitValContext)
}

func (s *VarDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitVarDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) VarDef() (localctx IVarDefContext) {
	localctx = NewVarDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, sysyParserRULE_varDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(sysyParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == sysyParserLeftBracket {
		{
			p.SetState(130)
			p.Match(sysyParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.ConstExp()
		}
		{
			p.SetState(132)
			p.Match(sysyParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == sysyParserAssign {
		{
			p.SetState(139)
			p.Match(sysyParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.InitVal()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInitValContext is an interface to support dynamic dispatch.
type IInitValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp() []IExpContext
	Exp(i int) IExpContext
	LeftBrace() antlr.TerminalNode
	RightBrace() antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsInitValContext differentiates from other interfaces.
	IsInitValContext()
}

type InitValContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitValContext() *InitValContext {
	var p = new(InitValContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_initVal
	return p
}

func InitEmptyInitValContext(p *InitValContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_initVal
}

func (*InitValContext) IsInitValContext() {}

func NewInitValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitValContext {
	var p = new(InitValContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_initVal

	return p
}

func (s *InitValContext) GetParser() antlr.Parser { return s.parser }

func (s *InitValContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *InitValContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *InitValContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(sysyParserLeftBrace, 0)
}

func (s *InitValContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(sysyParserRightBrace, 0)
}

func (s *InitValContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(sysyParserComma)
}

func (s *InitValContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserComma, i)
}

func (s *InitValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitInitVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) InitVal() (localctx IInitValContext) {
	localctx = NewInitValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, sysyParserRULE_initVal)
	var _la int

	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case sysyParserLeftParen, sysyParserAdd, sysyParserSub, sysyParserNot, sysyParserIdentifier, sysyParserIntegerConstant, sysyParserFloatConstant:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Exp()
		}

	case sysyParserLeftBrace:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.Match(sysyParserLeftBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&64426083328) != 0 {
			{
				p.SetState(145)
				p.Exp()
			}
			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == sysyParserComma {
				{
					p.SetState(146)
					p.Match(sysyParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(147)
					p.Exp()
				}

				p.SetState(152)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(155)
			p.Match(sysyParserRightBrace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncType() IFuncTypeContext
	Identifier() antlr.TerminalNode
	LeftParen() antlr.TerminalNode
	RightParen() antlr.TerminalNode
	Block() IBlockContext
	FuncFParams() IFuncFParamsContext

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_funcDef
	return p
}

func InitEmptyFuncDefContext(p *FuncDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_funcDef
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) FuncType() IFuncTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncTypeContext)
}

func (s *FuncDefContext) Identifier() antlr.TerminalNode {
	return s.GetToken(sysyParserIdentifier, 0)
}

func (s *FuncDefContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(sysyParserLeftParen, 0)
}

func (s *FuncDefContext) RightParen() antlr.TerminalNode {
	return s.GetToken(sysyParserRightParen, 0)
}

func (s *FuncDefContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDefContext) FuncFParams() IFuncFParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncFParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncFParamsContext)
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitFuncDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) FuncDef() (localctx IFuncDefContext) {
	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, sysyParserRULE_funcDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.FuncType()
	}
	{
		p.SetState(159)
		p.Match(sysyParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(sysyParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == sysyParserType {
		{
			p.SetState(161)
			p.FuncFParams()
		}

	}
	{
		p.SetState(164)
		p.Match(sysyParserRightParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncTypeContext is an interface to support dynamic dispatch.
type IFuncTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type() antlr.TerminalNode
	Void() antlr.TerminalNode

	// IsFuncTypeContext differentiates from other interfaces.
	IsFuncTypeContext()
}

type FuncTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncTypeContext() *FuncTypeContext {
	var p = new(FuncTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_funcType
	return p
}

func InitEmptyFuncTypeContext(p *FuncTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_funcType
}

func (*FuncTypeContext) IsFuncTypeContext() {}

func NewFuncTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncTypeContext {
	var p = new(FuncTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_funcType

	return p
}

func (s *FuncTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncTypeContext) Type() antlr.TerminalNode {
	return s.GetToken(sysyParserType, 0)
}

func (s *FuncTypeContext) Void() antlr.TerminalNode {
	return s.GetToken(sysyParserVoid, 0)
}

func (s *FuncTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitFuncType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) FuncType() (localctx IFuncTypeContext) {
	localctx = NewFuncTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, sysyParserRULE_funcType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		_la = p.GetTokenStream().LA(1)

		if !(_la == sysyParserVoid || _la == sysyParserType) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncFParamsContext is an interface to support dynamic dispatch.
type IFuncFParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFuncFParam() []IFuncFParamContext
	FuncFParam(i int) IFuncFParamContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsFuncFParamsContext differentiates from other interfaces.
	IsFuncFParamsContext()
}

type FuncFParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncFParamsContext() *FuncFParamsContext {
	var p = new(FuncFParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_funcFParams
	return p
}

func InitEmptyFuncFParamsContext(p *FuncFParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_funcFParams
}

func (*FuncFParamsContext) IsFuncFParamsContext() {}

func NewFuncFParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncFParamsContext {
	var p = new(FuncFParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_funcFParams

	return p
}

func (s *FuncFParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncFParamsContext) AllFuncFParam() []IFuncFParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncFParamContext); ok {
			len++
		}
	}

	tst := make([]IFuncFParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncFParamContext); ok {
			tst[i] = t.(IFuncFParamContext)
			i++
		}
	}

	return tst
}

func (s *FuncFParamsContext) FuncFParam(i int) IFuncFParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncFParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncFParamContext)
}

func (s *FuncFParamsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(sysyParserComma)
}

func (s *FuncFParamsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserComma, i)
}

func (s *FuncFParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncFParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitFuncFParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) FuncFParams() (localctx IFuncFParamsContext) {
	localctx = NewFuncFParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, sysyParserRULE_funcFParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.FuncFParam()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == sysyParserComma {
		{
			p.SetState(170)
			p.Match(sysyParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.FuncFParam()
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncFParamContext is an interface to support dynamic dispatch.
type IFuncFParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	AllLeftBracket() []antlr.TerminalNode
	LeftBracket(i int) antlr.TerminalNode
	AllRightBracket() []antlr.TerminalNode
	RightBracket(i int) antlr.TerminalNode
	AllConstExp() []IConstExpContext
	ConstExp(i int) IConstExpContext

	// IsFuncFParamContext differentiates from other interfaces.
	IsFuncFParamContext()
}

type FuncFParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncFParamContext() *FuncFParamContext {
	var p = new(FuncFParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_funcFParam
	return p
}

func InitEmptyFuncFParamContext(p *FuncFParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_funcFParam
}

func (*FuncFParamContext) IsFuncFParamContext() {}

func NewFuncFParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncFParamContext {
	var p = new(FuncFParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_funcFParam

	return p
}

func (s *FuncFParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncFParamContext) Type() antlr.TerminalNode {
	return s.GetToken(sysyParserType, 0)
}

func (s *FuncFParamContext) Identifier() antlr.TerminalNode {
	return s.GetToken(sysyParserIdentifier, 0)
}

func (s *FuncFParamContext) AllLeftBracket() []antlr.TerminalNode {
	return s.GetTokens(sysyParserLeftBracket)
}

func (s *FuncFParamContext) LeftBracket(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserLeftBracket, i)
}

func (s *FuncFParamContext) AllRightBracket() []antlr.TerminalNode {
	return s.GetTokens(sysyParserRightBracket)
}

func (s *FuncFParamContext) RightBracket(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserRightBracket, i)
}

func (s *FuncFParamContext) AllConstExp() []IConstExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstExpContext); ok {
			len++
		}
	}

	tst := make([]IConstExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstExpContext); ok {
			tst[i] = t.(IConstExpContext)
			i++
		}
	}

	return tst
}

func (s *FuncFParamContext) ConstExp(i int) IConstExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstExpContext)
}

func (s *FuncFParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncFParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitFuncFParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) FuncFParam() (localctx IFuncFParamContext) {
	localctx = NewFuncFParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, sysyParserRULE_funcFParam)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(sysyParserType)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Match(sysyParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == sysyParserLeftBracket {
		{
			p.SetState(179)
			p.Match(sysyParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Match(sysyParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == sysyParserLeftBracket {
			{
				p.SetState(181)
				p.Match(sysyParserLeftBracket)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(182)
				p.ConstExp()
			}
			{
				p.SetState(183)
				p.Match(sysyParserRightBracket)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(189)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftBrace() antlr.TerminalNode
	RightBrace() antlr.TerminalNode
	AllBlockItem() []IBlockItemContext
	BlockItem(i int) IBlockItemContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(sysyParserLeftBrace, 0)
}

func (s *BlockContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(sysyParserRightBrace, 0)
}

func (s *BlockContext) AllBlockItem() []IBlockItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockItemContext); ok {
			len++
		}
	}

	tst := make([]IBlockItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockItemContext); ok {
			tst[i] = t.(IBlockItemContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) BlockItem(i int) IBlockItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockItemContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, sysyParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(sysyParserLeftBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&64426231770) != 0 {
		{
			p.SetState(193)
			p.BlockItem()
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(199)
		p.Match(sysyParserRightBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockItemContext is an interface to support dynamic dispatch.
type IBlockItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl() IDeclContext
	Stmt() IStmtContext

	// IsBlockItemContext differentiates from other interfaces.
	IsBlockItemContext()
}

type BlockItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockItemContext() *BlockItemContext {
	var p = new(BlockItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_blockItem
	return p
}

func InitEmptyBlockItemContext(p *BlockItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_blockItem
}

func (*BlockItemContext) IsBlockItemContext() {}

func NewBlockItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockItemContext {
	var p = new(BlockItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_blockItem

	return p
}

func (s *BlockItemContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockItemContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *BlockItemContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitBlockItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) BlockItem() (localctx IBlockItemContext) {
	localctx = NewBlockItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, sysyParserRULE_blockItem)
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case sysyParserConst, sysyParserType:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.Decl()
		}

	case sysyParserIf, sysyParserWhile, sysyParserBreak, sysyParserContinue, sysyParserReturn, sysyParserLeftParen, sysyParserLeftBrace, sysyParserSemicolon, sysyParserAdd, sysyParserSub, sysyParserNot, sysyParserIdentifier, sysyParserIntegerConstant, sysyParserFloatConstant:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Stmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LVal() ILValContext
	Assign() antlr.TerminalNode
	Exp() IExpContext
	Semicolon() antlr.TerminalNode
	Block() IBlockContext
	If() antlr.TerminalNode
	LeftParen() antlr.TerminalNode
	Cond() ICondContext
	RightParen() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	Else() antlr.TerminalNode
	While() antlr.TerminalNode
	Break() antlr.TerminalNode
	Continue() antlr.TerminalNode
	Return() antlr.TerminalNode

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) LVal() ILValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILValContext)
}

func (s *StmtContext) Assign() antlr.TerminalNode {
	return s.GetToken(sysyParserAssign, 0)
}

func (s *StmtContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StmtContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(sysyParserSemicolon, 0)
}

func (s *StmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StmtContext) If() antlr.TerminalNode {
	return s.GetToken(sysyParserIf, 0)
}

func (s *StmtContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(sysyParserLeftParen, 0)
}

func (s *StmtContext) Cond() ICondContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondContext)
}

func (s *StmtContext) RightParen() antlr.TerminalNode {
	return s.GetToken(sysyParserRightParen, 0)
}

func (s *StmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *StmtContext) Else() antlr.TerminalNode {
	return s.GetToken(sysyParserElse, 0)
}

func (s *StmtContext) While() antlr.TerminalNode {
	return s.GetToken(sysyParserWhile, 0)
}

func (s *StmtContext) Break() antlr.TerminalNode {
	return s.GetToken(sysyParserBreak, 0)
}

func (s *StmtContext) Continue() antlr.TerminalNode {
	return s.GetToken(sysyParserContinue, 0)
}

func (s *StmtContext) Return() antlr.TerminalNode {
	return s.GetToken(sysyParserReturn, 0)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, sysyParserRULE_stmt)
	var _la int

	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.LVal()
		}
		{
			p.SetState(206)
			p.Match(sysyParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Exp()
		}
		{
			p.SetState(208)
			p.Match(sysyParserSemicolon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&64426083328) != 0 {
			{
				p.SetState(210)
				p.Exp()
			}

		}
		{
			p.SetState(213)
			p.Match(sysyParserSemicolon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(214)
			p.Block()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(215)
			p.Match(sysyParserIf)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Match(sysyParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Cond()
		}
		{
			p.SetState(218)
			p.Match(sysyParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Stmt()
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(220)
				p.Match(sysyParserElse)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(221)
				p.Stmt()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(224)
			p.Match(sysyParserWhile)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(sysyParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Cond()
		}
		{
			p.SetState(227)
			p.Match(sysyParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(230)
			p.Match(sysyParserBreak)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(sysyParserSemicolon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(232)
			p.Match(sysyParserContinue)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(sysyParserSemicolon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(234)
			p.Match(sysyParserReturn)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&64426083328) != 0 {
			{
				p.SetState(235)
				p.Exp()
			}

		}
		{
			p.SetState(238)
			p.Match(sysyParserSemicolon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AddExp() IAddExpContext

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) AddExp() IAddExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddExpContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, sysyParserRULE_exp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.AddExp()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICondContext is an interface to support dynamic dispatch.
type ICondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOrExp() ILOrExpContext

	// IsCondContext differentiates from other interfaces.
	IsCondContext()
}

type CondContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondContext() *CondContext {
	var p = new(CondContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_cond
	return p
}

func InitEmptyCondContext(p *CondContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_cond
}

func (*CondContext) IsCondContext() {}

func NewCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondContext {
	var p = new(CondContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_cond

	return p
}

func (s *CondContext) GetParser() antlr.Parser { return s.parser }

func (s *CondContext) LOrExp() ILOrExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILOrExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILOrExpContext)
}

func (s *CondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitCond(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) Cond() (localctx ICondContext) {
	localctx = NewCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, sysyParserRULE_cond)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.LOrExp()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILValContext is an interface to support dynamic dispatch.
type ILValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	AllLeftBracket() []antlr.TerminalNode
	LeftBracket(i int) antlr.TerminalNode
	AllExp() []IExpContext
	Exp(i int) IExpContext
	AllRightBracket() []antlr.TerminalNode
	RightBracket(i int) antlr.TerminalNode

	// IsLValContext differentiates from other interfaces.
	IsLValContext()
}

type LValContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLValContext() *LValContext {
	var p = new(LValContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_lVal
	return p
}

func InitEmptyLValContext(p *LValContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_lVal
}

func (*LValContext) IsLValContext() {}

func NewLValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LValContext {
	var p = new(LValContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_lVal

	return p
}

func (s *LValContext) GetParser() antlr.Parser { return s.parser }

func (s *LValContext) Identifier() antlr.TerminalNode {
	return s.GetToken(sysyParserIdentifier, 0)
}

func (s *LValContext) AllLeftBracket() []antlr.TerminalNode {
	return s.GetTokens(sysyParserLeftBracket)
}

func (s *LValContext) LeftBracket(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserLeftBracket, i)
}

func (s *LValContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *LValContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *LValContext) AllRightBracket() []antlr.TerminalNode {
	return s.GetTokens(sysyParserRightBracket)
}

func (s *LValContext) RightBracket(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserRightBracket, i)
}

func (s *LValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitLVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) LVal() (localctx ILValContext) {
	localctx = NewLValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, sysyParserRULE_lVal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(sysyParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == sysyParserLeftBracket {
		{
			p.SetState(246)
			p.Match(sysyParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Exp()
		}
		{
			p.SetState(248)
			p.Match(sysyParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpContext is an interface to support dynamic dispatch.
type IPrimaryExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftParen() antlr.TerminalNode
	Exp() IExpContext
	RightParen() antlr.TerminalNode
	LVal() ILValContext
	Number() INumberContext

	// IsPrimaryExpContext differentiates from other interfaces.
	IsPrimaryExpContext()
}

type PrimaryExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpContext() *PrimaryExpContext {
	var p = new(PrimaryExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_primaryExp
	return p
}

func InitEmptyPrimaryExpContext(p *PrimaryExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_primaryExp
}

func (*PrimaryExpContext) IsPrimaryExpContext() {}

func NewPrimaryExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpContext {
	var p = new(PrimaryExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_primaryExp

	return p
}

func (s *PrimaryExpContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(sysyParserLeftParen, 0)
}

func (s *PrimaryExpContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *PrimaryExpContext) RightParen() antlr.TerminalNode {
	return s.GetToken(sysyParserRightParen, 0)
}

func (s *PrimaryExpContext) LVal() ILValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILValContext)
}

func (s *PrimaryExpContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *PrimaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitPrimaryExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) PrimaryExp() (localctx IPrimaryExpContext) {
	localctx = NewPrimaryExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, sysyParserRULE_primaryExp)
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case sysyParserLeftParen:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(255)
			p.Match(sysyParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Exp()
		}
		{
			p.SetState(257)
			p.Match(sysyParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case sysyParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)
			p.LVal()
		}

	case sysyParserIntegerConstant, sysyParserFloatConstant:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.Number()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerConstant() antlr.TerminalNode
	FloatConstant() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) IntegerConstant() antlr.TerminalNode {
	return s.GetToken(sysyParserIntegerConstant, 0)
}

func (s *NumberContext) FloatConstant() antlr.TerminalNode {
	return s.GetToken(sysyParserFloatConstant, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, sysyParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		_la = p.GetTokenStream().LA(1)

		if !(_la == sysyParserIntegerConstant || _la == sysyParserFloatConstant) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpContext is an interface to support dynamic dispatch.
type IUnaryExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExp() IPrimaryExpContext
	Identifier() antlr.TerminalNode
	LeftParen() antlr.TerminalNode
	FuncRParams() IFuncRParamsContext
	RightParen() antlr.TerminalNode
	UnaryOp() IUnaryOpContext
	UnaryExp() IUnaryExpContext

	// IsUnaryExpContext differentiates from other interfaces.
	IsUnaryExpContext()
}

type UnaryExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpContext() *UnaryExpContext {
	var p = new(UnaryExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_unaryExp
	return p
}

func InitEmptyUnaryExpContext(p *UnaryExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_unaryExp
}

func (*UnaryExpContext) IsUnaryExpContext() {}

func NewUnaryExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpContext {
	var p = new(UnaryExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_unaryExp

	return p
}

func (s *UnaryExpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpContext) PrimaryExp() IPrimaryExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpContext)
}

func (s *UnaryExpContext) Identifier() antlr.TerminalNode {
	return s.GetToken(sysyParserIdentifier, 0)
}

func (s *UnaryExpContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(sysyParserLeftParen, 0)
}

func (s *UnaryExpContext) FuncRParams() IFuncRParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncRParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncRParamsContext)
}

func (s *UnaryExpContext) RightParen() antlr.TerminalNode {
	return s.GetToken(sysyParserRightParen, 0)
}

func (s *UnaryExpContext) UnaryOp() IUnaryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *UnaryExpContext) UnaryExp() IUnaryExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpContext)
}

func (s *UnaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitUnaryExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) UnaryExp() (localctx IUnaryExpContext) {
	localctx = NewUnaryExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, sysyParserRULE_unaryExp)
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.PrimaryExp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(sysyParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(sysyParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.FuncRParams()
		}
		{
			p.SetState(269)
			p.Match(sysyParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(271)
			p.UnaryOp()
		}
		{
			p.SetState(272)
			p.UnaryExp()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Add() antlr.TerminalNode
	Sub() antlr.TerminalNode
	Not() antlr.TerminalNode

	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_unaryOp
	return p
}

func InitEmptyUnaryOpContext(p *UnaryOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_unaryOp
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpContext) Add() antlr.TerminalNode {
	return s.GetToken(sysyParserAdd, 0)
}

func (s *UnaryOpContext) Sub() antlr.TerminalNode {
	return s.GetToken(sysyParserSub, 0)
}

func (s *UnaryOpContext) Not() antlr.TerminalNode {
	return s.GetToken(sysyParserNot, 0)
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitUnaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, sysyParserRULE_unaryOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4296540160) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncRParamsContext is an interface to support dynamic dispatch.
type IFuncRParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp() []IExpContext
	Exp(i int) IExpContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsFuncRParamsContext differentiates from other interfaces.
	IsFuncRParamsContext()
}

type FuncRParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncRParamsContext() *FuncRParamsContext {
	var p = new(FuncRParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_funcRParams
	return p
}

func InitEmptyFuncRParamsContext(p *FuncRParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_funcRParams
}

func (*FuncRParamsContext) IsFuncRParamsContext() {}

func NewFuncRParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncRParamsContext {
	var p = new(FuncRParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_funcRParams

	return p
}

func (s *FuncRParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncRParamsContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *FuncRParamsContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *FuncRParamsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(sysyParserComma)
}

func (s *FuncRParamsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(sysyParserComma, i)
}

func (s *FuncRParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncRParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncRParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitFuncRParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) FuncRParams() (localctx IFuncRParamsContext) {
	localctx = NewFuncRParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, sysyParserRULE_funcRParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Exp()
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == sysyParserComma {
		{
			p.SetState(279)
			p.Match(sysyParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Exp()
		}

		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMulExpContext is an interface to support dynamic dispatch.
type IMulExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExp() IUnaryExpContext
	MulExp() IMulExpContext
	Mul() antlr.TerminalNode
	Div() antlr.TerminalNode
	Mod() antlr.TerminalNode

	// IsMulExpContext differentiates from other interfaces.
	IsMulExpContext()
}

type MulExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulExpContext() *MulExpContext {
	var p = new(MulExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_mulExp
	return p
}

func InitEmptyMulExpContext(p *MulExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_mulExp
}

func (*MulExpContext) IsMulExpContext() {}

func NewMulExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulExpContext {
	var p = new(MulExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_mulExp

	return p
}

func (s *MulExpContext) GetParser() antlr.Parser { return s.parser }

func (s *MulExpContext) UnaryExp() IUnaryExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpContext)
}

func (s *MulExpContext) MulExp() IMulExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulExpContext)
}

func (s *MulExpContext) Mul() antlr.TerminalNode {
	return s.GetToken(sysyParserMul, 0)
}

func (s *MulExpContext) Div() antlr.TerminalNode {
	return s.GetToken(sysyParserDiv, 0)
}

func (s *MulExpContext) Mod() antlr.TerminalNode {
	return s.GetToken(sysyParserMod, 0)
}

func (s *MulExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitMulExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) MulExp() (localctx IMulExpContext) {
	localctx = NewMulExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, sysyParserRULE_mulExp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.UnaryExp()
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14680064) != 0 {
		{
			p.SetState(287)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14680064) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(288)
			p.MulExp()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddExpContext is an interface to support dynamic dispatch.
type IAddExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MulExp() IMulExpContext
	AddExp() IAddExpContext
	Add() antlr.TerminalNode
	Sub() antlr.TerminalNode

	// IsAddExpContext differentiates from other interfaces.
	IsAddExpContext()
}

type AddExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddExpContext() *AddExpContext {
	var p = new(AddExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_addExp
	return p
}

func InitEmptyAddExpContext(p *AddExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_addExp
}

func (*AddExpContext) IsAddExpContext() {}

func NewAddExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddExpContext {
	var p = new(AddExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_addExp

	return p
}

func (s *AddExpContext) GetParser() antlr.Parser { return s.parser }

func (s *AddExpContext) MulExp() IMulExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulExpContext)
}

func (s *AddExpContext) AddExp() IAddExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddExpContext)
}

func (s *AddExpContext) Add() antlr.TerminalNode {
	return s.GetToken(sysyParserAdd, 0)
}

func (s *AddExpContext) Sub() antlr.TerminalNode {
	return s.GetToken(sysyParserSub, 0)
}

func (s *AddExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitAddExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) AddExp() (localctx IAddExpContext) {
	localctx = NewAddExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, sysyParserRULE_addExp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.MulExp()
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == sysyParserAdd || _la == sysyParserSub {
		{
			p.SetState(292)
			_la = p.GetTokenStream().LA(1)

			if !(_la == sysyParserAdd || _la == sysyParserSub) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(293)
			p.AddExp()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelExpContext is an interface to support dynamic dispatch.
type IRelExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AddExp() IAddExpContext
	RelExp() IRelExpContext
	Less() antlr.TerminalNode
	Greater() antlr.TerminalNode
	LessEqual() antlr.TerminalNode
	GreaterEqual() antlr.TerminalNode

	// IsRelExpContext differentiates from other interfaces.
	IsRelExpContext()
}

type RelExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelExpContext() *RelExpContext {
	var p = new(RelExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_relExp
	return p
}

func InitEmptyRelExpContext(p *RelExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_relExp
}

func (*RelExpContext) IsRelExpContext() {}

func NewRelExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelExpContext {
	var p = new(RelExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_relExp

	return p
}

func (s *RelExpContext) GetParser() antlr.Parser { return s.parser }

func (s *RelExpContext) AddExp() IAddExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddExpContext)
}

func (s *RelExpContext) RelExp() IRelExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelExpContext)
}

func (s *RelExpContext) Less() antlr.TerminalNode {
	return s.GetToken(sysyParserLess, 0)
}

func (s *RelExpContext) Greater() antlr.TerminalNode {
	return s.GetToken(sysyParserGreater, 0)
}

func (s *RelExpContext) LessEqual() antlr.TerminalNode {
	return s.GetToken(sysyParserLessEqual, 0)
}

func (s *RelExpContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(sysyParserGreaterEqual, 0)
}

func (s *RelExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitRelExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) RelExp() (localctx IRelExpContext) {
	localctx = NewRelExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, sysyParserRULE_relExp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.AddExp()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&855638016) != 0 {
		{
			p.SetState(297)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&855638016) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(298)
			p.RelExp()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqExpContext is an interface to support dynamic dispatch.
type IEqExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelExp() IRelExpContext
	EqExp() IEqExpContext
	Equal() antlr.TerminalNode
	NotEqual() antlr.TerminalNode

	// IsEqExpContext differentiates from other interfaces.
	IsEqExpContext()
}

type EqExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqExpContext() *EqExpContext {
	var p = new(EqExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_eqExp
	return p
}

func InitEmptyEqExpContext(p *EqExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_eqExp
}

func (*EqExpContext) IsEqExpContext() {}

func NewEqExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqExpContext {
	var p = new(EqExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_eqExp

	return p
}

func (s *EqExpContext) GetParser() antlr.Parser { return s.parser }

func (s *EqExpContext) RelExp() IRelExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelExpContext)
}

func (s *EqExpContext) EqExp() IEqExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqExpContext)
}

func (s *EqExpContext) Equal() antlr.TerminalNode {
	return s.GetToken(sysyParserEqual, 0)
}

func (s *EqExpContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(sysyParserNotEqual, 0)
}

func (s *EqExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitEqExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) EqExp() (localctx IEqExpContext) {
	localctx = NewEqExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, sysyParserRULE_eqExp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.RelExp()
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == sysyParserEqual || _la == sysyParserNotEqual {
		{
			p.SetState(302)
			_la = p.GetTokenStream().LA(1)

			if !(_la == sysyParserEqual || _la == sysyParserNotEqual) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(303)
			p.EqExp()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILAndExpContext is an interface to support dynamic dispatch.
type ILAndExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqExp() IEqExpContext
	And() antlr.TerminalNode
	LAndExp() ILAndExpContext

	// IsLAndExpContext differentiates from other interfaces.
	IsLAndExpContext()
}

type LAndExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLAndExpContext() *LAndExpContext {
	var p = new(LAndExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_lAndExp
	return p
}

func InitEmptyLAndExpContext(p *LAndExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_lAndExp
}

func (*LAndExpContext) IsLAndExpContext() {}

func NewLAndExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LAndExpContext {
	var p = new(LAndExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_lAndExp

	return p
}

func (s *LAndExpContext) GetParser() antlr.Parser { return s.parser }

func (s *LAndExpContext) EqExp() IEqExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqExpContext)
}

func (s *LAndExpContext) And() antlr.TerminalNode {
	return s.GetToken(sysyParserAnd, 0)
}

func (s *LAndExpContext) LAndExp() ILAndExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILAndExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILAndExpContext)
}

func (s *LAndExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LAndExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LAndExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitLAndExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) LAndExp() (localctx ILAndExpContext) {
	localctx = NewLAndExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, sysyParserRULE_lAndExp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.EqExp()
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == sysyParserAnd {
		{
			p.SetState(307)
			p.Match(sysyParserAnd)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.LAndExp()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILOrExpContext is an interface to support dynamic dispatch.
type ILOrExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LAndExp() ILAndExpContext
	Or() antlr.TerminalNode
	LOrExp() ILOrExpContext

	// IsLOrExpContext differentiates from other interfaces.
	IsLOrExpContext()
}

type LOrExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLOrExpContext() *LOrExpContext {
	var p = new(LOrExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_lOrExp
	return p
}

func InitEmptyLOrExpContext(p *LOrExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_lOrExp
}

func (*LOrExpContext) IsLOrExpContext() {}

func NewLOrExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LOrExpContext {
	var p = new(LOrExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_lOrExp

	return p
}

func (s *LOrExpContext) GetParser() antlr.Parser { return s.parser }

func (s *LOrExpContext) LAndExp() ILAndExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILAndExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILAndExpContext)
}

func (s *LOrExpContext) Or() antlr.TerminalNode {
	return s.GetToken(sysyParserOr, 0)
}

func (s *LOrExpContext) LOrExp() ILOrExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILOrExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILOrExpContext)
}

func (s *LOrExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LOrExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LOrExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitLOrExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) LOrExp() (localctx ILOrExpContext) {
	localctx = NewLOrExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, sysyParserRULE_lOrExp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.LAndExp()
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == sysyParserOr {
		{
			p.SetState(312)
			p.Match(sysyParserOr)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.LOrExp()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstExpContext is an interface to support dynamic dispatch.
type IConstExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AddExp() IAddExpContext

	// IsConstExpContext differentiates from other interfaces.
	IsConstExpContext()
}

type ConstExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstExpContext() *ConstExpContext {
	var p = new(ConstExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_constExp
	return p
}

func InitEmptyConstExpContext(p *ConstExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = sysyParserRULE_constExp
}

func (*ConstExpContext) IsConstExpContext() {}

func NewConstExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstExpContext {
	var p = new(ConstExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = sysyParserRULE_constExp

	return p
}

func (s *ConstExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstExpContext) AddExp() IAddExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddExpContext)
}

func (s *ConstExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case sysyParserVisitor:
		return t.VisitConstExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *sysyParser) ConstExp() (localctx IConstExpContext) {
	localctx = NewConstExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, sysyParserRULE_constExp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.AddExp()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
