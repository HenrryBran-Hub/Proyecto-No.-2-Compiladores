// Code generated from C3DGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parserc3d // C3DGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "Optimizar/interfacesc3d"
import "Optimizar/instructionsc3d"

// import "Optimizar/environmentc3d"
// import "Optimizar/expressionsc3d"
// import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type C3DGrammarParser struct {
	*antlr.BaseParser
}

var C3DGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func c3dgrammarParserInit() {
	staticData := &C3DGrammarParserStaticData
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
		"z", "encabezadoa", "temporales", "blocktemporales", "bloquetemps",
		"blockfuncions", "bloquefunciones", "funcionmain", "block", "instruction",
		"head_op", "stack_op", "printff", "embebida", "tipodata", "operaritme",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 251, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 38, 8, 0, 1, 0, 3, 0, 41, 8, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 74, 8, 3, 11, 3, 12, 3, 75,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 85, 8, 4, 1, 5, 4, 5, 88,
		8, 5, 11, 5, 12, 5, 89, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 4, 8, 116, 8, 8, 11, 8, 12, 8, 117, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		3, 9, 134, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 163,
		8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 181, 8, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 189, 8, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 201, 8, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 221, 8, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 3, 13, 235, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 243, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 0,
		0, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 3,
		2, 0, 16, 16, 18, 18, 2, 0, 15, 16, 18, 18, 1, 0, 35, 36, 256, 0, 32, 1,
		0, 0, 0, 2, 46, 1, 0, 0, 0, 4, 68, 1, 0, 0, 0, 6, 73, 1, 0, 0, 0, 8, 84,
		1, 0, 0, 0, 10, 87, 1, 0, 0, 0, 12, 93, 1, 0, 0, 0, 14, 102, 1, 0, 0, 0,
		16, 115, 1, 0, 0, 0, 18, 133, 1, 0, 0, 0, 20, 162, 1, 0, 0, 0, 22, 200,
		1, 0, 0, 0, 24, 220, 1, 0, 0, 0, 26, 234, 1, 0, 0, 0, 28, 242, 1, 0, 0,
		0, 30, 244, 1, 0, 0, 0, 32, 33, 6, 0, -1, 0, 33, 37, 3, 2, 1, 0, 34, 35,
		3, 4, 2, 0, 35, 36, 5, 23, 0, 0, 36, 38, 1, 0, 0, 0, 37, 34, 1, 0, 0, 0,
		37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 41, 3, 10, 5, 0, 40, 39, 1,
		0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43, 3, 14, 7, 0, 43,
		44, 5, 0, 0, 1, 44, 45, 6, 0, -1, 0, 45, 1, 1, 0, 0, 0, 46, 47, 5, 6, 0,
		0, 47, 48, 5, 7, 0, 0, 48, 49, 5, 3, 0, 0, 49, 50, 5, 8, 0, 0, 50, 51,
		5, 41, 0, 0, 51, 52, 5, 16, 0, 0, 52, 53, 5, 42, 0, 0, 53, 54, 5, 23, 0,
		0, 54, 55, 5, 3, 0, 0, 55, 56, 5, 9, 0, 0, 56, 57, 5, 41, 0, 0, 57, 58,
		5, 16, 0, 0, 58, 59, 5, 42, 0, 0, 59, 60, 5, 23, 0, 0, 60, 61, 5, 3, 0,
		0, 61, 62, 5, 15, 0, 0, 62, 63, 5, 23, 0, 0, 63, 64, 5, 3, 0, 0, 64, 65,
		5, 14, 0, 0, 65, 66, 5, 23, 0, 0, 66, 67, 6, 1, -1, 0, 67, 3, 1, 0, 0,
		0, 68, 69, 5, 3, 0, 0, 69, 70, 3, 6, 3, 0, 70, 71, 6, 2, -1, 0, 71, 5,
		1, 0, 0, 0, 72, 74, 3, 8, 4, 0, 73, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0,
		75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 6,
		3, -1, 0, 78, 7, 1, 0, 0, 0, 79, 80, 5, 21, 0, 0, 80, 81, 5, 18, 0, 0,
		81, 85, 6, 4, -1, 0, 82, 83, 5, 18, 0, 0, 83, 85, 6, 4, -1, 0, 84, 79,
		1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 9, 1, 0, 0, 0, 86, 88, 3, 12, 6, 0,
		87, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 6, 5, -1, 0, 92, 11, 1, 0, 0, 0, 93,
		94, 5, 5, 0, 0, 94, 95, 5, 18, 0, 0, 95, 96, 5, 37, 0, 0, 96, 97, 5, 38,
		0, 0, 97, 98, 5, 39, 0, 0, 98, 99, 3, 16, 8, 0, 99, 100, 5, 40, 0, 0, 100,
		101, 6, 6, -1, 0, 101, 13, 1, 0, 0, 0, 102, 103, 5, 1, 0, 0, 103, 104,
		5, 18, 0, 0, 104, 105, 5, 37, 0, 0, 105, 106, 5, 38, 0, 0, 106, 107, 5,
		39, 0, 0, 107, 108, 3, 16, 8, 0, 108, 109, 5, 12, 0, 0, 109, 110, 5, 16,
		0, 0, 110, 111, 5, 23, 0, 0, 111, 112, 5, 40, 0, 0, 112, 113, 6, 7, -1,
		0, 113, 15, 1, 0, 0, 0, 114, 116, 3, 18, 9, 0, 115, 114, 1, 0, 0, 0, 116,
		117, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119,
		1, 0, 0, 0, 119, 120, 6, 8, -1, 0, 120, 17, 1, 0, 0, 0, 121, 122, 3, 20,
		10, 0, 122, 123, 6, 9, -1, 0, 123, 134, 1, 0, 0, 0, 124, 125, 3, 22, 11,
		0, 125, 126, 6, 9, -1, 0, 126, 134, 1, 0, 0, 0, 127, 128, 3, 24, 12, 0,
		128, 129, 6, 9, -1, 0, 129, 134, 1, 0, 0, 0, 130, 131, 3, 30, 15, 0, 131,
		132, 6, 9, -1, 0, 132, 134, 1, 0, 0, 0, 133, 121, 1, 0, 0, 0, 133, 124,
		1, 0, 0, 0, 133, 127, 1, 0, 0, 0, 133, 130, 1, 0, 0, 0, 134, 19, 1, 0,
		0, 0, 135, 136, 5, 14, 0, 0, 136, 137, 5, 24, 0, 0, 137, 138, 5, 16, 0,
		0, 138, 139, 5, 23, 0, 0, 139, 163, 6, 10, -1, 0, 140, 141, 5, 18, 0, 0,
		141, 142, 5, 24, 0, 0, 142, 143, 5, 14, 0, 0, 143, 144, 5, 23, 0, 0, 144,
		163, 6, 10, -1, 0, 145, 146, 5, 8, 0, 0, 146, 147, 5, 41, 0, 0, 147, 148,
		3, 26, 13, 0, 148, 149, 5, 14, 0, 0, 149, 150, 5, 42, 0, 0, 150, 151, 5,
		24, 0, 0, 151, 152, 7, 0, 0, 0, 152, 153, 5, 23, 0, 0, 153, 154, 6, 10,
		-1, 0, 154, 163, 1, 0, 0, 0, 155, 156, 5, 14, 0, 0, 156, 157, 5, 24, 0,
		0, 157, 158, 5, 14, 0, 0, 158, 159, 5, 35, 0, 0, 159, 160, 5, 16, 0, 0,
		160, 161, 5, 23, 0, 0, 161, 163, 6, 10, -1, 0, 162, 135, 1, 0, 0, 0, 162,
		140, 1, 0, 0, 0, 162, 145, 1, 0, 0, 0, 162, 155, 1, 0, 0, 0, 163, 21, 1,
		0, 0, 0, 164, 165, 5, 15, 0, 0, 165, 166, 5, 24, 0, 0, 166, 167, 5, 16,
		0, 0, 167, 168, 5, 23, 0, 0, 168, 201, 6, 11, -1, 0, 169, 170, 5, 18, 0,
		0, 170, 171, 5, 24, 0, 0, 171, 172, 5, 15, 0, 0, 172, 173, 5, 23, 0, 0,
		173, 201, 6, 11, -1, 0, 174, 175, 5, 9, 0, 0, 175, 180, 5, 41, 0, 0, 176,
		181, 5, 16, 0, 0, 177, 178, 3, 26, 13, 0, 178, 179, 7, 1, 0, 0, 179, 181,
		1, 0, 0, 0, 180, 176, 1, 0, 0, 0, 180, 177, 1, 0, 0, 0, 181, 182, 1, 0,
		0, 0, 182, 183, 5, 42, 0, 0, 183, 188, 5, 24, 0, 0, 184, 185, 5, 37, 0,
		0, 185, 186, 3, 28, 14, 0, 186, 187, 5, 38, 0, 0, 187, 189, 1, 0, 0, 0,
		188, 184, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		191, 7, 0, 0, 0, 191, 192, 5, 23, 0, 0, 192, 201, 6, 11, -1, 0, 193, 194,
		5, 15, 0, 0, 194, 195, 5, 24, 0, 0, 195, 196, 5, 15, 0, 0, 196, 197, 7,
		2, 0, 0, 197, 198, 5, 16, 0, 0, 198, 199, 5, 23, 0, 0, 199, 201, 6, 11,
		-1, 0, 200, 164, 1, 0, 0, 0, 200, 169, 1, 0, 0, 0, 200, 174, 1, 0, 0, 0,
		200, 193, 1, 0, 0, 0, 201, 23, 1, 0, 0, 0, 202, 203, 5, 13, 0, 0, 203,
		204, 5, 37, 0, 0, 204, 205, 5, 17, 0, 0, 205, 206, 5, 21, 0, 0, 206, 207,
		5, 16, 0, 0, 207, 208, 5, 38, 0, 0, 208, 209, 5, 23, 0, 0, 209, 221, 6,
		12, -1, 0, 210, 211, 5, 13, 0, 0, 211, 212, 5, 37, 0, 0, 212, 213, 5, 17,
		0, 0, 213, 214, 5, 21, 0, 0, 214, 215, 3, 26, 13, 0, 215, 216, 7, 0, 0,
		0, 216, 217, 5, 38, 0, 0, 217, 218, 5, 23, 0, 0, 218, 219, 6, 12, -1, 0,
		219, 221, 1, 0, 0, 0, 220, 202, 1, 0, 0, 0, 220, 210, 1, 0, 0, 0, 221,
		25, 1, 0, 0, 0, 222, 223, 5, 37, 0, 0, 223, 224, 5, 1, 0, 0, 224, 225,
		5, 38, 0, 0, 225, 235, 6, 13, -1, 0, 226, 227, 5, 37, 0, 0, 227, 228, 5,
		2, 0, 0, 228, 229, 5, 38, 0, 0, 229, 235, 6, 13, -1, 0, 230, 231, 5, 37,
		0, 0, 231, 232, 5, 4, 0, 0, 232, 233, 5, 38, 0, 0, 233, 235, 6, 13, -1,
		0, 234, 222, 1, 0, 0, 0, 234, 226, 1, 0, 0, 0, 234, 230, 1, 0, 0, 0, 235,
		27, 1, 0, 0, 0, 236, 237, 5, 1, 0, 0, 237, 243, 6, 14, -1, 0, 238, 239,
		5, 2, 0, 0, 239, 243, 6, 14, -1, 0, 240, 241, 5, 4, 0, 0, 241, 243, 6,
		14, -1, 0, 242, 236, 1, 0, 0, 0, 242, 238, 1, 0, 0, 0, 242, 240, 1, 0,
		0, 0, 243, 29, 1, 0, 0, 0, 244, 245, 5, 18, 0, 0, 245, 246, 5, 24, 0, 0,
		246, 247, 7, 0, 0, 0, 247, 248, 5, 23, 0, 0, 248, 249, 6, 15, -1, 0, 249,
		31, 1, 0, 0, 0, 14, 37, 40, 75, 84, 89, 117, 133, 162, 180, 188, 200, 220,
		234, 242,
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

// C3DGrammarParserInit initializes any static state used to implement C3DGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewC3DGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func C3DGrammarParserInit() {
	staticData := &C3DGrammarParserStaticData
	staticData.once.Do(c3dgrammarParserInit)
}

