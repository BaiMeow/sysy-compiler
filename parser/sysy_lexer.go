// Code generated from sysyLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type sysyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SysyLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sysylexerLexerInit() {
	staticData := &SysyLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"Const", "Int", "Float", "Void", "Type", "If", "Else", "While", "Break",
		"Continue", "Return", "LeftParen", "RightParen", "LeftBracket", "RightBracket",
		"LeftBrace", "RightBrace", "Comma", "Semicolon", "Assign", "Add", "Sub",
		"Mul", "Div", "Mod", "Less", "Greater", "Equal", "NotEqual", "LessEqual",
		"GreaterEqual", "And", "Or", "Not", "Identifier", "NonDigit", "Digit",
		"IntegerConstant", "DecimalConstant", "NonZeroDigit", "HexadecimalConstant",
		"HexadecimalDigit", "OctalConstant", "OctalDigit", "FloatConstant",
		"WS", "LineComment", "BlockComment", "Preprocessor",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 39, 318, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 3, 4, 123, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 5, 34, 216,
		8, 34, 10, 34, 12, 34, 219, 9, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1,
		37, 1, 37, 3, 37, 228, 8, 37, 1, 38, 1, 38, 5, 38, 232, 8, 38, 10, 38,
		12, 38, 235, 9, 38, 1, 38, 3, 38, 238, 8, 38, 1, 39, 1, 39, 1, 40, 1, 40,
		1, 40, 4, 40, 245, 8, 40, 11, 40, 12, 40, 246, 1, 41, 1, 41, 3, 41, 251,
		8, 41, 1, 42, 1, 42, 5, 42, 255, 8, 42, 10, 42, 12, 42, 258, 9, 42, 1,
		43, 1, 43, 1, 44, 1, 44, 1, 44, 5, 44, 265, 8, 44, 10, 44, 12, 44, 268,
		9, 44, 1, 44, 1, 44, 4, 44, 272, 8, 44, 11, 44, 12, 44, 273, 3, 44, 276,
		8, 44, 1, 45, 4, 45, 279, 8, 45, 11, 45, 12, 45, 280, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 46, 1, 46, 5, 46, 289, 8, 46, 10, 46, 12, 46, 292, 9, 46,
		1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 300, 8, 47, 10, 47, 12,
		47, 303, 9, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 5, 48,
		312, 8, 48, 10, 48, 12, 48, 315, 9, 48, 1, 48, 1, 48, 1, 301, 0, 49, 1,
		1, 3, 0, 5, 0, 7, 2, 9, 3, 11, 4, 13, 5, 15, 6, 17, 7, 19, 8, 21, 9, 23,
		10, 25, 11, 27, 12, 29, 13, 31, 14, 33, 15, 35, 16, 37, 17, 39, 18, 41,
		19, 43, 20, 45, 21, 47, 22, 49, 23, 51, 24, 53, 25, 55, 26, 57, 27, 59,
		28, 61, 29, 63, 30, 65, 31, 67, 32, 69, 33, 71, 0, 73, 0, 75, 34, 77, 0,
		79, 0, 81, 0, 83, 0, 85, 0, 87, 0, 89, 35, 91, 36, 93, 37, 95, 38, 97,
		39, 1, 0, 8, 3, 0, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 49, 57,
		2, 0, 88, 88, 120, 120, 2, 0, 65, 70, 97, 102, 1, 0, 48, 55, 3, 0, 9, 10,
		13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 324, 0, 1, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69,
		1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0,
		93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 1, 99, 1, 0, 0, 0,
		3, 105, 1, 0, 0, 0, 5, 109, 1, 0, 0, 0, 7, 115, 1, 0, 0, 0, 9, 122, 1,
		0, 0, 0, 11, 124, 1, 0, 0, 0, 13, 127, 1, 0, 0, 0, 15, 132, 1, 0, 0, 0,
		17, 138, 1, 0, 0, 0, 19, 144, 1, 0, 0, 0, 21, 153, 1, 0, 0, 0, 23, 160,
		1, 0, 0, 0, 25, 162, 1, 0, 0, 0, 27, 164, 1, 0, 0, 0, 29, 166, 1, 0, 0,
		0, 31, 168, 1, 0, 0, 0, 33, 170, 1, 0, 0, 0, 35, 172, 1, 0, 0, 0, 37, 174,
		1, 0, 0, 0, 39, 176, 1, 0, 0, 0, 41, 178, 1, 0, 0, 0, 43, 180, 1, 0, 0,
		0, 45, 182, 1, 0, 0, 0, 47, 184, 1, 0, 0, 0, 49, 186, 1, 0, 0, 0, 51, 188,
		1, 0, 0, 0, 53, 190, 1, 0, 0, 0, 55, 192, 1, 0, 0, 0, 57, 195, 1, 0, 0,
		0, 59, 198, 1, 0, 0, 0, 61, 201, 1, 0, 0, 0, 63, 204, 1, 0, 0, 0, 65, 207,
		1, 0, 0, 0, 67, 210, 1, 0, 0, 0, 69, 212, 1, 0, 0, 0, 71, 220, 1, 0, 0,
		0, 73, 222, 1, 0, 0, 0, 75, 227, 1, 0, 0, 0, 77, 237, 1, 0, 0, 0, 79, 239,
		1, 0, 0, 0, 81, 241, 1, 0, 0, 0, 83, 250, 1, 0, 0, 0, 85, 252, 1, 0, 0,
		0, 87, 259, 1, 0, 0, 0, 89, 275, 1, 0, 0, 0, 91, 278, 1, 0, 0, 0, 93, 284,
		1, 0, 0, 0, 95, 295, 1, 0, 0, 0, 97, 309, 1, 0, 0, 0, 99, 100, 5, 99, 0,
		0, 100, 101, 5, 111, 0, 0, 101, 102, 5, 110, 0, 0, 102, 103, 5, 115, 0,
		0, 103, 104, 5, 116, 0, 0, 104, 2, 1, 0, 0, 0, 105, 106, 5, 105, 0, 0,
		106, 107, 5, 110, 0, 0, 107, 108, 5, 116, 0, 0, 108, 4, 1, 0, 0, 0, 109,
		110, 5, 102, 0, 0, 110, 111, 5, 108, 0, 0, 111, 112, 5, 111, 0, 0, 112,
		113, 5, 97, 0, 0, 113, 114, 5, 116, 0, 0, 114, 6, 1, 0, 0, 0, 115, 116,
		5, 118, 0, 0, 116, 117, 5, 111, 0, 0, 117, 118, 5, 105, 0, 0, 118, 119,
		5, 100, 0, 0, 119, 8, 1, 0, 0, 0, 120, 123, 3, 3, 1, 0, 121, 123, 3, 5,
		2, 0, 122, 120, 1, 0, 0, 0, 122, 121, 1, 0, 0, 0, 123, 10, 1, 0, 0, 0,
		124, 125, 5, 105, 0, 0, 125, 126, 5, 102, 0, 0, 126, 12, 1, 0, 0, 0, 127,
		128, 5, 101, 0, 0, 128, 129, 5, 108, 0, 0, 129, 130, 5, 115, 0, 0, 130,
		131, 5, 101, 0, 0, 131, 14, 1, 0, 0, 0, 132, 133, 5, 119, 0, 0, 133, 134,
		5, 104, 0, 0, 134, 135, 5, 105, 0, 0, 135, 136, 5, 108, 0, 0, 136, 137,
		5, 101, 0, 0, 137, 16, 1, 0, 0, 0, 138, 139, 5, 98, 0, 0, 139, 140, 5,
		114, 0, 0, 140, 141, 5, 101, 0, 0, 141, 142, 5, 97, 0, 0, 142, 143, 5,
		107, 0, 0, 143, 18, 1, 0, 0, 0, 144, 145, 5, 99, 0, 0, 145, 146, 5, 111,
		0, 0, 146, 147, 5, 110, 0, 0, 147, 148, 5, 116, 0, 0, 148, 149, 5, 105,
		0, 0, 149, 150, 5, 110, 0, 0, 150, 151, 5, 117, 0, 0, 151, 152, 5, 101,
		0, 0, 152, 20, 1, 0, 0, 0, 153, 154, 5, 114, 0, 0, 154, 155, 5, 101, 0,
		0, 155, 156, 5, 116, 0, 0, 156, 157, 5, 117, 0, 0, 157, 158, 5, 114, 0,
		0, 158, 159, 5, 110, 0, 0, 159, 22, 1, 0, 0, 0, 160, 161, 5, 40, 0, 0,
		161, 24, 1, 0, 0, 0, 162, 163, 5, 41, 0, 0, 163, 26, 1, 0, 0, 0, 164, 165,
		5, 91, 0, 0, 165, 28, 1, 0, 0, 0, 166, 167, 5, 93, 0, 0, 167, 30, 1, 0,
		0, 0, 168, 169, 5, 123, 0, 0, 169, 32, 1, 0, 0, 0, 170, 171, 5, 125, 0,
		0, 171, 34, 1, 0, 0, 0, 172, 173, 5, 44, 0, 0, 173, 36, 1, 0, 0, 0, 174,
		175, 5, 59, 0, 0, 175, 38, 1, 0, 0, 0, 176, 177, 5, 61, 0, 0, 177, 40,
		1, 0, 0, 0, 178, 179, 5, 43, 0, 0, 179, 42, 1, 0, 0, 0, 180, 181, 5, 45,
		0, 0, 181, 44, 1, 0, 0, 0, 182, 183, 5, 42, 0, 0, 183, 46, 1, 0, 0, 0,
		184, 185, 5, 47, 0, 0, 185, 48, 1, 0, 0, 0, 186, 187, 5, 37, 0, 0, 187,
		50, 1, 0, 0, 0, 188, 189, 5, 60, 0, 0, 189, 52, 1, 0, 0, 0, 190, 191, 5,
		62, 0, 0, 191, 54, 1, 0, 0, 0, 192, 193, 5, 61, 0, 0, 193, 194, 5, 61,
		0, 0, 194, 56, 1, 0, 0, 0, 195, 196, 5, 33, 0, 0, 196, 197, 5, 61, 0, 0,
		197, 58, 1, 0, 0, 0, 198, 199, 5, 60, 0, 0, 199, 200, 5, 61, 0, 0, 200,
		60, 1, 0, 0, 0, 201, 202, 5, 62, 0, 0, 202, 203, 5, 61, 0, 0, 203, 62,
		1, 0, 0, 0, 204, 205, 5, 38, 0, 0, 205, 206, 5, 38, 0, 0, 206, 64, 1, 0,
		0, 0, 207, 208, 5, 124, 0, 0, 208, 209, 5, 124, 0, 0, 209, 66, 1, 0, 0,
		0, 210, 211, 5, 33, 0, 0, 211, 68, 1, 0, 0, 0, 212, 217, 3, 71, 35, 0,
		213, 216, 3, 71, 35, 0, 214, 216, 3, 73, 36, 0, 215, 213, 1, 0, 0, 0, 215,
		214, 1, 0, 0, 0, 216, 219, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217, 218,
		1, 0, 0, 0, 218, 70, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220, 221, 7, 0,
		0, 0, 221, 72, 1, 0, 0, 0, 222, 223, 7, 1, 0, 0, 223, 74, 1, 0, 0, 0, 224,
		228, 3, 85, 42, 0, 225, 228, 3, 81, 40, 0, 226, 228, 3, 77, 38, 0, 227,
		224, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 226, 1, 0, 0, 0, 228, 76, 1,
		0, 0, 0, 229, 233, 3, 79, 39, 0, 230, 232, 3, 73, 36, 0, 231, 230, 1, 0,
		0, 0, 232, 235, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0,
		234, 238, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236, 238, 5, 48, 0, 0, 237,
		229, 1, 0, 0, 0, 237, 236, 1, 0, 0, 0, 238, 78, 1, 0, 0, 0, 239, 240, 7,
		2, 0, 0, 240, 80, 1, 0, 0, 0, 241, 242, 5, 48, 0, 0, 242, 244, 7, 3, 0,
		0, 243, 245, 3, 83, 41, 0, 244, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0,
		246, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 82, 1, 0, 0, 0, 248, 251,
		3, 73, 36, 0, 249, 251, 7, 4, 0, 0, 250, 248, 1, 0, 0, 0, 250, 249, 1,
		0, 0, 0, 251, 84, 1, 0, 0, 0, 252, 256, 5, 48, 0, 0, 253, 255, 3, 87, 43,
		0, 254, 253, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 256,
		257, 1, 0, 0, 0, 257, 86, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 259, 260, 7,
		5, 0, 0, 260, 88, 1, 0, 0, 0, 261, 262, 3, 77, 38, 0, 262, 266, 5, 46,
		0, 0, 263, 265, 3, 73, 36, 0, 264, 263, 1, 0, 0, 0, 265, 268, 1, 0, 0,
		0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 276, 1, 0, 0, 0, 268,
		266, 1, 0, 0, 0, 269, 271, 5, 46, 0, 0, 270, 272, 3, 73, 36, 0, 271, 270,
		1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0,
		0, 0, 274, 276, 1, 0, 0, 0, 275, 261, 1, 0, 0, 0, 275, 269, 1, 0, 0, 0,
		276, 90, 1, 0, 0, 0, 277, 279, 7, 6, 0, 0, 278, 277, 1, 0, 0, 0, 279, 280,
		1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 282, 1, 0,
		0, 0, 282, 283, 6, 45, 0, 0, 283, 92, 1, 0, 0, 0, 284, 285, 5, 47, 0, 0,
		285, 286, 5, 47, 0, 0, 286, 290, 1, 0, 0, 0, 287, 289, 8, 7, 0, 0, 288,
		287, 1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291,
		1, 0, 0, 0, 291, 293, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 294, 6, 46,
		0, 0, 294, 94, 1, 0, 0, 0, 295, 296, 5, 47, 0, 0, 296, 297, 5, 42, 0, 0,
		297, 301, 1, 0, 0, 0, 298, 300, 9, 0, 0, 0, 299, 298, 1, 0, 0, 0, 300,
		303, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 302, 304,
		1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 305, 5, 42, 0, 0, 305, 306, 5, 47,
		0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 6, 47, 0, 0, 308, 96, 1, 0, 0, 0,
		309, 313, 5, 35, 0, 0, 310, 312, 8, 7, 0, 0, 311, 310, 1, 0, 0, 0, 312,
		315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 316,
		1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 316, 317, 6, 48, 0, 0, 317, 98, 1, 0,
		0, 0, 17, 0, 122, 215, 217, 227, 233, 237, 246, 250, 256, 266, 273, 275,
		280, 290, 301, 313, 1, 6, 0, 0,
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

