// Code generated from AclRules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // AclRules
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

type AclRulesParser struct {
	*antlr.BaseParser
}

var AclRulesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func aclrulesParserInit() {
	staticData := &AclRulesParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'**'", "'{'", "'}'", "'('", "')'", "'$('", "'['", "']'",
		"'.'", "':'", "','", "';'", "'<'", "'<='", "'>='", "'>'", "'=='", "'!='",
		"'&&'", "'||'", "'!'", "'&'", "'|'", "'+'", "'-'", "'*'", "'^'", "'%'",
		"'in'", "'allow'", "'match'", "'if'", "'get'", "'exists'", "'true'",
		"'false'", "'list'", "'create'", "'update'", "'read'", "'write'", "'delete'",
		"'function'", "'return'", "'null'", "'service'", "", "", "", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "CurlyOpen", "CurlyClose", "BracketOpen", "BracketClose",
		"PathVariableBracket", "SquareBracketOpen", "SquareBracketClose", "Dot",
		"Colon", "Comma", "Semicolon", "LessThan", "LessOrEqual", "GreaterOrEqual",
		"GreaterThan", "Equals", "Unequal", "LogicalAnd", "LogicalOr", "LogicalNot",
		"BinaryAnd", "BinaryOr", "ArithmeticAdd", "ArithmeticSub", "ArithmeticMul",
		"ArithmeticExp", "ArithmeticModus", "InOperator", "Allow", "Match",
		"If", "Get", "Exists", "True", "False", "List", "Create", "Update",
		"Read", "Write", "Delete", "Function", "Return", "Null", "Service",
		"Number", "String", "Identifier", "Slash", "WS", "Comment",
	}
	staticData.RuleNames = []string{
		"service", "namespaceIdentifier", "namespaceBlock", "matchBlock", "allowKey",
		"comment", "matcher", "allow", "getPathVariable", "pathVariable", "arg",
		"arguments", "memberArg", "memberArguments", "argDeclaration", "argDeclarations",
		"functionDeclaration", "fieldReference", "id", "expression", "objectReference",
		"getPathExpressionVariable", "getPath", "ruleFunctionCall", "functionCall",
		"matchPath",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 53, 280, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 60, 8, 1, 10, 1, 12, 1, 63, 9,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 69, 8, 2, 10, 2, 12, 2, 72, 9, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 81, 8, 3, 10, 3, 12, 3, 84, 9,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 5, 7, 100, 8, 7, 10, 7, 12, 7, 103, 9, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 108, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 118, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 3, 11, 125, 8, 11, 1, 11,
		1, 11, 5, 11, 129, 8, 11, 10, 11, 12, 11, 132, 9, 11, 1, 12, 1, 12, 1,
		13, 3, 13, 137, 8, 13, 1, 13, 1, 13, 5, 13, 141, 8, 13, 10, 13, 12, 13,
		144, 9, 13, 1, 14, 1, 14, 1, 15, 3, 15, 149, 8, 15, 1, 15, 1, 15, 5, 15,
		153, 8, 15, 10, 15, 12, 15, 156, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 3, 17, 175, 8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 3, 19, 183, 8, 19, 1, 19, 1, 19, 5, 19, 187, 8, 19, 10, 19, 12, 19,
		190, 9, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 205, 8, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 237, 8, 19, 1, 19, 1,
		19, 5, 19, 241, 8, 19, 10, 19, 12, 19, 244, 9, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 3, 22, 255, 8, 22, 4, 22, 257,
		8, 22, 11, 22, 12, 22, 258, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 3, 25, 274, 8, 25, 4, 25,
		276, 8, 25, 11, 25, 12, 25, 277, 1, 25, 0, 1, 38, 26, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 0, 9, 2, 0, 34, 34, 38, 43, 3, 0, 31, 33, 35, 47, 50, 50, 2, 0,
		22, 22, 26, 26, 1, 0, 36, 37, 1, 0, 14, 19, 1, 0, 20, 21, 1, 0, 23, 24,
		2, 0, 25, 29, 51, 51, 1, 0, 34, 35, 295, 0, 52, 1, 0, 0, 0, 2, 56, 1, 0,
		0, 0, 4, 64, 1, 0, 0, 0, 6, 75, 1, 0, 0, 0, 8, 87, 1, 0, 0, 0, 10, 89,
		1, 0, 0, 0, 12, 91, 1, 0, 0, 0, 14, 95, 1, 0, 0, 0, 16, 111, 1, 0, 0, 0,
		18, 113, 1, 0, 0, 0, 20, 121, 1, 0, 0, 0, 22, 124, 1, 0, 0, 0, 24, 133,
		1, 0, 0, 0, 26, 136, 1, 0, 0, 0, 28, 145, 1, 0, 0, 0, 30, 148, 1, 0, 0,
		0, 32, 157, 1, 0, 0, 0, 34, 174, 1, 0, 0, 0, 36, 176, 1, 0, 0, 0, 38, 204,
		1, 0, 0, 0, 40, 245, 1, 0, 0, 0, 42, 247, 1, 0, 0, 0, 44, 256, 1, 0, 0,
		0, 46, 260, 1, 0, 0, 0, 48, 265, 1, 0, 0, 0, 50, 275, 1, 0, 0, 0, 52, 53,
		5, 47, 0, 0, 53, 54, 3, 2, 1, 0, 54, 55, 3, 4, 2, 0, 55, 1, 1, 0, 0, 0,
		56, 61, 3, 36, 18, 0, 57, 58, 5, 10, 0, 0, 58, 60, 3, 36, 18, 0, 59, 57,
		1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0,
		62, 3, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 70, 5, 3, 0, 0, 65, 69, 3, 12,
		6, 0, 66, 69, 3, 10, 5, 0, 67, 69, 3, 32, 16, 0, 68, 65, 1, 0, 0, 0, 68,
		66, 1, 0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0,
		0, 70, 71, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 74,
		5, 4, 0, 0, 74, 5, 1, 0, 0, 0, 75, 82, 5, 3, 0, 0, 76, 81, 3, 14, 7, 0,
		77, 81, 3, 12, 6, 0, 78, 81, 5, 53, 0, 0, 79, 81, 3, 32, 16, 0, 80, 76,
		1, 0, 0, 0, 80, 77, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0,
		81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 85, 1,
		0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 4, 0, 0, 86, 7, 1, 0, 0, 0, 87,
		88, 7, 0, 0, 0, 88, 9, 1, 0, 0, 0, 89, 90, 5, 53, 0, 0, 90, 11, 1, 0, 0,
		0, 91, 92, 5, 32, 0, 0, 92, 93, 3, 50, 25, 0, 93, 94, 3, 6, 3, 0, 94, 13,
		1, 0, 0, 0, 95, 96, 5, 31, 0, 0, 96, 101, 3, 8, 4, 0, 97, 98, 5, 12, 0,
		0, 98, 100, 3, 8, 4, 0, 99, 97, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101,
		99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 107, 1, 0, 0, 0, 103, 101, 1,
		0, 0, 0, 104, 105, 5, 11, 0, 0, 105, 106, 5, 33, 0, 0, 106, 108, 3, 38,
		19, 0, 107, 104, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0,
		109, 110, 5, 13, 0, 0, 110, 15, 1, 0, 0, 0, 111, 112, 5, 50, 0, 0, 112,
		17, 1, 0, 0, 0, 113, 114, 5, 3, 0, 0, 114, 117, 5, 50, 0, 0, 115, 116,
		5, 1, 0, 0, 116, 118, 5, 2, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0,
		0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 5, 4, 0, 0, 120, 19, 1, 0, 0, 0,
		121, 122, 3, 38, 19, 0, 122, 21, 1, 0, 0, 0, 123, 125, 3, 20, 10, 0, 124,
		123, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 130, 1, 0, 0, 0, 126, 127,
		5, 12, 0, 0, 127, 129, 3, 20, 10, 0, 128, 126, 1, 0, 0, 0, 129, 132, 1,
		0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 23, 1, 0, 0,
		0, 132, 130, 1, 0, 0, 0, 133, 134, 3, 38, 19, 0, 134, 25, 1, 0, 0, 0, 135,
		137, 3, 24, 12, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 142,
		1, 0, 0, 0, 138, 139, 5, 12, 0, 0, 139, 141, 3, 24, 12, 0, 140, 138, 1,
		0, 0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0,
		0, 143, 27, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 145, 146, 5, 50, 0, 0, 146,
		29, 1, 0, 0, 0, 147, 149, 3, 28, 14, 0, 148, 147, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 154, 1, 0, 0, 0, 150, 151, 5, 12, 0, 0, 151, 153, 3, 28,
		14, 0, 152, 150, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0,
		154, 155, 1, 0, 0, 0, 155, 31, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158,
		5, 44, 0, 0, 158, 159, 5, 50, 0, 0, 159, 160, 5, 5, 0, 0, 160, 161, 3,
		30, 15, 0, 161, 162, 5, 6, 0, 0, 162, 163, 5, 3, 0, 0, 163, 164, 5, 45,
		0, 0, 164, 165, 3, 38, 19, 0, 165, 166, 5, 13, 0, 0, 166, 167, 5, 4, 0,
		0, 167, 33, 1, 0, 0, 0, 168, 169, 5, 10, 0, 0, 169, 175, 3, 36, 18, 0,
		170, 171, 5, 8, 0, 0, 171, 172, 3, 38, 19, 0, 172, 173, 5, 9, 0, 0, 173,
		175, 1, 0, 0, 0, 174, 168, 1, 0, 0, 0, 174, 170, 1, 0, 0, 0, 175, 35, 1,
		0, 0, 0, 176, 177, 7, 1, 0, 0, 177, 37, 1, 0, 0, 0, 178, 179, 6, 19, -1,
		0, 179, 205, 5, 46, 0, 0, 180, 182, 5, 8, 0, 0, 181, 183, 3, 38, 19, 0,
		182, 181, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 188, 1, 0, 0, 0, 184,
		185, 5, 12, 0, 0, 185, 187, 3, 38, 19, 0, 186, 184, 1, 0, 0, 0, 187, 190,
		1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0,
		0, 0, 190, 188, 1, 0, 0, 0, 191, 205, 5, 9, 0, 0, 192, 193, 7, 2, 0, 0,
		193, 205, 3, 38, 19, 8, 194, 205, 5, 49, 0, 0, 195, 205, 5, 48, 0, 0, 196,
		205, 7, 3, 0, 0, 197, 205, 5, 50, 0, 0, 198, 205, 3, 46, 23, 0, 199, 205,
		3, 48, 24, 0, 200, 201, 5, 5, 0, 0, 201, 202, 3, 38, 19, 0, 202, 203, 5,
		6, 0, 0, 203, 205, 1, 0, 0, 0, 204, 178, 1, 0, 0, 0, 204, 180, 1, 0, 0,
		0, 204, 192, 1, 0, 0, 0, 204, 194, 1, 0, 0, 0, 204, 195, 1, 0, 0, 0, 204,
		196, 1, 0, 0, 0, 204, 197, 1, 0, 0, 0, 204, 198, 1, 0, 0, 0, 204, 199,
		1, 0, 0, 0, 204, 200, 1, 0, 0, 0, 205, 242, 1, 0, 0, 0, 206, 207, 10, 17,
		0, 0, 207, 208, 7, 4, 0, 0, 208, 241, 3, 38, 19, 18, 209, 210, 10, 16,
		0, 0, 210, 211, 7, 5, 0, 0, 211, 241, 3, 38, 19, 17, 212, 213, 10, 15,
		0, 0, 213, 214, 7, 6, 0, 0, 214, 241, 3, 38, 19, 16, 215, 216, 10, 14,
		0, 0, 216, 217, 7, 7, 0, 0, 217, 241, 3, 38, 19, 15, 218, 219, 10, 13,
		0, 0, 219, 220, 5, 30, 0, 0, 220, 241, 3, 38, 19, 14, 221, 222, 10, 12,
		0, 0, 222, 223, 5, 10, 0, 0, 223, 241, 3, 36, 18, 0, 224, 225, 10, 11,
		0, 0, 225, 226, 5, 10, 0, 0, 226, 227, 3, 36, 18, 0, 227, 228, 5, 5, 0,
		0, 228, 229, 3, 26, 13, 0, 229, 230, 5, 6, 0, 0, 230, 241, 1, 0, 0, 0,
		231, 232, 10, 10, 0, 0, 232, 233, 5, 8, 0, 0, 233, 236, 3, 38, 19, 0, 234,
		235, 5, 11, 0, 0, 235, 237, 3, 38, 19, 0, 236, 234, 1, 0, 0, 0, 236, 237,
		1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 5, 9, 0, 0, 239, 241, 1, 0,
		0, 0, 240, 206, 1, 0, 0, 0, 240, 209, 1, 0, 0, 0, 240, 212, 1, 0, 0, 0,
		240, 215, 1, 0, 0, 0, 240, 218, 1, 0, 0, 0, 240, 221, 1, 0, 0, 0, 240,
		224, 1, 0, 0, 0, 240, 231, 1, 0, 0, 0, 241, 244, 1, 0, 0, 0, 242, 240,
		1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 39, 1, 0, 0, 0, 244, 242, 1, 0,
		0, 0, 245, 246, 5, 50, 0, 0, 246, 41, 1, 0, 0, 0, 247, 248, 5, 7, 0, 0,
		248, 249, 3, 38, 19, 0, 249, 250, 5, 6, 0, 0, 250, 43, 1, 0, 0, 0, 251,
		254, 5, 51, 0, 0, 252, 255, 3, 16, 8, 0, 253, 255, 3, 42, 21, 0, 254, 252,
		1, 0, 0, 0, 254, 253, 1, 0, 0, 0, 255, 257, 1, 0, 0, 0, 256, 251, 1, 0,
		0, 0, 257, 258, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0,
		259, 45, 1, 0, 0, 0, 260, 261, 7, 8, 0, 0, 261, 262, 5, 5, 0, 0, 262, 263,
		3, 44, 22, 0, 263, 264, 5, 6, 0, 0, 264, 47, 1, 0, 0, 0, 265, 266, 5, 50,
		0, 0, 266, 267, 5, 5, 0, 0, 267, 268, 3, 22, 11, 0, 268, 269, 5, 6, 0,
		0, 269, 49, 1, 0, 0, 0, 270, 273, 5, 51, 0, 0, 271, 274, 5, 50, 0, 0, 272,
		274, 3, 18, 9, 0, 273, 271, 1, 0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 276,
		1, 0, 0, 0, 275, 270, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 275, 1, 0,
		0, 0, 277, 278, 1, 0, 0, 0, 278, 51, 1, 0, 0, 0, 25, 61, 68, 70, 80, 82,
		101, 107, 117, 124, 130, 136, 142, 148, 154, 174, 182, 188, 204, 236, 240,
		242, 254, 258, 273, 277,
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

// AclRulesParserInit initializes any static state used to implement AclRulesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAclRulesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AclRulesParserInit() {
	staticData := &AclRulesParserStaticData
	staticData.once.Do(aclrulesParserInit)
}

// NewAclRulesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAclRulesParser(input antlr.TokenStream) *AclRulesParser {
	AclRulesParserInit()
	this := new(AclRulesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AclRulesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "AclRules.g4"

	return this
}

// AclRulesParser tokens.
const (
	AclRulesParserEOF                 = antlr.TokenEOF
	AclRulesParserT__0                = 1
	AclRulesParserT__1                = 2
	AclRulesParserCurlyOpen           = 3
	AclRulesParserCurlyClose          = 4
	AclRulesParserBracketOpen         = 5
	AclRulesParserBracketClose        = 6
	AclRulesParserPathVariableBracket = 7
	AclRulesParserSquareBracketOpen   = 8
	AclRulesParserSquareBracketClose  = 9
	AclRulesParserDot                 = 10
	AclRulesParserColon               = 11
	AclRulesParserComma               = 12
	AclRulesParserSemicolon           = 13
	AclRulesParserLessThan            = 14
	AclRulesParserLessOrEqual         = 15
	AclRulesParserGreaterOrEqual      = 16
	AclRulesParserGreaterThan         = 17
	AclRulesParserEquals              = 18
	AclRulesParserUnequal             = 19
	AclRulesParserLogicalAnd          = 20
	AclRulesParserLogicalOr           = 21
	AclRulesParserLogicalNot          = 22
	AclRulesParserBinaryAnd           = 23
	AclRulesParserBinaryOr            = 24
	AclRulesParserArithmeticAdd       = 25
	AclRulesParserArithmeticSub       = 26
	AclRulesParserArithmeticMul       = 27
	AclRulesParserArithmeticExp       = 28
	AclRulesParserArithmeticModus     = 29
	AclRulesParserInOperator          = 30
	AclRulesParserAllow               = 31
	AclRulesParserMatch               = 32
	AclRulesParserIf                  = 33
	AclRulesParserGet                 = 34
	AclRulesParserExists              = 35
	AclRulesParserTrue                = 36
	AclRulesParserFalse               = 37
	AclRulesParserList                = 38
	AclRulesParserCreate              = 39
	AclRulesParserUpdate              = 40
	AclRulesParserRead                = 41
	AclRulesParserWrite               = 42
	AclRulesParserDelete              = 43
	AclRulesParserFunction            = 44
	AclRulesParserReturn              = 45
	AclRulesParserNull                = 46
	AclRulesParserService             = 47
	AclRulesParserNumber              = 48
	AclRulesParserString_             = 49
	AclRulesParserIdentifier          = 50
	AclRulesParserSlash               = 51
	AclRulesParserWS                  = 52
	AclRulesParserComment             = 53
)