// NewC3DGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewC3DGrammarParser(input antlr.TokenStream) *C3DGrammarParser {
	C3DGrammarParserInit()
	this := new(C3DGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &C3DGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "C3DGrammar.g4"

	return this
}

// C3DGrammarParser tokens.
const (
	C3DGrammarParserEOF          = antlr.TokenEOF
	C3DGrammarParserINT          = 1
	C3DGrammarParserFLOAT        = 2
	C3DGrammarParserDOUBLE       = 3
	C3DGrammarParserCHAR         = 4
	C3DGrammarParserVOID         = 5
	C3DGrammarParserINCLUDE      = 6
	C3DGrammarParserSTDIO        = 7
	C3DGrammarParserHEAP         = 8
	C3DGrammarParserSTACK        = 9
	C3DGrammarParserIF           = 10
	C3DGrammarParserGOTO         = 11
	C3DGrammarParserRETURN       = 12
	C3DGrammarParserPRINTF       = 13
	C3DGrammarParserPHEAD        = 14
	C3DGrammarParserPSTACK       = 15
	C3DGrammarParserNUMERO       = 16
	C3DGrammarParserCADENA       = 17
	C3DGrammarParserID_VALIDO    = 18
	C3DGrammarParserCHARACTER    = 19
	C3DGrammarParserWS           = 20
	C3DGrammarParserCOMMA        = 21
	C3DGrammarParserDOS_PUNTOS   = 22
	C3DGrammarParserPUNTOCOMA    = 23
	C3DGrammarParserIG           = 24
	C3DGrammarParserDIF          = 25
	C3DGrammarParserIG_IG        = 26
	C3DGrammarParserNOT          = 27
	C3DGrammarParserMAY_IG       = 28
	C3DGrammarParserMEN_IG       = 29
	C3DGrammarParserMAYOR        = 30
	C3DGrammarParserMENOR        = 31
	C3DGrammarParserMODULO       = 32
	C3DGrammarParserMUL          = 33
	C3DGrammarParserDIV          = 34
	C3DGrammarParserADD          = 35
	C3DGrammarParserSUB          = 36
	C3DGrammarParserPARIZQ       = 37
	C3DGrammarParserPARDER       = 38
	C3DGrammarParserLLAVEIZQ     = 39
	C3DGrammarParserLLAVEDER     = 40
	C3DGrammarParserCORCHIZQ     = 41
	C3DGrammarParserCORCHDER     = 42
	C3DGrammarParserWHITESPACE   = 43
	C3DGrammarParserCOMMENT      = 44
	C3DGrammarParserLINE_COMMENT = 45
)

// C3DGrammarParser rules.
const (
	C3DGrammarParserRULE_z               = 0
	C3DGrammarParserRULE_encabezadoa     = 1
	C3DGrammarParserRULE_temporales      = 2
	C3DGrammarParserRULE_blocktemporales = 3
	C3DGrammarParserRULE_bloquetemps     = 4
	C3DGrammarParserRULE_blockfuncions   = 5
	C3DGrammarParserRULE_bloquefunciones = 6
	C3DGrammarParserRULE_funcionmain     = 7
	C3DGrammarParserRULE_block           = 8
	C3DGrammarParserRULE_instruction     = 9
	C3DGrammarParserRULE_head_op         = 10
	C3DGrammarParserRULE_stack_op        = 11
	C3DGrammarParserRULE_printff         = 12
	C3DGrammarParserRULE_embebida        = 13
	C3DGrammarParserRULE_tipodata        = 14
	C3DGrammarParserRULE_operaritme      = 15
)

// IZContext is an interface to support dynamic dispatch.
type IZContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PUNTOCOMA returns the _PUNTOCOMA token.
	Get_PUNTOCOMA() antlr.Token

	// Set_PUNTOCOMA sets the _PUNTOCOMA token.
	Set_PUNTOCOMA(antlr.Token)

	// Get_encabezadoa returns the _encabezadoa rule contexts.
	Get_encabezadoa() IEncabezadoaContext

	// Get_temporales returns the _temporales rule contexts.
	Get_temporales() ITemporalesContext

	// Get_blockfuncions returns the _blockfuncions rule contexts.
	Get_blockfuncions() IBlockfuncionsContext

	// Get_funcionmain returns the _funcionmain rule contexts.
	Get_funcionmain() IFuncionmainContext

	// Set_encabezadoa sets the _encabezadoa rule contexts.
	Set_encabezadoa(IEncabezadoaContext)

	// Set_temporales sets the _temporales rule contexts.
	Set_temporales(ITemporalesContext)

	// Set_blockfuncions sets the _blockfuncions rule contexts.
	Set_blockfuncions(IBlockfuncionsContext)

	// Set_funcionmain sets the _funcionmain rule contexts.
	Set_funcionmain(IFuncionmainContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Encabezadoa() IEncabezadoaContext
	Funcionmain() IFuncionmainContext
	EOF() antlr.TerminalNode
	Temporales() ITemporalesContext
	PUNTOCOMA() antlr.TerminalNode
	Blockfuncions() IBlockfuncionsContext

	// IsZContext differentiates from other interfaces.
	IsZContext()
}

type ZContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	code           []interface{}
	_encabezadoa   IEncabezadoaContext
	_temporales    ITemporalesContext
	_PUNTOCOMA     antlr.Token
	_blockfuncions IBlockfuncionsContext
	_funcionmain   IFuncionmainContext
}

func NewEmptyZContext() *ZContext {
	var p = new(ZContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_z
	return p
}

func InitEmptyZContext(p *ZContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_z
}

func (*ZContext) IsZContext() {}

func NewZContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZContext {
	var p = new(ZContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_z

	return p
}

func (s *ZContext) GetParser() antlr.Parser { return s.parser }

func (s *ZContext) Get_PUNTOCOMA() antlr.Token { return s._PUNTOCOMA }

func (s *ZContext) Set_PUNTOCOMA(v antlr.Token) { s._PUNTOCOMA = v }

func (s *ZContext) Get_encabezadoa() IEncabezadoaContext { return s._encabezadoa }

func (s *ZContext) Get_temporales() ITemporalesContext { return s._temporales }

func (s *ZContext) Get_blockfuncions() IBlockfuncionsContext { return s._blockfuncions }

func (s *ZContext) Get_funcionmain() IFuncionmainContext { return s._funcionmain }

func (s *ZContext) Set_encabezadoa(v IEncabezadoaContext) { s._encabezadoa = v }

func (s *ZContext) Set_temporales(v ITemporalesContext) { s._temporales = v }

func (s *ZContext) Set_blockfuncions(v IBlockfuncionsContext) { s._blockfuncions = v }

func (s *ZContext) Set_funcionmain(v IFuncionmainContext) { s._funcionmain = v }

func (s *ZContext) GetCode() []interface{} { return s.code }

func (s *ZContext) SetCode(v []interface{}) { s.code = v }

func (s *ZContext) Encabezadoa() IEncabezadoaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEncabezadoaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEncabezadoaContext)
}

func (s *ZContext) Funcionmain() IFuncionmainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncionmainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncionmainContext)
}

func (s *ZContext) EOF() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserEOF, 0)
}

func (s *ZContext) Temporales() ITemporalesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITemporalesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITemporalesContext)
}

func (s *ZContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPUNTOCOMA, 0)
}

func (s *ZContext) Blockfuncions() IBlockfuncionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockfuncionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockfuncionsContext)
}

func (s *ZContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterZ(s)
	}
}

func (s *ZContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitZ(s)
	}
}

