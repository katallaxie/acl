// Code generated from AclRules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing

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

type AclRulesLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AclRulesLexerLexerStaticData struct {
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

func aclruleslexerLexerInit() {
	staticData := &AclRulesLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "CurlyOpen", "CurlyClose", "BracketOpen", "BracketClose",
		"PathVariableBracket", "SquareBracketOpen", "SquareBracketClose", "Dot",
		"Colon", "Comma", "Semicolon", "LessThan", "LessOrEqual", "GreaterOrEqual",
		"GreaterThan", "Equals", "Unequal", "LogicalAnd", "LogicalOr", "LogicalNot",
		"BinaryAnd", "BinaryOr", "ArithmeticAdd", "ArithmeticSub", "ArithmeticMul",
		"ArithmeticExp", "ArithmeticModus", "InOperator", "Allow", "Match",
		"If", "Get", "Exists", "True", "False", "List", "Create", "Update",
		"Read", "Write", "Delete", "Function", "Return", "Null", "Service",
		"Number", "String", "Identifier", "Slash", "WS", "Comment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 53, 332, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 1, 46, 1, 47, 4, 47, 281, 8, 47, 11, 47, 12, 47, 282, 1,
		47, 1, 47, 4, 47, 287, 8, 47, 11, 47, 12, 47, 288, 3, 47, 291, 8, 47, 1,
		48, 1, 48, 5, 48, 295, 8, 48, 10, 48, 12, 48, 298, 9, 48, 1, 48, 1, 48,
		1, 48, 5, 48, 303, 8, 48, 10, 48, 12, 48, 306, 9, 48, 1, 48, 3, 48, 309,
		8, 48, 1, 49, 1, 49, 5, 49, 313, 8, 49, 10, 49, 12, 49, 316, 9, 49, 1,
		50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52,
		328, 8, 52, 10, 52, 12, 52, 331, 9, 52, 0, 0, 53, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
		32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81,
		41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99,
		50, 101, 51, 103, 52, 105, 53, 1, 0, 6, 1, 0, 39, 39, 1, 0, 34, 34, 3,
		0, 65, 90, 95, 95, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122,
		3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 10, 10, 339, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0,
		0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0,
		0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0,
		0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1,
		0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103,
		1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 1, 107, 1, 0, 0, 0, 3, 109, 1, 0, 0, 0,
		5, 112, 1, 0, 0, 0, 7, 114, 1, 0, 0, 0, 9, 116, 1, 0, 0, 0, 11, 118, 1,
		0, 0, 0, 13, 120, 1, 0, 0, 0, 15, 123, 1, 0, 0, 0, 17, 125, 1, 0, 0, 0,
		19, 127, 1, 0, 0, 0, 21, 129, 1, 0, 0, 0, 23, 131, 1, 0, 0, 0, 25, 133,
		1, 0, 0, 0, 27, 135, 1, 0, 0, 0, 29, 137, 1, 0, 0, 0, 31, 140, 1, 0, 0,
		0, 33, 143, 1, 0, 0, 0, 35, 145, 1, 0, 0, 0, 37, 148, 1, 0, 0, 0, 39, 151,
		1, 0, 0, 0, 41, 154, 1, 0, 0, 0, 43, 157, 1, 0, 0, 0, 45, 159, 1, 0, 0,
		0, 47, 161, 1, 0, 0, 0, 49, 163, 1, 0, 0, 0, 51, 165, 1, 0, 0, 0, 53, 167,
		1, 0, 0, 0, 55, 169, 1, 0, 0, 0, 57, 171, 1, 0, 0, 0, 59, 173, 1, 0, 0,
		0, 61, 176, 1, 0, 0, 0, 63, 182, 1, 0, 0, 0, 65, 188, 1, 0, 0, 0, 67, 191,
		1, 0, 0, 0, 69, 195, 1, 0, 0, 0, 71, 202, 1, 0, 0, 0, 73, 207, 1, 0, 0,
		0, 75, 213, 1, 0, 0, 0, 77, 218, 1, 0, 0, 0, 79, 225, 1, 0, 0, 0, 81, 232,
		1, 0, 0, 0, 83, 237, 1, 0, 0, 0, 85, 243, 1, 0, 0, 0, 87, 250, 1, 0, 0,
		0, 89, 259, 1, 0, 0, 0, 91, 266, 1, 0, 0, 0, 93, 271, 1, 0, 0, 0, 95, 280,
		1, 0, 0, 0, 97, 308, 1, 0, 0, 0, 99, 310, 1, 0, 0, 0, 101, 317, 1, 0, 0,
		0, 103, 319, 1, 0, 0, 0, 105, 323, 1, 0, 0, 0, 107, 108, 5, 61, 0, 0, 108,
		2, 1, 0, 0, 0, 109, 110, 5, 42, 0, 0, 110, 111, 5, 42, 0, 0, 111, 4, 1,
		0, 0, 0, 112, 113, 5, 123, 0, 0, 113, 6, 1, 0, 0, 0, 114, 115, 5, 125,
		0, 0, 115, 8, 1, 0, 0, 0, 116, 117, 5, 40, 0, 0, 117, 10, 1, 0, 0, 0, 118,
		119, 5, 41, 0, 0, 119, 12, 1, 0, 0, 0, 120, 121, 5, 36, 0, 0, 121, 122,
		5, 40, 0, 0, 122, 14, 1, 0, 0, 0, 123, 124, 5, 91, 0, 0, 124, 16, 1, 0,
		0, 0, 125, 126, 5, 93, 0, 0, 126, 18, 1, 0, 0, 0, 127, 128, 5, 46, 0, 0,
		128, 20, 1, 0, 0, 0, 129, 130, 5, 58, 0, 0, 130, 22, 1, 0, 0, 0, 131, 132,
		5, 44, 0, 0, 132, 24, 1, 0, 0, 0, 133, 134, 5, 59, 0, 0, 134, 26, 1, 0,
		0, 0, 135, 136, 5, 60, 0, 0, 136, 28, 1, 0, 0, 0, 137, 138, 5, 60, 0, 0,
		138, 139, 5, 61, 0, 0, 139, 30, 1, 0, 0, 0, 140, 141, 5, 62, 0, 0, 141,
		142, 5, 61, 0, 0, 142, 32, 1, 0, 0, 0, 143, 144, 5, 62, 0, 0, 144, 34,
		1, 0, 0, 0, 145, 146, 5, 61, 0, 0, 146, 147, 5, 61, 0, 0, 147, 36, 1, 0,
		0, 0, 148, 149, 5, 33, 0, 0, 149, 150, 5, 61, 0, 0, 150, 38, 1, 0, 0, 0,
		151, 152, 5, 38, 0, 0, 152, 153, 5, 38, 0, 0, 153, 40, 1, 0, 0, 0, 154,
		155, 5, 124, 0, 0, 155, 156, 5, 124, 0, 0, 156, 42, 1, 0, 0, 0, 157, 158,
		5, 33, 0, 0, 158, 44, 1, 0, 0, 0, 159, 160, 5, 38, 0, 0, 160, 46, 1, 0,
		0, 0, 161, 162, 5, 124, 0, 0, 162, 48, 1, 0, 0, 0, 163, 164, 5, 43, 0,
		0, 164, 50, 1, 0, 0, 0, 165, 166, 5, 45, 0, 0, 166, 52, 1, 0, 0, 0, 167,
		168, 5, 42, 0, 0, 168, 54, 1, 0, 0, 0, 169, 170, 5, 94, 0, 0, 170, 56,
		1, 0, 0, 0, 171, 172, 5, 37, 0, 0, 172, 58, 1, 0, 0, 0, 173, 174, 5, 105,
		0, 0, 174, 175, 5, 110, 0, 0, 175, 60, 1, 0, 0, 0, 176, 177, 5, 97, 0,
		0, 177, 178, 5, 108, 0, 0, 178, 179, 5, 108, 0, 0, 179, 180, 5, 111, 0,
		0, 180, 181, 5, 119, 0, 0, 181, 62, 1, 0, 0, 0, 182, 183, 5, 109, 0, 0,
		183, 184, 5, 97, 0, 0, 184, 185, 5, 116, 0, 0, 185, 186, 5, 99, 0, 0, 186,
		187, 5, 104, 0, 0, 187, 64, 1, 0, 0, 0, 188, 189, 5, 105, 0, 0, 189, 190,
		5, 102, 0, 0, 190, 66, 1, 0, 0, 0, 191, 192, 5, 103, 0, 0, 192, 193, 5,
		101, 0, 0, 193, 194, 5, 116, 0, 0, 194, 68, 1, 0, 0, 0, 195, 196, 5, 101,
		0, 0, 196, 197, 5, 120, 0, 0, 197, 198, 5, 105, 0, 0, 198, 199, 5, 115,
		0, 0, 199, 200, 5, 116, 0, 0, 200, 201, 5, 115, 0, 0, 201, 70, 1, 0, 0,
		0, 202, 203, 5, 116, 0, 0, 203, 204, 5, 114, 0, 0, 204, 205, 5, 117, 0,
		0, 205, 206, 5, 101, 0, 0, 206, 72, 1, 0, 0, 0, 207, 208, 5, 102, 0, 0,
		208, 209, 5, 97, 0, 0, 209, 210, 5, 108, 0, 0, 210, 211, 5, 115, 0, 0,
		211, 212, 5, 101, 0, 0, 212, 74, 1, 0, 0, 0, 213, 214, 5, 108, 0, 0, 214,
		215, 5, 105, 0, 0, 215, 216, 5, 115, 0, 0, 216, 217, 5, 116, 0, 0, 217,
		76, 1, 0, 0, 0, 218, 219, 5, 99, 0, 0, 219, 220, 5, 114, 0, 0, 220, 221,
		5, 101, 0, 0, 221, 222, 5, 97, 0, 0, 222, 223, 5, 116, 0, 0, 223, 224,
		5, 101, 0, 0, 224, 78, 1, 0, 0, 0, 225, 226, 5, 117, 0, 0, 226, 227, 5,
		112, 0, 0, 227, 228, 5, 100, 0, 0, 228, 229, 5, 97, 0, 0, 229, 230, 5,
		116, 0, 0, 230, 231, 5, 101, 0, 0, 231, 80, 1, 0, 0, 0, 232, 233, 5, 114,
		0, 0, 233, 234, 5, 101, 0, 0, 234, 235, 5, 97, 0, 0, 235, 236, 5, 100,
		0, 0, 236, 82, 1, 0, 0, 0, 237, 238, 5, 119, 0, 0, 238, 239, 5, 114, 0,
		0, 239, 240, 5, 105, 0, 0, 240, 241, 5, 116, 0, 0, 241, 242, 5, 101, 0,
		0, 242, 84, 1, 0, 0, 0, 243, 244, 5, 100, 0, 0, 244, 245, 5, 101, 0, 0,
		245, 246, 5, 108, 0, 0, 246, 247, 5, 101, 0, 0, 247, 248, 5, 116, 0, 0,
		248, 249, 5, 101, 0, 0, 249, 86, 1, 0, 0, 0, 250, 251, 5, 102, 0, 0, 251,
		252, 5, 117, 0, 0, 252, 253, 5, 110, 0, 0, 253, 254, 5, 99, 0, 0, 254,
		255, 5, 116, 0, 0, 255, 256, 5, 105, 0, 0, 256, 257, 5, 111, 0, 0, 257,
		258, 5, 110, 0, 0, 258, 88, 1, 0, 0, 0, 259, 260, 5, 114, 0, 0, 260, 261,
		5, 101, 0, 0, 261, 262, 5, 116, 0, 0, 262, 263, 5, 117, 0, 0, 263, 264,
		5, 114, 0, 0, 264, 265, 5, 110, 0, 0, 265, 90, 1, 0, 0, 0, 266, 267, 5,
		110, 0, 0, 267, 268, 5, 117, 0, 0, 268, 269, 5, 108, 0, 0, 269, 270, 5,
		108, 0, 0, 270, 92, 1, 0, 0, 0, 271, 272, 5, 115, 0, 0, 272, 273, 5, 101,
		0, 0, 273, 274, 5, 114, 0, 0, 274, 275, 5, 118, 0, 0, 275, 276, 5, 105,
		0, 0, 276, 277, 5, 99, 0, 0, 277, 278, 5, 101, 0, 0, 278, 94, 1, 0, 0,
		0, 279, 281, 2, 48, 57, 0, 280, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0,
		282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 290, 1, 0, 0, 0, 284,
		286, 5, 46, 0, 0, 285, 287, 2, 48, 57, 0, 286, 285, 1, 0, 0, 0, 287, 288,
		1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 291, 1, 0,
		0, 0, 290, 284, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 96, 1, 0, 0, 0,
		292, 296, 5, 39, 0, 0, 293, 295, 8, 0, 0, 0, 294, 293, 1, 0, 0, 0, 295,
		298, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 299,
		1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 299, 309, 5, 39, 0, 0, 300, 304, 5, 34,
		0, 0, 301, 303, 8, 1, 0, 0, 302, 301, 1, 0, 0, 0, 303, 306, 1, 0, 0, 0,
		304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 307, 1, 0, 0, 0, 306,
		304, 1, 0, 0, 0, 307, 309, 5, 34, 0, 0, 308, 292, 1, 0, 0, 0, 308, 300,
		1, 0, 0, 0, 309, 98, 1, 0, 0, 0, 310, 314, 7, 2, 0, 0, 311, 313, 7, 3,
		0, 0, 312, 311, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0,
		314, 315, 1, 0, 0, 0, 315, 100, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317,
		318, 5, 47, 0, 0, 318, 102, 1, 0, 0, 0, 319, 320, 7, 4, 0, 0, 320, 321,
		1, 0, 0, 0, 321, 322, 6, 51, 0, 0, 322, 104, 1, 0, 0, 0, 323, 324, 5, 47,
		0, 0, 324, 325, 5, 47, 0, 0, 325, 329, 1, 0, 0, 0, 326, 328, 8, 5, 0, 0,
		327, 326, 1, 0, 0, 0, 328, 331, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 329,
		330, 1, 0, 0, 0, 330, 106, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 9, 0, 282,
		288, 290, 296, 304, 308, 314, 329, 1, 6, 0, 0,
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

// AclRulesLexerInit initializes any static state used to implement AclRulesLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAclRulesLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AclRulesLexerInit() {
	staticData := &AclRulesLexerLexerStaticData
	staticData.once.Do(aclruleslexerLexerInit)
}

// NewAclRulesLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAclRulesLexer(input antlr.CharStream) *AclRulesLexer {
	AclRulesLexerInit()
	l := new(AclRulesLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AclRulesLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "AclRules.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AclRulesLexer tokens.
const (
	AclRulesLexerT__0                = 1
	AclRulesLexerT__1                = 2
	AclRulesLexerCurlyOpen           = 3
	AclRulesLexerCurlyClose          = 4
	AclRulesLexerBracketOpen         = 5
	AclRulesLexerBracketClose        = 6
	AclRulesLexerPathVariableBracket = 7
	AclRulesLexerSquareBracketOpen   = 8
	AclRulesLexerSquareBracketClose  = 9
	AclRulesLexerDot                 = 10
	AclRulesLexerColon               = 11
	AclRulesLexerComma               = 12
	AclRulesLexerSemicolon           = 13
	AclRulesLexerLessThan            = 14
	AclRulesLexerLessOrEqual         = 15
	AclRulesLexerGreaterOrEqual      = 16
	AclRulesLexerGreaterThan         = 17
	AclRulesLexerEquals              = 18
	AclRulesLexerUnequal             = 19
	AclRulesLexerLogicalAnd          = 20
	AclRulesLexerLogicalOr           = 21
	AclRulesLexerLogicalNot          = 22
	AclRulesLexerBinaryAnd           = 23
	AclRulesLexerBinaryOr            = 24
	AclRulesLexerArithmeticAdd       = 25
	AclRulesLexerArithmeticSub       = 26
	AclRulesLexerArithmeticMul       = 27
	AclRulesLexerArithmeticExp       = 28
	AclRulesLexerArithmeticModus     = 29
	AclRulesLexerInOperator          = 30
	AclRulesLexerAllow               = 31
	AclRulesLexerMatch               = 32
	AclRulesLexerIf                  = 33
	AclRulesLexerGet                 = 34
	AclRulesLexerExists              = 35
	AclRulesLexerTrue                = 36
	AclRulesLexerFalse               = 37
	AclRulesLexerList                = 38
	AclRulesLexerCreate              = 39
	AclRulesLexerUpdate              = 40
	AclRulesLexerRead                = 41
	AclRulesLexerWrite               = 42
	AclRulesLexerDelete              = 43
	AclRulesLexerFunction            = 44
	AclRulesLexerReturn              = 45
	AclRulesLexerNull                = 46
	AclRulesLexerService             = 47
	AclRulesLexerNumber              = 48
	AclRulesLexerString_             = 49
	AclRulesLexerIdentifier          = 50
	AclRulesLexerSlash               = 51
	AclRulesLexerWS                  = 52
	AclRulesLexerComment             = 53
)