// AclRulesParser rules.
const (
	AclRulesParserRULE_service                   = 0
	AclRulesParserRULE_namespaceIdentifier       = 1
	AclRulesParserRULE_namespaceBlock            = 2
	AclRulesParserRULE_matchBlock                = 3
	AclRulesParserRULE_allowKey                  = 4
	AclRulesParserRULE_comment                   = 5
	AclRulesParserRULE_matcher                   = 6
	AclRulesParserRULE_allow                     = 7
	AclRulesParserRULE_getPathVariable           = 8
	AclRulesParserRULE_pathVariable              = 9
	AclRulesParserRULE_arg                       = 10
	AclRulesParserRULE_arguments                 = 11
	AclRulesParserRULE_memberArg                 = 12
	AclRulesParserRULE_memberArguments           = 13
	AclRulesParserRULE_argDeclaration            = 14
	AclRulesParserRULE_argDeclarations           = 15
	AclRulesParserRULE_functionDeclaration       = 16
	AclRulesParserRULE_fieldReference            = 17
	AclRulesParserRULE_id                        = 18
	AclRulesParserRULE_expression                = 19
	AclRulesParserRULE_objectReference           = 20
	AclRulesParserRULE_getPathExpressionVariable = 21
	AclRulesParserRULE_getPath                   = 22
	AclRulesParserRULE_ruleFunctionCall          = 23
	AclRulesParserRULE_functionCall              = 24
	AclRulesParserRULE_matchPath                 = 25
)

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Service() antlr.TerminalNode
	NamespaceIdentifier() INamespaceIdentifierContext
	NamespaceBlock() INamespaceBlockContext

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_service
	return p
}

func InitEmptyServiceContext(p *ServiceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_service
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) Service() antlr.TerminalNode {
	return s.GetToken(AclRulesParserService, 0)
}

func (s *ServiceContext) NamespaceIdentifier() INamespaceIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceIdentifierContext)
}

func (s *ServiceContext) NamespaceBlock() INamespaceBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceBlockContext)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitService(s)
	}
}

func (p *AclRulesParser) Service() (localctx IServiceContext) {
	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AclRulesParserRULE_service)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(AclRulesParserService)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.NamespaceIdentifier()
	}
	{
		p.SetState(54)
		p.NamespaceBlock()
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

// INamespaceIdentifierContext is an interface to support dynamic dispatch.
type INamespaceIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllId() []IIdContext
	Id(i int) IIdContext
	AllDot() []antlr.TerminalNode
	Dot(i int) antlr.TerminalNode

	// IsNamespaceIdentifierContext differentiates from other interfaces.
	IsNamespaceIdentifierContext()
}

type NamespaceIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceIdentifierContext() *NamespaceIdentifierContext {
	var p = new(NamespaceIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_namespaceIdentifier
	return p
}

func InitEmptyNamespaceIdentifierContext(p *NamespaceIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_namespaceIdentifier
}

func (*NamespaceIdentifierContext) IsNamespaceIdentifierContext() {}

func NewNamespaceIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceIdentifierContext {
	var p = new(NamespaceIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_namespaceIdentifier

	return p
}

func (s *NamespaceIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceIdentifierContext) AllId() []IIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdContext); ok {
			len++
		}
	}

	tst := make([]IIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdContext); ok {
			tst[i] = t.(IIdContext)
			i++
		}
	}

	return tst
}

func (s *NamespaceIdentifierContext) Id(i int) IIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
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

	return t.(IIdContext)
}

func (s *NamespaceIdentifierContext) AllDot() []antlr.TerminalNode {
	return s.GetTokens(AclRulesParserDot)
}

func (s *NamespaceIdentifierContext) Dot(i int) antlr.TerminalNode {
	return s.GetToken(AclRulesParserDot, i)
}

func (s *NamespaceIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterNamespaceIdentifier(s)
	}
}

func (s *NamespaceIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitNamespaceIdentifier(s)
	}
}

func (p *AclRulesParser) NamespaceIdentifier() (localctx INamespaceIdentifierContext) {
	localctx = NewNamespaceIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AclRulesParserRULE_namespaceIdentifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Id()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AclRulesParserDot {
		{
			p.SetState(57)
			p.Match(AclRulesParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Id()
		}

		p.SetState(63)
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

// INamespaceBlockContext is an interface to support dynamic dispatch.
type INamespaceBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CurlyOpen() antlr.TerminalNode
	CurlyClose() antlr.TerminalNode
	AllMatcher() []IMatcherContext
	Matcher(i int) IMatcherContext
	AllComment() []ICommentContext
	Comment(i int) ICommentContext
	AllFunctionDeclaration() []IFunctionDeclarationContext
	FunctionDeclaration(i int) IFunctionDeclarationContext

	// IsNamespaceBlockContext differentiates from other interfaces.
	IsNamespaceBlockContext()
}

type NamespaceBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceBlockContext() *NamespaceBlockContext {
	var p = new(NamespaceBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_namespaceBlock
	return p
}

func InitEmptyNamespaceBlockContext(p *NamespaceBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_namespaceBlock
}

func (*NamespaceBlockContext) IsNamespaceBlockContext() {}

func NewNamespaceBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceBlockContext {
	var p = new(NamespaceBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_namespaceBlock

	return p
}

func (s *NamespaceBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceBlockContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserCurlyOpen, 0)
}

func (s *NamespaceBlockContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserCurlyClose, 0)
}

func (s *NamespaceBlockContext) AllMatcher() []IMatcherContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatcherContext); ok {
			len++
		}
	}

	tst := make([]IMatcherContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatcherContext); ok {
			tst[i] = t.(IMatcherContext)
			i++
		}
	}

	return tst
}

func (s *NamespaceBlockContext) Matcher(i int) IMatcherContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatcherContext); ok {
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

	return t.(IMatcherContext)
}

func (s *NamespaceBlockContext) AllComment() []ICommentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommentContext); ok {
			len++
		}
	}

	tst := make([]ICommentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommentContext); ok {
			tst[i] = t.(ICommentContext)
			i++
		}
	}

	return tst
}

func (s *NamespaceBlockContext) Comment(i int) ICommentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
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

	return t.(ICommentContext)
}

func (s *NamespaceBlockContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDeclarationContext); ok {
			tst[i] = t.(IFunctionDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *NamespaceBlockContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
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

	return t.(IFunctionDeclarationContext)
}

func (s *NamespaceBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterNamespaceBlock(s)
	}
}

func (s *NamespaceBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitNamespaceBlock(s)
	}
}

func (p *AclRulesParser) NamespaceBlock() (localctx INamespaceBlockContext) {
	localctx = NewNamespaceBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AclRulesParserRULE_namespaceBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(AclRulesParserCurlyOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9024795735752704) != 0 {
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AclRulesParserMatch:
			{
				p.SetState(65)
				p.Matcher()
			}

		case AclRulesParserComment:
			{
				p.SetState(66)
				p.Comment()
			}

		case AclRulesParserFunction:
			{
				p.SetState(67)
				p.FunctionDeclaration()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(73)
		p.Match(AclRulesParserCurlyClose)
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

// IMatchBlockContext is an interface to support dynamic dispatch.
type IMatchBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CurlyOpen() antlr.TerminalNode
	CurlyClose() antlr.TerminalNode
	AllAllow() []IAllowContext
	Allow(i int) IAllowContext
	AllMatcher() []IMatcherContext
	Matcher(i int) IMatcherContext
	AllComment() []antlr.TerminalNode
	Comment(i int) antlr.TerminalNode
	AllFunctionDeclaration() []IFunctionDeclarationContext
	FunctionDeclaration(i int) IFunctionDeclarationContext

	// IsMatchBlockContext differentiates from other interfaces.
	IsMatchBlockContext()
}

type MatchBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchBlockContext() *MatchBlockContext {
	var p = new(MatchBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_matchBlock
	return p
}

func InitEmptyMatchBlockContext(p *MatchBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_matchBlock
}

func (*MatchBlockContext) IsMatchBlockContext() {}

func NewMatchBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchBlockContext {
	var p = new(MatchBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_matchBlock

	return p
}

func (s *MatchBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchBlockContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserCurlyOpen, 0)
}

func (s *MatchBlockContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserCurlyClose, 0)
}

func (s *MatchBlockContext) AllAllow() []IAllowContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAllowContext); ok {
			len++
		}
	}

	tst := make([]IAllowContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAllowContext); ok {
			tst[i] = t.(IAllowContext)
			i++
		}
	}

	return tst
}

func (s *MatchBlockContext) Allow(i int) IAllowContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllowContext); ok {
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

	return t.(IAllowContext)
}

func (s *MatchBlockContext) AllMatcher() []IMatcherContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatcherContext); ok {
			len++
		}
	}

	tst := make([]IMatcherContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatcherContext); ok {
			tst[i] = t.(IMatcherContext)
			i++
		}
	}

	return tst
}

func (s *MatchBlockContext) Matcher(i int) IMatcherContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatcherContext); ok {
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

	return t.(IMatcherContext)
}

func (s *MatchBlockContext) AllComment() []antlr.TerminalNode {
	return s.GetTokens(AclRulesParserComment)
}

func (s *MatchBlockContext) Comment(i int) antlr.TerminalNode {
	return s.GetToken(AclRulesParserComment, i)
}