func (p *C3DGrammarParser) Z() (localctx IZContext) {
	localctx = NewZContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, C3DGrammarParserRULE_z)
	var _la int

	p.EnterOuterAlt(localctx, 1)

	var mySlice []interface{}
	mySlice = make([]interface{}, 0) // Inicializa el slice vacío

	{
		p.SetState(33)

		var _x = p.Encabezadoa()

		localctx.(*ZContext)._encabezadoa = _x
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == C3DGrammarParserDOUBLE {
		{
			p.SetState(34)

			var _x = p.Temporales()

			localctx.(*ZContext)._temporales = _x
		}
		{
			p.SetState(35)

			var _m = p.Match(C3DGrammarParserPUNTOCOMA)

			localctx.(*ZContext)._PUNTOCOMA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == C3DGrammarParserVOID {
		{
			p.SetState(39)

			var _x = p.Blockfuncions()

			localctx.(*ZContext)._blockfuncions = _x
		}

	}
	{
		p.SetState(42)

		var _x = p.Funcionmain()

		localctx.(*ZContext)._funcionmain = _x
	}
	{
		p.SetState(43)
		p.Match(C3DGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	mySlice = append(mySlice, localctx.(*ZContext).Get_encabezadoa().GetEncaa())
	if localctx.(*ZContext).Get_PUNTOCOMA() != nil {
		mySlice = append(mySlice, localctx.(*ZContext).Get_temporales().GetTinst())
	}

	for _, item := range localctx.(*ZContext).Get_blockfuncions().GetBlkfunc() {
		mySlice = append(mySlice, item)
	}

	mySlice = append(mySlice, localctx.(*ZContext).Get_funcionmain().GetFunmain())

	localctx.(*ZContext).code = mySlice

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

// IEncabezadoaContext is an interface to support dynamic dispatch.
type IEncabezadoaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetH returns the h token.
	GetH() antlr.Token

	// GetS returns the s token.
	GetS() antlr.Token

	// SetH sets the h token.
	SetH(antlr.Token)

	// SetS sets the s token.
	SetS(antlr.Token)

	// GetEncaa returns the encaa attribute.
	GetEncaa() interfacesc3d.Instruction

	// SetEncaa sets the encaa attribute.
	SetEncaa(interfacesc3d.Instruction)

	// Getter signatures
	INCLUDE() antlr.TerminalNode
	STDIO() antlr.TerminalNode
	AllDOUBLE() []antlr.TerminalNode
	DOUBLE(i int) antlr.TerminalNode
	HEAP() antlr.TerminalNode
	AllCORCHIZQ() []antlr.TerminalNode
	CORCHIZQ(i int) antlr.TerminalNode
	AllCORCHDER() []antlr.TerminalNode
	CORCHDER(i int) antlr.TerminalNode
	AllPUNTOCOMA() []antlr.TerminalNode
	PUNTOCOMA(i int) antlr.TerminalNode
	STACK() antlr.TerminalNode
	PSTACK() antlr.TerminalNode
	PHEAD() antlr.TerminalNode
	AllNUMERO() []antlr.TerminalNode
	NUMERO(i int) antlr.TerminalNode

	// IsEncabezadoaContext differentiates from other interfaces.
	IsEncabezadoaContext()
}

type EncabezadoaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	encaa  interfacesc3d.Instruction
	h      antlr.Token
	s      antlr.Token
}

func NewEmptyEncabezadoaContext() *EncabezadoaContext {
	var p = new(EncabezadoaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_encabezadoa
	return p
}

func InitEmptyEncabezadoaContext(p *EncabezadoaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_encabezadoa
}

func (*EncabezadoaContext) IsEncabezadoaContext() {}

func NewEncabezadoaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncabezadoaContext {
	var p = new(EncabezadoaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_encabezadoa

	return p
}

func (s *EncabezadoaContext) GetParser() antlr.Parser { return s.parser }

func (s *EncabezadoaContext) GetH() antlr.Token { return s.h }

func (s *EncabezadoaContext) GetS() antlr.Token { return s.s }

func (s *EncabezadoaContext) SetH(v antlr.Token) { s.h = v }

func (s *EncabezadoaContext) SetS(v antlr.Token) { s.s = v }

func (s *EncabezadoaContext) GetEncaa() interfacesc3d.Instruction { return s.encaa }

func (s *EncabezadoaContext) SetEncaa(v interfacesc3d.Instruction) { s.encaa = v }

func (s *EncabezadoaContext) INCLUDE() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserINCLUDE, 0)
}

func (s *EncabezadoaContext) STDIO() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserSTDIO, 0)
}

func (s *EncabezadoaContext) AllDOUBLE() []antlr.TerminalNode {
	return s.GetTokens(C3DGrammarParserDOUBLE)
}

func (s *EncabezadoaContext) DOUBLE(i int) antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserDOUBLE, i)
}

func (s *EncabezadoaContext) HEAP() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserHEAP, 0)
}

func (s *EncabezadoaContext) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(C3DGrammarParserCORCHIZQ)
}

func (s *EncabezadoaContext) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCORCHIZQ, i)
}

func (s *EncabezadoaContext) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(C3DGrammarParserCORCHDER)
}

func (s *EncabezadoaContext) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCORCHDER, i)
}

func (s *EncabezadoaContext) AllPUNTOCOMA() []antlr.TerminalNode {
	return s.GetTokens(C3DGrammarParserPUNTOCOMA)
}

func (s *EncabezadoaContext) PUNTOCOMA(i int) antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPUNTOCOMA, i)
}

func (s *EncabezadoaContext) STACK() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserSTACK, 0)
}

func (s *EncabezadoaContext) PSTACK() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPSTACK, 0)
}

func (s *EncabezadoaContext) PHEAD() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPHEAD, 0)
}

func (s *EncabezadoaContext) AllNUMERO() []antlr.TerminalNode {
	return s.GetTokens(C3DGrammarParserNUMERO)
}

func (s *EncabezadoaContext) NUMERO(i int) antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserNUMERO, i)
}

func (s *EncabezadoaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncabezadoaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncabezadoaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterEncabezadoa(s)
	}
}

func (s *EncabezadoaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitEncabezadoa(s)
	}
}