// sysyLexerInit initializes any static state used to implement sysyLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewsysyLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SysyLexerInit() {
	staticData := &SysyLexerLexerStaticData
	staticData.once.Do(sysylexerLexerInit)
}

// NewsysyLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewsysyLexer(input antlr.CharStream) *sysyLexer {
	SysyLexerInit()
	l := new(sysyLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SysyLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "sysyLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// sysyLexer tokens.
const (
	sysyLexerConst           = 1
	sysyLexerVoid            = 2
	sysyLexerType            = 3
	sysyLexerIf              = 4
	sysyLexerElse            = 5
	sysyLexerWhile           = 6
	sysyLexerBreak           = 7
	sysyLexerContinue        = 8
	sysyLexerReturn          = 9
	sysyLexerLeftParen       = 10
	sysyLexerRightParen      = 11
	sysyLexerLeftBracket     = 12
	sysyLexerRightBracket    = 13
	sysyLexerLeftBrace       = 14
	sysyLexerRightBrace      = 15
	sysyLexerComma           = 16
	sysyLexerSemicolon       = 17
	sysyLexerAssign          = 18
	sysyLexerAdd             = 19
	sysyLexerSub             = 20
	sysyLexerMul             = 21
	sysyLexerDiv             = 22
	sysyLexerMod             = 23
	sysyLexerLess            = 24
	sysyLexerGreater         = 25
	sysyLexerEqual           = 26
	sysyLexerNotEqual        = 27
	sysyLexerLessEqual       = 28
	sysyLexerGreaterEqual    = 29
	sysyLexerAnd             = 30
	sysyLexerOr              = 31
	sysyLexerNot             = 32
	sysyLexerIdentifier      = 33
	sysyLexerIntegerConstant = 34
	sysyLexerFloatConstant   = 35
	sysyLexerWS              = 36
	sysyLexerLineComment     = 37
	sysyLexerBlockComment    = 38
	sysyLexerPreprocessor    = 39
)