func (s *MatchBlockContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDeclarationContext); ok {
			tst[i] = t.(IFunctionDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *MatchBlockContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
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

	return t.(IFunctionDeclarationContext)
}

func (s *MatchBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterMatchBlock(s)
	}
}

func (s *MatchBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitMatchBlock(s)
	}
}

func (p *AclRulesParser) MatchBlock() (localctx IMatchBlockContext) {
	localctx = NewMatchBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AclRulesParserRULE_matchBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(AclRulesParserCurlyOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9024797883236352) != 0 {
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AclRulesParserAllow:
			{
				p.SetState(76)
				p.Allow()
			}

		case AclRulesParserMatch:
			{
				p.SetState(77)
				p.Matcher()
			}

		case AclRulesParserComment:
			{
				p.SetState(78)
				p.Match(AclRulesParserComment)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case AclRulesParserFunction:
			{
				p.SetState(79)
				p.FunctionDeclaration()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(85)
		p.Match(AclRulesParserCurlyClose)
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

// IAllowKeyContext is an interface to support dynamic dispatch.
type IAllowKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Read() antlr.TerminalNode
	Write() antlr.TerminalNode
	Update() antlr.TerminalNode
	Delete() antlr.TerminalNode
	Create() antlr.TerminalNode
	List() antlr.TerminalNode
	Get() antlr.TerminalNode

	// IsAllowKeyContext differentiates from other interfaces.
	IsAllowKeyContext()
}

type AllowKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllowKeyContext() *AllowKeyContext {
	var p = new(AllowKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_allowKey
	return p
}

func InitEmptyAllowKeyContext(p *AllowKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_allowKey
}

func (*AllowKeyContext) IsAllowKeyContext() {}

func NewAllowKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllowKeyContext {
	var p = new(AllowKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_allowKey

	return p
}

func (s *AllowKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *AllowKeyContext) Read() antlr.TerminalNode {
	return s.GetToken(AclRulesParserRead, 0)
}

func (s *AllowKeyContext) Write() antlr.TerminalNode {
	return s.GetToken(AclRulesParserWrite, 0)
}

func (s *AllowKeyContext) Update() antlr.TerminalNode {
	return s.GetToken(AclRulesParserUpdate, 0)
}

func (s *AllowKeyContext) Delete() antlr.TerminalNode {
	return s.GetToken(AclRulesParserDelete, 0)
}

func (s *AllowKeyContext) Create() antlr.TerminalNode {
	return s.GetToken(AclRulesParserCreate, 0)
}

func (s *AllowKeyContext) List() antlr.TerminalNode {
	return s.GetToken(AclRulesParserList, 0)
}

func (s *AllowKeyContext) Get() antlr.TerminalNode {
	return s.GetToken(AclRulesParserGet, 0)
}

func (s *AllowKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllowKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllowKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterAllowKey(s)
	}
}

func (s *AllowKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitAllowKey(s)
	}
}

func (p *AclRulesParser) AllowKey() (localctx IAllowKeyContext) {
	localctx = NewAllowKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AclRulesParserRULE_allowKey)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17334488006656) != 0) {
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

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Comment() antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) Comment() antlr.TerminalNode {
	return s.GetToken(AclRulesParserComment, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *AclRulesParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AclRulesParserRULE_comment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(AclRulesParserComment)
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

// IMatcherContext is an interface to support dynamic dispatch.
type IMatcherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Match() antlr.TerminalNode
	MatchPath() IMatchPathContext
	MatchBlock() IMatchBlockContext

	// IsMatcherContext differentiates from other interfaces.
	IsMatcherContext()
}

type MatcherContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatcherContext() *MatcherContext {
	var p = new(MatcherContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_matcher
	return p
}

func InitEmptyMatcherContext(p *MatcherContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_matcher
}

func (*MatcherContext) IsMatcherContext() {}

func NewMatcherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatcherContext {
	var p = new(MatcherContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_matcher

	return p
}

func (s *MatcherContext) GetParser() antlr.Parser { return s.parser }

func (s *MatcherContext) Match() antlr.TerminalNode {
	return s.GetToken(AclRulesParserMatch, 0)
}

func (s *MatcherContext) MatchPath() IMatchPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchPathContext)
}

func (s *MatcherContext) MatchBlock() IMatchBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchBlockContext)
}

func (s *MatcherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatcherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatcherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterMatcher(s)
	}
}

func (s *MatcherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitMatcher(s)
	}
}

func (p *AclRulesParser) Matcher() (localctx IMatcherContext) {
	localctx = NewMatcherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AclRulesParserRULE_matcher)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(AclRulesParserMatch)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.MatchPath()
	}
	{
		p.SetState(93)
		p.MatchBlock()
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

// IAllowContext is an interface to support dynamic dispatch.
type IAllowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Allow() antlr.TerminalNode
	AllAllowKey() []IAllowKeyContext
	AllowKey(i int) IAllowKeyContext
	Semicolon() antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode
	Colon() antlr.TerminalNode
	If() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAllowContext differentiates from other interfaces.
	IsAllowContext()
}

type AllowContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllowContext() *AllowContext {
	var p = new(AllowContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_allow
	return p
}

func InitEmptyAllowContext(p *AllowContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_allow
}

func (*AllowContext) IsAllowContext() {}

func NewAllowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllowContext {
	var p = new(AllowContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_allow

	return p
}

func (s *AllowContext) GetParser() antlr.Parser { return s.parser }

func (s *AllowContext) Allow() antlr.TerminalNode {
	return s.GetToken(AclRulesParserAllow, 0)
}

func (s *AllowContext) AllAllowKey() []IAllowKeyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAllowKeyContext); ok {
			len++
		}
	}

	tst := make([]IAllowKeyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAllowKeyContext); ok {
			tst[i] = t.(IAllowKeyContext)
			i++
		}
	}

	return tst
}

func (s *AllowContext) AllowKey(i int) IAllowKeyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllowKeyContext); ok {
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

	return t.(IAllowKeyContext)
}

func (s *AllowContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AclRulesParserSemicolon, 0)
}

func (s *AllowContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AclRulesParserComma)
}

func (s *AllowContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AclRulesParserComma, i)
}

func (s *AllowContext) Colon() antlr.TerminalNode {
	return s.GetToken(AclRulesParserColon, 0)
}

func (s *AllowContext) If() antlr.TerminalNode {
	return s.GetToken(AclRulesParserIf, 0)
}

func (s *AllowContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AllowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterAllow(s)
	}
}

func (s *AllowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitAllow(s)
	}
}

func (p *AclRulesParser) Allow() (localctx IAllowContext) {
	localctx = NewAllowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AclRulesParserRULE_allow)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(AclRulesParserAllow)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.AllowKey()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AclRulesParserComma {
		{
			p.SetState(97)
			p.Match(AclRulesParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.AllowKey()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AclRulesParserColon {
		{
			p.SetState(104)
			p.Match(AclRulesParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(AclRulesParserIf)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.expression(0)
		}

	}
	{
		p.SetState(109)
		p.Match(AclRulesParserSemicolon)
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

// IGetPathVariableContext is an interface to support dynamic dispatch.
type IGetPathVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsGetPathVariableContext differentiates from other interfaces.
	IsGetPathVariableContext()
}

type GetPathVariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetPathVariableContext() *GetPathVariableContext {
	var p = new(GetPathVariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_getPathVariable
	return p
}

func InitEmptyGetPathVariableContext(p *GetPathVariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_getPathVariable
}

func (*GetPathVariableContext) IsGetPathVariableContext() {}

func NewGetPathVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetPathVariableContext {
	var p = new(GetPathVariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_getPathVariable

	return p
}

func (s *GetPathVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *GetPathVariableContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AclRulesParserIdentifier, 0)
}

func (s *GetPathVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetPathVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetPathVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterGetPathVariable(s)
	}
}

func (s *GetPathVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitGetPathVariable(s)
	}
}

func (p *AclRulesParser) GetPathVariable() (localctx IGetPathVariableContext) {
	localctx = NewGetPathVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AclRulesParserRULE_getPathVariable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(AclRulesParserIdentifier)
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

// IPathVariableContext is an interface to support dynamic dispatch.
type IPathVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CurlyOpen() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	CurlyClose() antlr.TerminalNode

	// IsPathVariableContext differentiates from other interfaces.
	IsPathVariableContext()
}

type PathVariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathVariableContext() *PathVariableContext {
	var p = new(PathVariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_pathVariable
	return p
}

func InitEmptyPathVariableContext(p *PathVariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_pathVariable
}

func (*PathVariableContext) IsPathVariableContext() {}

func NewPathVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathVariableContext {
	var p = new(PathVariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_pathVariable

	return p
}

func (s *PathVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *PathVariableContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserCurlyOpen, 0)
}

func (s *PathVariableContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AclRulesParserIdentifier, 0)
}

func (s *PathVariableContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserCurlyClose, 0)
}

func (s *PathVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterPathVariable(s)
	}
}

func (s *PathVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitPathVariable(s)
	}
}