func (p *C3DGrammarParser) Encabezadoa() (localctx IEncabezadoaContext) {
	localctx = NewEncabezadoaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, C3DGrammarParserRULE_encabezadoa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(C3DGrammarParserINCLUDE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(C3DGrammarParserSTDIO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(C3DGrammarParserDOUBLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(C3DGrammarParserHEAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(C3DGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)

		var _m = p.Match(C3DGrammarParserNUMERO)

		localctx.(*EncabezadoaContext).h = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(C3DGrammarParserCORCHDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(C3DGrammarParserPUNTOCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(C3DGrammarParserDOUBLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(C3DGrammarParserSTACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(C3DGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)

		var _m = p.Match(C3DGrammarParserNUMERO)

		localctx.(*EncabezadoaContext).s = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(C3DGrammarParserCORCHDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(C3DGrammarParserPUNTOCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(C3DGrammarParserDOUBLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(C3DGrammarParserPSTACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Match(C3DGrammarParserPUNTOCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(C3DGrammarParserDOUBLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(C3DGrammarParserPHEAD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(C3DGrammarParserPUNTOCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*EncabezadoaContext).encaa = instructionsc3d.NewAcumuladorEncabezado((func() string {
		if localctx.(*EncabezadoaContext).GetH() == nil {
			return ""
		} else {
			return localctx.(*EncabezadoaContext).GetH().GetText()
		}
	}()), (func() string {
		if localctx.(*EncabezadoaContext).GetS() == nil {
			return ""
		} else {
			return localctx.(*EncabezadoaContext).GetS().GetText()
		}
	}()))

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

// ITemporalesContext is an interface to support dynamic dispatch.
type ITemporalesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_blocktemporales returns the _blocktemporales rule contexts.
	Get_blocktemporales() IBlocktemporalesContext

	// Set_blocktemporales sets the _blocktemporales rule contexts.
	Set_blocktemporales(IBlocktemporalesContext)

	// GetTinst returns the tinst attribute.
	GetTinst() interfacesc3d.Instruction

	// SetTinst sets the tinst attribute.
	SetTinst(interfacesc3d.Instruction)

	// Getter signatures
	DOUBLE() antlr.TerminalNode
	Blocktemporales() IBlocktemporalesContext

	// IsTemporalesContext differentiates from other interfaces.
	IsTemporalesContext()
}

type TemporalesContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	tinst            interfacesc3d.Instruction
	_blocktemporales IBlocktemporalesContext
}

func NewEmptyTemporalesContext() *TemporalesContext {
	var p = new(TemporalesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_temporales
	return p
}

func InitEmptyTemporalesContext(p *TemporalesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_temporales
}

func (*TemporalesContext) IsTemporalesContext() {}

func NewTemporalesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemporalesContext {
	var p = new(TemporalesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_temporales

	return p
}

func (s *TemporalesContext) GetParser() antlr.Parser { return s.parser }

func (s *TemporalesContext) Get_blocktemporales() IBlocktemporalesContext { return s._blocktemporales }

func (s *TemporalesContext) Set_blocktemporales(v IBlocktemporalesContext) { s._blocktemporales = v }

func (s *TemporalesContext) GetTinst() interfacesc3d.Instruction { return s.tinst }

func (s *TemporalesContext) SetTinst(v interfacesc3d.Instruction) { s.tinst = v }

func (s *TemporalesContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserDOUBLE, 0)
}

func (s *TemporalesContext) Blocktemporales() IBlocktemporalesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlocktemporalesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlocktemporalesContext)
}

func (s *TemporalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemporalesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemporalesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterTemporales(s)
	}
}

func (s *TemporalesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitTemporales(s)
	}
}

func (p *C3DGrammarParser) Temporales() (localctx ITemporalesContext) {
	localctx = NewTemporalesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, C3DGrammarParserRULE_temporales)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(C3DGrammarParserDOUBLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)

		var _x = p.Blocktemporales()

		localctx.(*TemporalesContext)._blocktemporales = _x
	}

	localctx.(*TemporalesContext).tinst = instructionsc3d.NewEjecucionTemporales(localctx.(*TemporalesContext).Get_blocktemporales().GetBlktmps())

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

// IBlocktemporalesContext is an interface to support dynamic dispatch.
type IBlocktemporalesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_bloquetemps returns the _bloquetemps rule contexts.
	Get_bloquetemps() IBloquetempsContext

	// Set_bloquetemps sets the _bloquetemps rule contexts.
	Set_bloquetemps(IBloquetempsContext)

	// GetTemps returns the temps rule context list.
	GetTemps() []IBloquetempsContext

	// SetTemps sets the temps rule context list.
	SetTemps([]IBloquetempsContext)

	// GetBlktmps returns the blktmps attribute.
	GetBlktmps() []interface{}

	// SetBlktmps sets the blktmps attribute.
	SetBlktmps([]interface{})

	// Getter signatures
	AllBloquetemps() []IBloquetempsContext
	Bloquetemps(i int) IBloquetempsContext

	// IsBlocktemporalesContext differentiates from other interfaces.
	IsBlocktemporalesContext()
}

type BlocktemporalesContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blktmps      []interface{}
	_bloquetemps IBloquetempsContext
	temps        []IBloquetempsContext
}

func NewEmptyBlocktemporalesContext() *BlocktemporalesContext {
	var p = new(BlocktemporalesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_blocktemporales
	return p
}

func InitEmptyBlocktemporalesContext(p *BlocktemporalesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_blocktemporales
}

func (*BlocktemporalesContext) IsBlocktemporalesContext() {}

func NewBlocktemporalesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlocktemporalesContext {
	var p = new(BlocktemporalesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_blocktemporales

	return p
}

func (s *BlocktemporalesContext) GetParser() antlr.Parser { return s.parser }

func (s *BlocktemporalesContext) Get_bloquetemps() IBloquetempsContext { return s._bloquetemps }

func (s *BlocktemporalesContext) Set_bloquetemps(v IBloquetempsContext) { s._bloquetemps = v }

func (s *BlocktemporalesContext) GetTemps() []IBloquetempsContext { return s.temps }

func (s *BlocktemporalesContext) SetTemps(v []IBloquetempsContext) { s.temps = v }

func (s *BlocktemporalesContext) GetBlktmps() []interface{} { return s.blktmps }

func (s *BlocktemporalesContext) SetBlktmps(v []interface{}) { s.blktmps = v }

func (s *BlocktemporalesContext) AllBloquetemps() []IBloquetempsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBloquetempsContext); ok {
			len++
		}
	}

	tst := make([]IBloquetempsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBloquetempsContext); ok {
			tst[i] = t.(IBloquetempsContext)
			i++
		}
	}

	return tst
}

func (s *BlocktemporalesContext) Bloquetemps(i int) IBloquetempsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloquetempsContext); ok {
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

	return t.(IBloquetempsContext)
}

func (s *BlocktemporalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlocktemporalesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlocktemporalesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterBlocktemporales(s)
	}
}

func (s *BlocktemporalesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitBlocktemporales(s)
	}
}

func (p *C3DGrammarParser) Blocktemporales() (localctx IBlocktemporalesContext) {
	localctx = NewBlocktemporalesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, C3DGrammarParserRULE_blocktemporales)

	localctx.(*BlocktemporalesContext).blktmps = []interface{}{}
	var listTemp []IBloquetempsContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == C3DGrammarParserID_VALIDO || _la == C3DGrammarParserCOMMA {
		{
			p.SetState(72)

			var _x = p.Bloquetemps()

			localctx.(*BlocktemporalesContext)._bloquetemps = _x
		}
		localctx.(*BlocktemporalesContext).temps = append(localctx.(*BlocktemporalesContext).temps, localctx.(*BlocktemporalesContext)._bloquetemps)

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listTemp = localctx.(*BlocktemporalesContext).GetTemps()
	for _, e := range listTemp {
		localctx.(*BlocktemporalesContext).blktmps = append(localctx.(*BlocktemporalesContext).blktmps, e.GetTemps())
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

// IBloquetempsContext is an interface to support dynamic dispatch.
type IBloquetempsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// GetTemps returns the temps attribute.
	GetTemps() interfacesc3d.Instruction

	// SetTemps sets the temps attribute.
	SetTemps(interfacesc3d.Instruction)

	// Getter signatures
	COMMA() antlr.TerminalNode
	ID_VALIDO() antlr.TerminalNode

	// IsBloquetempsContext differentiates from other interfaces.
	IsBloquetempsContext()
}

type BloquetempsContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	temps      interfacesc3d.Instruction
	_ID_VALIDO antlr.Token
}

func NewEmptyBloquetempsContext() *BloquetempsContext {
	var p = new(BloquetempsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_bloquetemps
	return p
}

func InitEmptyBloquetempsContext(p *BloquetempsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_bloquetemps
}

func (*BloquetempsContext) IsBloquetempsContext() {}

func NewBloquetempsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloquetempsContext {
	var p = new(BloquetempsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_bloquetemps

	return p
}

func (s *BloquetempsContext) GetParser() antlr.Parser { return s.parser }

func (s *BloquetempsContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *BloquetempsContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *BloquetempsContext) GetTemps() interfacesc3d.Instruction { return s.temps }

func (s *BloquetempsContext) SetTemps(v interfacesc3d.Instruction) { s.temps = v }

func (s *BloquetempsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCOMMA, 0)
}

func (s *BloquetempsContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserID_VALIDO, 0)
}

func (s *BloquetempsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloquetempsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloquetempsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterBloquetemps(s)
	}
}

func (s *BloquetempsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitBloquetemps(s)
	}
}

func (p *C3DGrammarParser) Bloquetemps() (localctx IBloquetempsContext) {
	localctx = NewBloquetempsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, C3DGrammarParserRULE_bloquetemps)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case C3DGrammarParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Match(C3DGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)

			var _m = p.Match(C3DGrammarParserID_VALIDO)

			localctx.(*BloquetempsContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*BloquetempsContext).temps = instructionsc3d.NewArregloTemporales((func() string {
			if localctx.(*BloquetempsContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*BloquetempsContext).Get_ID_VALIDO().GetText()
			}
		}()))

	case C3DGrammarParserID_VALIDO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)

			var _m = p.Match(C3DGrammarParserID_VALIDO)

			localctx.(*BloquetempsContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*BloquetempsContext).temps = instructionsc3d.NewArregloTemporales((func() string {
			if localctx.(*BloquetempsContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*BloquetempsContext).Get_ID_VALIDO().GetText()
			}
		}()))

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

// IBlockfuncionsContext is an interface to support dynamic dispatch.
type IBlockfuncionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_bloquefunciones returns the _bloquefunciones rule contexts.
	Get_bloquefunciones() IBloquefuncionesContext

	// Set_bloquefunciones sets the _bloquefunciones rule contexts.
	Set_bloquefunciones(IBloquefuncionesContext)

	// GetFunc_ returns the func_ rule context list.
	GetFunc_() []IBloquefuncionesContext

	// SetFunc_ sets the func_ rule context list.
	SetFunc_([]IBloquefuncionesContext)

	// GetBlkfunc returns the blkfunc attribute.
	GetBlkfunc() []interface{}

	// SetBlkfunc sets the blkfunc attribute.
	SetBlkfunc([]interface{})

	// Getter signatures
	AllBloquefunciones() []IBloquefuncionesContext
	Bloquefunciones(i int) IBloquefuncionesContext

	// IsBlockfuncionsContext differentiates from other interfaces.
	IsBlockfuncionsContext()
}

type BlockfuncionsContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	blkfunc          []interface{}
	_bloquefunciones IBloquefuncionesContext
	func_            []IBloquefuncionesContext
}

func NewEmptyBlockfuncionsContext() *BlockfuncionsContext {
	var p = new(BlockfuncionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_blockfuncions
	return p
}

func InitEmptyBlockfuncionsContext(p *BlockfuncionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_blockfuncions
}

func (*BlockfuncionsContext) IsBlockfuncionsContext() {}

func NewBlockfuncionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockfuncionsContext {
	var p = new(BlockfuncionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_blockfuncions

	return p
}

func (s *BlockfuncionsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockfuncionsContext) Get_bloquefunciones() IBloquefuncionesContext {
	return s._bloquefunciones
}

func (s *BlockfuncionsContext) Set_bloquefunciones(v IBloquefuncionesContext) { s._bloquefunciones = v }

func (s *BlockfuncionsContext) GetFunc_() []IBloquefuncionesContext { return s.func_ }

func (s *BlockfuncionsContext) SetFunc_(v []IBloquefuncionesContext) { s.func_ = v }

func (s *BlockfuncionsContext) GetBlkfunc() []interface{} { return s.blkfunc }

func (s *BlockfuncionsContext) SetBlkfunc(v []interface{}) { s.blkfunc = v }

func (s *BlockfuncionsContext) AllBloquefunciones() []IBloquefuncionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBloquefuncionesContext); ok {
			len++
		}
	}

	tst := make([]IBloquefuncionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBloquefuncionesContext); ok {
			tst[i] = t.(IBloquefuncionesContext)
			i++
		}
	}

	return tst
}

func (s *BlockfuncionsContext) Bloquefunciones(i int) IBloquefuncionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloquefuncionesContext); ok {
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

	return t.(IBloquefuncionesContext)
}

func (s *BlockfuncionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockfuncionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockfuncionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterBlockfuncions(s)
	}
}

func (s *BlockfuncionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitBlockfuncions(s)
	}
}

