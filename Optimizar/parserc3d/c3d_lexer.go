// Code generated from C3DLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parserc3d

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

type C3DLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var C3DLexerLexerStaticData struct {
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

func c3dlexerLexerInit() {
	staticData := &C3DLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'int'", "'float'", "'double'", "'char'", "'void'", "'#include'",
		"'<stdio.h>'", "'heap'", "'stack'", "'if'", "'goto'", "'return'", "'printf'",
		"'H'", "'P'", "", "", "", "", "", "','", "':'", "';'", "'='", "'!='",
		"'=='", "'!'", "'\\u0009\\u0009>='", "'<='", "'>'", "'<'", "'%'", "'*'",
		"'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "DOUBLE", "CHAR", "VOID", "INCLUDE", "STDIO", "HEAP",
		"STACK", "IF", "GOTO", "RETURN", "PRINTF", "PHEAD", "PSTACK", "NUMERO",
		"CADENA", "ID_VALIDO", "CHARACTER", "WS", "COMMA", "DOS_PUNTOS", "PUNTOCOMA",
		"IG", "DIF", "IG_IG", "NOT", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MODULO",
		"MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER",
		"CORCHIZQ", "CORCHDER", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "DOUBLE", "CHAR", "VOID", "INCLUDE", "STDIO", "HEAP",
		"STACK", "IF", "GOTO", "RETURN", "PRINTF", "PHEAD", "PSTACK", "NUMERO",
		"CADENA", "ID_VALIDO", "CHARACTER", "ESCAPE", "WS", "COMMA", "DOS_PUNTOS",
		"PUNTOCOMA", "IG", "DIF", "IG_IG", "NOT", "MAY_IG", "MEN_IG", "MAYOR",
		"MENOR", "MODULO", "MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ",
		"LLAVEDER", "CORCHIZQ", "CORCHDER", "WHITESPACE", "COMMENT", "LINE_COMMENT",
		"ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 45, 312, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 3, 15, 180, 8, 15, 1, 15, 4, 15, 183, 8, 15, 11, 15,
		12, 15, 184, 1, 15, 1, 15, 4, 15, 189, 8, 15, 11, 15, 12, 15, 190, 3, 15,
		193, 8, 15, 1, 16, 1, 16, 5, 16, 197, 8, 16, 10, 16, 12, 16, 200, 9, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 5, 17, 206, 8, 17, 10, 17, 12, 17, 209, 9,
		17, 1, 18, 1, 18, 1, 18, 3, 18, 214, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 20, 4, 20, 222, 8, 20, 11, 20, 12, 20, 223, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 4, 43, 279,
		8, 43, 11, 43, 12, 43, 280, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 5,
		44, 289, 8, 44, 10, 44, 12, 44, 292, 9, 44, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 303, 8, 45, 10, 45, 12, 45, 306,
		9, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 290, 0, 47, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 0, 41, 20, 43,
		21, 45, 22, 47, 23, 49, 24, 51, 25, 53, 26, 55, 27, 57, 28, 59, 29, 61,
		30, 63, 31, 65, 32, 67, 33, 69, 34, 71, 35, 73, 36, 75, 37, 77, 38, 79,
		39, 81, 40, 83, 41, 85, 42, 87, 43, 89, 44, 91, 45, 93, 0, 1, 0, 10, 1,
		0, 48, 57, 1, 0, 34, 34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 5, 0, 39, 39,
		92, 92, 110, 110, 114, 114, 116, 116, 3, 0, 9, 10, 13, 13, 32, 32, 4, 0,
		9, 10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33, 35,
		35, 43, 43, 45, 46, 58, 58, 64, 64, 91, 93, 320, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0,
		0, 0, 0, 91, 1, 0, 0, 0, 1, 95, 1, 0, 0, 0, 3, 99, 1, 0, 0, 0, 5, 105,
		1, 0, 0, 0, 7, 112, 1, 0, 0, 0, 9, 117, 1, 0, 0, 0, 11, 122, 1, 0, 0, 0,
		13, 131, 1, 0, 0, 0, 15, 141, 1, 0, 0, 0, 17, 146, 1, 0, 0, 0, 19, 152,
		1, 0, 0, 0, 21, 155, 1, 0, 0, 0, 23, 160, 1, 0, 0, 0, 25, 167, 1, 0, 0,
		0, 27, 174, 1, 0, 0, 0, 29, 176, 1, 0, 0, 0, 31, 179, 1, 0, 0, 0, 33, 194,
		1, 0, 0, 0, 35, 203, 1, 0, 0, 0, 37, 210, 1, 0, 0, 0, 39, 217, 1, 0, 0,
		0, 41, 221, 1, 0, 0, 0, 43, 227, 1, 0, 0, 0, 45, 229, 1, 0, 0, 0, 47, 231,
		1, 0, 0, 0, 49, 233, 1, 0, 0, 0, 51, 235, 1, 0, 0, 0, 53, 238, 1, 0, 0,
		0, 55, 241, 1, 0, 0, 0, 57, 243, 1, 0, 0, 0, 59, 248, 1, 0, 0, 0, 61, 251,
		1, 0, 0, 0, 63, 253, 1, 0, 0, 0, 65, 255, 1, 0, 0, 0, 67, 257, 1, 0, 0,
		0, 69, 259, 1, 0, 0, 0, 71, 261, 1, 0, 0, 0, 73, 263, 1, 0, 0, 0, 75, 265,
		1, 0, 0, 0, 77, 267, 1, 0, 0, 0, 79, 269, 1, 0, 0, 0, 81, 271, 1, 0, 0,
		0, 83, 273, 1, 0, 0, 0, 85, 275, 1, 0, 0, 0, 87, 278, 1, 0, 0, 0, 89, 284,
		1, 0, 0, 0, 91, 298, 1, 0, 0, 0, 93, 309, 1, 0, 0, 0, 95, 96, 5, 105, 0,
		0, 96, 97, 5, 110, 0, 0, 97, 98, 5, 116, 0, 0, 98, 2, 1, 0, 0, 0, 99, 100,
		5, 102, 0, 0, 100, 101, 5, 108, 0, 0, 101, 102, 5, 111, 0, 0, 102, 103,
		5, 97, 0, 0, 103, 104, 5, 116, 0, 0, 104, 4, 1, 0, 0, 0, 105, 106, 5, 100,
		0, 0, 106, 107, 5, 111, 0, 0, 107, 108, 5, 117, 0, 0, 108, 109, 5, 98,
		0, 0, 109, 110, 5, 108, 0, 0, 110, 111, 5, 101, 0, 0, 111, 6, 1, 0, 0,
		0, 112, 113, 5, 99, 0, 0, 113, 114, 5, 104, 0, 0, 114, 115, 5, 97, 0, 0,
		115, 116, 5, 114, 0, 0, 116, 8, 1, 0, 0, 0, 117, 118, 5, 118, 0, 0, 118,
		119, 5, 111, 0, 0, 119, 120, 5, 105, 0, 0, 120, 121, 5, 100, 0, 0, 121,
		10, 1, 0, 0, 0, 122, 123, 5, 35, 0, 0, 123, 124, 5, 105, 0, 0, 124, 125,
		5, 110, 0, 0, 125, 126, 5, 99, 0, 0, 126, 127, 5, 108, 0, 0, 127, 128,
		5, 117, 0, 0, 128, 129, 5, 100, 0, 0, 129, 130, 5, 101, 0, 0, 130, 12,
		1, 0, 0, 0, 131, 132, 5, 60, 0, 0, 132, 133, 5, 115, 0, 0, 133, 134, 5,
		116, 0, 0, 134, 135, 5, 100, 0, 0, 135, 136, 5, 105, 0, 0, 136, 137, 5,
		111, 0, 0, 137, 138, 5, 46, 0, 0, 138, 139, 5, 104, 0, 0, 139, 140, 5,
		62, 0, 0, 140, 14, 1, 0, 0, 0, 141, 142, 5, 104, 0, 0, 142, 143, 5, 101,
		0, 0, 143, 144, 5, 97, 0, 0, 144, 145, 5, 112, 0, 0, 145, 16, 1, 0, 0,
		0, 146, 147, 5, 115, 0, 0, 147, 148, 5, 116, 0, 0, 148, 149, 5, 97, 0,
		0, 149, 150, 5, 99, 0, 0, 150, 151, 5, 107, 0, 0, 151, 18, 1, 0, 0, 0,
		152, 153, 5, 105, 0, 0, 153, 154, 5, 102, 0, 0, 154, 20, 1, 0, 0, 0, 155,
		156, 5, 103, 0, 0, 156, 157, 5, 111, 0, 0, 157, 158, 5, 116, 0, 0, 158,
		159, 5, 111, 0, 0, 159, 22, 1, 0, 0, 0, 160, 161, 5, 114, 0, 0, 161, 162,
		5, 101, 0, 0, 162, 163, 5, 116, 0, 0, 163, 164, 5, 117, 0, 0, 164, 165,
		5, 114, 0, 0, 165, 166, 5, 110, 0, 0, 166, 24, 1, 0, 0, 0, 167, 168, 5,
		112, 0, 0, 168, 169, 5, 114, 0, 0, 169, 170, 5, 105, 0, 0, 170, 171, 5,
		110, 0, 0, 171, 172, 5, 116, 0, 0, 172, 173, 5, 102, 0, 0, 173, 26, 1,
		0, 0, 0, 174, 175, 5, 72, 0, 0, 175, 28, 1, 0, 0, 0, 176, 177, 5, 80, 0,
		0, 177, 30, 1, 0, 0, 0, 178, 180, 5, 45, 0, 0, 179, 178, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 182, 1, 0, 0, 0, 181, 183, 7, 0, 0, 0, 182, 181,
		1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0,
		0, 0, 185, 192, 1, 0, 0, 0, 186, 188, 5, 46, 0, 0, 187, 189, 7, 0, 0, 0,
		188, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 186, 1, 0, 0, 0, 192, 193,
		1, 0, 0, 0, 193, 32, 1, 0, 0, 0, 194, 198, 5, 34, 0, 0, 195, 197, 8, 1,
		0, 0, 196, 195, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0,
		198, 199, 1, 0, 0, 0, 199, 201, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 201,
		202, 5, 34, 0, 0, 202, 34, 1, 0, 0, 0, 203, 207, 7, 2, 0, 0, 204, 206,
		7, 3, 0, 0, 205, 204, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0,
		0, 0, 207, 208, 1, 0, 0, 0, 208, 36, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0,
		210, 213, 5, 39, 0, 0, 211, 214, 3, 39, 19, 0, 212, 214, 8, 4, 0, 0, 213,
		211, 1, 0, 0, 0, 213, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216,
		5, 39, 0, 0, 216, 38, 1, 0, 0, 0, 217, 218, 5, 92, 0, 0, 218, 219, 7, 5,
		0, 0, 219, 40, 1, 0, 0, 0, 220, 222, 7, 6, 0, 0, 221, 220, 1, 0, 0, 0,
		222, 223, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224,
		225, 1, 0, 0, 0, 225, 226, 6, 20, 0, 0, 226, 42, 1, 0, 0, 0, 227, 228,
		5, 44, 0, 0, 228, 44, 1, 0, 0, 0, 229, 230, 5, 58, 0, 0, 230, 46, 1, 0,
		0, 0, 231, 232, 5, 59, 0, 0, 232, 48, 1, 0, 0, 0, 233, 234, 5, 61, 0, 0,
		234, 50, 1, 0, 0, 0, 235, 236, 5, 33, 0, 0, 236, 237, 5, 61, 0, 0, 237,
		52, 1, 0, 0, 0, 238, 239, 5, 61, 0, 0, 239, 240, 5, 61, 0, 0, 240, 54,
		1, 0, 0, 0, 241, 242, 5, 33, 0, 0, 242, 56, 1, 0, 0, 0, 243, 244, 5, 9,
		0, 0, 244, 245, 5, 9, 0, 0, 245, 246, 5, 62, 0, 0, 246, 247, 5, 61, 0,
		0, 247, 58, 1, 0, 0, 0, 248, 249, 5, 60, 0, 0, 249, 250, 5, 61, 0, 0, 250,
		60, 1, 0, 0, 0, 251, 252, 5, 62, 0, 0, 252, 62, 1, 0, 0, 0, 253, 254, 5,
		60, 0, 0, 254, 64, 1, 0, 0, 0, 255, 256, 5, 37, 0, 0, 256, 66, 1, 0, 0,
		0, 257, 258, 5, 42, 0, 0, 258, 68, 1, 0, 0, 0, 259, 260, 5, 47, 0, 0, 260,
		70, 1, 0, 0, 0, 261, 262, 5, 43, 0, 0, 262, 72, 1, 0, 0, 0, 263, 264, 5,
		45, 0, 0, 264, 74, 1, 0, 0, 0, 265, 266, 5, 40, 0, 0, 266, 76, 1, 0, 0,
		0, 267, 268, 5, 41, 0, 0, 268, 78, 1, 0, 0, 0, 269, 270, 5, 123, 0, 0,
		270, 80, 1, 0, 0, 0, 271, 272, 5, 125, 0, 0, 272, 82, 1, 0, 0, 0, 273,
		274, 5, 91, 0, 0, 274, 84, 1, 0, 0, 0, 275, 276, 5, 93, 0, 0, 276, 86,
		1, 0, 0, 0, 277, 279, 7, 7, 0, 0, 278, 277, 1, 0, 0, 0, 279, 280, 1, 0,
		0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0,
		282, 283, 6, 43, 0, 0, 283, 88, 1, 0, 0, 0, 284, 285, 5, 47, 0, 0, 285,
		286, 5, 42, 0, 0, 286, 290, 1, 0, 0, 0, 287, 289, 9, 0, 0, 0, 288, 287,
		1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 290, 288, 1, 0,
		0, 0, 291, 293, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 294, 5, 42, 0, 0,
		294, 295, 5, 47, 0, 0, 295, 296, 1, 0, 0, 0, 296, 297, 6, 44, 0, 0, 297,
		90, 1, 0, 0, 0, 298, 299, 5, 47, 0, 0, 299, 300, 5, 47, 0, 0, 300, 304,
		1, 0, 0, 0, 301, 303, 8, 8, 0, 0, 302, 301, 1, 0, 0, 0, 303, 306, 1, 0,
		0, 0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 307, 1, 0, 0, 0,
		306, 304, 1, 0, 0, 0, 307, 308, 6, 45, 0, 0, 308, 92, 1, 0, 0, 0, 309,
		310, 5, 92, 0, 0, 310, 311, 7, 9, 0, 0, 311, 94, 1, 0, 0, 0, 12, 0, 179,
		184, 190, 192, 198, 207, 213, 223, 280, 290, 304, 1, 6, 0, 0,
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

// C3DLexerInit initializes any static state used to implement C3DLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewC3DLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func C3DLexerInit() {
	staticData := &C3DLexerLexerStaticData
	staticData.once.Do(c3dlexerLexerInit)
}

// NewC3DLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewC3DLexer(input antlr.CharStream) *C3DLexer {
	C3DLexerInit()
	l := new(C3DLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &C3DLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "C3DLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// C3DLexer tokens.
const (
	C3DLexerINT          = 1
	C3DLexerFLOAT        = 2
	C3DLexerDOUBLE       = 3
	C3DLexerCHAR         = 4
	C3DLexerVOID         = 5
	C3DLexerINCLUDE      = 6
	C3DLexerSTDIO        = 7
	C3DLexerHEAP         = 8
	C3DLexerSTACK        = 9
	C3DLexerIF           = 10
	C3DLexerGOTO         = 11
	C3DLexerRETURN       = 12
	C3DLexerPRINTF       = 13
	C3DLexerPHEAD        = 14
	C3DLexerPSTACK       = 15
	C3DLexerNUMERO       = 16
	C3DLexerCADENA       = 17
	C3DLexerID_VALIDO    = 18
	C3DLexerCHARACTER    = 19
	C3DLexerWS           = 20
	C3DLexerCOMMA        = 21
	C3DLexerDOS_PUNTOS   = 22
	C3DLexerPUNTOCOMA    = 23
	C3DLexerIG           = 24
	C3DLexerDIF          = 25
	C3DLexerIG_IG        = 26
	C3DLexerNOT          = 27
	C3DLexerMAY_IG       = 28
	C3DLexerMEN_IG       = 29
	C3DLexerMAYOR        = 30
	C3DLexerMENOR        = 31
	C3DLexerMODULO       = 32
	C3DLexerMUL          = 33
	C3DLexerDIV          = 34
	C3DLexerADD          = 35
	C3DLexerSUB          = 36
	C3DLexerPARIZQ       = 37
	C3DLexerPARDER       = 38
	C3DLexerLLAVEIZQ     = 39
	C3DLexerLLAVEDER     = 40
	C3DLexerCORCHIZQ     = 41
	C3DLexerCORCHDER     = 42
	C3DLexerWHITESPACE   = 43
	C3DLexerCOMMENT      = 44
	C3DLexerLINE_COMMENT = 45
)