func (p *AclRulesParser) PathVariable() (localctx IPathVariableContext) {
	localctx = NewPathVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AclRulesParserRULE_pathVariable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(AclRulesParserCurlyOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Match(AclRulesParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AclRulesParserT__0 {
		{
			p.SetState(115)
			p.Match(AclRulesParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(AclRulesParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(119)
		p.Match(AclRulesParserCurlyClose)
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

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_arg
	return p
}

func InitEmptyArgContext(p *ArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_arg
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *AclRulesParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AclRulesParserRULE_arg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.expression(0)
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

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArg() []IArgContext
	Arg(i int) IArgContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
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

	return t.(IArgContext)
}

func (s *ArgumentsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AclRulesParserComma)
}

func (s *ArgumentsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AclRulesParserComma, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *AclRulesParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AclRulesParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2040951350493472) != 0 {
		{
			p.SetState(123)
			p.Arg()
		}

	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AclRulesParserComma {
		{
			p.SetState(126)
			p.Match(AclRulesParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Arg()
		}

		p.SetState(132)
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

// IMemberArgContext is an interface to support dynamic dispatch.
type IMemberArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsMemberArgContext differentiates from other interfaces.
	IsMemberArgContext()
}

type MemberArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberArgContext() *MemberArgContext {
	var p = new(MemberArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_memberArg
	return p
}

func InitEmptyMemberArgContext(p *MemberArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_memberArg
}

func (*MemberArgContext) IsMemberArgContext() {}

func NewMemberArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberArgContext {
	var p = new(MemberArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_memberArg

	return p
}

func (s *MemberArgContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberArgContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterMemberArg(s)
	}
}

func (s *MemberArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitMemberArg(s)
	}
}

func (p *AclRulesParser) MemberArg() (localctx IMemberArgContext) {
	localctx = NewMemberArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AclRulesParserRULE_memberArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.expression(0)
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

// IMemberArgumentsContext is an interface to support dynamic dispatch.
type IMemberArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMemberArg() []IMemberArgContext
	MemberArg(i int) IMemberArgContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsMemberArgumentsContext differentiates from other interfaces.
	IsMemberArgumentsContext()
}

type MemberArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberArgumentsContext() *MemberArgumentsContext {
	var p = new(MemberArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_memberArguments
	return p
}

func InitEmptyMemberArgumentsContext(p *MemberArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_memberArguments
}

func (*MemberArgumentsContext) IsMemberArgumentsContext() {}

func NewMemberArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberArgumentsContext {
	var p = new(MemberArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_memberArguments

	return p
}

func (s *MemberArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberArgumentsContext) AllMemberArg() []IMemberArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMemberArgContext); ok {
			len++
		}
	}

	tst := make([]IMemberArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMemberArgContext); ok {
			tst[i] = t.(IMemberArgContext)
			i++
		}
	}

	return tst
}

func (s *MemberArgumentsContext) MemberArg(i int) IMemberArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberArgContext); ok {
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

	return t.(IMemberArgContext)
}

func (s *MemberArgumentsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AclRulesParserComma)
}

func (s *MemberArgumentsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AclRulesParserComma, i)
}

func (s *MemberArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterMemberArguments(s)
	}
}

func (s *MemberArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitMemberArguments(s)
	}
}

func (p *AclRulesParser) MemberArguments() (localctx IMemberArgumentsContext) {
	localctx = NewMemberArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AclRulesParserRULE_memberArguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2040951350493472) != 0 {
		{
			p.SetState(135)
			p.MemberArg()
		}

	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AclRulesParserComma {
		{
			p.SetState(138)
			p.Match(AclRulesParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.MemberArg()
		}

		p.SetState(144)
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

// IArgDeclarationContext is an interface to support dynamic dispatch.
type IArgDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsArgDeclarationContext differentiates from other interfaces.
	IsArgDeclarationContext()
}

type ArgDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgDeclarationContext() *ArgDeclarationContext {
	var p = new(ArgDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_argDeclaration
	return p
}

func InitEmptyArgDeclarationContext(p *ArgDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_argDeclaration
}

func (*ArgDeclarationContext) IsArgDeclarationContext() {}

func NewArgDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgDeclarationContext {
	var p = new(ArgDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_argDeclaration

	return p
}

func (s *ArgDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AclRulesParserIdentifier, 0)
}

func (s *ArgDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterArgDeclaration(s)
	}
}

func (s *ArgDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitArgDeclaration(s)
	}
}

func (p *AclRulesParser) ArgDeclaration() (localctx IArgDeclarationContext) {
	localctx = NewArgDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AclRulesParserRULE_argDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(AclRulesParserIdentifier)
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

// IArgDeclarationsContext is an interface to support dynamic dispatch.
type IArgDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArgDeclaration() []IArgDeclarationContext
	ArgDeclaration(i int) IArgDeclarationContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsArgDeclarationsContext differentiates from other interfaces.
	IsArgDeclarationsContext()
}

type ArgDeclarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgDeclarationsContext() *ArgDeclarationsContext {
	var p = new(ArgDeclarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_argDeclarations
	return p
}

func InitEmptyArgDeclarationsContext(p *ArgDeclarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_argDeclarations
}

func (*ArgDeclarationsContext) IsArgDeclarationsContext() {}

func NewArgDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgDeclarationsContext {
	var p = new(ArgDeclarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_argDeclarations

	return p
}

func (s *ArgDeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgDeclarationsContext) AllArgDeclaration() []IArgDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IArgDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgDeclarationContext); ok {
			tst[i] = t.(IArgDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ArgDeclarationsContext) ArgDeclaration(i int) IArgDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgDeclarationContext); ok {
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

	return t.(IArgDeclarationContext)
}

func (s *ArgDeclarationsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AclRulesParserComma)
}

func (s *ArgDeclarationsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AclRulesParserComma, i)
}

func (s *ArgDeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgDeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgDeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterArgDeclarations(s)
	}
}

func (s *ArgDeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitArgDeclarations(s)
	}
}

func (p *AclRulesParser) ArgDeclarations() (localctx IArgDeclarationsContext) {
	localctx = NewArgDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AclRulesParserRULE_argDeclarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AclRulesParserIdentifier {
		{
			p.SetState(147)
			p.ArgDeclaration()
		}

	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AclRulesParserComma {
		{
			p.SetState(150)
			p.Match(AclRulesParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.ArgDeclaration()
		}

		p.SetState(156)
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

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	BracketOpen() antlr.TerminalNode
	ArgDeclarations() IArgDeclarationsContext
	BracketClose() antlr.TerminalNode
	CurlyOpen() antlr.TerminalNode
	Return() antlr.TerminalNode
	Expression() IExpressionContext
	Semicolon() antlr.TerminalNode
	CurlyClose() antlr.TerminalNode

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) Function() antlr.TerminalNode {
	return s.GetToken(AclRulesParserFunction, 0)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AclRulesParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketOpen, 0)
}

func (s *FunctionDeclarationContext) ArgDeclarations() IArgDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgDeclarationsContext)
}

func (s *FunctionDeclarationContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketClose, 0)
}

func (s *FunctionDeclarationContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserCurlyOpen, 0)
}

func (s *FunctionDeclarationContext) Return() antlr.TerminalNode {
	return s.GetToken(AclRulesParserReturn, 0)
}

func (s *FunctionDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionDeclarationContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(AclRulesParserSemicolon, 0)
}

func (s *FunctionDeclarationContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserCurlyClose, 0)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *AclRulesParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AclRulesParserRULE_functionDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(AclRulesParserFunction)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(AclRulesParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(AclRulesParserBracketOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.ArgDeclarations()
	}
	{
		p.SetState(161)
		p.Match(AclRulesParserBracketClose)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(AclRulesParserCurlyOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(AclRulesParserReturn)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.expression(0)
	}
	{
		p.SetState(165)
		p.Match(AclRulesParserSemicolon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(AclRulesParserCurlyClose)
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

// IFieldReferenceContext is an interface to support dynamic dispatch.
type IFieldReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFieldReferenceContext differentiates from other interfaces.
	IsFieldReferenceContext()
}

type FieldReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldReferenceContext() *FieldReferenceContext {
	var p = new(FieldReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_fieldReference
	return p
}

func InitEmptyFieldReferenceContext(p *FieldReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_fieldReference
}

func (*FieldReferenceContext) IsFieldReferenceContext() {}

func NewFieldReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldReferenceContext {
	var p = new(FieldReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_fieldReference

	return p
}

func (s *FieldReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldReferenceContext) CopyAll(ctx *FieldReferenceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FieldReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldReferenceWithIdentifierContext struct {
	FieldReferenceContext
}

func NewFieldReferenceWithIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldReferenceWithIdentifierContext {
	var p = new(FieldReferenceWithIdentifierContext)

	InitEmptyFieldReferenceContext(&p.FieldReferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldReferenceContext))

	return p
}

func (s *FieldReferenceWithIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldReferenceWithIdentifierContext) Dot() antlr.TerminalNode {
	return s.GetToken(AclRulesParserDot, 0)
}

func (s *FieldReferenceWithIdentifierContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *FieldReferenceWithIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterFieldReferenceWithIdentifier(s)
	}
}

func (s *FieldReferenceWithIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitFieldReferenceWithIdentifier(s)
	}
}

type FieldReferenceWithMemberRefContext struct {
	FieldReferenceContext
}

func NewFieldReferenceWithMemberRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldReferenceWithMemberRefContext {
	var p = new(FieldReferenceWithMemberRefContext)

	InitEmptyFieldReferenceContext(&p.FieldReferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldReferenceContext))

	return p
}

func (s *FieldReferenceWithMemberRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldReferenceWithMemberRefContext) SquareBracketOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserSquareBracketOpen, 0)
}

func (s *FieldReferenceWithMemberRefContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldReferenceWithMemberRefContext) SquareBracketClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserSquareBracketClose, 0)
}

func (s *FieldReferenceWithMemberRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterFieldReferenceWithMemberRef(s)
	}
}

func (s *FieldReferenceWithMemberRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitFieldReferenceWithMemberRef(s)
	}
}

func (p *AclRulesParser) FieldReference() (localctx IFieldReferenceContext) {
	localctx = NewFieldReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AclRulesParserRULE_fieldReference)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AclRulesParserDot:
		localctx = NewFieldReferenceWithIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Match(AclRulesParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Id()
		}

	case AclRulesParserSquareBracketOpen:
		localctx = NewFieldReferenceWithMemberRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.Match(AclRulesParserSquareBracketOpen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.expression(0)
		}
		{
			p.SetState(172)
			p.Match(AclRulesParserSquareBracketClose)
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

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Allow() antlr.TerminalNode
	Match() antlr.TerminalNode
	If() antlr.TerminalNode
	Exists() antlr.TerminalNode
	True() antlr.TerminalNode
	False() antlr.TerminalNode
	List() antlr.TerminalNode
	Create() antlr.TerminalNode
	Update() antlr.TerminalNode
	Read() antlr.TerminalNode
	Write() antlr.TerminalNode
	Delete() antlr.TerminalNode
	Function() antlr.TerminalNode
	Return() antlr.TerminalNode
	Null() antlr.TerminalNode
	Service() antlr.TerminalNode

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_id
	return p
}

func InitEmptyIdContext(p *IdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_id
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AclRulesParserIdentifier, 0)
}

func (s *IdContext) Allow() antlr.TerminalNode {
	return s.GetToken(AclRulesParserAllow, 0)
}

func (s *IdContext) Match() antlr.TerminalNode {
	return s.GetToken(AclRulesParserMatch, 0)
}

func (s *IdContext) If() antlr.TerminalNode {
	return s.GetToken(AclRulesParserIf, 0)
}

func (s *IdContext) Exists() antlr.TerminalNode {
	return s.GetToken(AclRulesParserExists, 0)
}

func (s *IdContext) True() antlr.TerminalNode {
	return s.GetToken(AclRulesParserTrue, 0)
}

func (s *IdContext) False() antlr.TerminalNode {
	return s.GetToken(AclRulesParserFalse, 0)
}

func (s *IdContext) List() antlr.TerminalNode {
	return s.GetToken(AclRulesParserList, 0)
}

func (s *IdContext) Create() antlr.TerminalNode {
	return s.GetToken(AclRulesParserCreate, 0)
}

func (s *IdContext) Update() antlr.TerminalNode {
	return s.GetToken(AclRulesParserUpdate, 0)
}

func (s *IdContext) Read() antlr.TerminalNode {
	return s.GetToken(AclRulesParserRead, 0)
}

func (s *IdContext) Write() antlr.TerminalNode {
	return s.GetToken(AclRulesParserWrite, 0)
}

func (s *IdContext) Delete() antlr.TerminalNode {
	return s.GetToken(AclRulesParserDelete, 0)
}

func (s *IdContext) Function() antlr.TerminalNode {
	return s.GetToken(AclRulesParserFunction, 0)
}

func (s *IdContext) Return() antlr.TerminalNode {
	return s.GetToken(AclRulesParserReturn, 0)
}

func (s *IdContext) Null() antlr.TerminalNode {
	return s.GetToken(AclRulesParserNull, 0)
}

func (s *IdContext) Service() antlr.TerminalNode {
	return s.GetToken(AclRulesParserService, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitId(s)
	}
}

func (p *AclRulesParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AclRulesParserRULE_id)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1407355556200448) != 0) {
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayExpressionContext struct {
	ExpressionContext
}

func NewArrayExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExpressionContext {
	var p = new(ArrayExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExpressionContext) SquareBracketOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserSquareBracketOpen, 0)
}

func (s *ArrayExpressionContext) SquareBracketClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserSquareBracketClose, 0)
}

func (s *ArrayExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArrayExpressionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(AclRulesParserComma)
}

func (s *ArrayExpressionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(AclRulesParserComma, i)
}

func (s *ArrayExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterArrayExpression(s)
	}
}

func (s *ArrayExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitArrayExpression(s)
	}
}

type NumberExpressionContext struct {
	ExpressionContext
}

func NewNumberExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExpressionContext {
	var p = new(NumberExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(AclRulesParserNumber, 0)
}

func (s *NumberExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterNumberExpression(s)
	}
}

func (s *NumberExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitNumberExpression(s)
	}
}

type ObjectReferenceExpressionContext struct {
	ExpressionContext
}

func NewObjectReferenceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectReferenceExpressionContext {
	var p = new(ObjectReferenceExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ObjectReferenceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectReferenceExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AclRulesParserIdentifier, 0)
}

func (s *ObjectReferenceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterObjectReferenceExpression(s)
	}
}

func (s *ObjectReferenceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitObjectReferenceExpression(s)
	}
}

type ParenthesisExpressionContext struct {
	ExpressionContext
}

func NewParenthesisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisExpressionContext {
	var p = new(ParenthesisExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisExpressionContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketOpen, 0)
}

func (s *ParenthesisExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesisExpressionContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketClose, 0)
}

func (s *ParenthesisExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterParenthesisExpression(s)
	}
}

func (s *ParenthesisExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitParenthesisExpression(s)
	}
}

type ArithmeticExpressionContext struct {
	ExpressionContext
}

func NewArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticExpressionContext {
	var p = new(ArithmeticExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArithmeticExpressionContext) ArithmeticAdd() antlr.TerminalNode {
	return s.GetToken(AclRulesParserArithmeticAdd, 0)
}

func (s *ArithmeticExpressionContext) ArithmeticSub() antlr.TerminalNode {
	return s.GetToken(AclRulesParserArithmeticSub, 0)
}

func (s *ArithmeticExpressionContext) ArithmeticMul() antlr.TerminalNode {
	return s.GetToken(AclRulesParserArithmeticMul, 0)
}

func (s *ArithmeticExpressionContext) Slash() antlr.TerminalNode {
	return s.GetToken(AclRulesParserSlash, 0)
}

func (s *ArithmeticExpressionContext) ArithmeticExp() antlr.TerminalNode {
	return s.GetToken(AclRulesParserArithmeticExp, 0)
}

func (s *ArithmeticExpressionContext) ArithmeticModus() antlr.TerminalNode {
	return s.GetToken(AclRulesParserArithmeticModus, 0)
}

func (s *ArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterArithmeticExpression(s)
	}
}

func (s *ArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitArithmeticExpression(s)
	}
}

type MemberReferenceExpressionContext struct {
	ExpressionContext
}

func NewMemberReferenceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberReferenceExpressionContext {
	var p = new(MemberReferenceExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MemberReferenceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberReferenceExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberReferenceExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(AclRulesParserDot, 0)
}

func (s *MemberReferenceExpressionContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *MemberReferenceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterMemberReferenceExpression(s)
	}
}

func (s *MemberReferenceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitMemberReferenceExpression(s)
	}
}

type BooleanExpressionContext struct {
	ExpressionContext
}

func NewBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExpressionContext) True() antlr.TerminalNode {
	return s.GetToken(AclRulesParserTrue, 0)
}

func (s *BooleanExpressionContext) False() antlr.TerminalNode {
	return s.GetToken(AclRulesParserFalse, 0)
}

func (s *BooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterBooleanExpression(s)
	}
}

func (s *BooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitBooleanExpression(s)
	}
}