func (p *C3DGrammarParser) Blockfuncions() (localctx IBlockfuncionsContext) {
	localctx = NewBlockfuncionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, C3DGrammarParserRULE_blockfuncions)

	localctx.(*BlockfuncionsContext).blkfunc = []interface{}{}
	var listFunc []IBloquefuncionesContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == C3DGrammarParserVOID {
		{
			p.SetState(86)

			var _x = p.Bloquefunciones()

			localctx.(*BlockfuncionsContext)._bloquefunciones = _x
		}
		localctx.(*BlockfuncionsContext).func_ = append(localctx.(*BlockfuncionsContext).func_, localctx.(*BlockfuncionsContext)._bloquefunciones)

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listFunc = localctx.(*BlockfuncionsContext).GetFunc_()
	for _, e := range listFunc {
		localctx.(*BlockfuncionsContext).blkfunc = append(localctx.(*BlockfuncionsContext).blkfunc, e.GetFunc_())
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

// IBloquefuncionesContext is an interface to support dynamic dispatch.
type IBloquefuncionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetFunc_ returns the func_ attribute.
	GetFunc_() interfacesc3d.Instruction

	// SetFunc_ sets the func_ attribute.
	SetFunc_(interfacesc3d.Instruction)

	// Getter signatures
	VOID() antlr.TerminalNode
	ID_VALIDO() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsBloquefuncionesContext differentiates from other interfaces.
	IsBloquefuncionesContext()
}

type BloquefuncionesContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	func_      interfacesc3d.Instruction
	_ID_VALIDO antlr.Token
	_block     IBlockContext
}

func NewEmptyBloquefuncionesContext() *BloquefuncionesContext {
	var p = new(BloquefuncionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_bloquefunciones
	return p
}

func InitEmptyBloquefuncionesContext(p *BloquefuncionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_bloquefunciones
}

func (*BloquefuncionesContext) IsBloquefuncionesContext() {}

func NewBloquefuncionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloquefuncionesContext {
	var p = new(BloquefuncionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_bloquefunciones

	return p
}

func (s *BloquefuncionesContext) GetParser() antlr.Parser { return s.parser }

func (s *BloquefuncionesContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *BloquefuncionesContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *BloquefuncionesContext) Get_block() IBlockContext { return s._block }

func (s *BloquefuncionesContext) Set_block(v IBlockContext) { s._block = v }

func (s *BloquefuncionesContext) GetFunc_() interfacesc3d.Instruction { return s.func_ }

func (s *BloquefuncionesContext) SetFunc_(v interfacesc3d.Instruction) { s.func_ = v }

func (s *BloquefuncionesContext) VOID() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserVOID, 0)
}

func (s *BloquefuncionesContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserID_VALIDO, 0)
}

func (s *BloquefuncionesContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPARIZQ, 0)
}

func (s *BloquefuncionesContext) PARDER() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPARDER, 0)
}

func (s *BloquefuncionesContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserLLAVEIZQ, 0)
}

func (s *BloquefuncionesContext) Block() IBlockContext {
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

func (s *BloquefuncionesContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserLLAVEDER, 0)
}

func (s *BloquefuncionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloquefuncionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloquefuncionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterBloquefunciones(s)
	}
}

func (s *BloquefuncionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitBloquefunciones(s)
	}
}

func (p *C3DGrammarParser) Bloquefunciones() (localctx IBloquefuncionesContext) {
	localctx = NewBloquefuncionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, C3DGrammarParserRULE_bloquefunciones)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(C3DGrammarParserVOID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)

		var _m = p.Match(C3DGrammarParserID_VALIDO)

		localctx.(*BloquefuncionesContext)._ID_VALIDO = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(C3DGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(C3DGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(C3DGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)

		var _x = p.Block()

		localctx.(*BloquefuncionesContext)._block = _x
	}
	{
		p.SetState(99)
		p.Match(C3DGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*BloquefuncionesContext).func_ = instructionsc3d.NewFuncionVoid((func() string {
		if localctx.(*BloquefuncionesContext).Get_ID_VALIDO() == nil {
			return ""
		} else {
			return localctx.(*BloquefuncionesContext).Get_ID_VALIDO().GetText()
		}
	}()), localctx.(*BloquefuncionesContext).Get_block().GetBlk())

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

// IFuncionmainContext is an interface to support dynamic dispatch.
type IFuncionmainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetFunmain returns the funmain attribute.
	GetFunmain() interfacesc3d.Instruction

	// SetFunmain sets the funmain attribute.
	SetFunmain(interfacesc3d.Instruction)

	// Getter signatures
	INT() antlr.TerminalNode
	ID_VALIDO() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	RETURN() antlr.TerminalNode
	NUMERO() antlr.TerminalNode
	PUNTOCOMA() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode

	// IsFuncionmainContext differentiates from other interfaces.
	IsFuncionmainContext()
}

type FuncionmainContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	funmain interfacesc3d.Instruction
	_block  IBlockContext
}

func NewEmptyFuncionmainContext() *FuncionmainContext {
	var p = new(FuncionmainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_funcionmain
	return p
}

func InitEmptyFuncionmainContext(p *FuncionmainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_funcionmain
}

func (*FuncionmainContext) IsFuncionmainContext() {}

func NewFuncionmainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionmainContext {
	var p = new(FuncionmainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_funcionmain

	return p
}

func (s *FuncionmainContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionmainContext) Get_block() IBlockContext { return s._block }

func (s *FuncionmainContext) Set_block(v IBlockContext) { s._block = v }

func (s *FuncionmainContext) GetFunmain() interfacesc3d.Instruction { return s.funmain }

func (s *FuncionmainContext) SetFunmain(v interfacesc3d.Instruction) { s.funmain = v }

func (s *FuncionmainContext) INT() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserINT, 0)
}

func (s *FuncionmainContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserID_VALIDO, 0)
}

func (s *FuncionmainContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPARIZQ, 0)
}

func (s *FuncionmainContext) PARDER() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPARDER, 0)
}

func (s *FuncionmainContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserLLAVEIZQ, 0)
}

func (s *FuncionmainContext) Block() IBlockContext {
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

func (s *FuncionmainContext) RETURN() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserRETURN, 0)
}

func (s *FuncionmainContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserNUMERO, 0)
}

func (s *FuncionmainContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPUNTOCOMA, 0)
}

func (s *FuncionmainContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserLLAVEDER, 0)
}

func (s *FuncionmainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionmainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionmainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterFuncionmain(s)
	}
}

func (s *FuncionmainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitFuncionmain(s)
	}
}

func (p *C3DGrammarParser) Funcionmain() (localctx IFuncionmainContext) {
	localctx = NewFuncionmainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, C3DGrammarParserRULE_funcionmain)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(C3DGrammarParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(C3DGrammarParserID_VALIDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(C3DGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(C3DGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(C3DGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)

		var _x = p.Block()

		localctx.(*FuncionmainContext)._block = _x
	}
	{
		p.SetState(108)
		p.Match(C3DGrammarParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(C3DGrammarParserNUMERO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(C3DGrammarParserPUNTOCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(C3DGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*FuncionmainContext).funmain = instructionsc3d.NewFuncionMain(localctx.(*FuncionmainContext).Get_block().GetBlk())

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

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IInstructionContext

	// SetIns sets the ins rule context list.
	SetIns([]IInstructionContext)

	// GetBlk returns the blk attribute.
	GetBlk() []interface{}

	// SetBlk sets the blk attribute.
	SetBlk([]interface{})

	// Getter signatures
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blk          []interface{}
	_instruction IInstructionContext
	ins          []IInstructionContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *BlockContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *BlockContext) GetIns() []IInstructionContext { return s.ins }

func (s *BlockContext) SetIns(v []IInstructionContext) { s.ins = v }

func (s *BlockContext) GetBlk() []interface{} { return s.blk }

func (s *BlockContext) SetBlk(v []interface{}) { s.blk = v }

func (s *BlockContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
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

	return t.(IInstructionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *C3DGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, C3DGrammarParserRULE_block)

	localctx.(*BlockContext).blk = []interface{}{}
	var listBlo []IInstructionContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&320256) != 0) {
		{
			p.SetState(114)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listBlo = localctx.(*BlockContext).GetIns()
	for _, e := range listBlo {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInstr())
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

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_head_op returns the _head_op rule contexts.
	Get_head_op() IHead_opContext

	// Get_stack_op returns the _stack_op rule contexts.
	Get_stack_op() IStack_opContext

	// Get_printff returns the _printff rule contexts.
	Get_printff() IPrintffContext

	// Get_operaritme returns the _operaritme rule contexts.
	Get_operaritme() IOperaritmeContext

	// Set_head_op sets the _head_op rule contexts.
	Set_head_op(IHead_opContext)

	// Set_stack_op sets the _stack_op rule contexts.
	Set_stack_op(IStack_opContext)

	// Set_printff sets the _printff rule contexts.
	Set_printff(IPrintffContext)

	// Set_operaritme sets the _operaritme rule contexts.
	Set_operaritme(IOperaritmeContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfacesc3d.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfacesc3d.Instruction)

	// Getter signatures
	Head_op() IHead_opContext
	Stack_op() IStack_opContext
	Printff() IPrintffContext
	Operaritme() IOperaritmeContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       interfacesc3d.Instruction
	_head_op    IHead_opContext
	_stack_op   IStack_opContext
	_printff    IPrintffContext
	_operaritme IOperaritmeContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_head_op() IHead_opContext { return s._head_op }

func (s *InstructionContext) Get_stack_op() IStack_opContext { return s._stack_op }

func (s *InstructionContext) Get_printff() IPrintffContext { return s._printff }

func (s *InstructionContext) Get_operaritme() IOperaritmeContext { return s._operaritme }

func (s *InstructionContext) Set_head_op(v IHead_opContext) { s._head_op = v }

func (s *InstructionContext) Set_stack_op(v IStack_opContext) { s._stack_op = v }

func (s *InstructionContext) Set_printff(v IPrintffContext) { s._printff = v }

func (s *InstructionContext) Set_operaritme(v IOperaritmeContext) { s._operaritme = v }

func (s *InstructionContext) GetInstr() interfacesc3d.Instruction { return s.instr }

func (s *InstructionContext) SetInstr(v interfacesc3d.Instruction) { s.instr = v }

func (s *InstructionContext) Head_op() IHead_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHead_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHead_opContext)
}

func (s *InstructionContext) Stack_op() IStack_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStack_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStack_opContext)
}

func (s *InstructionContext) Printff() IPrintffContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintffContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintffContext)
}

func (s *InstructionContext) Operaritme() IOperaritmeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperaritmeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperaritmeContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *C3DGrammarParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, C3DGrammarParserRULE_instruction)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)

			var _x = p.Head_op()

			localctx.(*InstructionContext)._head_op = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_head_op().GetHeapop()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)

			var _x = p.Stack_op()

			localctx.(*InstructionContext)._stack_op = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_stack_op().GetStackop()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)

			var _x = p.Printff()

			localctx.(*InstructionContext)._printff = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_printff().GetPrtff()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)

			var _x = p.Operaritme()

			localctx.(*InstructionContext)._operaritme = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_operaritme().GetOparit()

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

// IHead_opContext is an interface to support dynamic dispatch.
type IHead_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Get_embebida returns the _embebida rule contexts.
	Get_embebida() IEmbebidaContext

	// Set_embebida sets the _embebida rule contexts.
	Set_embebida(IEmbebidaContext)

	// GetHeapop returns the heapop attribute.
	GetHeapop() interfacesc3d.Instruction

	// SetHeapop sets the heapop attribute.
	SetHeapop(interfacesc3d.Instruction)

	// Getter signatures
	AllPHEAD() []antlr.TerminalNode
	PHEAD(i int) antlr.TerminalNode
	IG() antlr.TerminalNode
	NUMERO() antlr.TerminalNode
	PUNTOCOMA() antlr.TerminalNode
	ID_VALIDO() antlr.TerminalNode
	HEAP() antlr.TerminalNode
	CORCHIZQ() antlr.TerminalNode
	Embebida() IEmbebidaContext
	CORCHDER() antlr.TerminalNode
	ADD() antlr.TerminalNode

	// IsHead_opContext differentiates from other interfaces.
	IsHead_opContext()
}

type Head_opContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	heapop     interfacesc3d.Instruction
	_NUMERO    antlr.Token
	_ID_VALIDO antlr.Token
	_embebida  IEmbebidaContext
	op         antlr.Token
}

func NewEmptyHead_opContext() *Head_opContext {
	var p = new(Head_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_head_op
	return p
}

func InitEmptyHead_opContext(p *Head_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_head_op
}

func (*Head_opContext) IsHead_opContext() {}

func NewHead_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Head_opContext {
	var p = new(Head_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_head_op

	return p
}

func (s *Head_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Head_opContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *Head_opContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *Head_opContext) GetOp() antlr.Token { return s.op }

func (s *Head_opContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *Head_opContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *Head_opContext) SetOp(v antlr.Token) { s.op = v }

func (s *Head_opContext) Get_embebida() IEmbebidaContext { return s._embebida }

func (s *Head_opContext) Set_embebida(v IEmbebidaContext) { s._embebida = v }

func (s *Head_opContext) GetHeapop() interfacesc3d.Instruction { return s.heapop }

func (s *Head_opContext) SetHeapop(v interfacesc3d.Instruction) { s.heapop = v }

func (s *Head_opContext) AllPHEAD() []antlr.TerminalNode {
	return s.GetTokens(C3DGrammarParserPHEAD)
}

func (s *Head_opContext) PHEAD(i int) antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPHEAD, i)
}

func (s *Head_opContext) IG() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserIG, 0)
}

func (s *Head_opContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserNUMERO, 0)
}

func (s *Head_opContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPUNTOCOMA, 0)
}

func (s *Head_opContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserID_VALIDO, 0)
}

func (s *Head_opContext) HEAP() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserHEAP, 0)
}

func (s *Head_opContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCORCHIZQ, 0)
}

func (s *Head_opContext) Embebida() IEmbebidaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmbebidaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmbebidaContext)
}

func (s *Head_opContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCORCHDER, 0)
}

func (s *Head_opContext) ADD() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserADD, 0)
}

func (s *Head_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Head_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Head_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterHead_op(s)
	}
}

func (s *Head_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitHead_op(s)
	}
}

func (p *C3DGrammarParser) Head_op() (localctx IHead_opContext) {
	localctx = NewHead_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, C3DGrammarParserRULE_head_op)
	var _la int

	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(C3DGrammarParserPHEAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(C3DGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)

			var _m = p.Match(C3DGrammarParserNUMERO)

			localctx.(*Head_opContext)._NUMERO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(C3DGrammarParserPUNTOCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Head_opContext).heapop = instructionsc3d.NewHeapop1((func() string {
			if localctx.(*Head_opContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*Head_opContext).Get_NUMERO().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)

			var _m = p.Match(C3DGrammarParserID_VALIDO)

			localctx.(*Head_opContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(C3DGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(C3DGrammarParserPHEAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(C3DGrammarParserPUNTOCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Head_opContext).heapop = instructionsc3d.NewHeapop2((func() string {
			if localctx.(*Head_opContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*Head_opContext).Get_ID_VALIDO().GetText()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)
			p.Match(C3DGrammarParserHEAP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(C3DGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)

			var _x = p.Embebida()

			localctx.(*Head_opContext)._embebida = _x
		}
		{
			p.SetState(148)
			p.Match(C3DGrammarParserPHEAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(C3DGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(C3DGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Head_opContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == C3DGrammarParserNUMERO || _la == C3DGrammarParserID_VALIDO) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Head_opContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(152)
			p.Match(C3DGrammarParserPUNTOCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Head_opContext).heapop = instructionsc3d.NewHeapop3(localctx.(*Head_opContext).Get_embebida().GetEmbe(), (func() string {
			if localctx.(*Head_opContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Head_opContext).GetOp().GetText()
			}
		}()))

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(155)
			p.Match(C3DGrammarParserPHEAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Match(C3DGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(C3DGrammarParserPHEAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(C3DGrammarParserADD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)

			var _m = p.Match(C3DGrammarParserNUMERO)

			localctx.(*Head_opContext)._NUMERO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(C3DGrammarParserPUNTOCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Head_opContext).heapop = instructionsc3d.NewHeapop4((func() string {
			if localctx.(*Head_opContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*Head_opContext).Get_NUMERO().GetText()
			}
		}()))

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

// IStack_opContext is an interface to support dynamic dispatch.
type IStack_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// GetOp1 returns the op1 token.
	GetOp1() antlr.Token

	// GetOp2 returns the op2 token.
	GetOp2() antlr.Token

	// Get_PARIZQ returns the _PARIZQ token.
	Get_PARIZQ() antlr.Token

	// GetOp3 returns the op3 token.
	GetOp3() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// SetOp1 sets the op1 token.
	SetOp1(antlr.Token)

	// SetOp2 sets the op2 token.
	SetOp2(antlr.Token)

	// Set_PARIZQ sets the _PARIZQ token.
	Set_PARIZQ(antlr.Token)

	// SetOp3 sets the op3 token.
	SetOp3(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetTip1 returns the tip1 rule contexts.
	GetTip1() IEmbebidaContext

	// Get_tipodata returns the _tipodata rule contexts.
	Get_tipodata() ITipodataContext

	// SetTip1 sets the tip1 rule contexts.
	SetTip1(IEmbebidaContext)

	// Set_tipodata sets the _tipodata rule contexts.
	Set_tipodata(ITipodataContext)

	// GetStackop returns the stackop attribute.
	GetStackop() interfacesc3d.Instruction

	// SetStackop sets the stackop attribute.
	SetStackop(interfacesc3d.Instruction)

	// Getter signatures
	AllPSTACK() []antlr.TerminalNode
	PSTACK(i int) antlr.TerminalNode
	IG() antlr.TerminalNode
	AllNUMERO() []antlr.TerminalNode
	NUMERO(i int) antlr.TerminalNode
	PUNTOCOMA() antlr.TerminalNode
	AllID_VALIDO() []antlr.TerminalNode
	ID_VALIDO(i int) antlr.TerminalNode
	STACK() antlr.TerminalNode
	CORCHIZQ() antlr.TerminalNode
	CORCHDER() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Tipodata() ITipodataContext
	PARDER() antlr.TerminalNode
	Embebida() IEmbebidaContext
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode

	// IsStack_opContext differentiates from other interfaces.
	IsStack_opContext()
}

type Stack_opContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	stackop    interfacesc3d.Instruction
	_NUMERO    antlr.Token
	_ID_VALIDO antlr.Token
	op1        antlr.Token
	tip1       IEmbebidaContext
	op2        antlr.Token
	_PARIZQ    antlr.Token
	_tipodata  ITipodataContext
	op3        antlr.Token
	op         antlr.Token
}

func NewEmptyStack_opContext() *Stack_opContext {
	var p = new(Stack_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_stack_op
	return p
}

func InitEmptyStack_opContext(p *Stack_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_stack_op
}

func (*Stack_opContext) IsStack_opContext() {}

func NewStack_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stack_opContext {
	var p = new(Stack_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_stack_op

	return p
}

func (s *Stack_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Stack_opContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *Stack_opContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *Stack_opContext) GetOp1() antlr.Token { return s.op1 }

func (s *Stack_opContext) GetOp2() antlr.Token { return s.op2 }

func (s *Stack_opContext) Get_PARIZQ() antlr.Token { return s._PARIZQ }

func (s *Stack_opContext) GetOp3() antlr.Token { return s.op3 }

func (s *Stack_opContext) GetOp() antlr.Token { return s.op }

func (s *Stack_opContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *Stack_opContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *Stack_opContext) SetOp1(v antlr.Token) { s.op1 = v }

func (s *Stack_opContext) SetOp2(v antlr.Token) { s.op2 = v }

func (s *Stack_opContext) Set_PARIZQ(v antlr.Token) { s._PARIZQ = v }

func (s *Stack_opContext) SetOp3(v antlr.Token) { s.op3 = v }

func (s *Stack_opContext) SetOp(v antlr.Token) { s.op = v }

func (s *Stack_opContext) GetTip1() IEmbebidaContext { return s.tip1 }

func (s *Stack_opContext) Get_tipodata() ITipodataContext { return s._tipodata }

func (s *Stack_opContext) SetTip1(v IEmbebidaContext) { s.tip1 = v }

func (s *Stack_opContext) Set_tipodata(v ITipodataContext) { s._tipodata = v }

func (s *Stack_opContext) GetStackop() interfacesc3d.Instruction { return s.stackop }

func (s *Stack_opContext) SetStackop(v interfacesc3d.Instruction) { s.stackop = v }

func (s *Stack_opContext) AllPSTACK() []antlr.TerminalNode {
	return s.GetTokens(C3DGrammarParserPSTACK)
}

func (s *Stack_opContext) PSTACK(i int) antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPSTACK, i)
}

func (s *Stack_opContext) IG() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserIG, 0)
}

func (s *Stack_opContext) AllNUMERO() []antlr.TerminalNode {
	return s.GetTokens(C3DGrammarParserNUMERO)
}

func (s *Stack_opContext) NUMERO(i int) antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserNUMERO, i)
}

func (s *Stack_opContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPUNTOCOMA, 0)
}

func (s *Stack_opContext) AllID_VALIDO() []antlr.TerminalNode {
	return s.GetTokens(C3DGrammarParserID_VALIDO)
}

func (s *Stack_opContext) ID_VALIDO(i int) antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserID_VALIDO, i)
}

func (s *Stack_opContext) STACK() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserSTACK, 0)
}

func (s *Stack_opContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCORCHIZQ, 0)
}

func (s *Stack_opContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCORCHDER, 0)
}