type FunctionExpressionContext struct {
	ExpressionContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

type CompareExpressionContext struct {
	ExpressionContext
}

func NewCompareExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpressionContext {
	var p = new(CompareExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CompareExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompareExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *CompareExpressionContext) LessThan() antlr.TerminalNode {
	return s.GetToken(AclRulesParserLessThan, 0)
}

func (s *CompareExpressionContext) LessOrEqual() antlr.TerminalNode {
	return s.GetToken(AclRulesParserLessOrEqual, 0)
}

func (s *CompareExpressionContext) Equals() antlr.TerminalNode {
	return s.GetToken(AclRulesParserEquals, 0)
}

func (s *CompareExpressionContext) Unequal() antlr.TerminalNode {
	return s.GetToken(AclRulesParserUnequal, 0)
}

func (s *CompareExpressionContext) GreaterThan() antlr.TerminalNode {
	return s.GetToken(AclRulesParserGreaterThan, 0)
}

func (s *CompareExpressionContext) GreaterOrEqual() antlr.TerminalNode {
	return s.GetToken(AclRulesParserGreaterOrEqual, 0)
}

func (s *CompareExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterCompareExpression(s)
	}
}

func (s *CompareExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitCompareExpression(s)
	}
}

type BinaryExpressionContext struct {
	ExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) BinaryAnd() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBinaryAnd, 0)
}

func (s *BinaryExpressionContext) BinaryOr() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBinaryOr, 0)
}

func (s *BinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitBinaryExpression(s)
	}
}

type LogicalExpressionContext struct {
	ExpressionContext
}

func NewLogicalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *LogicalExpressionContext) LogicalAnd() antlr.TerminalNode {
	return s.GetToken(AclRulesParserLogicalAnd, 0)
}

func (s *LogicalExpressionContext) LogicalOr() antlr.TerminalNode {
	return s.GetToken(AclRulesParserLogicalOr, 0)
}

func (s *LogicalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitLogicalExpression(s)
	}
}

type GetExpressionContext struct {
	ExpressionContext
}

func NewGetExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetExpressionContext {
	var p = new(GetExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *GetExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetExpressionContext) RuleFunctionCall() IRuleFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleFunctionCallContext)
}

func (s *GetExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterGetExpression(s)
	}
}

func (s *GetExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitGetExpression(s)
	}
}

type InExpressionContext struct {
	ExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *InExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *InExpressionContext) InOperator() antlr.TerminalNode {
	return s.GetToken(AclRulesParserInOperator, 0)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitInExpression(s)
	}
}

type StringExpressionContext struct {
	ExpressionContext
}

func NewStringExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExpressionContext {
	var p = new(StringExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StringExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExpressionContext) String_() antlr.TerminalNode {
	return s.GetToken(AclRulesParserString_, 0)
}

func (s *StringExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterStringExpression(s)
	}
}

func (s *StringExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitStringExpression(s)
	}
}

type NullExpressionContext struct {
	ExpressionContext
}

func NewNullExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullExpressionContext {
	var p = new(NullExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NullExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullExpressionContext) Null() antlr.TerminalNode {
	return s.GetToken(AclRulesParserNull, 0)
}

func (s *NullExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterNullExpression(s)
	}
}

func (s *NullExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitNullExpression(s)
	}
}

type RangeExpressionContext struct {
	ExpressionContext
}

func NewRangeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeExpressionContext {
	var p = new(RangeExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RangeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RangeExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *RangeExpressionContext) SquareBracketOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserSquareBracketOpen, 0)
}

func (s *RangeExpressionContext) SquareBracketClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserSquareBracketClose, 0)
}

func (s *RangeExpressionContext) Colon() antlr.TerminalNode {
	return s.GetToken(AclRulesParserColon, 0)
}

func (s *RangeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterRangeExpression(s)
	}
}

func (s *RangeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitRangeExpression(s)
	}
}

type UnaryExpressionContext struct {
	ExpressionContext
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) LogicalNot() antlr.TerminalNode {
	return s.GetToken(AclRulesParserLogicalNot, 0)
}

func (s *UnaryExpressionContext) ArithmeticSub() antlr.TerminalNode {
	return s.GetToken(AclRulesParserArithmeticSub, 0)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

type MemberFunctionExpressionContext struct {
	ExpressionContext
}

func NewMemberFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberFunctionExpressionContext {
	var p = new(MemberFunctionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MemberFunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberFunctionExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberFunctionExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(AclRulesParserDot, 0)
}

func (s *MemberFunctionExpressionContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *MemberFunctionExpressionContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketOpen, 0)
}

func (s *MemberFunctionExpressionContext) MemberArguments() IMemberArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberArgumentsContext)
}

func (s *MemberFunctionExpressionContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketClose, 0)
}

func (s *MemberFunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterMemberFunctionExpression(s)
	}
}

func (s *MemberFunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitMemberFunctionExpression(s)
	}
}

func (p *AclRulesParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *AclRulesParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, AclRulesParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNullExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(179)
			p.Match(AclRulesParserNull)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewArrayExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(180)
			p.Match(AclRulesParserSquareBracketOpen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2040951350493472) != 0 {
			{
				p.SetState(181)
				p.expression(0)
			}

		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == AclRulesParserComma {
			{
				p.SetState(184)
				p.Match(AclRulesParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(185)
				p.expression(0)
			}

			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(191)
			p.Match(AclRulesParserSquareBracketClose)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(192)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AclRulesParserLogicalNot || _la == AclRulesParserArithmeticSub) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(193)
			p.expression(8)
		}

	case 4:
		localctx = NewStringExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(194)
			p.Match(AclRulesParserString_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewNumberExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(195)
			p.Match(AclRulesParserNumber)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(196)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AclRulesParserTrue || _la == AclRulesParserFalse) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 7:
		localctx = NewObjectReferenceExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(197)
			p.Match(AclRulesParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewGetExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(198)
			p.RuleFunctionCall()
		}

	case 9:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(199)
			p.FunctionCall()
		}

	case 10:
		localctx = NewParenthesisExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(200)
			p.Match(AclRulesParserBracketOpen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.expression(0)
		}
		{
			p.SetState(202)
			p.Match(AclRulesParserBracketClose)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(240)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCompareExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AclRulesParserRULE_expression)
				p.SetState(206)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(207)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1032192) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(208)
					p.expression(18)
				}

			case 2:
				localctx = NewLogicalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AclRulesParserRULE_expression)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(210)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AclRulesParserLogicalAnd || _la == AclRulesParserLogicalOr) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(211)
					p.expression(17)
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AclRulesParserRULE_expression)
				p.SetState(212)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(213)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AclRulesParserBinaryAnd || _la == AclRulesParserBinaryOr) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(214)
					p.expression(16)
				}

			case 4:
				localctx = NewArithmeticExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AclRulesParserRULE_expression)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(216)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2251800853872640) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(217)
					p.expression(15)
				}

			case 5:
				localctx = NewInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AclRulesParserRULE_expression)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(219)
					p.Match(AclRulesParserInOperator)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(220)
					p.expression(14)
				}

			case 6:
				localctx = NewMemberReferenceExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AclRulesParserRULE_expression)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(222)
					p.Match(AclRulesParserDot)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(223)
					p.Id()
				}

			case 7:
				localctx = NewMemberFunctionExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AclRulesParserRULE_expression)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(225)
					p.Match(AclRulesParserDot)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(226)
					p.Id()
				}
				{
					p.SetState(227)
					p.Match(AclRulesParserBracketOpen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(228)
					p.MemberArguments()
				}
				{
					p.SetState(229)
					p.Match(AclRulesParserBracketClose)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 8:
				localctx = NewRangeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AclRulesParserRULE_expression)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(232)
					p.Match(AclRulesParserSquareBracketOpen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(233)
					p.expression(0)
				}
				p.SetState(236)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == AclRulesParserColon {
					{
						p.SetState(234)
						p.Match(AclRulesParserColon)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(235)
						p.expression(0)
					}

				}
				{
					p.SetState(238)
					p.Match(AclRulesParserSquareBracketClose)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectReferenceContext is an interface to support dynamic dispatch.
type IObjectReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsObjectReferenceContext differentiates from other interfaces.
	IsObjectReferenceContext()
}

type ObjectReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectReferenceContext() *ObjectReferenceContext {
	var p = new(ObjectReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_objectReference
	return p
}

func InitEmptyObjectReferenceContext(p *ObjectReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_objectReference
}

func (*ObjectReferenceContext) IsObjectReferenceContext() {}

func NewObjectReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectReferenceContext {
	var p = new(ObjectReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_objectReference

	return p
}

func (s *ObjectReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AclRulesParserIdentifier, 0)
}

func (s *ObjectReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterObjectReference(s)
	}
}

func (s *ObjectReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitObjectReference(s)
	}
}

func (p *AclRulesParser) ObjectReference() (localctx IObjectReferenceContext) {
	localctx = NewObjectReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AclRulesParserRULE_objectReference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(AclRulesParserIdentifier)
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

// IGetPathExpressionVariableContext is an interface to support dynamic dispatch.
type IGetPathExpressionVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PathVariableBracket() antlr.TerminalNode
	Expression() IExpressionContext
	BracketClose() antlr.TerminalNode

	// IsGetPathExpressionVariableContext differentiates from other interfaces.
	IsGetPathExpressionVariableContext()
}

type GetPathExpressionVariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetPathExpressionVariableContext() *GetPathExpressionVariableContext {
	var p = new(GetPathExpressionVariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_getPathExpressionVariable
	return p
}

func InitEmptyGetPathExpressionVariableContext(p *GetPathExpressionVariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_getPathExpressionVariable
}

func (*GetPathExpressionVariableContext) IsGetPathExpressionVariableContext() {}

func NewGetPathExpressionVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetPathExpressionVariableContext {
	var p = new(GetPathExpressionVariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_getPathExpressionVariable

	return p
}

func (s *GetPathExpressionVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *GetPathExpressionVariableContext) PathVariableBracket() antlr.TerminalNode {
	return s.GetToken(AclRulesParserPathVariableBracket, 0)
}

func (s *GetPathExpressionVariableContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GetPathExpressionVariableContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketClose, 0)
}

func (s *GetPathExpressionVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetPathExpressionVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetPathExpressionVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterGetPathExpressionVariable(s)
	}
}

func (s *GetPathExpressionVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitGetPathExpressionVariable(s)
	}
}

func (p *AclRulesParser) GetPathExpressionVariable() (localctx IGetPathExpressionVariableContext) {
	localctx = NewGetPathExpressionVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AclRulesParserRULE_getPathExpressionVariable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(AclRulesParserPathVariableBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.expression(0)
	}
	{
		p.SetState(249)
		p.Match(AclRulesParserBracketClose)
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

// IGetPathContext is an interface to support dynamic dispatch.
type IGetPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSlash() []antlr.TerminalNode
	Slash(i int) antlr.TerminalNode
	AllGetPathVariable() []IGetPathVariableContext
	GetPathVariable(i int) IGetPathVariableContext
	AllGetPathExpressionVariable() []IGetPathExpressionVariableContext
	GetPathExpressionVariable(i int) IGetPathExpressionVariableContext

	// IsGetPathContext differentiates from other interfaces.
	IsGetPathContext()
}

type GetPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetPathContext() *GetPathContext {
	var p = new(GetPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_getPath
	return p
}

func InitEmptyGetPathContext(p *GetPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_getPath
}

func (*GetPathContext) IsGetPathContext() {}

func NewGetPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetPathContext {
	var p = new(GetPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_getPath

	return p
}

func (s *GetPathContext) GetParser() antlr.Parser { return s.parser }

func (s *GetPathContext) AllSlash() []antlr.TerminalNode {
	return s.GetTokens(AclRulesParserSlash)
}

func (s *GetPathContext) Slash(i int) antlr.TerminalNode {
	return s.GetToken(AclRulesParserSlash, i)
}

func (s *GetPathContext) AllGetPathVariable() []IGetPathVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGetPathVariableContext); ok {
			len++
		}
	}

	tst := make([]IGetPathVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGetPathVariableContext); ok {
			tst[i] = t.(IGetPathVariableContext)
			i++
		}
	}

	return tst
}

func (s *GetPathContext) GetPathVariable(i int) IGetPathVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetPathVariableContext); ok {
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

	return t.(IGetPathVariableContext)
}

func (s *GetPathContext) AllGetPathExpressionVariable() []IGetPathExpressionVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGetPathExpressionVariableContext); ok {
			len++
		}
	}

	tst := make([]IGetPathExpressionVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGetPathExpressionVariableContext); ok {
			tst[i] = t.(IGetPathExpressionVariableContext)
			i++
		}
	}

	return tst
}

func (s *GetPathContext) GetPathExpressionVariable(i int) IGetPathExpressionVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetPathExpressionVariableContext); ok {
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

	return t.(IGetPathExpressionVariableContext)
}

func (s *GetPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterGetPath(s)
	}
}

func (s *GetPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitGetPath(s)
	}
}

func (p *AclRulesParser) GetPath() (localctx IGetPathContext) {
	localctx = NewGetPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AclRulesParserRULE_getPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AclRulesParserSlash {
		{
			p.SetState(251)
			p.Match(AclRulesParserSlash)
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

		switch p.GetTokenStream().LA(1) {
		case AclRulesParserIdentifier:
			{
				p.SetState(252)
				p.GetPathVariable()
			}

		case AclRulesParserPathVariableBracket:
			{
				p.SetState(253)
				p.GetPathExpressionVariable()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(258)
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

// IRuleFunctionCallContext is an interface to support dynamic dispatch.
type IRuleFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BracketOpen() antlr.TerminalNode
	GetPath() IGetPathContext
	BracketClose() antlr.TerminalNode
	Get() antlr.TerminalNode
	Exists() antlr.TerminalNode

	// IsRuleFunctionCallContext differentiates from other interfaces.
	IsRuleFunctionCallContext()
}

type RuleFunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleFunctionCallContext() *RuleFunctionCallContext {
	var p = new(RuleFunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_ruleFunctionCall
	return p
}

func InitEmptyRuleFunctionCallContext(p *RuleFunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_ruleFunctionCall
}

func (*RuleFunctionCallContext) IsRuleFunctionCallContext() {}

func NewRuleFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleFunctionCallContext {
	var p = new(RuleFunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_ruleFunctionCall

	return p
}

func (s *RuleFunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleFunctionCallContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketOpen, 0)
}

func (s *RuleFunctionCallContext) GetPath() IGetPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGetPathContext)
}

func (s *RuleFunctionCallContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketClose, 0)
}

func (s *RuleFunctionCallContext) Get() antlr.TerminalNode {
	return s.GetToken(AclRulesParserGet, 0)
}

func (s *RuleFunctionCallContext) Exists() antlr.TerminalNode {
	return s.GetToken(AclRulesParserExists, 0)
}

func (s *RuleFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleFunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterRuleFunctionCall(s)
	}
}

func (s *RuleFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitRuleFunctionCall(s)
	}
}

func (p *AclRulesParser) RuleFunctionCall() (localctx IRuleFunctionCallContext) {
	localctx = NewRuleFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AclRulesParserRULE_ruleFunctionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AclRulesParserGet || _la == AclRulesParserExists) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(261)
		p.Match(AclRulesParserBracketOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.GetPath()
	}
	{
		p.SetState(263)
		p.Match(AclRulesParserBracketClose)
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

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	BracketOpen() antlr.TerminalNode
	Arguments() IArgumentsContext
	BracketClose() antlr.TerminalNode

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(AclRulesParserIdentifier, 0)
}

func (s *FunctionCallContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketOpen, 0)
}

func (s *FunctionCallContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FunctionCallContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(AclRulesParserBracketClose, 0)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *AclRulesParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AclRulesParserRULE_functionCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(AclRulesParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.Match(AclRulesParserBracketOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Arguments()
	}
	{
		p.SetState(268)
		p.Match(AclRulesParserBracketClose)
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

// IMatchPathContext is an interface to support dynamic dispatch.
type IMatchPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSlash() []antlr.TerminalNode
	Slash(i int) antlr.TerminalNode
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllPathVariable() []IPathVariableContext
	PathVariable(i int) IPathVariableContext

	// IsMatchPathContext differentiates from other interfaces.
	IsMatchPathContext()
}

type MatchPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchPathContext() *MatchPathContext {
	var p = new(MatchPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_matchPath
	return p
}

func InitEmptyMatchPathContext(p *MatchPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AclRulesParserRULE_matchPath
}

func (*MatchPathContext) IsMatchPathContext() {}

func NewMatchPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchPathContext {
	var p = new(MatchPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AclRulesParserRULE_matchPath

	return p
}

func (s *MatchPathContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchPathContext) AllSlash() []antlr.TerminalNode {
	return s.GetTokens(AclRulesParserSlash)
}

func (s *MatchPathContext) Slash(i int) antlr.TerminalNode {
	return s.GetToken(AclRulesParserSlash, i)
}

func (s *MatchPathContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(AclRulesParserIdentifier)
}

func (s *MatchPathContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(AclRulesParserIdentifier, i)
}

func (s *MatchPathContext) AllPathVariable() []IPathVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPathVariableContext); ok {
			len++
		}
	}

	tst := make([]IPathVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPathVariableContext); ok {
			tst[i] = t.(IPathVariableContext)
			i++
		}
	}

	return tst
}

func (s *MatchPathContext) PathVariable(i int) IPathVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathVariableContext); ok {
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

	return t.(IPathVariableContext)
}

func (s *MatchPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.EnterMatchPath(s)
	}
}

func (s *MatchPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AclRulesListener); ok {
		listenerT.ExitMatchPath(s)
	}
}

func (p *AclRulesParser) MatchPath() (localctx IMatchPathContext) {
	localctx = NewMatchPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, AclRulesParserRULE_matchPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AclRulesParserSlash {
		{
			p.SetState(270)
			p.Match(AclRulesParserSlash)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AclRulesParserIdentifier:
			{
				p.SetState(271)
				p.Match(AclRulesParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case AclRulesParserCurlyOpen:
			{
				p.SetState(272)
				p.PathVariable()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(277)
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

func (p *AclRulesParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 19:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AclRulesParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