func (s *Stack_opContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPARIZQ, 0)
}

func (s *Stack_opContext) Tipodata() ITipodataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipodataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipodataContext)
}

func (s *Stack_opContext) PARDER() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPARDER, 0)
}

func (s *Stack_opContext) Embebida() IEmbebidaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmbebidaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmbebidaContext)
}

func (s *Stack_opContext) ADD() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserADD, 0)
}

func (s *Stack_opContext) SUB() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserSUB, 0)
}

func (s *Stack_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stack_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stack_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterStack_op(s)
	}
}

func (s *Stack_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitStack_op(s)
	}
}

func (p *C3DGrammarParser) Stack_op() (localctx IStack_opContext) {
	localctx = NewStack_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, C3DGrammarParserRULE_stack_op)
	var _la int

	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Match(C3DGrammarParserPSTACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Match(C3DGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)

			var _m = p.Match(C3DGrammarParserNUMERO)

			localctx.(*Stack_opContext)._NUMERO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Match(C3DGrammarParserPUNTOCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Stack_opContext).stackop = instructionsc3d.NewStackop1((func() string {
			if localctx.(*Stack_opContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*Stack_opContext).Get_NUMERO().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)

			var _m = p.Match(C3DGrammarParserID_VALIDO)

			localctx.(*Stack_opContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Match(C3DGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(C3DGrammarParserPSTACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(C3DGrammarParserPUNTOCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Stack_opContext).stackop = instructionsc3d.NewStackop2((func() string {
			if localctx.(*Stack_opContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*Stack_opContext).Get_ID_VALIDO().GetText()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(174)
			p.Match(C3DGrammarParserSTACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Match(C3DGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case C3DGrammarParserNUMERO:
			{
				p.SetState(176)

				var _m = p.Match(C3DGrammarParserNUMERO)

				localctx.(*Stack_opContext).op1 = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case C3DGrammarParserPARIZQ:
			{
				p.SetState(177)

				var _x = p.Embebida()

				localctx.(*Stack_opContext).tip1 = _x
			}
			{
				p.SetState(178)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Stack_opContext).op2 = _lt

				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&360448) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Stack_opContext).op2 = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(182)
			p.Match(C3DGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Match(C3DGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == C3DGrammarParserPARIZQ {
			{
				p.SetState(184)

				var _m = p.Match(C3DGrammarParserPARIZQ)

				localctx.(*Stack_opContext)._PARIZQ = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(185)

				var _x = p.Tipodata()

				localctx.(*Stack_opContext)._tipodata = _x
			}
			{
				p.SetState(186)
				p.Match(C3DGrammarParserPARDER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(190)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Stack_opContext).op3 = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == C3DGrammarParserNUMERO || _la == C3DGrammarParserID_VALIDO) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Stack_opContext).op3 = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(191)
			p.Match(C3DGrammarParserPUNTOCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if localctx.(*Stack_opContext).GetOp1() != nil {
			if localctx.(*Stack_opContext).Get_PARIZQ() != nil {
				localctx.(*Stack_opContext).stackop = instructionsc3d.NewStackop31((func() string {
					if localctx.(*Stack_opContext).GetOp1() == nil {
						return ""
					} else {
						return localctx.(*Stack_opContext).GetOp1().GetText()
					}
				}()), localctx.(*Stack_opContext).Get_tipodata().GetTipdat(), (func() string {
					if localctx.(*Stack_opContext).GetOp3() == nil {
						return ""
					} else {
						return localctx.(*Stack_opContext).GetOp3().GetText()
					}
				}()))
			} else {
				localctx.(*Stack_opContext).stackop = instructionsc3d.NewStackop32((func() string {
					if localctx.(*Stack_opContext).GetOp1() == nil {
						return ""
					} else {
						return localctx.(*Stack_opContext).GetOp1().GetText()
					}
				}()), (func() string {
					if localctx.(*Stack_opContext).GetOp3() == nil {
						return ""
					} else {
						return localctx.(*Stack_opContext).GetOp3().GetText()
					}
				}()))
			}
		} else {
			if localctx.(*Stack_opContext).Get_PARIZQ() != nil {
				localctx.(*Stack_opContext).stackop = instructionsc3d.NewStackop33(localctx.(*Stack_opContext).GetTip1().GetEmbe(), (func() string {
					if localctx.(*Stack_opContext).GetOp2() == nil {
						return ""
					} else {
						return localctx.(*Stack_opContext).GetOp2().GetText()
					}
				}()), localctx.(*Stack_opContext).Get_tipodata().GetTipdat(), (func() string {
					if localctx.(*Stack_opContext).GetOp3() == nil {
						return ""
					} else {
						return localctx.(*Stack_opContext).GetOp3().GetText()
					}
				}()))
			} else {
				localctx.(*Stack_opContext).stackop = instructionsc3d.NewStackop34(localctx.(*Stack_opContext).GetTip1().GetEmbe(), (func() string {
					if localctx.(*Stack_opContext).GetOp2() == nil {
						return ""
					} else {
						return localctx.(*Stack_opContext).GetOp2().GetText()
					}
				}()), (func() string {
					if localctx.(*Stack_opContext).GetOp3() == nil {
						return ""
					} else {
						return localctx.(*Stack_opContext).GetOp3().GetText()
					}
				}()))
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(193)
			p.Match(C3DGrammarParserPSTACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Match(C3DGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Match(C3DGrammarParserPSTACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Stack_opContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == C3DGrammarParserADD || _la == C3DGrammarParserSUB) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Stack_opContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(197)

			var _m = p.Match(C3DGrammarParserNUMERO)

			localctx.(*Stack_opContext)._NUMERO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(C3DGrammarParserPUNTOCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*Stack_opContext).stackop = instructionsc3d.NewStackop4((func() string {
			if localctx.(*Stack_opContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Stack_opContext).GetOp().GetText()
			}
		}()), (func() string {
			if localctx.(*Stack_opContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*Stack_opContext).Get_NUMERO().GetText()
			}
		}()))

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

// IPrintffContext is an interface to support dynamic dispatch.
type IPrintffContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Get_embebida returns the _embebida rule contexts.
	Get_embebida() IEmbebidaContext

	// Set_embebida sets the _embebida rule contexts.
	Set_embebida(IEmbebidaContext)

	// GetPrtff returns the prtff attribute.
	GetPrtff() interfacesc3d.Instruction

	// SetPrtff sets the prtff attribute.
	SetPrtff(interfacesc3d.Instruction)

	// Getter signatures
	PRINTF() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	CADENA() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	NUMERO() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	PUNTOCOMA() antlr.TerminalNode
	Embebida() IEmbebidaContext
	ID_VALIDO() antlr.TerminalNode

	// IsPrintffContext differentiates from other interfaces.
	IsPrintffContext()
}

type PrintffContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	prtff     interfacesc3d.Instruction
	_CADENA   antlr.Token
	_NUMERO   antlr.Token
	_embebida IEmbebidaContext
	op        antlr.Token
}

func NewEmptyPrintffContext() *PrintffContext {
	var p = new(PrintffContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_printff
	return p
}

func InitEmptyPrintffContext(p *PrintffContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_printff
}

func (*PrintffContext) IsPrintffContext() {}

func NewPrintffContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintffContext {
	var p = new(PrintffContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_printff

	return p
}

func (s *PrintffContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintffContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *PrintffContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *PrintffContext) GetOp() antlr.Token { return s.op }

func (s *PrintffContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrintffContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *PrintffContext) SetOp(v antlr.Token) { s.op = v }

func (s *PrintffContext) Get_embebida() IEmbebidaContext { return s._embebida }

func (s *PrintffContext) Set_embebida(v IEmbebidaContext) { s._embebida = v }

func (s *PrintffContext) GetPrtff() interfacesc3d.Instruction { return s.prtff }

func (s *PrintffContext) SetPrtff(v interfacesc3d.Instruction) { s.prtff = v }

func (s *PrintffContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPRINTF, 0)
}

func (s *PrintffContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPARIZQ, 0)
}

func (s *PrintffContext) CADENA() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCADENA, 0)
}

func (s *PrintffContext) COMMA() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCOMMA, 0)
}

func (s *PrintffContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserNUMERO, 0)
}

func (s *PrintffContext) PARDER() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPARDER, 0)
}

func (s *PrintffContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPUNTOCOMA, 0)
}

func (s *PrintffContext) Embebida() IEmbebidaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmbebidaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmbebidaContext)
}

func (s *PrintffContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserID_VALIDO, 0)
}

func (s *PrintffContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintffContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintffContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterPrintff(s)
	}
}

func (s *PrintffContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitPrintff(s)
	}
}

func (p *C3DGrammarParser) Printff() (localctx IPrintffContext) {
	localctx = NewPrintffContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, C3DGrammarParserRULE_printff)
	var _la int

	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Match(C3DGrammarParserPRINTF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(C3DGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)

			var _m = p.Match(C3DGrammarParserCADENA)

			localctx.(*PrintffContext)._CADENA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.Match(C3DGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)

			var _m = p.Match(C3DGrammarParserNUMERO)

			localctx.(*PrintffContext)._NUMERO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Match(C3DGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(C3DGrammarParserPUNTOCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*PrintffContext).prtff = instructionsc3d.NewPrint1((func() string {
			if localctx.(*PrintffContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrintffContext).Get_CADENA().GetText()
			}
		}()), (func() string {
			if localctx.(*PrintffContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*PrintffContext).Get_NUMERO().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.Match(C3DGrammarParserPRINTF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Match(C3DGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)

			var _m = p.Match(C3DGrammarParserCADENA)

			localctx.(*PrintffContext)._CADENA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(C3DGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)

			var _x = p.Embebida()

			localctx.(*PrintffContext)._embebida = _x
		}
		{
			p.SetState(215)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*PrintffContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == C3DGrammarParserNUMERO || _la == C3DGrammarParserID_VALIDO) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*PrintffContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(216)
			p.Match(C3DGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Match(C3DGrammarParserPUNTOCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*PrintffContext).prtff = instructionsc3d.NewPrint2((func() string {
			if localctx.(*PrintffContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrintffContext).Get_CADENA().GetText()
			}
		}()), localctx.(*PrintffContext).Get_embebida().GetEmbe(), (func() string {
			if localctx.(*PrintffContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*PrintffContext).GetOp().GetText()
			}
		}()))

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

// IEmbebidaContext is an interface to support dynamic dispatch.
type IEmbebidaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEmbe returns the embe attribute.
	GetEmbe() string

	// SetEmbe sets the embe attribute.
	SetEmbe(string)

	// Getter signatures
	PARIZQ() antlr.TerminalNode
	INT() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	CHAR() antlr.TerminalNode

	// IsEmbebidaContext differentiates from other interfaces.
	IsEmbebidaContext()
}

type EmbebidaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	embe   string
}

func NewEmptyEmbebidaContext() *EmbebidaContext {
	var p = new(EmbebidaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_embebida
	return p
}

func InitEmptyEmbebidaContext(p *EmbebidaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_embebida
}

func (*EmbebidaContext) IsEmbebidaContext() {}

func NewEmbebidaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmbebidaContext {
	var p = new(EmbebidaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_embebida

	return p
}

func (s *EmbebidaContext) GetParser() antlr.Parser { return s.parser }

func (s *EmbebidaContext) GetEmbe() string { return s.embe }

func (s *EmbebidaContext) SetEmbe(v string) { s.embe = v }

func (s *EmbebidaContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPARIZQ, 0)
}

func (s *EmbebidaContext) INT() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserINT, 0)
}

func (s *EmbebidaContext) PARDER() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPARDER, 0)
}

func (s *EmbebidaContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserFLOAT, 0)
}

func (s *EmbebidaContext) CHAR() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCHAR, 0)
}

func (s *EmbebidaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmbebidaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmbebidaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterEmbebida(s)
	}
}

func (s *EmbebidaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitEmbebida(s)
	}
}

func (p *C3DGrammarParser) Embebida() (localctx IEmbebidaContext) {
	localctx = NewEmbebidaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, C3DGrammarParserRULE_embebida)
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.Match(C3DGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Match(C3DGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(C3DGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := "(int)"
		localctx.(*EmbebidaContext).embe = str

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.Match(C3DGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Match(C3DGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(C3DGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := "(float)"
		localctx.(*EmbebidaContext).embe = str

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.Match(C3DGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(C3DGrammarParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(C3DGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := "(char)"
		localctx.(*EmbebidaContext).embe = str

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

// ITipodataContext is an interface to support dynamic dispatch.
type ITipodataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipdat returns the tipdat attribute.
	GetTipdat() string

	// SetTipdat sets the tipdat attribute.
	SetTipdat(string)

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	CHAR() antlr.TerminalNode

	// IsTipodataContext differentiates from other interfaces.
	IsTipodataContext()
}

type TipodataContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	tipdat string
}

func NewEmptyTipodataContext() *TipodataContext {
	var p = new(TipodataContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_tipodata
	return p
}

func InitEmptyTipodataContext(p *TipodataContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_tipodata
}

func (*TipodataContext) IsTipodataContext() {}

func NewTipodataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipodataContext {
	var p = new(TipodataContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_tipodata

	return p
}

func (s *TipodataContext) GetParser() antlr.Parser { return s.parser }

func (s *TipodataContext) GetTipdat() string { return s.tipdat }

func (s *TipodataContext) SetTipdat(v string) { s.tipdat = v }

func (s *TipodataContext) INT() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserINT, 0)
}

func (s *TipodataContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserFLOAT, 0)
}

func (s *TipodataContext) CHAR() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserCHAR, 0)
}

func (s *TipodataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipodataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipodataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterTipodata(s)
	}
}

func (s *TipodataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitTipodata(s)
	}
}

func (p *C3DGrammarParser) Tipodata() (localctx ITipodataContext) {
	localctx = NewTipodataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, C3DGrammarParserRULE_tipodata)
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case C3DGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(236)
			p.Match(C3DGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := "int"
		localctx.(*TipodataContext).tipdat = str

	case C3DGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Match(C3DGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := "float"
		localctx.(*TipodataContext).tipdat = str

	case C3DGrammarParserCHAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(240)
			p.Match(C3DGrammarParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := "char"
		localctx.(*TipodataContext).tipdat = str

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

// IOperaritmeContext is an interface to support dynamic dispatch.
type IOperaritmeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOparit returns the oparit attribute.
	GetOparit() interfacesc3d.Instruction

	// SetOparit sets the oparit attribute.
	SetOparit(interfacesc3d.Instruction)

	// Getter signatures
	AllID_VALIDO() []antlr.TerminalNode
	ID_VALIDO(i int) antlr.TerminalNode
	IG() antlr.TerminalNode
	PUNTOCOMA() antlr.TerminalNode
	NUMERO() antlr.TerminalNode

	// IsOperaritmeContext differentiates from other interfaces.
	IsOperaritmeContext()
}

type OperaritmeContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	oparit     interfacesc3d.Instruction
	_ID_VALIDO antlr.Token
	op         antlr.Token
}

func NewEmptyOperaritmeContext() *OperaritmeContext {
	var p = new(OperaritmeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_operaritme
	return p
}

func InitEmptyOperaritmeContext(p *OperaritmeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = C3DGrammarParserRULE_operaritme
}

func (*OperaritmeContext) IsOperaritmeContext() {}

func NewOperaritmeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperaritmeContext {
	var p = new(OperaritmeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = C3DGrammarParserRULE_operaritme

	return p
}

func (s *OperaritmeContext) GetParser() antlr.Parser { return s.parser }

func (s *OperaritmeContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *OperaritmeContext) GetOp() antlr.Token { return s.op }

func (s *OperaritmeContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *OperaritmeContext) SetOp(v antlr.Token) { s.op = v }

func (s *OperaritmeContext) GetOparit() interfacesc3d.Instruction { return s.oparit }

func (s *OperaritmeContext) SetOparit(v interfacesc3d.Instruction) { s.oparit = v }

func (s *OperaritmeContext) AllID_VALIDO() []antlr.TerminalNode {
	return s.GetTokens(C3DGrammarParserID_VALIDO)
}

func (s *OperaritmeContext) ID_VALIDO(i int) antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserID_VALIDO, i)
}

func (s *OperaritmeContext) IG() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserIG, 0)
}

func (s *OperaritmeContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserPUNTOCOMA, 0)
}

func (s *OperaritmeContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(C3DGrammarParserNUMERO, 0)
}

func (s *OperaritmeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperaritmeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperaritmeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.EnterOperaritme(s)
	}
}

func (s *OperaritmeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(C3DGrammarListener); ok {
		listenerT.ExitOperaritme(s)
	}
}

func (p *C3DGrammarParser) Operaritme() (localctx IOperaritmeContext) {
	localctx = NewOperaritmeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, C3DGrammarParserRULE_operaritme)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)

		var _m = p.Match(C3DGrammarParserID_VALIDO)

		localctx.(*OperaritmeContext)._ID_VALIDO = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Match(C3DGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OperaritmeContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == C3DGrammarParserNUMERO || _la == C3DGrammarParserID_VALIDO) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OperaritmeContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(247)
		p.Match(C3DGrammarParserPUNTOCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*OperaritmeContext).oparit = instructionsc3d.NewOparit1((func() string {
		if localctx.(*OperaritmeContext).Get_ID_VALIDO() == nil {
			return ""
		} else {
			return localctx.(*OperaritmeContext).Get_ID_VALIDO().GetText()
		}
	}()), (func() string {
		if localctx.(*OperaritmeContext).GetOp() == nil {
			return ""
		} else {
			return localctx.(*OperaritmeContext).GetOp().GetText()
		}
	}()))

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
