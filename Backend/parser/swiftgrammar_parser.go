// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "Backend/interfaces"
import "Backend/environment"
import "Backend/expressions"
import "Backend/instructions/datoscompuestos"
import "Backend/instructions/datosprimitivos"
import "Backend/instructions/funciones"
import "Backend/instructions/sentencias"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftGrammarParser struct {
	*antlr.BaseParser
}

var SwiftGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'true'",
		"'false'", "'var'", "'let'", "'nil'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'break'", "'continue'", "'for'", "'in'", "'...'",
		"'while'", "'guard'", "'return'", "'func'", "'print'", "'inout'", "'append'",
		"'remove'", "'removeLast'", "'count'", "'isEmpty'", "'at'", "'repeating'",
		"'struct'", "'mutating'", "'self'", "", "", "", "", "", "'='", "':'",
		"';'", "'?'", "'('", "')'", "'!='", "'=='", "'!'", "'||'", "'&&'", "'>='",
		"'<='", "'>'", "'<'", "'%'", "'*'", "'/'", "'+'", "'-'", "'+='", "'-='",
		"'{'", "'}'", "'->'", "','", "'.'", "'_'", "'['", "']'", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "STRING", "BOOL", "CHARACT", "TRU", "FAL", "VAR",
		"LET", "NULO", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "BREAK", "CONTINUE",
		"FOR", "IN", "RANGO", "WHILE", "GUARD", "RETURN", "FUNCION", "PRINT",
		"INOUT", "APPEND", "REMOVE", "REMOVELAST", "COUNT", "ISEMPTY", "AT",
		"REPEATING", "STRUCT", "MUTATING", "SELF", "NUMBER", "CADENA", "ID_VALIDO",
		"CHARACTER", "WS", "IG", "DOS_PUNTOS", "PUNTOCOMA", "CIERRE_INTE", "PARIZQ",
		"PARDER", "DIF", "IG_IG", "NOT", "OR", "AND", "MAY_IG", "MEN_IG", "MAYOR",
		"MENOR", "MODULO", "MUL", "DIV", "ADD", "SUB", "SUMA", "RESTA", "LLAVEIZQ",
		"LLAVEDER", "RETORNO", "COMA", "PUNTO", "GUIONBAJO", "CORCHIZQ", "CORCHDER",
		"DIRME", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "blockinterno", "instructionint", "declavarible",
		"declaconstante", "asignacionvariable", "tipodato", "expr", "sentenciaifelse",
		"switchcontrol", "blockcase", "bloquecase", "whilecontrol", "forcontrol",
		"guardcontrol", "continuee", "breakk", "retornos", "vectorcontrol",
		"blockparams", "bloqueparams", "vectoragregar", "vectorremover", "vectorvacio",
		"vectorcount", "vectoraccess", "matrizcontrol", "tipomatriz", "defmatriz",
		"listavaloresmat", "listavaloresmat2", "listaexpresions", "listaexpresion",
		"simplematriz", "printstmt",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 75, 716, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 80, 8, 1, 11, 1, 12, 1, 81, 1, 1, 1,
		1, 1, 2, 1, 2, 3, 2, 88, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 94, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 100, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 121, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 133, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 139, 8,
		2, 1, 2, 1, 2, 3, 2, 143, 8, 2, 1, 3, 4, 3, 146, 8, 3, 11, 3, 12, 3, 147,
		1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 154, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		160, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 166, 8, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 187, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 193,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 199, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 205, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 211, 8, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 3, 4, 217, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 223, 8, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 3, 4, 229, 8, 4, 1, 4, 1, 4, 3, 4, 233, 8, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 256, 8, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 3, 6, 272, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 289, 8, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 301, 8, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 339,
		8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 381, 8, 9, 10, 9, 12, 9, 384, 9, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 413, 8, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 422, 8, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 4, 12, 428, 8, 12, 11, 12, 12, 12, 429, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3,
		15, 476, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 3, 19, 498, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 533, 8, 20, 1, 21, 4,
		21, 536, 8, 21, 11, 21, 12, 21, 537, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 549, 8, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 578, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 3, 24, 596, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 28, 1, 28, 3, 28, 618, 8, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 634,
		8, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 3, 31, 647, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 3, 32, 656, 8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 663,
		8, 32, 10, 32, 12, 32, 666, 9, 32, 1, 33, 4, 33, 669, 8, 33, 11, 33, 12,
		33, 670, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		3, 34, 682, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 708, 8, 35, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 0, 2, 18, 64, 37, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 0, 5, 1, 0, 58,
		59, 1, 0, 60, 61, 2, 0, 53, 53, 55, 55, 2, 0, 54, 54, 56, 56, 1, 0, 48,
		49, 775, 0, 74, 1, 0, 0, 0, 2, 79, 1, 0, 0, 0, 4, 142, 1, 0, 0, 0, 6, 145,
		1, 0, 0, 0, 8, 232, 1, 0, 0, 0, 10, 255, 1, 0, 0, 0, 12, 271, 1, 0, 0,
		0, 14, 288, 1, 0, 0, 0, 16, 300, 1, 0, 0, 0, 18, 338, 1, 0, 0, 0, 20, 412,
		1, 0, 0, 0, 22, 414, 1, 0, 0, 0, 24, 427, 1, 0, 0, 0, 26, 433, 1, 0, 0,
		0, 28, 439, 1, 0, 0, 0, 30, 475, 1, 0, 0, 0, 32, 477, 1, 0, 0, 0, 34, 485,
		1, 0, 0, 0, 36, 488, 1, 0, 0, 0, 38, 497, 1, 0, 0, 0, 40, 532, 1, 0, 0,
		0, 42, 535, 1, 0, 0, 0, 44, 548, 1, 0, 0, 0, 46, 577, 1, 0, 0, 0, 48, 595,
		1, 0, 0, 0, 50, 597, 1, 0, 0, 0, 52, 602, 1, 0, 0, 0, 54, 607, 1, 0, 0,
		0, 56, 613, 1, 0, 0, 0, 58, 633, 1, 0, 0, 0, 60, 635, 1, 0, 0, 0, 62, 646,
		1, 0, 0, 0, 64, 655, 1, 0, 0, 0, 66, 668, 1, 0, 0, 0, 68, 681, 1, 0, 0,
		0, 70, 707, 1, 0, 0, 0, 72, 709, 1, 0, 0, 0, 74, 75, 3, 2, 1, 0, 75, 76,
		5, 0, 0, 1, 76, 77, 6, 0, -1, 0, 77, 1, 1, 0, 0, 0, 78, 80, 3, 4, 2, 0,
		79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1,
		0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 6, 1, -1, 0, 84, 3, 1, 0, 0, 0, 85,
		87, 3, 10, 5, 0, 86, 88, 5, 44, 0, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0,
		0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 6, 2, -1, 0, 90, 143, 1, 0, 0, 0, 91,
		93, 3, 12, 6, 0, 92, 94, 5, 44, 0, 0, 93, 92, 1, 0, 0, 0, 93, 94, 1, 0,
		0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 6, 2, -1, 0, 96, 143, 1, 0, 0, 0, 97,
		99, 3, 14, 7, 0, 98, 100, 5, 44, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1,
		0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 6, 2, -1, 0, 102, 143, 1, 0, 0,
		0, 103, 104, 3, 20, 10, 0, 104, 105, 6, 2, -1, 0, 105, 143, 1, 0, 0, 0,
		106, 107, 3, 22, 11, 0, 107, 108, 6, 2, -1, 0, 108, 143, 1, 0, 0, 0, 109,
		110, 3, 28, 14, 0, 110, 111, 6, 2, -1, 0, 111, 143, 1, 0, 0, 0, 112, 113,
		3, 30, 15, 0, 113, 114, 6, 2, -1, 0, 114, 143, 1, 0, 0, 0, 115, 116, 3,
		32, 16, 0, 116, 117, 6, 2, -1, 0, 117, 143, 1, 0, 0, 0, 118, 120, 3, 40,
		20, 0, 119, 121, 5, 44, 0, 0, 120, 119, 1, 0, 0, 0, 120, 121, 1, 0, 0,
		0, 121, 122, 1, 0, 0, 0, 122, 123, 6, 2, -1, 0, 123, 143, 1, 0, 0, 0, 124,
		125, 3, 46, 23, 0, 125, 126, 6, 2, -1, 0, 126, 143, 1, 0, 0, 0, 127, 128,
		3, 48, 24, 0, 128, 129, 6, 2, -1, 0, 129, 143, 1, 0, 0, 0, 130, 132, 3,
		72, 36, 0, 131, 133, 5, 44, 0, 0, 132, 131, 1, 0, 0, 0, 132, 133, 1, 0,
		0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 6, 2, -1, 0, 135, 143, 1, 0, 0, 0,
		136, 138, 3, 56, 28, 0, 137, 139, 5, 44, 0, 0, 138, 137, 1, 0, 0, 0, 138,
		139, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 6, 2, -1, 0, 141, 143,
		1, 0, 0, 0, 142, 85, 1, 0, 0, 0, 142, 91, 1, 0, 0, 0, 142, 97, 1, 0, 0,
		0, 142, 103, 1, 0, 0, 0, 142, 106, 1, 0, 0, 0, 142, 109, 1, 0, 0, 0, 142,
		112, 1, 0, 0, 0, 142, 115, 1, 0, 0, 0, 142, 118, 1, 0, 0, 0, 142, 124,
		1, 0, 0, 0, 142, 127, 1, 0, 0, 0, 142, 130, 1, 0, 0, 0, 142, 136, 1, 0,
		0, 0, 143, 5, 1, 0, 0, 0, 144, 146, 3, 8, 4, 0, 145, 144, 1, 0, 0, 0, 146,
		147, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 150, 6, 3, -1, 0, 150, 7, 1, 0, 0, 0, 151, 153, 3, 10,
		5, 0, 152, 154, 5, 44, 0, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0,
		154, 155, 1, 0, 0, 0, 155, 156, 6, 4, -1, 0, 156, 233, 1, 0, 0, 0, 157,
		159, 3, 12, 6, 0, 158, 160, 5, 44, 0, 0, 159, 158, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 6, 4, -1, 0, 162, 233, 1, 0,
		0, 0, 163, 165, 3, 14, 7, 0, 164, 166, 5, 44, 0, 0, 165, 164, 1, 0, 0,
		0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 6, 4, -1, 0, 168,
		233, 1, 0, 0, 0, 169, 170, 3, 20, 10, 0, 170, 171, 6, 4, -1, 0, 171, 233,
		1, 0, 0, 0, 172, 173, 3, 22, 11, 0, 173, 174, 6, 4, -1, 0, 174, 233, 1,
		0, 0, 0, 175, 176, 3, 28, 14, 0, 176, 177, 6, 4, -1, 0, 177, 233, 1, 0,
		0, 0, 178, 179, 3, 30, 15, 0, 179, 180, 6, 4, -1, 0, 180, 233, 1, 0, 0,
		0, 181, 182, 3, 32, 16, 0, 182, 183, 6, 4, -1, 0, 183, 233, 1, 0, 0, 0,
		184, 186, 3, 34, 17, 0, 185, 187, 5, 44, 0, 0, 186, 185, 1, 0, 0, 0, 186,
		187, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189, 6, 4, -1, 0, 189, 233,
		1, 0, 0, 0, 190, 192, 3, 36, 18, 0, 191, 193, 5, 44, 0, 0, 192, 191, 1,
		0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 6, 4, -1,
		0, 195, 233, 1, 0, 0, 0, 196, 198, 3, 38, 19, 0, 197, 199, 5, 44, 0, 0,
		198, 197, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200,
		201, 6, 4, -1, 0, 201, 233, 1, 0, 0, 0, 202, 204, 3, 40, 20, 0, 203, 205,
		5, 44, 0, 0, 204, 203, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206, 1, 0,
		0, 0, 206, 207, 6, 4, -1, 0, 207, 233, 1, 0, 0, 0, 208, 210, 3, 46, 23,
		0, 209, 211, 5, 44, 0, 0, 210, 209, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211,
		212, 1, 0, 0, 0, 212, 213, 6, 4, -1, 0, 213, 233, 1, 0, 0, 0, 214, 216,
		3, 48, 24, 0, 215, 217, 5, 44, 0, 0, 216, 215, 1, 0, 0, 0, 216, 217, 1,
		0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 219, 6, 4, -1, 0, 219, 233, 1, 0, 0,
		0, 220, 222, 3, 72, 36, 0, 221, 223, 5, 44, 0, 0, 222, 221, 1, 0, 0, 0,
		222, 223, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 6, 4, -1, 0, 225,
		233, 1, 0, 0, 0, 226, 228, 3, 56, 28, 0, 227, 229, 5, 44, 0, 0, 228, 227,
		1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 6, 4,
		-1, 0, 231, 233, 1, 0, 0, 0, 232, 151, 1, 0, 0, 0, 232, 157, 1, 0, 0, 0,
		232, 163, 1, 0, 0, 0, 232, 169, 1, 0, 0, 0, 232, 172, 1, 0, 0, 0, 232,
		175, 1, 0, 0, 0, 232, 178, 1, 0, 0, 0, 232, 181, 1, 0, 0, 0, 232, 184,
		1, 0, 0, 0, 232, 190, 1, 0, 0, 0, 232, 196, 1, 0, 0, 0, 232, 202, 1, 0,
		0, 0, 232, 208, 1, 0, 0, 0, 232, 214, 1, 0, 0, 0, 232, 220, 1, 0, 0, 0,
		232, 226, 1, 0, 0, 0, 233, 9, 1, 0, 0, 0, 234, 235, 5, 8, 0, 0, 235, 236,
		5, 39, 0, 0, 236, 237, 5, 43, 0, 0, 237, 238, 3, 16, 8, 0, 238, 239, 5,
		42, 0, 0, 239, 240, 3, 18, 9, 0, 240, 241, 6, 5, -1, 0, 241, 256, 1, 0,
		0, 0, 242, 243, 5, 8, 0, 0, 243, 244, 5, 39, 0, 0, 244, 245, 5, 42, 0,
		0, 245, 246, 3, 18, 9, 0, 246, 247, 6, 5, -1, 0, 247, 256, 1, 0, 0, 0,
		248, 249, 5, 8, 0, 0, 249, 250, 5, 39, 0, 0, 250, 251, 5, 43, 0, 0, 251,
		252, 3, 16, 8, 0, 252, 253, 5, 45, 0, 0, 253, 254, 6, 5, -1, 0, 254, 256,
		1, 0, 0, 0, 255, 234, 1, 0, 0, 0, 255, 242, 1, 0, 0, 0, 255, 248, 1, 0,
		0, 0, 256, 11, 1, 0, 0, 0, 257, 258, 5, 9, 0, 0, 258, 259, 5, 39, 0, 0,
		259, 260, 5, 43, 0, 0, 260, 261, 3, 16, 8, 0, 261, 262, 5, 42, 0, 0, 262,
		263, 3, 18, 9, 0, 263, 264, 6, 6, -1, 0, 264, 272, 1, 0, 0, 0, 265, 266,
		5, 9, 0, 0, 266, 267, 5, 39, 0, 0, 267, 268, 5, 42, 0, 0, 268, 269, 3,
		18, 9, 0, 269, 270, 6, 6, -1, 0, 270, 272, 1, 0, 0, 0, 271, 257, 1, 0,
		0, 0, 271, 265, 1, 0, 0, 0, 272, 13, 1, 0, 0, 0, 273, 274, 5, 39, 0, 0,
		274, 275, 5, 42, 0, 0, 275, 276, 3, 18, 9, 0, 276, 277, 6, 7, -1, 0, 277,
		289, 1, 0, 0, 0, 278, 279, 5, 39, 0, 0, 279, 280, 5, 62, 0, 0, 280, 281,
		3, 18, 9, 0, 281, 282, 6, 7, -1, 0, 282, 289, 1, 0, 0, 0, 283, 284, 5,
		39, 0, 0, 284, 285, 5, 63, 0, 0, 285, 286, 3, 18, 9, 0, 286, 287, 6, 7,
		-1, 0, 287, 289, 1, 0, 0, 0, 288, 273, 1, 0, 0, 0, 288, 278, 1, 0, 0, 0,
		288, 283, 1, 0, 0, 0, 289, 15, 1, 0, 0, 0, 290, 291, 5, 1, 0, 0, 291, 301,
		6, 8, -1, 0, 292, 293, 5, 2, 0, 0, 293, 301, 6, 8, -1, 0, 294, 295, 5,
		3, 0, 0, 295, 301, 6, 8, -1, 0, 296, 297, 5, 4, 0, 0, 297, 301, 6, 8, -1,
		0, 298, 299, 5, 5, 0, 0, 299, 301, 6, 8, -1, 0, 300, 290, 1, 0, 0, 0, 300,
		292, 1, 0, 0, 0, 300, 294, 1, 0, 0, 0, 300, 296, 1, 0, 0, 0, 300, 298,
		1, 0, 0, 0, 301, 17, 1, 0, 0, 0, 302, 303, 6, 9, -1, 0, 303, 304, 5, 50,
		0, 0, 304, 305, 3, 18, 9, 21, 305, 306, 6, 9, -1, 0, 306, 339, 1, 0, 0,
		0, 307, 308, 5, 46, 0, 0, 308, 309, 3, 18, 9, 0, 309, 310, 5, 47, 0, 0,
		310, 311, 6, 9, -1, 0, 311, 339, 1, 0, 0, 0, 312, 313, 5, 61, 0, 0, 313,
		314, 5, 37, 0, 0, 314, 339, 6, 9, -1, 0, 315, 316, 5, 37, 0, 0, 316, 339,
		6, 9, -1, 0, 317, 318, 5, 38, 0, 0, 318, 339, 6, 9, -1, 0, 319, 320, 5,
		6, 0, 0, 320, 339, 6, 9, -1, 0, 321, 322, 5, 7, 0, 0, 322, 339, 6, 9, -1,
		0, 323, 324, 5, 40, 0, 0, 324, 339, 6, 9, -1, 0, 325, 326, 5, 39, 0, 0,
		326, 339, 6, 9, -1, 0, 327, 328, 5, 10, 0, 0, 328, 339, 6, 9, -1, 0, 329,
		330, 3, 50, 25, 0, 330, 331, 6, 9, -1, 0, 331, 339, 1, 0, 0, 0, 332, 333,
		3, 52, 26, 0, 333, 334, 6, 9, -1, 0, 334, 339, 1, 0, 0, 0, 335, 336, 3,
		54, 27, 0, 336, 337, 6, 9, -1, 0, 337, 339, 1, 0, 0, 0, 338, 302, 1, 0,
		0, 0, 338, 307, 1, 0, 0, 0, 338, 312, 1, 0, 0, 0, 338, 315, 1, 0, 0, 0,
		338, 317, 1, 0, 0, 0, 338, 319, 1, 0, 0, 0, 338, 321, 1, 0, 0, 0, 338,
		323, 1, 0, 0, 0, 338, 325, 1, 0, 0, 0, 338, 327, 1, 0, 0, 0, 338, 329,
		1, 0, 0, 0, 338, 332, 1, 0, 0, 0, 338, 335, 1, 0, 0, 0, 339, 382, 1, 0,
		0, 0, 340, 341, 10, 20, 0, 0, 341, 342, 5, 57, 0, 0, 342, 343, 3, 18, 9,
		21, 343, 344, 6, 9, -1, 0, 344, 381, 1, 0, 0, 0, 345, 346, 10, 19, 0, 0,
		346, 347, 7, 0, 0, 0, 347, 348, 3, 18, 9, 20, 348, 349, 6, 9, -1, 0, 349,
		381, 1, 0, 0, 0, 350, 351, 10, 18, 0, 0, 351, 352, 7, 1, 0, 0, 352, 353,
		3, 18, 9, 19, 353, 354, 6, 9, -1, 0, 354, 381, 1, 0, 0, 0, 355, 356, 10,
		17, 0, 0, 356, 357, 7, 2, 0, 0, 357, 358, 3, 18, 9, 18, 358, 359, 6, 9,
		-1, 0, 359, 381, 1, 0, 0, 0, 360, 361, 10, 16, 0, 0, 361, 362, 7, 3, 0,
		0, 362, 363, 3, 18, 9, 17, 363, 364, 6, 9, -1, 0, 364, 381, 1, 0, 0, 0,
		365, 366, 10, 15, 0, 0, 366, 367, 7, 4, 0, 0, 367, 368, 3, 18, 9, 16, 368,
		369, 6, 9, -1, 0, 369, 381, 1, 0, 0, 0, 370, 371, 10, 14, 0, 0, 371, 372,
		5, 52, 0, 0, 372, 373, 3, 18, 9, 15, 373, 374, 6, 9, -1, 0, 374, 381, 1,
		0, 0, 0, 375, 376, 10, 13, 0, 0, 376, 377, 5, 51, 0, 0, 377, 378, 3, 18,
		9, 14, 378, 379, 6, 9, -1, 0, 379, 381, 1, 0, 0, 0, 380, 340, 1, 0, 0,
		0, 380, 345, 1, 0, 0, 0, 380, 350, 1, 0, 0, 0, 380, 355, 1, 0, 0, 0, 380,
		360, 1, 0, 0, 0, 380, 365, 1, 0, 0, 0, 380, 370, 1, 0, 0, 0, 380, 375,
		1, 0, 0, 0, 381, 384, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 382, 383, 1, 0,
		0, 0, 383, 19, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 385, 386, 5, 11, 0, 0,
		386, 387, 3, 18, 9, 0, 387, 388, 5, 64, 0, 0, 388, 389, 3, 6, 3, 0, 389,
		390, 5, 65, 0, 0, 390, 391, 6, 10, -1, 0, 391, 413, 1, 0, 0, 0, 392, 393,
		5, 11, 0, 0, 393, 394, 3, 18, 9, 0, 394, 395, 5, 64, 0, 0, 395, 396, 3,
		6, 3, 0, 396, 397, 5, 65, 0, 0, 397, 398, 5, 12, 0, 0, 398, 399, 5, 64,
		0, 0, 399, 400, 3, 6, 3, 0, 400, 401, 5, 65, 0, 0, 401, 402, 6, 10, -1,
		0, 402, 413, 1, 0, 0, 0, 403, 404, 5, 11, 0, 0, 404, 405, 3, 18, 9, 0,
		405, 406, 5, 64, 0, 0, 406, 407, 3, 6, 3, 0, 407, 408, 5, 65, 0, 0, 408,
		409, 5, 12, 0, 0, 409, 410, 3, 20, 10, 0, 410, 411, 6, 10, -1, 0, 411,
		413, 1, 0, 0, 0, 412, 385, 1, 0, 0, 0, 412, 392, 1, 0, 0, 0, 412, 403,
		1, 0, 0, 0, 413, 21, 1, 0, 0, 0, 414, 415, 5, 13, 0, 0, 415, 416, 3, 18,
		9, 0, 416, 417, 5, 64, 0, 0, 417, 421, 3, 24, 12, 0, 418, 419, 5, 15, 0,
		0, 419, 420, 5, 43, 0, 0, 420, 422, 3, 6, 3, 0, 421, 418, 1, 0, 0, 0, 421,
		422, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 424, 5, 65, 0, 0, 424, 425,
		6, 11, -1, 0, 425, 23, 1, 0, 0, 0, 426, 428, 3, 26, 13, 0, 427, 426, 1,
		0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 429, 430, 1, 0, 0,
		0, 430, 431, 1, 0, 0, 0, 431, 432, 6, 12, -1, 0, 432, 25, 1, 0, 0, 0, 433,
		434, 5, 14, 0, 0, 434, 435, 3, 18, 9, 0, 435, 436, 5, 43, 0, 0, 436, 437,
		3, 6, 3, 0, 437, 438, 6, 13, -1, 0, 438, 27, 1, 0, 0, 0, 439, 440, 5, 21,
		0, 0, 440, 441, 3, 18, 9, 0, 441, 442, 5, 64, 0, 0, 442, 443, 3, 6, 3,
		0, 443, 444, 5, 65, 0, 0, 444, 445, 6, 14, -1, 0, 445, 29, 1, 0, 0, 0,
		446, 447, 5, 18, 0, 0, 447, 448, 5, 39, 0, 0, 448, 449, 5, 19, 0, 0, 449,
		450, 3, 18, 9, 0, 450, 451, 5, 20, 0, 0, 451, 452, 3, 18, 9, 0, 452, 453,
		5, 64, 0, 0, 453, 454, 3, 6, 3, 0, 454, 455, 5, 65, 0, 0, 455, 456, 6,
		15, -1, 0, 456, 476, 1, 0, 0, 0, 457, 458, 5, 18, 0, 0, 458, 459, 5, 39,
		0, 0, 459, 460, 5, 19, 0, 0, 460, 461, 5, 39, 0, 0, 461, 462, 5, 64, 0,
		0, 462, 463, 3, 6, 3, 0, 463, 464, 5, 65, 0, 0, 464, 465, 6, 15, -1, 0,
		465, 476, 1, 0, 0, 0, 466, 467, 5, 18, 0, 0, 467, 468, 5, 39, 0, 0, 468,
		469, 5, 19, 0, 0, 469, 470, 3, 18, 9, 0, 470, 471, 5, 64, 0, 0, 471, 472,
		3, 6, 3, 0, 472, 473, 5, 65, 0, 0, 473, 474, 6, 15, -1, 0, 474, 476, 1,
		0, 0, 0, 475, 446, 1, 0, 0, 0, 475, 457, 1, 0, 0, 0, 475, 466, 1, 0, 0,
		0, 476, 31, 1, 0, 0, 0, 477, 478, 5, 22, 0, 0, 478, 479, 3, 18, 9, 0, 479,
		480, 5, 12, 0, 0, 480, 481, 5, 64, 0, 0, 481, 482, 3, 6, 3, 0, 482, 483,
		5, 65, 0, 0, 483, 484, 6, 16, -1, 0, 484, 33, 1, 0, 0, 0, 485, 486, 5,
		17, 0, 0, 486, 487, 6, 17, -1, 0, 487, 35, 1, 0, 0, 0, 488, 489, 5, 16,
		0, 0, 489, 490, 6, 18, -1, 0, 490, 37, 1, 0, 0, 0, 491, 492, 5, 23, 0,
		0, 492, 493, 3, 18, 9, 0, 493, 494, 6, 19, -1, 0, 494, 498, 1, 0, 0, 0,
		495, 496, 5, 23, 0, 0, 496, 498, 6, 19, -1, 0, 497, 491, 1, 0, 0, 0, 497,
		495, 1, 0, 0, 0, 498, 39, 1, 0, 0, 0, 499, 500, 5, 8, 0, 0, 500, 501, 5,
		39, 0, 0, 501, 502, 5, 43, 0, 0, 502, 503, 5, 70, 0, 0, 503, 504, 3, 16,
		8, 0, 504, 505, 5, 71, 0, 0, 505, 506, 5, 42, 0, 0, 506, 507, 5, 70, 0,
		0, 507, 508, 3, 42, 21, 0, 508, 509, 5, 71, 0, 0, 509, 510, 6, 20, -1,
		0, 510, 533, 1, 0, 0, 0, 511, 512, 5, 8, 0, 0, 512, 513, 5, 39, 0, 0, 513,
		514, 5, 43, 0, 0, 514, 515, 5, 70, 0, 0, 515, 516, 3, 16, 8, 0, 516, 517,
		5, 71, 0, 0, 517, 518, 5, 42, 0, 0, 518, 519, 5, 70, 0, 0, 519, 520, 5,
		71, 0, 0, 520, 521, 6, 20, -1, 0, 521, 533, 1, 0, 0, 0, 522, 523, 5, 8,
		0, 0, 523, 524, 5, 39, 0, 0, 524, 525, 5, 43, 0, 0, 525, 526, 5, 70, 0,
		0, 526, 527, 3, 16, 8, 0, 527, 528, 5, 71, 0, 0, 528, 529, 5, 42, 0, 0,
		529, 530, 5, 39, 0, 0, 530, 531, 6, 20, -1, 0, 531, 533, 1, 0, 0, 0, 532,
		499, 1, 0, 0, 0, 532, 511, 1, 0, 0, 0, 532, 522, 1, 0, 0, 0, 533, 41, 1,
		0, 0, 0, 534, 536, 3, 44, 22, 0, 535, 534, 1, 0, 0, 0, 536, 537, 1, 0,
		0, 0, 537, 535, 1, 0, 0, 0, 537, 538, 1, 0, 0, 0, 538, 539, 1, 0, 0, 0,
		539, 540, 6, 21, -1, 0, 540, 43, 1, 0, 0, 0, 541, 542, 5, 67, 0, 0, 542,
		543, 3, 18, 9, 0, 543, 544, 6, 22, -1, 0, 544, 549, 1, 0, 0, 0, 545, 546,
		3, 18, 9, 0, 546, 547, 6, 22, -1, 0, 547, 549, 1, 0, 0, 0, 548, 541, 1,
		0, 0, 0, 548, 545, 1, 0, 0, 0, 549, 45, 1, 0, 0, 0, 550, 551, 5, 39, 0,
		0, 551, 552, 5, 68, 0, 0, 552, 553, 5, 27, 0, 0, 553, 554, 5, 46, 0, 0,
		554, 555, 3, 18, 9, 0, 555, 556, 5, 47, 0, 0, 556, 557, 6, 23, -1, 0, 557,
		578, 1, 0, 0, 0, 558, 559, 5, 39, 0, 0, 559, 560, 5, 70, 0, 0, 560, 561,
		3, 18, 9, 0, 561, 562, 5, 71, 0, 0, 562, 563, 5, 42, 0, 0, 563, 564, 5,
		39, 0, 0, 564, 565, 5, 70, 0, 0, 565, 566, 3, 18, 9, 0, 566, 567, 5, 71,
		0, 0, 567, 568, 6, 23, -1, 0, 568, 578, 1, 0, 0, 0, 569, 570, 5, 39, 0,
		0, 570, 571, 5, 70, 0, 0, 571, 572, 3, 18, 9, 0, 572, 573, 5, 71, 0, 0,
		573, 574, 5, 42, 0, 0, 574, 575, 3, 18, 9, 0, 575, 576, 6, 23, -1, 0, 576,
		578, 1, 0, 0, 0, 577, 550, 1, 0, 0, 0, 577, 558, 1, 0, 0, 0, 577, 569,
		1, 0, 0, 0, 578, 47, 1, 0, 0, 0, 579, 580, 5, 39, 0, 0, 580, 581, 5, 68,
		0, 0, 581, 582, 5, 29, 0, 0, 582, 583, 5, 46, 0, 0, 583, 584, 5, 47, 0,
		0, 584, 596, 6, 24, -1, 0, 585, 586, 5, 39, 0, 0, 586, 587, 5, 68, 0, 0,
		587, 588, 5, 28, 0, 0, 588, 589, 5, 46, 0, 0, 589, 590, 5, 32, 0, 0, 590,
		591, 5, 43, 0, 0, 591, 592, 3, 18, 9, 0, 592, 593, 5, 47, 0, 0, 593, 594,
		6, 24, -1, 0, 594, 596, 1, 0, 0, 0, 595, 579, 1, 0, 0, 0, 595, 585, 1,
		0, 0, 0, 596, 49, 1, 0, 0, 0, 597, 598, 5, 39, 0, 0, 598, 599, 5, 68, 0,
		0, 599, 600, 5, 31, 0, 0, 600, 601, 6, 25, -1, 0, 601, 51, 1, 0, 0, 0,
		602, 603, 5, 39, 0, 0, 603, 604, 5, 68, 0, 0, 604, 605, 5, 30, 0, 0, 605,
		606, 6, 26, -1, 0, 606, 53, 1, 0, 0, 0, 607, 608, 5, 39, 0, 0, 608, 609,
		5, 70, 0, 0, 609, 610, 3, 18, 9, 0, 610, 611, 5, 71, 0, 0, 611, 612, 6,
		27, -1, 0, 612, 55, 1, 0, 0, 0, 613, 614, 5, 8, 0, 0, 614, 617, 5, 39,
		0, 0, 615, 616, 5, 43, 0, 0, 616, 618, 3, 58, 29, 0, 617, 615, 1, 0, 0,
		0, 617, 618, 1, 0, 0, 0, 618, 619, 1, 0, 0, 0, 619, 620, 5, 42, 0, 0, 620,
		621, 3, 60, 30, 0, 621, 622, 6, 28, -1, 0, 622, 57, 1, 0, 0, 0, 623, 624,
		5, 70, 0, 0, 624, 625, 3, 58, 29, 0, 625, 626, 5, 71, 0, 0, 626, 627, 6,
		29, -1, 0, 627, 634, 1, 0, 0, 0, 628, 629, 5, 70, 0, 0, 629, 630, 3, 16,
		8, 0, 630, 631, 5, 71, 0, 0, 631, 632, 6, 29, -1, 0, 632, 634, 1, 0, 0,
		0, 633, 623, 1, 0, 0, 0, 633, 628, 1, 0, 0, 0, 634, 59, 1, 0, 0, 0, 635,
		636, 3, 62, 31, 0, 636, 637, 6, 30, -1, 0, 637, 61, 1, 0, 0, 0, 638, 639,
		5, 70, 0, 0, 639, 640, 3, 64, 32, 0, 640, 641, 5, 71, 0, 0, 641, 642, 6,
		31, -1, 0, 642, 647, 1, 0, 0, 0, 643, 644, 3, 70, 35, 0, 644, 645, 6, 31,
		-1, 0, 645, 647, 1, 0, 0, 0, 646, 638, 1, 0, 0, 0, 646, 643, 1, 0, 0, 0,
		647, 63, 1, 0, 0, 0, 648, 649, 6, 32, -1, 0, 649, 650, 3, 62, 31, 0, 650,
		651, 6, 32, -1, 0, 651, 656, 1, 0, 0, 0, 652, 653, 3, 66, 33, 0, 653, 654,
		6, 32, -1, 0, 654, 656, 1, 0, 0, 0, 655, 648, 1, 0, 0, 0, 655, 652, 1,
		0, 0, 0, 656, 664, 1, 0, 0, 0, 657, 658, 10, 3, 0, 0, 658, 659, 5, 67,
		0, 0, 659, 660, 3, 62, 31, 0, 660, 661, 6, 32, -1, 0, 661, 663, 1, 0, 0,
		0, 662, 657, 1, 0, 0, 0, 663, 666, 1, 0, 0, 0, 664, 662, 1, 0, 0, 0, 664,
		665, 1, 0, 0, 0, 665, 65, 1, 0, 0, 0, 666, 664, 1, 0, 0, 0, 667, 669, 3,
		68, 34, 0, 668, 667, 1, 0, 0, 0, 669, 670, 1, 0, 0, 0, 670, 668, 1, 0,
		0, 0, 670, 671, 1, 0, 0, 0, 671, 672, 1, 0, 0, 0, 672, 673, 6, 33, -1,
		0, 673, 67, 1, 0, 0, 0, 674, 675, 5, 67, 0, 0, 675, 676, 3, 18, 9, 0, 676,
		677, 6, 34, -1, 0, 677, 682, 1, 0, 0, 0, 678, 679, 3, 18, 9, 0, 679, 680,
		6, 34, -1, 0, 680, 682, 1, 0, 0, 0, 681, 674, 1, 0, 0, 0, 681, 678, 1,
		0, 0, 0, 682, 69, 1, 0, 0, 0, 683, 684, 3, 58, 29, 0, 684, 685, 5, 46,
		0, 0, 685, 686, 5, 33, 0, 0, 686, 687, 5, 43, 0, 0, 687, 688, 3, 70, 35,
		0, 688, 689, 5, 67, 0, 0, 689, 690, 5, 30, 0, 0, 690, 691, 5, 43, 0, 0,
		691, 692, 5, 37, 0, 0, 692, 693, 5, 47, 0, 0, 693, 694, 6, 35, -1, 0, 694,
		708, 1, 0, 0, 0, 695, 696, 3, 58, 29, 0, 696, 697, 5, 46, 0, 0, 697, 698,
		5, 33, 0, 0, 698, 699, 5, 43, 0, 0, 699, 700, 3, 18, 9, 0, 700, 701, 5,
		67, 0, 0, 701, 702, 5, 30, 0, 0, 702, 703, 5, 43, 0, 0, 703, 704, 5, 37,
		0, 0, 704, 705, 5, 47, 0, 0, 705, 706, 6, 35, -1, 0, 706, 708, 1, 0, 0,
		0, 707, 683, 1, 0, 0, 0, 707, 695, 1, 0, 0, 0, 708, 71, 1, 0, 0, 0, 709,
		710, 5, 25, 0, 0, 710, 711, 5, 46, 0, 0, 711, 712, 3, 66, 33, 0, 712, 713,
		5, 47, 0, 0, 713, 714, 6, 36, -1, 0, 714, 73, 1, 0, 0, 0, 46, 81, 87, 93,
		99, 120, 132, 138, 142, 147, 153, 159, 165, 186, 192, 198, 204, 210, 216,
		222, 228, 232, 255, 271, 288, 300, 338, 380, 382, 412, 421, 429, 475, 497,
		532, 537, 548, 577, 595, 617, 633, 646, 655, 664, 670, 681, 707,
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

// SwiftGrammarParserInit initializes any static state used to implement SwiftGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftGrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.once.Do(swiftgrammarParserInit)
}

// NewSwiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftGrammarParser(input antlr.TokenStream) *SwiftGrammarParser {
	SwiftGrammarParserInit()
	this := new(SwiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SwiftGrammar.g4"

	return this
}

// SwiftGrammarParser tokens.
const (
	SwiftGrammarParserEOF          = antlr.TokenEOF
	SwiftGrammarParserINT          = 1
	SwiftGrammarParserFLOAT        = 2
	SwiftGrammarParserSTRING       = 3
	SwiftGrammarParserBOOL         = 4
	SwiftGrammarParserCHARACT      = 5
	SwiftGrammarParserTRU          = 6
	SwiftGrammarParserFAL          = 7
	SwiftGrammarParserVAR          = 8
	SwiftGrammarParserLET          = 9
	SwiftGrammarParserNULO         = 10
	SwiftGrammarParserIF           = 11
	SwiftGrammarParserELSE         = 12
	SwiftGrammarParserSWITCH       = 13
	SwiftGrammarParserCASE         = 14
	SwiftGrammarParserDEFAULT      = 15
	SwiftGrammarParserBREAK        = 16
	SwiftGrammarParserCONTINUE     = 17
	SwiftGrammarParserFOR          = 18
	SwiftGrammarParserIN           = 19
	SwiftGrammarParserRANGO        = 20
	SwiftGrammarParserWHILE        = 21
	SwiftGrammarParserGUARD        = 22
	SwiftGrammarParserRETURN       = 23
	SwiftGrammarParserFUNCION      = 24
	SwiftGrammarParserPRINT        = 25
	SwiftGrammarParserINOUT        = 26
	SwiftGrammarParserAPPEND       = 27
	SwiftGrammarParserREMOVE       = 28
	SwiftGrammarParserREMOVELAST   = 29
	SwiftGrammarParserCOUNT        = 30
	SwiftGrammarParserISEMPTY      = 31
	SwiftGrammarParserAT           = 32
	SwiftGrammarParserREPEATING    = 33
	SwiftGrammarParserSTRUCT       = 34
	SwiftGrammarParserMUTATING     = 35
	SwiftGrammarParserSELF         = 36
	SwiftGrammarParserNUMBER       = 37
	SwiftGrammarParserCADENA       = 38
	SwiftGrammarParserID_VALIDO    = 39
	SwiftGrammarParserCHARACTER    = 40
	SwiftGrammarParserWS           = 41
	SwiftGrammarParserIG           = 42
	SwiftGrammarParserDOS_PUNTOS   = 43
	SwiftGrammarParserPUNTOCOMA    = 44
	SwiftGrammarParserCIERRE_INTE  = 45
	SwiftGrammarParserPARIZQ       = 46
	SwiftGrammarParserPARDER       = 47
	SwiftGrammarParserDIF          = 48
	SwiftGrammarParserIG_IG        = 49
	SwiftGrammarParserNOT          = 50
	SwiftGrammarParserOR           = 51
	SwiftGrammarParserAND          = 52
	SwiftGrammarParserMAY_IG       = 53
	SwiftGrammarParserMEN_IG       = 54
	SwiftGrammarParserMAYOR        = 55
	SwiftGrammarParserMENOR        = 56
	SwiftGrammarParserMODULO       = 57
	SwiftGrammarParserMUL          = 58
	SwiftGrammarParserDIV          = 59
	SwiftGrammarParserADD          = 60
	SwiftGrammarParserSUB          = 61
	SwiftGrammarParserSUMA         = 62
	SwiftGrammarParserRESTA        = 63
	SwiftGrammarParserLLAVEIZQ     = 64
	SwiftGrammarParserLLAVEDER     = 65
	SwiftGrammarParserRETORNO      = 66
	SwiftGrammarParserCOMA         = 67
	SwiftGrammarParserPUNTO        = 68
	SwiftGrammarParserGUIONBAJO    = 69
	SwiftGrammarParserCORCHIZQ     = 70
	SwiftGrammarParserCORCHDER     = 71
	SwiftGrammarParserDIRME        = 72
	SwiftGrammarParserWHITESPACE   = 73
	SwiftGrammarParserCOMMENT      = 74
	SwiftGrammarParserLINE_COMMENT = 75
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s                  = 0
	SwiftGrammarParserRULE_block              = 1
	SwiftGrammarParserRULE_instruction        = 2
	SwiftGrammarParserRULE_blockinterno       = 3
	SwiftGrammarParserRULE_instructionint     = 4
	SwiftGrammarParserRULE_declavarible       = 5
	SwiftGrammarParserRULE_declaconstante     = 6
	SwiftGrammarParserRULE_asignacionvariable = 7
	SwiftGrammarParserRULE_tipodato           = 8
	SwiftGrammarParserRULE_expr               = 9
	SwiftGrammarParserRULE_sentenciaifelse    = 10
	SwiftGrammarParserRULE_switchcontrol      = 11
	SwiftGrammarParserRULE_blockcase          = 12
	SwiftGrammarParserRULE_bloquecase         = 13
	SwiftGrammarParserRULE_whilecontrol       = 14
	SwiftGrammarParserRULE_forcontrol         = 15
	SwiftGrammarParserRULE_guardcontrol       = 16
	SwiftGrammarParserRULE_continuee          = 17
	SwiftGrammarParserRULE_breakk             = 18
	SwiftGrammarParserRULE_retornos           = 19
	SwiftGrammarParserRULE_vectorcontrol      = 20
	SwiftGrammarParserRULE_blockparams        = 21
	SwiftGrammarParserRULE_bloqueparams       = 22
	SwiftGrammarParserRULE_vectoragregar      = 23
	SwiftGrammarParserRULE_vectorremover      = 24
	SwiftGrammarParserRULE_vectorvacio        = 25
	SwiftGrammarParserRULE_vectorcount        = 26
	SwiftGrammarParserRULE_vectoraccess       = 27
	SwiftGrammarParserRULE_matrizcontrol      = 28
	SwiftGrammarParserRULE_tipomatriz         = 29
	SwiftGrammarParserRULE_defmatriz          = 30
	SwiftGrammarParserRULE_listavaloresmat    = 31
	SwiftGrammarParserRULE_listavaloresmat2   = 32
	SwiftGrammarParserRULE_listaexpresions    = 33
	SwiftGrammarParserRULE_listaexpresion     = 34
	SwiftGrammarParserRULE_simplematriz       = 35
	SwiftGrammarParserRULE_printstmt          = 36
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	code   []interface{}
	_block IBlockContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Get_block() IBlockContext { return s._block }

func (s *SContext) Set_block(v IBlockContext) { s._block = v }

func (s *SContext) GetCode() []interface{} { return s.code }

func (s *SContext) SetCode(v []interface{}) { s.code = v }

func (s *SContext) Block() IBlockContext {
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

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *SwiftGrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftGrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(75)
		p.Match(SwiftGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*SContext).code = localctx.(*SContext).Get_block().GetBlk()

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
	p.RuleIndex = SwiftGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_block

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
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *SwiftGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftGrammarParserRULE_block)

	localctx.(*BlockContext).blk = []interface{}{}
	var listInt []IInstructionContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549795932928) != 0) {
		{
			p.SetState(78)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt = localctx.(*BlockContext).GetIns()
	for _, e := range listInt {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInst())
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

	// Get_declavarible returns the _declavarible rule contexts.
	Get_declavarible() IDeclavaribleContext

	// Get_declaconstante returns the _declaconstante rule contexts.
	Get_declaconstante() IDeclaconstanteContext

	// Get_asignacionvariable returns the _asignacionvariable rule contexts.
	Get_asignacionvariable() IAsignacionvariableContext

	// Get_sentenciaifelse returns the _sentenciaifelse rule contexts.
	Get_sentenciaifelse() ISentenciaifelseContext

	// Get_switchcontrol returns the _switchcontrol rule contexts.
	Get_switchcontrol() ISwitchcontrolContext

	// Get_whilecontrol returns the _whilecontrol rule contexts.
	Get_whilecontrol() IWhilecontrolContext

	// Get_forcontrol returns the _forcontrol rule contexts.
	Get_forcontrol() IForcontrolContext

	// Get_guardcontrol returns the _guardcontrol rule contexts.
	Get_guardcontrol() IGuardcontrolContext

	// Get_vectorcontrol returns the _vectorcontrol rule contexts.
	Get_vectorcontrol() IVectorcontrolContext

	// Get_vectoragregar returns the _vectoragregar rule contexts.
	Get_vectoragregar() IVectoragregarContext

	// Get_vectorremover returns the _vectorremover rule contexts.
	Get_vectorremover() IVectorremoverContext

	// Get_printstmt returns the _printstmt rule contexts.
	Get_printstmt() IPrintstmtContext

	// Get_matrizcontrol returns the _matrizcontrol rule contexts.
	Get_matrizcontrol() IMatrizcontrolContext

	// Set_declavarible sets the _declavarible rule contexts.
	Set_declavarible(IDeclavaribleContext)

	// Set_declaconstante sets the _declaconstante rule contexts.
	Set_declaconstante(IDeclaconstanteContext)

	// Set_asignacionvariable sets the _asignacionvariable rule contexts.
	Set_asignacionvariable(IAsignacionvariableContext)

	// Set_sentenciaifelse sets the _sentenciaifelse rule contexts.
	Set_sentenciaifelse(ISentenciaifelseContext)

	// Set_switchcontrol sets the _switchcontrol rule contexts.
	Set_switchcontrol(ISwitchcontrolContext)

	// Set_whilecontrol sets the _whilecontrol rule contexts.
	Set_whilecontrol(IWhilecontrolContext)

	// Set_forcontrol sets the _forcontrol rule contexts.
	Set_forcontrol(IForcontrolContext)

	// Set_guardcontrol sets the _guardcontrol rule contexts.
	Set_guardcontrol(IGuardcontrolContext)

	// Set_vectorcontrol sets the _vectorcontrol rule contexts.
	Set_vectorcontrol(IVectorcontrolContext)

	// Set_vectoragregar sets the _vectoragregar rule contexts.
	Set_vectoragregar(IVectoragregarContext)

	// Set_vectorremover sets the _vectorremover rule contexts.
	Set_vectorremover(IVectorremoverContext)

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_matrizcontrol sets the _matrizcontrol rule contexts.
	Set_matrizcontrol(IMatrizcontrolContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Declavarible() IDeclavaribleContext
	PUNTOCOMA() antlr.TerminalNode
	Declaconstante() IDeclaconstanteContext
	Asignacionvariable() IAsignacionvariableContext
	Sentenciaifelse() ISentenciaifelseContext
	Switchcontrol() ISwitchcontrolContext
	Whilecontrol() IWhilecontrolContext
	Forcontrol() IForcontrolContext
	Guardcontrol() IGuardcontrolContext
	Vectorcontrol() IVectorcontrolContext
	Vectoragregar() IVectoragregarContext
	Vectorremover() IVectorremoverContext
	Printstmt() IPrintstmtContext
	Matrizcontrol() IMatrizcontrolContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser              antlr.Parser
	inst                interfaces.Instruction
	_declavarible       IDeclavaribleContext
	_declaconstante     IDeclaconstanteContext
	_asignacionvariable IAsignacionvariableContext
	_sentenciaifelse    ISentenciaifelseContext
	_switchcontrol      ISwitchcontrolContext
	_whilecontrol       IWhilecontrolContext
	_forcontrol         IForcontrolContext
	_guardcontrol       IGuardcontrolContext
	_vectorcontrol      IVectorcontrolContext
	_vectoragregar      IVectoragregarContext
	_vectorremover      IVectorremoverContext
	_printstmt          IPrintstmtContext
	_matrizcontrol      IMatrizcontrolContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_declavarible() IDeclavaribleContext { return s._declavarible }

func (s *InstructionContext) Get_declaconstante() IDeclaconstanteContext { return s._declaconstante }

func (s *InstructionContext) Get_asignacionvariable() IAsignacionvariableContext {
	return s._asignacionvariable
}

func (s *InstructionContext) Get_sentenciaifelse() ISentenciaifelseContext { return s._sentenciaifelse }

func (s *InstructionContext) Get_switchcontrol() ISwitchcontrolContext { return s._switchcontrol }

func (s *InstructionContext) Get_whilecontrol() IWhilecontrolContext { return s._whilecontrol }

func (s *InstructionContext) Get_forcontrol() IForcontrolContext { return s._forcontrol }

func (s *InstructionContext) Get_guardcontrol() IGuardcontrolContext { return s._guardcontrol }

func (s *InstructionContext) Get_vectorcontrol() IVectorcontrolContext { return s._vectorcontrol }

func (s *InstructionContext) Get_vectoragregar() IVectoragregarContext { return s._vectoragregar }

func (s *InstructionContext) Get_vectorremover() IVectorremoverContext { return s._vectorremover }

func (s *InstructionContext) Get_printstmt() IPrintstmtContext { return s._printstmt }

func (s *InstructionContext) Get_matrizcontrol() IMatrizcontrolContext { return s._matrizcontrol }

func (s *InstructionContext) Set_declavarible(v IDeclavaribleContext) { s._declavarible = v }

func (s *InstructionContext) Set_declaconstante(v IDeclaconstanteContext) { s._declaconstante = v }

func (s *InstructionContext) Set_asignacionvariable(v IAsignacionvariableContext) {
	s._asignacionvariable = v
}

func (s *InstructionContext) Set_sentenciaifelse(v ISentenciaifelseContext) { s._sentenciaifelse = v }

func (s *InstructionContext) Set_switchcontrol(v ISwitchcontrolContext) { s._switchcontrol = v }

func (s *InstructionContext) Set_whilecontrol(v IWhilecontrolContext) { s._whilecontrol = v }

func (s *InstructionContext) Set_forcontrol(v IForcontrolContext) { s._forcontrol = v }

func (s *InstructionContext) Set_guardcontrol(v IGuardcontrolContext) { s._guardcontrol = v }

func (s *InstructionContext) Set_vectorcontrol(v IVectorcontrolContext) { s._vectorcontrol = v }

func (s *InstructionContext) Set_vectoragregar(v IVectoragregarContext) { s._vectoragregar = v }

func (s *InstructionContext) Set_vectorremover(v IVectorremoverContext) { s._vectorremover = v }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_matrizcontrol(v IMatrizcontrolContext) { s._matrizcontrol = v }

func (s *InstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstructionContext) Declavarible() IDeclavaribleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclavaribleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclavaribleContext)
}

func (s *InstructionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTOCOMA, 0)
}

func (s *InstructionContext) Declaconstante() IDeclaconstanteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaconstanteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaconstanteContext)
}

func (s *InstructionContext) Asignacionvariable() IAsignacionvariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionvariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionvariableContext)
}

func (s *InstructionContext) Sentenciaifelse() ISentenciaifelseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaifelseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaifelseContext)
}

func (s *InstructionContext) Switchcontrol() ISwitchcontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchcontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchcontrolContext)
}

func (s *InstructionContext) Whilecontrol() IWhilecontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilecontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilecontrolContext)
}

func (s *InstructionContext) Forcontrol() IForcontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForcontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForcontrolContext)
}

func (s *InstructionContext) Guardcontrol() IGuardcontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardcontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardcontrolContext)
}

func (s *InstructionContext) Vectorcontrol() IVectorcontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorcontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorcontrolContext)
}

func (s *InstructionContext) Vectoragregar() IVectoragregarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectoragregarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectoragregarContext)
}

func (s *InstructionContext) Vectorremover() IVectorremoverContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorremoverContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorremoverContext)
}

func (s *InstructionContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *InstructionContext) Matrizcontrol() IMatrizcontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrizcontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrizcontrolContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *SwiftGrammarParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftGrammarParserRULE_instruction)
	var _la int

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)

			var _x = p.Declavarible()

			localctx.(*InstructionContext)._declavarible = _x
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(86)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declavarible().GetDecvbl()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)

			var _x = p.Declaconstante()

			localctx.(*InstructionContext)._declaconstante = _x
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(92)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declaconstante().GetDeccon()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)

			var _x = p.Asignacionvariable()

			localctx.(*InstructionContext)._asignacionvariable = _x
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(98)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_asignacionvariable().GetAsgvbl()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)

			var _x = p.Sentenciaifelse()

			localctx.(*InstructionContext)._sentenciaifelse = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_sentenciaifelse().GetMyIfElse()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(106)

			var _x = p.Switchcontrol()

			localctx.(*InstructionContext)._switchcontrol = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchcontrol().GetMySwitch()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(109)

			var _x = p.Whilecontrol()

			localctx.(*InstructionContext)._whilecontrol = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilecontrol().GetWhict()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(112)

			var _x = p.Forcontrol()

			localctx.(*InstructionContext)._forcontrol = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forcontrol().GetForct()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(115)

			var _x = p.Guardcontrol()

			localctx.(*InstructionContext)._guardcontrol = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_guardcontrol().GetGuct()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(118)

			var _x = p.Vectorcontrol()

			localctx.(*InstructionContext)._vectorcontrol = _x
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(119)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vectorcontrol().GetVect()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(124)

			var _x = p.Vectoragregar()

			localctx.(*InstructionContext)._vectoragregar = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vectoragregar().GetVeadct()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(127)

			var _x = p.Vectorremover()

			localctx.(*InstructionContext)._vectorremover = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vectorremover().GetVermct()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(130)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(131)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(136)

			var _x = p.Matrizcontrol()

			localctx.(*InstructionContext)._matrizcontrol = _x
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(137)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_matrizcontrol().GetMatct()

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

// IBlockinternoContext is an interface to support dynamic dispatch.
type IBlockinternoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instructionint returns the _instructionint rule contexts.
	Get_instructionint() IInstructionintContext

	// Set_instructionint sets the _instructionint rule contexts.
	Set_instructionint(IInstructionintContext)

	// GetInsint returns the insint rule context list.
	GetInsint() []IInstructionintContext

	// SetInsint sets the insint rule context list.
	SetInsint([]IInstructionintContext)

	// GetBlkint returns the blkint attribute.
	GetBlkint() []interface{}

	// SetBlkint sets the blkint attribute.
	SetBlkint([]interface{})

	// Getter signatures
	AllInstructionint() []IInstructionintContext
	Instructionint(i int) IInstructionintContext

	// IsBlockinternoContext differentiates from other interfaces.
	IsBlockinternoContext()
}

type BlockinternoContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	blkint          []interface{}
	_instructionint IInstructionintContext
	insint          []IInstructionintContext
}

func NewEmptyBlockinternoContext() *BlockinternoContext {
	var p = new(BlockinternoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_blockinterno
	return p
}

func InitEmptyBlockinternoContext(p *BlockinternoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_blockinterno
}

func (*BlockinternoContext) IsBlockinternoContext() {}

func NewBlockinternoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockinternoContext {
	var p = new(BlockinternoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_blockinterno

	return p
}

func (s *BlockinternoContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockinternoContext) Get_instructionint() IInstructionintContext { return s._instructionint }

func (s *BlockinternoContext) Set_instructionint(v IInstructionintContext) { s._instructionint = v }

func (s *BlockinternoContext) GetInsint() []IInstructionintContext { return s.insint }

func (s *BlockinternoContext) SetInsint(v []IInstructionintContext) { s.insint = v }

func (s *BlockinternoContext) GetBlkint() []interface{} { return s.blkint }

func (s *BlockinternoContext) SetBlkint(v []interface{}) { s.blkint = v }

func (s *BlockinternoContext) AllInstructionint() []IInstructionintContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionintContext); ok {
			len++
		}
	}

	tst := make([]IInstructionintContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionintContext); ok {
			tst[i] = t.(IInstructionintContext)
			i++
		}
	}

	return tst
}

func (s *BlockinternoContext) Instructionint(i int) IInstructionintContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionintContext); ok {
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

	return t.(IInstructionintContext)
}

func (s *BlockinternoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockinternoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockinternoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlockinterno(s)
	}
}

func (s *BlockinternoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlockinterno(s)
	}
}

func (p *SwiftGrammarParser) Blockinterno() (localctx IBlockinternoContext) {
	localctx = NewBlockinternoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_blockinterno)

	localctx.(*BlockinternoContext).blkint = []interface{}{}
	var listInt []IInstructionintContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549804518144) != 0) {
		{
			p.SetState(144)

			var _x = p.Instructionint()

			localctx.(*BlockinternoContext)._instructionint = _x
		}
		localctx.(*BlockinternoContext).insint = append(localctx.(*BlockinternoContext).insint, localctx.(*BlockinternoContext)._instructionint)

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt = localctx.(*BlockinternoContext).GetInsint()
	for _, e := range listInt {
		localctx.(*BlockinternoContext).blkint = append(localctx.(*BlockinternoContext).blkint, e.GetInsint())
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

// IInstructionintContext is an interface to support dynamic dispatch.
type IInstructionintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_declavarible returns the _declavarible rule contexts.
	Get_declavarible() IDeclavaribleContext

	// Get_declaconstante returns the _declaconstante rule contexts.
	Get_declaconstante() IDeclaconstanteContext

	// Get_asignacionvariable returns the _asignacionvariable rule contexts.
	Get_asignacionvariable() IAsignacionvariableContext

	// Get_sentenciaifelse returns the _sentenciaifelse rule contexts.
	Get_sentenciaifelse() ISentenciaifelseContext

	// Get_switchcontrol returns the _switchcontrol rule contexts.
	Get_switchcontrol() ISwitchcontrolContext

	// Get_whilecontrol returns the _whilecontrol rule contexts.
	Get_whilecontrol() IWhilecontrolContext

	// Get_forcontrol returns the _forcontrol rule contexts.
	Get_forcontrol() IForcontrolContext

	// Get_guardcontrol returns the _guardcontrol rule contexts.
	Get_guardcontrol() IGuardcontrolContext

	// Get_continuee returns the _continuee rule contexts.
	Get_continuee() IContinueeContext

	// Get_breakk returns the _breakk rule contexts.
	Get_breakk() IBreakkContext

	// Get_retornos returns the _retornos rule contexts.
	Get_retornos() IRetornosContext

	// Get_vectorcontrol returns the _vectorcontrol rule contexts.
	Get_vectorcontrol() IVectorcontrolContext

	// Get_vectoragregar returns the _vectoragregar rule contexts.
	Get_vectoragregar() IVectoragregarContext

	// Get_vectorremover returns the _vectorremover rule contexts.
	Get_vectorremover() IVectorremoverContext

	// Get_printstmt returns the _printstmt rule contexts.
	Get_printstmt() IPrintstmtContext

	// Get_matrizcontrol returns the _matrizcontrol rule contexts.
	Get_matrizcontrol() IMatrizcontrolContext

	// Set_declavarible sets the _declavarible rule contexts.
	Set_declavarible(IDeclavaribleContext)

	// Set_declaconstante sets the _declaconstante rule contexts.
	Set_declaconstante(IDeclaconstanteContext)

	// Set_asignacionvariable sets the _asignacionvariable rule contexts.
	Set_asignacionvariable(IAsignacionvariableContext)

	// Set_sentenciaifelse sets the _sentenciaifelse rule contexts.
	Set_sentenciaifelse(ISentenciaifelseContext)

	// Set_switchcontrol sets the _switchcontrol rule contexts.
	Set_switchcontrol(ISwitchcontrolContext)

	// Set_whilecontrol sets the _whilecontrol rule contexts.
	Set_whilecontrol(IWhilecontrolContext)

	// Set_forcontrol sets the _forcontrol rule contexts.
	Set_forcontrol(IForcontrolContext)

	// Set_guardcontrol sets the _guardcontrol rule contexts.
	Set_guardcontrol(IGuardcontrolContext)

	// Set_continuee sets the _continuee rule contexts.
	Set_continuee(IContinueeContext)

	// Set_breakk sets the _breakk rule contexts.
	Set_breakk(IBreakkContext)

	// Set_retornos sets the _retornos rule contexts.
	Set_retornos(IRetornosContext)

	// Set_vectorcontrol sets the _vectorcontrol rule contexts.
	Set_vectorcontrol(IVectorcontrolContext)

	// Set_vectoragregar sets the _vectoragregar rule contexts.
	Set_vectoragregar(IVectoragregarContext)

	// Set_vectorremover sets the _vectorremover rule contexts.
	Set_vectorremover(IVectorremoverContext)

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_matrizcontrol sets the _matrizcontrol rule contexts.
	Set_matrizcontrol(IMatrizcontrolContext)

	// GetInsint returns the insint attribute.
	GetInsint() interfaces.Instruction

	// SetInsint sets the insint attribute.
	SetInsint(interfaces.Instruction)

	// Getter signatures
	Declavarible() IDeclavaribleContext
	PUNTOCOMA() antlr.TerminalNode
	Declaconstante() IDeclaconstanteContext
	Asignacionvariable() IAsignacionvariableContext
	Sentenciaifelse() ISentenciaifelseContext
	Switchcontrol() ISwitchcontrolContext
	Whilecontrol() IWhilecontrolContext
	Forcontrol() IForcontrolContext
	Guardcontrol() IGuardcontrolContext
	Continuee() IContinueeContext
	Breakk() IBreakkContext
	Retornos() IRetornosContext
	Vectorcontrol() IVectorcontrolContext
	Vectoragregar() IVectoragregarContext
	Vectorremover() IVectorremoverContext
	Printstmt() IPrintstmtContext
	Matrizcontrol() IMatrizcontrolContext

	// IsInstructionintContext differentiates from other interfaces.
	IsInstructionintContext()
}

type InstructionintContext struct {
	antlr.BaseParserRuleContext
	parser              antlr.Parser
	insint              interfaces.Instruction
	_declavarible       IDeclavaribleContext
	_declaconstante     IDeclaconstanteContext
	_asignacionvariable IAsignacionvariableContext
	_sentenciaifelse    ISentenciaifelseContext
	_switchcontrol      ISwitchcontrolContext
	_whilecontrol       IWhilecontrolContext
	_forcontrol         IForcontrolContext
	_guardcontrol       IGuardcontrolContext
	_continuee          IContinueeContext
	_breakk             IBreakkContext
	_retornos           IRetornosContext
	_vectorcontrol      IVectorcontrolContext
	_vectoragregar      IVectoragregarContext
	_vectorremover      IVectorremoverContext
	_printstmt          IPrintstmtContext
	_matrizcontrol      IMatrizcontrolContext
}

func NewEmptyInstructionintContext() *InstructionintContext {
	var p = new(InstructionintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instructionint
	return p
}

func InitEmptyInstructionintContext(p *InstructionintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instructionint
}

func (*InstructionintContext) IsInstructionintContext() {}

func NewInstructionintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionintContext {
	var p = new(InstructionintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_instructionint

	return p
}

func (s *InstructionintContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionintContext) Get_declavarible() IDeclavaribleContext { return s._declavarible }

func (s *InstructionintContext) Get_declaconstante() IDeclaconstanteContext { return s._declaconstante }

func (s *InstructionintContext) Get_asignacionvariable() IAsignacionvariableContext {
	return s._asignacionvariable
}

func (s *InstructionintContext) Get_sentenciaifelse() ISentenciaifelseContext {
	return s._sentenciaifelse
}

func (s *InstructionintContext) Get_switchcontrol() ISwitchcontrolContext { return s._switchcontrol }

func (s *InstructionintContext) Get_whilecontrol() IWhilecontrolContext { return s._whilecontrol }

func (s *InstructionintContext) Get_forcontrol() IForcontrolContext { return s._forcontrol }

func (s *InstructionintContext) Get_guardcontrol() IGuardcontrolContext { return s._guardcontrol }

func (s *InstructionintContext) Get_continuee() IContinueeContext { return s._continuee }

func (s *InstructionintContext) Get_breakk() IBreakkContext { return s._breakk }

func (s *InstructionintContext) Get_retornos() IRetornosContext { return s._retornos }

func (s *InstructionintContext) Get_vectorcontrol() IVectorcontrolContext { return s._vectorcontrol }

func (s *InstructionintContext) Get_vectoragregar() IVectoragregarContext { return s._vectoragregar }

func (s *InstructionintContext) Get_vectorremover() IVectorremoverContext { return s._vectorremover }

func (s *InstructionintContext) Get_printstmt() IPrintstmtContext { return s._printstmt }

func (s *InstructionintContext) Get_matrizcontrol() IMatrizcontrolContext { return s._matrizcontrol }

func (s *InstructionintContext) Set_declavarible(v IDeclavaribleContext) { s._declavarible = v }

func (s *InstructionintContext) Set_declaconstante(v IDeclaconstanteContext) { s._declaconstante = v }

func (s *InstructionintContext) Set_asignacionvariable(v IAsignacionvariableContext) {
	s._asignacionvariable = v
}

func (s *InstructionintContext) Set_sentenciaifelse(v ISentenciaifelseContext) {
	s._sentenciaifelse = v
}

func (s *InstructionintContext) Set_switchcontrol(v ISwitchcontrolContext) { s._switchcontrol = v }

func (s *InstructionintContext) Set_whilecontrol(v IWhilecontrolContext) { s._whilecontrol = v }

func (s *InstructionintContext) Set_forcontrol(v IForcontrolContext) { s._forcontrol = v }

func (s *InstructionintContext) Set_guardcontrol(v IGuardcontrolContext) { s._guardcontrol = v }

func (s *InstructionintContext) Set_continuee(v IContinueeContext) { s._continuee = v }

func (s *InstructionintContext) Set_breakk(v IBreakkContext) { s._breakk = v }

func (s *InstructionintContext) Set_retornos(v IRetornosContext) { s._retornos = v }

func (s *InstructionintContext) Set_vectorcontrol(v IVectorcontrolContext) { s._vectorcontrol = v }

func (s *InstructionintContext) Set_vectoragregar(v IVectoragregarContext) { s._vectoragregar = v }

func (s *InstructionintContext) Set_vectorremover(v IVectorremoverContext) { s._vectorremover = v }

func (s *InstructionintContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionintContext) Set_matrizcontrol(v IMatrizcontrolContext) { s._matrizcontrol = v }

func (s *InstructionintContext) GetInsint() interfaces.Instruction { return s.insint }

func (s *InstructionintContext) SetInsint(v interfaces.Instruction) { s.insint = v }

func (s *InstructionintContext) Declavarible() IDeclavaribleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclavaribleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclavaribleContext)
}

func (s *InstructionintContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTOCOMA, 0)
}

func (s *InstructionintContext) Declaconstante() IDeclaconstanteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaconstanteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaconstanteContext)
}

func (s *InstructionintContext) Asignacionvariable() IAsignacionvariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionvariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionvariableContext)
}

func (s *InstructionintContext) Sentenciaifelse() ISentenciaifelseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaifelseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaifelseContext)
}

func (s *InstructionintContext) Switchcontrol() ISwitchcontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchcontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchcontrolContext)
}

func (s *InstructionintContext) Whilecontrol() IWhilecontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilecontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilecontrolContext)
}

func (s *InstructionintContext) Forcontrol() IForcontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForcontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForcontrolContext)
}

func (s *InstructionintContext) Guardcontrol() IGuardcontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardcontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardcontrolContext)
}

func (s *InstructionintContext) Continuee() IContinueeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueeContext)
}

func (s *InstructionintContext) Breakk() IBreakkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakkContext)
}

func (s *InstructionintContext) Retornos() IRetornosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetornosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetornosContext)
}

func (s *InstructionintContext) Vectorcontrol() IVectorcontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorcontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorcontrolContext)
}

func (s *InstructionintContext) Vectoragregar() IVectoragregarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectoragregarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectoragregarContext)
}

func (s *InstructionintContext) Vectorremover() IVectorremoverContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorremoverContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorremoverContext)
}

func (s *InstructionintContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *InstructionintContext) Matrizcontrol() IMatrizcontrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrizcontrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrizcontrolContext)
}

func (s *InstructionintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterInstructionint(s)
	}
}

func (s *InstructionintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitInstructionint(s)
	}
}

func (p *SwiftGrammarParser) Instructionint() (localctx IInstructionintContext) {
	localctx = NewInstructionintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_instructionint)
	var _la int

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)

			var _x = p.Declavarible()

			localctx.(*InstructionintContext)._declavarible = _x
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(152)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_declavarible().GetDecvbl()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)

			var _x = p.Declaconstante()

			localctx.(*InstructionintContext)._declaconstante = _x
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(158)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_declaconstante().GetDeccon()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)

			var _x = p.Asignacionvariable()

			localctx.(*InstructionintContext)._asignacionvariable = _x
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(164)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_asignacionvariable().GetAsgvbl()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(169)

			var _x = p.Sentenciaifelse()

			localctx.(*InstructionintContext)._sentenciaifelse = _x
		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_sentenciaifelse().GetMyIfElse()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(172)

			var _x = p.Switchcontrol()

			localctx.(*InstructionintContext)._switchcontrol = _x
		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_switchcontrol().GetMySwitch()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(175)

			var _x = p.Whilecontrol()

			localctx.(*InstructionintContext)._whilecontrol = _x
		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_whilecontrol().GetWhict()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(178)

			var _x = p.Forcontrol()

			localctx.(*InstructionintContext)._forcontrol = _x
		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_forcontrol().GetForct()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(181)

			var _x = p.Guardcontrol()

			localctx.(*InstructionintContext)._guardcontrol = _x
		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_guardcontrol().GetGuct()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(184)

			var _x = p.Continuee()

			localctx.(*InstructionintContext)._continuee = _x
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(185)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_continuee().GetCoct()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(190)

			var _x = p.Breakk()

			localctx.(*InstructionintContext)._breakk = _x
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(191)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_breakk().GetBrkct()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(196)

			var _x = p.Retornos()

			localctx.(*InstructionintContext)._retornos = _x
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(197)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_retornos().GetRect()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(202)

			var _x = p.Vectorcontrol()

			localctx.(*InstructionintContext)._vectorcontrol = _x
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(203)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_vectorcontrol().GetVect()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(208)

			var _x = p.Vectoragregar()

			localctx.(*InstructionintContext)._vectoragregar = _x
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(209)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_vectoragregar().GetVeadct()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(214)

			var _x = p.Vectorremover()

			localctx.(*InstructionintContext)._vectorremover = _x
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(215)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_vectorremover().GetVermct()

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(220)

			var _x = p.Printstmt()

			localctx.(*InstructionintContext)._printstmt = _x
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(221)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_printstmt().GetPrnt()

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(226)

			var _x = p.Matrizcontrol()

			localctx.(*InstructionintContext)._matrizcontrol = _x
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(227)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionintContext).insint = localctx.(*InstructionintContext).Get_matrizcontrol().GetMatct()

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

// IDeclavaribleContext is an interface to support dynamic dispatch.
type IDeclavaribleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// Get_tipodato returns the _tipodato rule contexts.
	Get_tipodato() ITipodatoContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_tipodato sets the _tipodato rule contexts.
	Set_tipodato(ITipodatoContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetDecvbl returns the decvbl attribute.
	GetDecvbl() interfaces.Instruction

	// SetDecvbl sets the decvbl attribute.
	SetDecvbl(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID_VALIDO() antlr.TerminalNode
	DOS_PUNTOS() antlr.TerminalNode
	Tipodato() ITipodatoContext
	IG() antlr.TerminalNode
	Expr() IExprContext
	CIERRE_INTE() antlr.TerminalNode

	// IsDeclavaribleContext differentiates from other interfaces.
	IsDeclavaribleContext()
}

type DeclavaribleContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	decvbl     interfaces.Instruction
	_VAR       antlr.Token
	_ID_VALIDO antlr.Token
	_tipodato  ITipodatoContext
	_expr      IExprContext
}

func NewEmptyDeclavaribleContext() *DeclavaribleContext {
	var p = new(DeclavaribleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declavarible
	return p
}

func InitEmptyDeclavaribleContext(p *DeclavaribleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declavarible
}

func (*DeclavaribleContext) IsDeclavaribleContext() {}

func NewDeclavaribleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclavaribleContext {
	var p = new(DeclavaribleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_declavarible

	return p
}

func (s *DeclavaribleContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclavaribleContext) Get_VAR() antlr.Token { return s._VAR }

func (s *DeclavaribleContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *DeclavaribleContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *DeclavaribleContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *DeclavaribleContext) Get_tipodato() ITipodatoContext { return s._tipodato }

func (s *DeclavaribleContext) Get_expr() IExprContext { return s._expr }

func (s *DeclavaribleContext) Set_tipodato(v ITipodatoContext) { s._tipodato = v }

func (s *DeclavaribleContext) Set_expr(v IExprContext) { s._expr = v }

func (s *DeclavaribleContext) GetDecvbl() interfaces.Instruction { return s.decvbl }

func (s *DeclavaribleContext) SetDecvbl(v interfaces.Instruction) { s.decvbl = v }

func (s *DeclavaribleContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *DeclavaribleContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, 0)
}

func (s *DeclavaribleContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOS_PUNTOS, 0)
}

func (s *DeclavaribleContext) Tipodato() ITipodatoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipodatoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipodatoContext)
}

func (s *DeclavaribleContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *DeclavaribleContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclavaribleContext) CIERRE_INTE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCIERRE_INTE, 0)
}

func (s *DeclavaribleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclavaribleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclavaribleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclavarible(s)
	}
}

func (s *DeclavaribleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclavarible(s)
	}
}

func (p *SwiftGrammarParser) Declavarible() (localctx IDeclavaribleContext) {
	localctx = NewDeclavaribleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_declavarible)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclavaribleContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*DeclavaribleContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)

			var _x = p.Tipodato()

			localctx.(*DeclavaribleContext)._tipodato = _x
		}
		{
			p.SetState(238)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)

			var _x = p.expr(0)

			localctx.(*DeclavaribleContext)._expr = _x
		}
		localctx.(*DeclavaribleContext).decvbl = datosprimitivos.NewVariableDeclaration((func() int {
			if localctx.(*DeclavaribleContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclavaribleContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclavaribleContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclavaribleContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclavaribleContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*DeclavaribleContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*DeclavaribleContext).Get_tipodato().GetTipo(), localctx.(*DeclavaribleContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclavaribleContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*DeclavaribleContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)

			var _x = p.expr(0)

			localctx.(*DeclavaribleContext)._expr = _x
		}
		localctx.(*DeclavaribleContext).decvbl = datosprimitivos.NewVariableDeclaracionSinTipo((func() int {
			if localctx.(*DeclavaribleContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclavaribleContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclavaribleContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclavaribleContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclavaribleContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*DeclavaribleContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*DeclavaribleContext).Get_expr().GetE())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(248)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclavaribleContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*DeclavaribleContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)

			var _x = p.Tipodato()

			localctx.(*DeclavaribleContext)._tipodato = _x
		}
		{
			p.SetState(252)
			p.Match(SwiftGrammarParserCIERRE_INTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*DeclavaribleContext).decvbl = datosprimitivos.NewVariableDeclaracionSinExp((func() int {
			if localctx.(*DeclavaribleContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclavaribleContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclavaribleContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclavaribleContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclavaribleContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*DeclavaribleContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*DeclavaribleContext).Get_tipodato().GetTipo())

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

// IDeclaconstanteContext is an interface to support dynamic dispatch.
type IDeclaconstanteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// Get_tipodato returns the _tipodato rule contexts.
	Get_tipodato() ITipodatoContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_tipodato sets the _tipodato rule contexts.
	Set_tipodato(ITipodatoContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetDeccon returns the deccon attribute.
	GetDeccon() interfaces.Instruction

	// SetDeccon sets the deccon attribute.
	SetDeccon(interfaces.Instruction)

	// Getter signatures
	LET() antlr.TerminalNode
	ID_VALIDO() antlr.TerminalNode
	DOS_PUNTOS() antlr.TerminalNode
	Tipodato() ITipodatoContext
	IG() antlr.TerminalNode
	Expr() IExprContext

	// IsDeclaconstanteContext differentiates from other interfaces.
	IsDeclaconstanteContext()
}

type DeclaconstanteContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	deccon     interfaces.Instruction
	_LET       antlr.Token
	_ID_VALIDO antlr.Token
	_tipodato  ITipodatoContext
	_expr      IExprContext
}

func NewEmptyDeclaconstanteContext() *DeclaconstanteContext {
	var p = new(DeclaconstanteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declaconstante
	return p
}

func InitEmptyDeclaconstanteContext(p *DeclaconstanteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declaconstante
}

func (*DeclaconstanteContext) IsDeclaconstanteContext() {}

func NewDeclaconstanteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaconstanteContext {
	var p = new(DeclaconstanteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_declaconstante

	return p
}

func (s *DeclaconstanteContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaconstanteContext) Get_LET() antlr.Token { return s._LET }

func (s *DeclaconstanteContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *DeclaconstanteContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *DeclaconstanteContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *DeclaconstanteContext) Get_tipodato() ITipodatoContext { return s._tipodato }

func (s *DeclaconstanteContext) Get_expr() IExprContext { return s._expr }

func (s *DeclaconstanteContext) Set_tipodato(v ITipodatoContext) { s._tipodato = v }

func (s *DeclaconstanteContext) Set_expr(v IExprContext) { s._expr = v }

func (s *DeclaconstanteContext) GetDeccon() interfaces.Instruction { return s.deccon }

func (s *DeclaconstanteContext) SetDeccon(v interfaces.Instruction) { s.deccon = v }

func (s *DeclaconstanteContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
}

func (s *DeclaconstanteContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, 0)
}

func (s *DeclaconstanteContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOS_PUNTOS, 0)
}

func (s *DeclaconstanteContext) Tipodato() ITipodatoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipodatoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipodatoContext)
}

func (s *DeclaconstanteContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *DeclaconstanteContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclaconstanteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaconstanteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaconstanteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclaconstante(s)
	}
}

func (s *DeclaconstanteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclaconstante(s)
	}
}

func (p *SwiftGrammarParser) Declaconstante() (localctx IDeclaconstanteContext) {
	localctx = NewDeclaconstanteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_declaconstante)
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclaconstanteContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*DeclaconstanteContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)

			var _x = p.Tipodato()

			localctx.(*DeclaconstanteContext)._tipodato = _x
		}
		{
			p.SetState(261)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)

			var _x = p.expr(0)

			localctx.(*DeclaconstanteContext)._expr = _x
		}
		localctx.(*DeclaconstanteContext).deccon = datosprimitivos.NewConstanteDeclaration((func() int {
			if localctx.(*DeclaconstanteContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaconstanteContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclaconstanteContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaconstanteContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclaconstanteContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*DeclaconstanteContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*DeclaconstanteContext).Get_tipodato().GetTipo(), localctx.(*DeclaconstanteContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclaconstanteContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*DeclaconstanteContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)

			var _x = p.expr(0)

			localctx.(*DeclaconstanteContext)._expr = _x
		}
		localctx.(*DeclaconstanteContext).deccon = datosprimitivos.NewConstanteDeclaracionSinTipo((func() int {
			if localctx.(*DeclaconstanteContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaconstanteContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclaconstanteContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaconstanteContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclaconstanteContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*DeclaconstanteContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*DeclaconstanteContext).Get_expr().GetE())

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

// IAsignacionvariableContext is an interface to support dynamic dispatch.
type IAsignacionvariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetAsgvbl returns the asgvbl attribute.
	GetAsgvbl() interfaces.Instruction

	// SetAsgvbl sets the asgvbl attribute.
	SetAsgvbl(interfaces.Instruction)

	// Getter signatures
	ID_VALIDO() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext
	SUMA() antlr.TerminalNode
	RESTA() antlr.TerminalNode

	// IsAsignacionvariableContext differentiates from other interfaces.
	IsAsignacionvariableContext()
}

type AsignacionvariableContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	asgvbl     interfaces.Instruction
	_ID_VALIDO antlr.Token
	_expr      IExprContext
}

func NewEmptyAsignacionvariableContext() *AsignacionvariableContext {
	var p = new(AsignacionvariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignacionvariable
	return p
}

func InitEmptyAsignacionvariableContext(p *AsignacionvariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignacionvariable
}

func (*AsignacionvariableContext) IsAsignacionvariableContext() {}

func NewAsignacionvariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionvariableContext {
	var p = new(AsignacionvariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_asignacionvariable

	return p
}

func (s *AsignacionvariableContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionvariableContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *AsignacionvariableContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *AsignacionvariableContext) Get_expr() IExprContext { return s._expr }

func (s *AsignacionvariableContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AsignacionvariableContext) GetAsgvbl() interfaces.Instruction { return s.asgvbl }

func (s *AsignacionvariableContext) SetAsgvbl(v interfaces.Instruction) { s.asgvbl = v }

func (s *AsignacionvariableContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, 0)
}

func (s *AsignacionvariableContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *AsignacionvariableContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionvariableContext) SUMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUMA, 0)
}

func (s *AsignacionvariableContext) RESTA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRESTA, 0)
}

func (s *AsignacionvariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionvariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionvariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAsignacionvariable(s)
	}
}

func (s *AsignacionvariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAsignacionvariable(s)
	}
}

func (p *SwiftGrammarParser) Asignacionvariable() (localctx IAsignacionvariableContext) {
	localctx = NewAsignacionvariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_asignacionvariable)
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*AsignacionvariableContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)

			var _x = p.expr(0)

			localctx.(*AsignacionvariableContext)._expr = _x
		}
		localctx.(*AsignacionvariableContext).asgvbl = sentencias.NewAsignacionVariable((func() int {
			if localctx.(*AsignacionvariableContext).Get_ID_VALIDO() == nil {
				return 0
			} else {
				return localctx.(*AsignacionvariableContext).Get_ID_VALIDO().GetLine()
			}
		}()), (func() int {
			if localctx.(*AsignacionvariableContext).Get_ID_VALIDO() == nil {
				return 0
			} else {
				return localctx.(*AsignacionvariableContext).Get_ID_VALIDO().GetColumn()
			}
		}()), (func() string {
			if localctx.(*AsignacionvariableContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*AsignacionvariableContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*AsignacionvariableContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*AsignacionvariableContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Match(SwiftGrammarParserSUMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)

			var _x = p.expr(0)

			localctx.(*AsignacionvariableContext)._expr = _x
		}
		localctx.(*AsignacionvariableContext).asgvbl = sentencias.NewAsignacionSuma((func() int {
			if localctx.(*AsignacionvariableContext).Get_ID_VALIDO() == nil {
				return 0
			} else {
				return localctx.(*AsignacionvariableContext).Get_ID_VALIDO().GetLine()
			}
		}()), (func() int {
			if localctx.(*AsignacionvariableContext).Get_ID_VALIDO() == nil {
				return 0
			} else {
				return localctx.(*AsignacionvariableContext).Get_ID_VALIDO().GetColumn()
			}
		}()), (func() string {
			if localctx.(*AsignacionvariableContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*AsignacionvariableContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*AsignacionvariableContext).Get_expr().GetE())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(283)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*AsignacionvariableContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(SwiftGrammarParserRESTA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)

			var _x = p.expr(0)

			localctx.(*AsignacionvariableContext)._expr = _x
		}
		localctx.(*AsignacionvariableContext).asgvbl = sentencias.NewAsignacionResta((func() int {
			if localctx.(*AsignacionvariableContext).Get_ID_VALIDO() == nil {
				return 0
			} else {
				return localctx.(*AsignacionvariableContext).Get_ID_VALIDO().GetLine()
			}
		}()), (func() int {
			if localctx.(*AsignacionvariableContext).Get_ID_VALIDO() == nil {
				return 0
			} else {
				return localctx.(*AsignacionvariableContext).Get_ID_VALIDO().GetColumn()
			}
		}()), (func() string {
			if localctx.(*AsignacionvariableContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*AsignacionvariableContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*AsignacionvariableContext).Get_expr().GetE())

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

// ITipodatoContext is an interface to support dynamic dispatch.
type ITipodatoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipo returns the tipo attribute.
	GetTipo() environment.TipoExpresion

	// SetTipo sets the tipo attribute.
	SetTipo(environment.TipoExpresion)

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHARACT() antlr.TerminalNode

	// IsTipodatoContext differentiates from other interfaces.
	IsTipodatoContext()
}

type TipodatoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	tipo   environment.TipoExpresion
}

func NewEmptyTipodatoContext() *TipodatoContext {
	var p = new(TipodatoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_tipodato
	return p
}

func InitEmptyTipodatoContext(p *TipodatoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_tipodato
}

func (*TipodatoContext) IsTipodatoContext() {}

func NewTipodatoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipodatoContext {
	var p = new(TipodatoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_tipodato

	return p
}

func (s *TipodatoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipodatoContext) GetTipo() environment.TipoExpresion { return s.tipo }

func (s *TipodatoContext) SetTipo(v environment.TipoExpresion) { s.tipo = v }

func (s *TipodatoContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *TipodatoContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *TipodatoContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *TipodatoContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBOOL, 0)
}

func (s *TipodatoContext) CHARACT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCHARACT, 0)
}

func (s *TipodatoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipodatoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipodatoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTipodato(s)
	}
}

func (s *TipodatoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTipodato(s)
	}
}

func (p *SwiftGrammarParser) Tipodato() (localctx ITipodatoContext) {
	localctx = NewTipodatoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_tipodato)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.Match(SwiftGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TipodatoContext).tipo = environment.INTEGER

	case SwiftGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Match(SwiftGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TipodatoContext).tipo = environment.FLOAT

	case SwiftGrammarParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(294)
			p.Match(SwiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TipodatoContext).tipo = environment.STRING

	case SwiftGrammarParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(296)
			p.Match(SwiftGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TipodatoContext).tipo = environment.BOOLEAN

	case SwiftGrammarParserCHARACT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(298)
			p.Match(SwiftGrammarParserCHARACT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TipodatoContext).tipo = environment.CHARACTER

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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_TRU returns the _TRU token.
	Get_TRU() antlr.Token

	// Get_FAL returns the _FAL token.
	Get_FAL() antlr.Token

	// Get_CHARACTER returns the _CHARACTER token.
	Get_CHARACTER() antlr.Token

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Get_NULO returns the _NULO token.
	Get_NULO() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_TRU sets the _TRU token.
	Set_TRU(antlr.Token)

	// Set_FAL sets the _FAL token.
	Set_FAL(antlr.Token)

	// Set_CHARACTER sets the _CHARACTER token.
	Set_CHARACTER(antlr.Token)

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// Set_NULO sets the _NULO token.
	Set_NULO(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_vectorvacio returns the _vectorvacio rule contexts.
	Get_vectorvacio() IVectorvacioContext

	// Get_vectorcount returns the _vectorcount rule contexts.
	Get_vectorcount() IVectorcountContext

	// Get_vectoraccess returns the _vectoraccess rule contexts.
	Get_vectoraccess() IVectoraccessContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_vectorvacio sets the _vectorvacio rule contexts.
	Set_vectorvacio(IVectorvacioContext)

	// Set_vectorcount sets the _vectorcount rule contexts.
	Set_vectorcount(IVectorcountContext)

	// Set_vectoraccess sets the _vectoraccess rule contexts.
	Set_vectoraccess(IVectoraccessContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	NOT() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	SUB() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	CADENA() antlr.TerminalNode
	TRU() antlr.TerminalNode
	FAL() antlr.TerminalNode
	CHARACTER() antlr.TerminalNode
	ID_VALIDO() antlr.TerminalNode
	NULO() antlr.TerminalNode
	Vectorvacio() IVectorvacioContext
	Vectorcount() IVectorcountContext
	Vectoraccess() IVectoraccessContext
	MODULO() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	ADD() antlr.TerminalNode
	MAY_IG() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	MEN_IG() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	IG_IG() antlr.TerminalNode
	DIF() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	e             interfaces.Expression
	left          IExprContext
	op            antlr.Token
	right         IExprContext
	_expr         IExprContext
	_NUMBER       antlr.Token
	_CADENA       antlr.Token
	_TRU          antlr.Token
	_FAL          antlr.Token
	_CHARACTER    antlr.Token
	_ID_VALIDO    antlr.Token
	_NULO         antlr.Token
	_vectorvacio  IVectorvacioContext
	_vectorcount  IVectorcountContext
	_vectoraccess IVectoraccessContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExprContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ExprContext) Get_TRU() antlr.Token { return s._TRU }

func (s *ExprContext) Get_FAL() antlr.Token { return s._FAL }

func (s *ExprContext) Get_CHARACTER() antlr.Token { return s._CHARACTER }

func (s *ExprContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *ExprContext) Get_NULO() antlr.Token { return s._NULO }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ExprContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *ExprContext) Set_FAL(v antlr.Token) { s._FAL = v }

func (s *ExprContext) Set_CHARACTER(v antlr.Token) { s._CHARACTER = v }

func (s *ExprContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *ExprContext) Set_NULO(v antlr.Token) { s._NULO = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) Get_vectorvacio() IVectorvacioContext { return s._vectorvacio }

func (s *ExprContext) Get_vectorcount() IVectorcountContext { return s._vectorcount }

func (s *ExprContext) Get_vectoraccess() IVectoraccessContext { return s._vectoraccess }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) Set_vectorvacio(v IVectorvacioContext) { s._vectorvacio = v }

func (s *ExprContext) Set_vectorcount(v IVectorcountContext) { s._vectorcount = v }

func (s *ExprContext) Set_vectoraccess(v IVectoraccessContext) { s._vectoraccess = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNOT, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
}

func (s *ExprContext) CADENA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCADENA, 0)
}

func (s *ExprContext) TRU() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTRU, 0)
}

func (s *ExprContext) FAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFAL, 0)
}

func (s *ExprContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCHARACTER, 0)
}

func (s *ExprContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, 0)
}

func (s *ExprContext) NULO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNULO, 0)
}

func (s *ExprContext) Vectorvacio() IVectorvacioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorvacioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorvacioContext)
}

func (s *ExprContext) Vectorcount() IVectorcountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorcountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorcountContext)
}

func (s *ExprContext) Vectoraccess() IVectoraccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectoraccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectoraccessContext)
}

func (s *ExprContext) MODULO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMODULO, 0)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIV, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *ExprContext) MAY_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAY_IG, 0)
}

func (s *ExprContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYOR, 0)
}

func (s *ExprContext) MEN_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMEN_IG, 0)
}

func (s *ExprContext) MENOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOR, 0)
}

func (s *ExprContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG_IG, 0)
}

func (s *ExprContext) DIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIF, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *SwiftGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(303)

			var _m = p.Match(SwiftGrammarParserNOT)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)

			var _x = p.expr(21)

			localctx.(*ExprContext).right = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
			if localctx.(*ExprContext).GetRight() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetRight().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ExprContext).GetRight() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetRight().GetStart()
			}
		}()).GetColumn(), localctx.(*ExprContext).GetRight().GetE(), (func() string {
			if localctx.(*ExprContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetOp().GetText()
			}
		}()), localctx.(*ExprContext).GetRight().GetE())

	case 2:
		{
			p.SetState(307)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(309)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 3:
		{
			p.SetState(312)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*ExprContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if strings.Contains((func() string {
			if localctx.(*ExprContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			num2 := fmt.Sprintf("%.6f", num)
			num3, err := strconv.ParseFloat(num2, 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), -num3, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), -num, environment.INTEGER)
		}

	case 4:
		{
			p.SetState(315)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*ExprContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if strings.Contains((func() string {
			if localctx.(*ExprContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			num2 := fmt.Sprintf("%.6f", num)
			num3, err := strconv.ParseFloat(num2, 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num3, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.INTEGER)
		}

	case 5:
		{
			p.SetState(317)

			var _m = p.Match(SwiftGrammarParserCADENA)

			localctx.(*ExprContext)._CADENA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*ExprContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_CADENA().GetText()
			}
		}())
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_CADENA() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CADENA().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_CADENA() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CADENA().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case 6:
		{
			p.SetState(319)

			var _m = p.Match(SwiftGrammarParserTRU)

			localctx.(*ExprContext)._TRU = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetColumn()
			}
		}()), true, environment.BOOLEAN)

	case 7:
		{
			p.SetState(321)

			var _m = p.Match(SwiftGrammarParserFAL)

			localctx.(*ExprContext)._FAL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetColumn()
			}
		}()), false, environment.BOOLEAN)

	case 8:
		{
			p.SetState(323)

			var _m = p.Match(SwiftGrammarParserCHARACTER)

			localctx.(*ExprContext)._CHARACTER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*ExprContext).Get_CHARACTER() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_CHARACTER().GetText()
			}
		}())
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_CHARACTER() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CHARACTER().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_CHARACTER() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CHARACTER().GetColumn()
			}
		}()), str[1:len(str)-1], environment.CHARACTER)

	case 9:
		{
			p.SetState(325)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*ExprContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		id := (func() string {
			if localctx.(*ExprContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID_VALIDO().GetText()
			}
		}())
		localctx.(*ExprContext).e = sentencias.NewCallid((func() int {
			if localctx.(*ExprContext).Get_ID_VALIDO() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID_VALIDO().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID_VALIDO() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID_VALIDO().GetColumn()
			}
		}()), id)

	case 10:
		{
			p.SetState(327)

			var _m = p.Match(SwiftGrammarParserNULO)

			localctx.(*ExprContext)._NULO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_NULO() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NULO().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_NULO() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NULO().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_NULO() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NULO().GetText()
			}
		}()), environment.NULL)

	case 11:
		{
			p.SetState(329)

			var _x = p.Vectorvacio()

			localctx.(*ExprContext)._vectorvacio = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_vectorvacio().GetVeemct()

	case 12:
		{
			p.SetState(332)

			var _x = p.Vectorcount()

			localctx.(*ExprContext)._vectorcount = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_vectorcount().GetVecnct()

	case 13:
		{
			p.SetState(335)

			var _x = p.Vectoraccess()

			localctx.(*ExprContext)._vectoraccess = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_vectoraccess().GetVepposct()

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(380)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(340)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(341)

					var _m = p.Match(SwiftGrammarParserMODULO)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(342)

					var _x = p.expr(21)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(345)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(346)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMUL || _la == SwiftGrammarParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(347)

					var _x = p.expr(20)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(350)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(351)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserADD || _la == SwiftGrammarParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(352)

					var _x = p.expr(19)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(355)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(356)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMAY_IG || _la == SwiftGrammarParserMAYOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(357)

					var _x = p.expr(18)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(360)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(361)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMEN_IG || _la == SwiftGrammarParserMENOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(362)

					var _x = p.expr(17)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(365)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(366)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserDIF || _la == SwiftGrammarParserIG_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(367)

					var _x = p.expr(16)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(370)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(371)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(372)

					var _x = p.expr(15)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(375)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(376)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(377)

					var _x = p.expr(14)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
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

// ISentenciaifelseContext is an interface to support dynamic dispatch.
type ISentenciaifelseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_blockinterno returns the _blockinterno rule contexts.
	Get_blockinterno() IBlockinternoContext

	// GetIfop returns the ifop rule contexts.
	GetIfop() IBlockinternoContext

	// GetElseop returns the elseop rule contexts.
	GetElseop() IBlockinternoContext

	// Get_sentenciaifelse returns the _sentenciaifelse rule contexts.
	Get_sentenciaifelse() ISentenciaifelseContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_blockinterno sets the _blockinterno rule contexts.
	Set_blockinterno(IBlockinternoContext)

	// SetIfop sets the ifop rule contexts.
	SetIfop(IBlockinternoContext)

	// SetElseop sets the elseop rule contexts.
	SetElseop(IBlockinternoContext)

	// Set_sentenciaifelse sets the _sentenciaifelse rule contexts.
	Set_sentenciaifelse(ISentenciaifelseContext)

	// GetMyIfElse returns the myIfElse attribute.
	GetMyIfElse() interfaces.Instruction

	// SetMyIfElse sets the myIfElse attribute.
	SetMyIfElse(interfaces.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllLLAVEIZQ() []antlr.TerminalNode
	LLAVEIZQ(i int) antlr.TerminalNode
	AllBlockinterno() []IBlockinternoContext
	Blockinterno(i int) IBlockinternoContext
	AllLLAVEDER() []antlr.TerminalNode
	LLAVEDER(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode
	Sentenciaifelse() ISentenciaifelseContext

	// IsSentenciaifelseContext differentiates from other interfaces.
	IsSentenciaifelseContext()
}

type SentenciaifelseContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	myIfElse         interfaces.Instruction
	_IF              antlr.Token
	_expr            IExprContext
	_blockinterno    IBlockinternoContext
	ifop             IBlockinternoContext
	elseop           IBlockinternoContext
	_sentenciaifelse ISentenciaifelseContext
}

func NewEmptySentenciaifelseContext() *SentenciaifelseContext {
	var p = new(SentenciaifelseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_sentenciaifelse
	return p
}

func InitEmptySentenciaifelseContext(p *SentenciaifelseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_sentenciaifelse
}

func (*SentenciaifelseContext) IsSentenciaifelseContext() {}

func NewSentenciaifelseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaifelseContext {
	var p = new(SentenciaifelseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_sentenciaifelse

	return p
}

func (s *SentenciaifelseContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaifelseContext) Get_IF() antlr.Token { return s._IF }

func (s *SentenciaifelseContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *SentenciaifelseContext) Get_expr() IExprContext { return s._expr }

func (s *SentenciaifelseContext) Get_blockinterno() IBlockinternoContext { return s._blockinterno }

func (s *SentenciaifelseContext) GetIfop() IBlockinternoContext { return s.ifop }

func (s *SentenciaifelseContext) GetElseop() IBlockinternoContext { return s.elseop }

func (s *SentenciaifelseContext) Get_sentenciaifelse() ISentenciaifelseContext {
	return s._sentenciaifelse
}

func (s *SentenciaifelseContext) Set_expr(v IExprContext) { s._expr = v }

func (s *SentenciaifelseContext) Set_blockinterno(v IBlockinternoContext) { s._blockinterno = v }

func (s *SentenciaifelseContext) SetIfop(v IBlockinternoContext) { s.ifop = v }

func (s *SentenciaifelseContext) SetElseop(v IBlockinternoContext) { s.elseop = v }

func (s *SentenciaifelseContext) Set_sentenciaifelse(v ISentenciaifelseContext) {
	s._sentenciaifelse = v
}

func (s *SentenciaifelseContext) GetMyIfElse() interfaces.Instruction { return s.myIfElse }

func (s *SentenciaifelseContext) SetMyIfElse(v interfaces.Instruction) { s.myIfElse = v }

func (s *SentenciaifelseContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *SentenciaifelseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SentenciaifelseContext) AllLLAVEIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEIZQ)
}

func (s *SentenciaifelseContext) LLAVEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, i)
}

func (s *SentenciaifelseContext) AllBlockinterno() []IBlockinternoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockinternoContext); ok {
			len++
		}
	}

	tst := make([]IBlockinternoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockinternoContext); ok {
			tst[i] = t.(IBlockinternoContext)
			i++
		}
	}

	return tst
}

func (s *SentenciaifelseContext) Blockinterno(i int) IBlockinternoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockinternoContext); ok {
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

	return t.(IBlockinternoContext)
}

func (s *SentenciaifelseContext) AllLLAVEDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEDER)
}

func (s *SentenciaifelseContext) LLAVEDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, i)
}

func (s *SentenciaifelseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *SentenciaifelseContext) Sentenciaifelse() ISentenciaifelseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaifelseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaifelseContext)
}

func (s *SentenciaifelseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaifelseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenciaifelseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterSentenciaifelse(s)
	}
}

func (s *SentenciaifelseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitSentenciaifelse(s)
	}
}

func (p *SwiftGrammarParser) Sentenciaifelse() (localctx ISentenciaifelseContext) {
	localctx = NewSentenciaifelseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_sentenciaifelse)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*SentenciaifelseContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)

			var _x = p.expr(0)

			localctx.(*SentenciaifelseContext)._expr = _x
		}
		{
			p.SetState(387)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)

			var _x = p.Blockinterno()

			localctx.(*SentenciaifelseContext)._blockinterno = _x
		}
		{
			p.SetState(389)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SentenciaifelseContext).myIfElse = sentencias.NewSentenciaIf((func() int {
			if localctx.(*SentenciaifelseContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*SentenciaifelseContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*SentenciaifelseContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*SentenciaifelseContext).Get_IF().GetColumn()
			}
		}()), localctx.(*SentenciaifelseContext).Get_expr().GetE(), localctx.(*SentenciaifelseContext).Get_blockinterno().GetBlkint())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(392)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*SentenciaifelseContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)

			var _x = p.expr(0)

			localctx.(*SentenciaifelseContext)._expr = _x
		}
		{
			p.SetState(394)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)

			var _x = p.Blockinterno()

			localctx.(*SentenciaifelseContext).ifop = _x
		}
		{
			p.SetState(396)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)

			var _x = p.Blockinterno()

			localctx.(*SentenciaifelseContext).elseop = _x
		}
		{
			p.SetState(400)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SentenciaifelseContext).myIfElse = sentencias.NewSentenciaIfElse((func() int {
			if localctx.(*SentenciaifelseContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*SentenciaifelseContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*SentenciaifelseContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*SentenciaifelseContext).Get_IF().GetColumn()
			}
		}()), localctx.(*SentenciaifelseContext).Get_expr().GetE(), localctx.(*SentenciaifelseContext).GetIfop().GetBlkint(), localctx.(*SentenciaifelseContext).GetElseop().GetBlkint())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(403)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*SentenciaifelseContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)

			var _x = p.expr(0)

			localctx.(*SentenciaifelseContext)._expr = _x
		}
		{
			p.SetState(405)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)

			var _x = p.Blockinterno()

			localctx.(*SentenciaifelseContext)._blockinterno = _x
		}
		{
			p.SetState(407)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)

			var _x = p.Sentenciaifelse()

			localctx.(*SentenciaifelseContext)._sentenciaifelse = _x
		}
		localctx.(*SentenciaifelseContext).myIfElse = sentencias.NewSentenciaIfElseIf((func() int {
			if localctx.(*SentenciaifelseContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*SentenciaifelseContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*SentenciaifelseContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*SentenciaifelseContext).Get_IF().GetColumn()
			}
		}()), localctx.(*SentenciaifelseContext).Get_expr().GetE(), localctx.(*SentenciaifelseContext).Get_blockinterno().GetBlkint(), localctx.(*SentenciaifelseContext).Get_sentenciaifelse().GetMyIfElse())

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

// ISwitchcontrolContext is an interface to support dynamic dispatch.
type ISwitchcontrolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SWITCH returns the _SWITCH token.
	Get_SWITCH() antlr.Token

	// Get_DEFAULT returns the _DEFAULT token.
	Get_DEFAULT() antlr.Token

	// Set_SWITCH sets the _SWITCH token.
	Set_SWITCH(antlr.Token)

	// Set_DEFAULT sets the _DEFAULT token.
	Set_DEFAULT(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_blockcase returns the _blockcase rule contexts.
	Get_blockcase() IBlockcaseContext

	// Get_blockinterno returns the _blockinterno rule contexts.
	Get_blockinterno() IBlockinternoContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_blockcase sets the _blockcase rule contexts.
	Set_blockcase(IBlockcaseContext)

	// Set_blockinterno sets the _blockinterno rule contexts.
	Set_blockinterno(IBlockinternoContext)

	// GetMySwitch returns the mySwitch attribute.
	GetMySwitch() interfaces.Instruction

	// SetMySwitch sets the mySwitch attribute.
	SetMySwitch(interfaces.Instruction)

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Blockcase() IBlockcaseContext
	LLAVEDER() antlr.TerminalNode
	DEFAULT() antlr.TerminalNode
	DOS_PUNTOS() antlr.TerminalNode
	Blockinterno() IBlockinternoContext

	// IsSwitchcontrolContext differentiates from other interfaces.
	IsSwitchcontrolContext()
}

type SwitchcontrolContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	mySwitch      interfaces.Instruction
	_SWITCH       antlr.Token
	_expr         IExprContext
	_blockcase    IBlockcaseContext
	_DEFAULT      antlr.Token
	_blockinterno IBlockinternoContext
}

func NewEmptySwitchcontrolContext() *SwitchcontrolContext {
	var p = new(SwitchcontrolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchcontrol
	return p
}

func InitEmptySwitchcontrolContext(p *SwitchcontrolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchcontrol
}

func (*SwitchcontrolContext) IsSwitchcontrolContext() {}

func NewSwitchcontrolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchcontrolContext {
	var p = new(SwitchcontrolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_switchcontrol

	return p
}

func (s *SwitchcontrolContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchcontrolContext) Get_SWITCH() antlr.Token { return s._SWITCH }

func (s *SwitchcontrolContext) Get_DEFAULT() antlr.Token { return s._DEFAULT }

func (s *SwitchcontrolContext) Set_SWITCH(v antlr.Token) { s._SWITCH = v }

func (s *SwitchcontrolContext) Set_DEFAULT(v antlr.Token) { s._DEFAULT = v }

func (s *SwitchcontrolContext) Get_expr() IExprContext { return s._expr }

func (s *SwitchcontrolContext) Get_blockcase() IBlockcaseContext { return s._blockcase }

func (s *SwitchcontrolContext) Get_blockinterno() IBlockinternoContext { return s._blockinterno }

func (s *SwitchcontrolContext) Set_expr(v IExprContext) { s._expr = v }

func (s *SwitchcontrolContext) Set_blockcase(v IBlockcaseContext) { s._blockcase = v }

func (s *SwitchcontrolContext) Set_blockinterno(v IBlockinternoContext) { s._blockinterno = v }

func (s *SwitchcontrolContext) GetMySwitch() interfaces.Instruction { return s.mySwitch }

func (s *SwitchcontrolContext) SetMySwitch(v interfaces.Instruction) { s.mySwitch = v }

func (s *SwitchcontrolContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSWITCH, 0)
}

func (s *SwitchcontrolContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchcontrolContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *SwitchcontrolContext) Blockcase() IBlockcaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockcaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockcaseContext)
}

func (s *SwitchcontrolContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *SwitchcontrolContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDEFAULT, 0)
}

func (s *SwitchcontrolContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOS_PUNTOS, 0)
}

func (s *SwitchcontrolContext) Blockinterno() IBlockinternoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockinternoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockinternoContext)
}

func (s *SwitchcontrolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchcontrolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchcontrolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterSwitchcontrol(s)
	}
}

func (s *SwitchcontrolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitSwitchcontrol(s)
	}
}

func (p *SwiftGrammarParser) Switchcontrol() (localctx ISwitchcontrolContext) {
	localctx = NewSwitchcontrolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_switchcontrol)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)

		var _m = p.Match(SwiftGrammarParserSWITCH)

		localctx.(*SwitchcontrolContext)._SWITCH = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)

		var _x = p.expr(0)

		localctx.(*SwitchcontrolContext)._expr = _x
	}
	{
		p.SetState(416)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(417)

		var _x = p.Blockcase()

		localctx.(*SwitchcontrolContext)._blockcase = _x
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserDEFAULT {
		{
			p.SetState(418)

			var _m = p.Match(SwiftGrammarParserDEFAULT)

			localctx.(*SwitchcontrolContext)._DEFAULT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)

			var _x = p.Blockinterno()

			localctx.(*SwitchcontrolContext)._blockinterno = _x
		}

	}
	{
		p.SetState(423)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	if localctx.(*SwitchcontrolContext).Get_DEFAULT() != nil {
		localctx.(*SwitchcontrolContext).mySwitch = sentencias.NewSentenciaSwitchDefault((func() int {
			if localctx.(*SwitchcontrolContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchcontrolContext).Get_SWITCH().GetLine()
			}
		}()), (func() int {
			if localctx.(*SwitchcontrolContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchcontrolContext).Get_SWITCH().GetColumn()
			}
		}()), localctx.(*SwitchcontrolContext).Get_expr().GetE(), localctx.(*SwitchcontrolContext).Get_blockcase().GetBlkcase(), localctx.(*SwitchcontrolContext).Get_blockinterno().GetBlkint())
	} else {
		localctx.(*SwitchcontrolContext).mySwitch = sentencias.NewSentenciaSwitch((func() int {
			if localctx.(*SwitchcontrolContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchcontrolContext).Get_SWITCH().GetLine()
			}
		}()), (func() int {
			if localctx.(*SwitchcontrolContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchcontrolContext).Get_SWITCH().GetColumn()
			}
		}()), localctx.(*SwitchcontrolContext).Get_expr().GetE(), localctx.(*SwitchcontrolContext).Get_blockcase().GetBlkcase())
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

// IBlockcaseContext is an interface to support dynamic dispatch.
type IBlockcaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_bloquecase returns the _bloquecase rule contexts.
	Get_bloquecase() IBloquecaseContext

	// Set_bloquecase sets the _bloquecase rule contexts.
	Set_bloquecase(IBloquecaseContext)

	// GetBlocas returns the blocas rule context list.
	GetBlocas() []IBloquecaseContext

	// SetBlocas sets the blocas rule context list.
	SetBlocas([]IBloquecaseContext)

	// GetBlkcase returns the blkcase attribute.
	GetBlkcase() []interface{}

	// SetBlkcase sets the blkcase attribute.
	SetBlkcase([]interface{})

	// Getter signatures
	AllBloquecase() []IBloquecaseContext
	Bloquecase(i int) IBloquecaseContext

	// IsBlockcaseContext differentiates from other interfaces.
	IsBlockcaseContext()
}

type BlockcaseContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	blkcase     []interface{}
	_bloquecase IBloquecaseContext
	blocas      []IBloquecaseContext
}

func NewEmptyBlockcaseContext() *BlockcaseContext {
	var p = new(BlockcaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_blockcase
	return p
}

func InitEmptyBlockcaseContext(p *BlockcaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_blockcase
}

func (*BlockcaseContext) IsBlockcaseContext() {}

func NewBlockcaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockcaseContext {
	var p = new(BlockcaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_blockcase

	return p
}

func (s *BlockcaseContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockcaseContext) Get_bloquecase() IBloquecaseContext { return s._bloquecase }

func (s *BlockcaseContext) Set_bloquecase(v IBloquecaseContext) { s._bloquecase = v }

func (s *BlockcaseContext) GetBlocas() []IBloquecaseContext { return s.blocas }

func (s *BlockcaseContext) SetBlocas(v []IBloquecaseContext) { s.blocas = v }

func (s *BlockcaseContext) GetBlkcase() []interface{} { return s.blkcase }

func (s *BlockcaseContext) SetBlkcase(v []interface{}) { s.blkcase = v }

func (s *BlockcaseContext) AllBloquecase() []IBloquecaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBloquecaseContext); ok {
			len++
		}
	}

	tst := make([]IBloquecaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBloquecaseContext); ok {
			tst[i] = t.(IBloquecaseContext)
			i++
		}
	}

	return tst
}

func (s *BlockcaseContext) Bloquecase(i int) IBloquecaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloquecaseContext); ok {
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

	return t.(IBloquecaseContext)
}

func (s *BlockcaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockcaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockcaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlockcase(s)
	}
}

func (s *BlockcaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlockcase(s)
	}
}

func (p *SwiftGrammarParser) Blockcase() (localctx IBlockcaseContext) {
	localctx = NewBlockcaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_blockcase)

	localctx.(*BlockcaseContext).blkcase = []interface{}{}
	var listInt []IBloquecaseContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftGrammarParserCASE {
		{
			p.SetState(426)

			var _x = p.Bloquecase()

			localctx.(*BlockcaseContext)._bloquecase = _x
		}
		localctx.(*BlockcaseContext).blocas = append(localctx.(*BlockcaseContext).blocas, localctx.(*BlockcaseContext)._bloquecase)

		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt = localctx.(*BlockcaseContext).GetBlocas()
	for _, e := range listInt {
		localctx.(*BlockcaseContext).blkcase = append(localctx.(*BlockcaseContext).blkcase, e.GetBlocas())
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

// IBloquecaseContext is an interface to support dynamic dispatch.
type IBloquecaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CASE returns the _CASE token.
	Get_CASE() antlr.Token

	// Set_CASE sets the _CASE token.
	Set_CASE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_blockinterno returns the _blockinterno rule contexts.
	Get_blockinterno() IBlockinternoContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_blockinterno sets the _blockinterno rule contexts.
	Set_blockinterno(IBlockinternoContext)

	// GetBlocas returns the blocas attribute.
	GetBlocas() interfaces.Instruction

	// SetBlocas sets the blocas attribute.
	SetBlocas(interfaces.Instruction)

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	DOS_PUNTOS() antlr.TerminalNode
	Blockinterno() IBlockinternoContext

	// IsBloquecaseContext differentiates from other interfaces.
	IsBloquecaseContext()
}

type BloquecaseContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	blocas        interfaces.Instruction
	_CASE         antlr.Token
	_expr         IExprContext
	_blockinterno IBlockinternoContext
}

func NewEmptyBloquecaseContext() *BloquecaseContext {
	var p = new(BloquecaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_bloquecase
	return p
}

func InitEmptyBloquecaseContext(p *BloquecaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_bloquecase
}

func (*BloquecaseContext) IsBloquecaseContext() {}

func NewBloquecaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloquecaseContext {
	var p = new(BloquecaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_bloquecase

	return p
}

func (s *BloquecaseContext) GetParser() antlr.Parser { return s.parser }

func (s *BloquecaseContext) Get_CASE() antlr.Token { return s._CASE }

func (s *BloquecaseContext) Set_CASE(v antlr.Token) { s._CASE = v }

func (s *BloquecaseContext) Get_expr() IExprContext { return s._expr }

func (s *BloquecaseContext) Get_blockinterno() IBlockinternoContext { return s._blockinterno }

func (s *BloquecaseContext) Set_expr(v IExprContext) { s._expr = v }

func (s *BloquecaseContext) Set_blockinterno(v IBlockinternoContext) { s._blockinterno = v }

func (s *BloquecaseContext) GetBlocas() interfaces.Instruction { return s.blocas }

func (s *BloquecaseContext) SetBlocas(v interfaces.Instruction) { s.blocas = v }

func (s *BloquecaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCASE, 0)
}

func (s *BloquecaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BloquecaseContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOS_PUNTOS, 0)
}

func (s *BloquecaseContext) Blockinterno() IBlockinternoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockinternoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockinternoContext)
}

func (s *BloquecaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloquecaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloquecaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBloquecase(s)
	}
}

func (s *BloquecaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBloquecase(s)
	}
}

func (p *SwiftGrammarParser) Bloquecase() (localctx IBloquecaseContext) {
	localctx = NewBloquecaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_bloquecase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(433)

		var _m = p.Match(SwiftGrammarParserCASE)

		localctx.(*BloquecaseContext)._CASE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(434)

		var _x = p.expr(0)

		localctx.(*BloquecaseContext)._expr = _x
	}
	{
		p.SetState(435)
		p.Match(SwiftGrammarParserDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(436)

		var _x = p.Blockinterno()

		localctx.(*BloquecaseContext)._blockinterno = _x
	}

	localctx.(*BloquecaseContext).blocas = sentencias.NewSentenciaSwitchCase((func() int {
		if localctx.(*BloquecaseContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*BloquecaseContext).Get_CASE().GetLine()
		}
	}()), (func() int {
		if localctx.(*BloquecaseContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*BloquecaseContext).Get_CASE().GetColumn()
		}
	}()), localctx.(*BloquecaseContext).Get_expr().GetE(), localctx.(*BloquecaseContext).Get_blockinterno().GetBlkint())

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

// IWhilecontrolContext is an interface to support dynamic dispatch.
type IWhilecontrolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_blockinterno returns the _blockinterno rule contexts.
	Get_blockinterno() IBlockinternoContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_blockinterno sets the _blockinterno rule contexts.
	Set_blockinterno(IBlockinternoContext)

	// GetWhict returns the whict attribute.
	GetWhict() interfaces.Instruction

	// SetWhict sets the whict attribute.
	SetWhict(interfaces.Instruction)

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Blockinterno() IBlockinternoContext
	LLAVEDER() antlr.TerminalNode

	// IsWhilecontrolContext differentiates from other interfaces.
	IsWhilecontrolContext()
}

type WhilecontrolContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	whict         interfaces.Instruction
	_WHILE        antlr.Token
	_expr         IExprContext
	_blockinterno IBlockinternoContext
}

func NewEmptyWhilecontrolContext() *WhilecontrolContext {
	var p = new(WhilecontrolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilecontrol
	return p
}

func InitEmptyWhilecontrolContext(p *WhilecontrolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilecontrol
}

func (*WhilecontrolContext) IsWhilecontrolContext() {}

func NewWhilecontrolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilecontrolContext {
	var p = new(WhilecontrolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_whilecontrol

	return p
}

func (s *WhilecontrolContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilecontrolContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *WhilecontrolContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *WhilecontrolContext) Get_expr() IExprContext { return s._expr }

func (s *WhilecontrolContext) Get_blockinterno() IBlockinternoContext { return s._blockinterno }

func (s *WhilecontrolContext) Set_expr(v IExprContext) { s._expr = v }

func (s *WhilecontrolContext) Set_blockinterno(v IBlockinternoContext) { s._blockinterno = v }

func (s *WhilecontrolContext) GetWhict() interfaces.Instruction { return s.whict }

func (s *WhilecontrolContext) SetWhict(v interfaces.Instruction) { s.whict = v }

func (s *WhilecontrolContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserWHILE, 0)
}

func (s *WhilecontrolContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhilecontrolContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *WhilecontrolContext) Blockinterno() IBlockinternoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockinternoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockinternoContext)
}

func (s *WhilecontrolContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *WhilecontrolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilecontrolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilecontrolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterWhilecontrol(s)
	}
}

func (s *WhilecontrolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitWhilecontrol(s)
	}
}

func (p *SwiftGrammarParser) Whilecontrol() (localctx IWhilecontrolContext) {
	localctx = NewWhilecontrolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_whilecontrol)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)

		var _m = p.Match(SwiftGrammarParserWHILE)

		localctx.(*WhilecontrolContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)

		var _x = p.expr(0)

		localctx.(*WhilecontrolContext)._expr = _x
	}
	{
		p.SetState(441)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(442)

		var _x = p.Blockinterno()

		localctx.(*WhilecontrolContext)._blockinterno = _x
	}
	{
		p.SetState(443)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*WhilecontrolContext).whict = sentencias.NewSentenciaWhile((func() int {
		if localctx.(*WhilecontrolContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilecontrolContext).Get_WHILE().GetLine()
		}
	}()), (func() int {
		if localctx.(*WhilecontrolContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilecontrolContext).Get_WHILE().GetColumn()
		}
	}()), localctx.(*WhilecontrolContext).Get_expr().GetE(), localctx.(*WhilecontrolContext).Get_blockinterno().GetBlkint())

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

// IForcontrolContext is an interface to support dynamic dispatch.
type IForcontrolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FOR returns the _FOR token.
	Get_FOR() antlr.Token

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// GetOp1 returns the op1 token.
	GetOp1() antlr.Token

	// GetOp2 returns the op2 token.
	GetOp2() antlr.Token

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// SetOp1 sets the op1 token.
	SetOp1(antlr.Token)

	// SetOp2 sets the op2 token.
	SetOp2(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// Get_blockinterno returns the _blockinterno rule contexts.
	Get_blockinterno() IBlockinternoContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// Set_blockinterno sets the _blockinterno rule contexts.
	Set_blockinterno(IBlockinternoContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetForct returns the forct attribute.
	GetForct() interfaces.Instruction

	// SetForct sets the forct attribute.
	SetForct(interfaces.Instruction)

	// Getter signatures
	FOR() antlr.TerminalNode
	AllID_VALIDO() []antlr.TerminalNode
	ID_VALIDO(i int) antlr.TerminalNode
	IN() antlr.TerminalNode
	RANGO() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Blockinterno() IBlockinternoContext
	LLAVEDER() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsForcontrolContext differentiates from other interfaces.
	IsForcontrolContext()
}

type ForcontrolContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	forct         interfaces.Instruction
	_FOR          antlr.Token
	_ID_VALIDO    antlr.Token
	left          IExprContext
	right         IExprContext
	_blockinterno IBlockinternoContext
	op1           antlr.Token
	op2           antlr.Token
	_expr         IExprContext
}

func NewEmptyForcontrolContext() *ForcontrolContext {
	var p = new(ForcontrolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forcontrol
	return p
}

func InitEmptyForcontrolContext(p *ForcontrolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forcontrol
}

func (*ForcontrolContext) IsForcontrolContext() {}

func NewForcontrolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForcontrolContext {
	var p = new(ForcontrolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_forcontrol

	return p
}

func (s *ForcontrolContext) GetParser() antlr.Parser { return s.parser }

func (s *ForcontrolContext) Get_FOR() antlr.Token { return s._FOR }

func (s *ForcontrolContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *ForcontrolContext) GetOp1() antlr.Token { return s.op1 }

func (s *ForcontrolContext) GetOp2() antlr.Token { return s.op2 }

func (s *ForcontrolContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *ForcontrolContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *ForcontrolContext) SetOp1(v antlr.Token) { s.op1 = v }

func (s *ForcontrolContext) SetOp2(v antlr.Token) { s.op2 = v }

func (s *ForcontrolContext) GetLeft() IExprContext { return s.left }

func (s *ForcontrolContext) GetRight() IExprContext { return s.right }

func (s *ForcontrolContext) Get_blockinterno() IBlockinternoContext { return s._blockinterno }

func (s *ForcontrolContext) Get_expr() IExprContext { return s._expr }

func (s *ForcontrolContext) SetLeft(v IExprContext) { s.left = v }

func (s *ForcontrolContext) SetRight(v IExprContext) { s.right = v }

func (s *ForcontrolContext) Set_blockinterno(v IBlockinternoContext) { s._blockinterno = v }

func (s *ForcontrolContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ForcontrolContext) GetForct() interfaces.Instruction { return s.forct }

func (s *ForcontrolContext) SetForct(v interfaces.Instruction) { s.forct = v }

func (s *ForcontrolContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFOR, 0)
}

func (s *ForcontrolContext) AllID_VALIDO() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID_VALIDO)
}

func (s *ForcontrolContext) ID_VALIDO(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, i)
}

func (s *ForcontrolContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIN, 0)
}

func (s *ForcontrolContext) RANGO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRANGO, 0)
}

func (s *ForcontrolContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ForcontrolContext) Blockinterno() IBlockinternoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockinternoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockinternoContext)
}

func (s *ForcontrolContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ForcontrolContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForcontrolContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ForcontrolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForcontrolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForcontrolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterForcontrol(s)
	}
}

func (s *ForcontrolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitForcontrol(s)
	}
}

func (p *SwiftGrammarParser) Forcontrol() (localctx IForcontrolContext) {
	localctx = NewForcontrolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_forcontrol)
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(446)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForcontrolContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*ForcontrolContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)

			var _x = p.expr(0)

			localctx.(*ForcontrolContext).left = _x
		}
		{
			p.SetState(450)
			p.Match(SwiftGrammarParserRANGO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)

			var _x = p.expr(0)

			localctx.(*ForcontrolContext).right = _x
		}
		{
			p.SetState(452)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)

			var _x = p.Blockinterno()

			localctx.(*ForcontrolContext)._blockinterno = _x
		}
		{
			p.SetState(454)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForcontrolContext).forct = sentencias.NewSentenciaForRango((func() int {
			if localctx.(*ForcontrolContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForcontrolContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForcontrolContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForcontrolContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForcontrolContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*ForcontrolContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*ForcontrolContext).GetLeft().GetE(), localctx.(*ForcontrolContext).GetRight().GetE(), localctx.(*ForcontrolContext).Get_blockinterno().GetBlkint())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(457)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForcontrolContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*ForcontrolContext).op1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*ForcontrolContext).op2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)

			var _x = p.Blockinterno()

			localctx.(*ForcontrolContext)._blockinterno = _x
		}
		{
			p.SetState(463)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForcontrolContext).forct = sentencias.NewSentenciaForId((func() int {
			if localctx.(*ForcontrolContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForcontrolContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForcontrolContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForcontrolContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForcontrolContext).GetOp1() == nil {
				return ""
			} else {
				return localctx.(*ForcontrolContext).GetOp1().GetText()
			}
		}()), (func() string {
			if localctx.(*ForcontrolContext).GetOp2() == nil {
				return ""
			} else {
				return localctx.(*ForcontrolContext).GetOp2().GetText()
			}
		}()), localctx.(*ForcontrolContext).Get_blockinterno().GetBlkint())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(466)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForcontrolContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*ForcontrolContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)

			var _x = p.expr(0)

			localctx.(*ForcontrolContext)._expr = _x
		}
		{
			p.SetState(470)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)

			var _x = p.Blockinterno()

			localctx.(*ForcontrolContext)._blockinterno = _x
		}
		{
			p.SetState(472)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForcontrolContext).forct = sentencias.NewSentenciaForCadena((func() int {
			if localctx.(*ForcontrolContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForcontrolContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForcontrolContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForcontrolContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForcontrolContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*ForcontrolContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*ForcontrolContext).Get_expr().GetE(), localctx.(*ForcontrolContext).Get_blockinterno().GetBlkint())

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

// IGuardcontrolContext is an interface to support dynamic dispatch.
type IGuardcontrolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_GUARD returns the _GUARD token.
	Get_GUARD() antlr.Token

	// Set_GUARD sets the _GUARD token.
	Set_GUARD(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_blockinterno returns the _blockinterno rule contexts.
	Get_blockinterno() IBlockinternoContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_blockinterno sets the _blockinterno rule contexts.
	Set_blockinterno(IBlockinternoContext)

	// GetGuct returns the guct attribute.
	GetGuct() interfaces.Instruction

	// SetGuct sets the guct attribute.
	SetGuct(interfaces.Instruction)

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Blockinterno() IBlockinternoContext
	LLAVEDER() antlr.TerminalNode

	// IsGuardcontrolContext differentiates from other interfaces.
	IsGuardcontrolContext()
}

type GuardcontrolContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	guct          interfaces.Instruction
	_GUARD        antlr.Token
	_expr         IExprContext
	_blockinterno IBlockinternoContext
}

func NewEmptyGuardcontrolContext() *GuardcontrolContext {
	var p = new(GuardcontrolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardcontrol
	return p
}

func InitEmptyGuardcontrolContext(p *GuardcontrolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardcontrol
}

func (*GuardcontrolContext) IsGuardcontrolContext() {}

func NewGuardcontrolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardcontrolContext {
	var p = new(GuardcontrolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_guardcontrol

	return p
}

func (s *GuardcontrolContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardcontrolContext) Get_GUARD() antlr.Token { return s._GUARD }

func (s *GuardcontrolContext) Set_GUARD(v antlr.Token) { s._GUARD = v }

func (s *GuardcontrolContext) Get_expr() IExprContext { return s._expr }

func (s *GuardcontrolContext) Get_blockinterno() IBlockinternoContext { return s._blockinterno }

func (s *GuardcontrolContext) Set_expr(v IExprContext) { s._expr = v }

func (s *GuardcontrolContext) Set_blockinterno(v IBlockinternoContext) { s._blockinterno = v }

func (s *GuardcontrolContext) GetGuct() interfaces.Instruction { return s.guct }

func (s *GuardcontrolContext) SetGuct(v interfaces.Instruction) { s.guct = v }

func (s *GuardcontrolContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGUARD, 0)
}

func (s *GuardcontrolContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardcontrolContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *GuardcontrolContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *GuardcontrolContext) Blockinterno() IBlockinternoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockinternoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockinternoContext)
}

func (s *GuardcontrolContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *GuardcontrolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardcontrolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardcontrolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterGuardcontrol(s)
	}
}

func (s *GuardcontrolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitGuardcontrol(s)
	}
}

func (p *SwiftGrammarParser) Guardcontrol() (localctx IGuardcontrolContext) {
	localctx = NewGuardcontrolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_guardcontrol)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)

		var _m = p.Match(SwiftGrammarParserGUARD)

		localctx.(*GuardcontrolContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(478)

		var _x = p.expr(0)

		localctx.(*GuardcontrolContext)._expr = _x
	}
	{
		p.SetState(479)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(480)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)

		var _x = p.Blockinterno()

		localctx.(*GuardcontrolContext)._blockinterno = _x
	}
	{
		p.SetState(482)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*GuardcontrolContext).guct = sentencias.NewSentenciaGuard((func() int {
		if localctx.(*GuardcontrolContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardcontrolContext).Get_GUARD().GetLine()
		}
	}()), (func() int {
		if localctx.(*GuardcontrolContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardcontrolContext).Get_GUARD().GetColumn()
		}
	}()), localctx.(*GuardcontrolContext).Get_expr().GetE(), localctx.(*GuardcontrolContext).Get_blockinterno().GetBlkint())

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

// IContinueeContext is an interface to support dynamic dispatch.
type IContinueeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CONTINUE returns the _CONTINUE token.
	Get_CONTINUE() antlr.Token

	// Set_CONTINUE sets the _CONTINUE token.
	Set_CONTINUE(antlr.Token)

	// GetCoct returns the coct attribute.
	GetCoct() interfaces.Instruction

	// SetCoct sets the coct attribute.
	SetCoct(interfaces.Instruction)

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueeContext differentiates from other interfaces.
	IsContinueeContext()
}

type ContinueeContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	coct      interfaces.Instruction
	_CONTINUE antlr.Token
}

func NewEmptyContinueeContext() *ContinueeContext {
	var p = new(ContinueeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_continuee
	return p
}

func InitEmptyContinueeContext(p *ContinueeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_continuee
}

func (*ContinueeContext) IsContinueeContext() {}

func NewContinueeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueeContext {
	var p = new(ContinueeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_continuee

	return p
}

func (s *ContinueeContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueeContext) Get_CONTINUE() antlr.Token { return s._CONTINUE }

func (s *ContinueeContext) Set_CONTINUE(v antlr.Token) { s._CONTINUE = v }

func (s *ContinueeContext) GetCoct() interfaces.Instruction { return s.coct }

func (s *ContinueeContext) SetCoct(v interfaces.Instruction) { s.coct = v }

func (s *ContinueeContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCONTINUE, 0)
}

func (s *ContinueeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterContinuee(s)
	}
}

func (s *ContinueeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitContinuee(s)
	}
}

func (p *SwiftGrammarParser) Continuee() (localctx IContinueeContext) {
	localctx = NewContinueeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_continuee)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)

		var _m = p.Match(SwiftGrammarParserCONTINUE)

		localctx.(*ContinueeContext)._CONTINUE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ContinueeContext).coct = sentencias.NewTransferenciaContinue((func() int {
		if localctx.(*ContinueeContext).Get_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*ContinueeContext).Get_CONTINUE().GetLine()
		}
	}()), (func() int {
		if localctx.(*ContinueeContext).Get_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*ContinueeContext).Get_CONTINUE().GetColumn()
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

// IBreakkContext is an interface to support dynamic dispatch.
type IBreakkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_BREAK returns the _BREAK token.
	Get_BREAK() antlr.Token

	// Set_BREAK sets the _BREAK token.
	Set_BREAK(antlr.Token)

	// GetBrkct returns the brkct attribute.
	GetBrkct() interfaces.Instruction

	// SetBrkct sets the brkct attribute.
	SetBrkct(interfaces.Instruction)

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakkContext differentiates from other interfaces.
	IsBreakkContext()
}

type BreakkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	brkct  interfaces.Instruction
	_BREAK antlr.Token
}

func NewEmptyBreakkContext() *BreakkContext {
	var p = new(BreakkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_breakk
	return p
}

func InitEmptyBreakkContext(p *BreakkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_breakk
}

func (*BreakkContext) IsBreakkContext() {}

func NewBreakkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakkContext {
	var p = new(BreakkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_breakk

	return p
}

func (s *BreakkContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakkContext) Get_BREAK() antlr.Token { return s._BREAK }

func (s *BreakkContext) Set_BREAK(v antlr.Token) { s._BREAK = v }

func (s *BreakkContext) GetBrkct() interfaces.Instruction { return s.brkct }

func (s *BreakkContext) SetBrkct(v interfaces.Instruction) { s.brkct = v }

func (s *BreakkContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBREAK, 0)
}

func (s *BreakkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBreakk(s)
	}
}

func (s *BreakkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBreakk(s)
	}
}

func (p *SwiftGrammarParser) Breakk() (localctx IBreakkContext) {
	localctx = NewBreakkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_breakk)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(488)

		var _m = p.Match(SwiftGrammarParserBREAK)

		localctx.(*BreakkContext)._BREAK = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*BreakkContext).brkct = sentencias.NewTransferenciaBreak((func() int {
		if localctx.(*BreakkContext).Get_BREAK() == nil {
			return 0
		} else {
			return localctx.(*BreakkContext).Get_BREAK().GetLine()
		}
	}()), (func() int {
		if localctx.(*BreakkContext).Get_BREAK() == nil {
			return 0
		} else {
			return localctx.(*BreakkContext).Get_BREAK().GetColumn()
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

// IRetornosContext is an interface to support dynamic dispatch.
type IRetornosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// GetOp returns the op rule contexts.
	GetOp() IExprContext

	// SetOp sets the op rule contexts.
	SetOp(IExprContext)

	// GetRect returns the rect attribute.
	GetRect() interfaces.Instruction

	// SetRect sets the rect attribute.
	SetRect(interfaces.Instruction)

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsRetornosContext differentiates from other interfaces.
	IsRetornosContext()
}

type RetornosContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	rect    interfaces.Instruction
	_RETURN antlr.Token
	op      IExprContext
}

func NewEmptyRetornosContext() *RetornosContext {
	var p = new(RetornosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_retornos
	return p
}

func InitEmptyRetornosContext(p *RetornosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_retornos
}

func (*RetornosContext) IsRetornosContext() {}

func NewRetornosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetornosContext {
	var p = new(RetornosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_retornos

	return p
}

func (s *RetornosContext) GetParser() antlr.Parser { return s.parser }

func (s *RetornosContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *RetornosContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *RetornosContext) GetOp() IExprContext { return s.op }

func (s *RetornosContext) SetOp(v IExprContext) { s.op = v }

func (s *RetornosContext) GetRect() interfaces.Instruction { return s.rect }

func (s *RetornosContext) SetRect(v interfaces.Instruction) { s.rect = v }

func (s *RetornosContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRETURN, 0)
}

func (s *RetornosContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RetornosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetornosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetornosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRetornos(s)
	}
}

func (s *RetornosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRetornos(s)
	}
}

func (p *SwiftGrammarParser) Retornos() (localctx IRetornosContext) {
	localctx = NewRetornosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_retornos)
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(491)

			var _m = p.Match(SwiftGrammarParserRETURN)

			localctx.(*RetornosContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(492)

			var _x = p.expr(0)

			localctx.(*RetornosContext).op = _x
		}

		localctx.(*RetornosContext).SetRect(sentencias.NewTransferenciaReturnExp((func() int {
			if localctx.(*RetornosContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*RetornosContext).Get_RETURN().GetLine()
			}
		}()), (func() int {
			if localctx.(*RetornosContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*RetornosContext).Get_RETURN().GetColumn()
			}
		}()), localctx.(*RetornosContext).GetOp().GetE()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(495)

			var _m = p.Match(SwiftGrammarParserRETURN)

			localctx.(*RetornosContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*RetornosContext).SetRect(sentencias.NewTransferenciaReturn((func() int {
			if localctx.(*RetornosContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*RetornosContext).Get_RETURN().GetLine()
			}
		}()), (func() int {
			if localctx.(*RetornosContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*RetornosContext).Get_RETURN().GetColumn()
			}
		}())))

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

// IVectorcontrolContext is an interface to support dynamic dispatch.
type IVectorcontrolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// GetPrin returns the prin token.
	GetPrin() antlr.Token

	// GetSecu returns the secu token.
	GetSecu() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// SetPrin sets the prin token.
	SetPrin(antlr.Token)

	// SetSecu sets the secu token.
	SetSecu(antlr.Token)

	// Get_tipodato returns the _tipodato rule contexts.
	Get_tipodato() ITipodatoContext

	// Get_blockparams returns the _blockparams rule contexts.
	Get_blockparams() IBlockparamsContext

	// Set_tipodato sets the _tipodato rule contexts.
	Set_tipodato(ITipodatoContext)

	// Set_blockparams sets the _blockparams rule contexts.
	Set_blockparams(IBlockparamsContext)

	// GetVect returns the vect attribute.
	GetVect() interfaces.Instruction

	// SetVect sets the vect attribute.
	SetVect(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	AllID_VALIDO() []antlr.TerminalNode
	ID_VALIDO(i int) antlr.TerminalNode
	DOS_PUNTOS() antlr.TerminalNode
	AllCORCHIZQ() []antlr.TerminalNode
	CORCHIZQ(i int) antlr.TerminalNode
	Tipodato() ITipodatoContext
	AllCORCHDER() []antlr.TerminalNode
	CORCHDER(i int) antlr.TerminalNode
	IG() antlr.TerminalNode
	Blockparams() IBlockparamsContext

	// IsVectorcontrolContext differentiates from other interfaces.
	IsVectorcontrolContext()
}

type VectorcontrolContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	vect         interfaces.Instruction
	_VAR         antlr.Token
	_ID_VALIDO   antlr.Token
	_tipodato    ITipodatoContext
	_blockparams IBlockparamsContext
	prin         antlr.Token
	secu         antlr.Token
}

func NewEmptyVectorcontrolContext() *VectorcontrolContext {
	var p = new(VectorcontrolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectorcontrol
	return p
}

func InitEmptyVectorcontrolContext(p *VectorcontrolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectorcontrol
}

func (*VectorcontrolContext) IsVectorcontrolContext() {}

func NewVectorcontrolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorcontrolContext {
	var p = new(VectorcontrolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vectorcontrol

	return p
}

func (s *VectorcontrolContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorcontrolContext) Get_VAR() antlr.Token { return s._VAR }

func (s *VectorcontrolContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *VectorcontrolContext) GetPrin() antlr.Token { return s.prin }

func (s *VectorcontrolContext) GetSecu() antlr.Token { return s.secu }

func (s *VectorcontrolContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *VectorcontrolContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *VectorcontrolContext) SetPrin(v antlr.Token) { s.prin = v }

func (s *VectorcontrolContext) SetSecu(v antlr.Token) { s.secu = v }

func (s *VectorcontrolContext) Get_tipodato() ITipodatoContext { return s._tipodato }

func (s *VectorcontrolContext) Get_blockparams() IBlockparamsContext { return s._blockparams }

func (s *VectorcontrolContext) Set_tipodato(v ITipodatoContext) { s._tipodato = v }

func (s *VectorcontrolContext) Set_blockparams(v IBlockparamsContext) { s._blockparams = v }

func (s *VectorcontrolContext) GetVect() interfaces.Instruction { return s.vect }

func (s *VectorcontrolContext) SetVect(v interfaces.Instruction) { s.vect = v }

func (s *VectorcontrolContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *VectorcontrolContext) AllID_VALIDO() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID_VALIDO)
}

func (s *VectorcontrolContext) ID_VALIDO(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, i)
}

func (s *VectorcontrolContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOS_PUNTOS, 0)
}

func (s *VectorcontrolContext) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHIZQ)
}

func (s *VectorcontrolContext) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, i)
}

func (s *VectorcontrolContext) Tipodato() ITipodatoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipodatoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipodatoContext)
}

func (s *VectorcontrolContext) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHDER)
}

func (s *VectorcontrolContext) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, i)
}

func (s *VectorcontrolContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *VectorcontrolContext) Blockparams() IBlockparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockparamsContext)
}

func (s *VectorcontrolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorcontrolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorcontrolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVectorcontrol(s)
	}
}

func (s *VectorcontrolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVectorcontrol(s)
	}
}

func (p *SwiftGrammarParser) Vectorcontrol() (localctx IVectorcontrolContext) {
	localctx = NewVectorcontrolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_vectorcontrol)
	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VectorcontrolContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(500)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*VectorcontrolContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(501)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(502)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(503)

			var _x = p.Tipodato()

			localctx.(*VectorcontrolContext)._tipodato = _x
		}
		{
			p.SetState(504)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(506)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(507)

			var _x = p.Blockparams()

			localctx.(*VectorcontrolContext)._blockparams = _x
		}
		{
			p.SetState(508)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VectorcontrolContext).vect = datoscompuestos.NewArregloDeclaracionLista((func() int {
			if localctx.(*VectorcontrolContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VectorcontrolContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VectorcontrolContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VectorcontrolContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VectorcontrolContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*VectorcontrolContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*VectorcontrolContext).Get_tipodato().GetTipo(), localctx.(*VectorcontrolContext).Get_blockparams().GetBlkpar())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(511)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VectorcontrolContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*VectorcontrolContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(513)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(514)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(515)

			var _x = p.Tipodato()

			localctx.(*VectorcontrolContext)._tipodato = _x
		}
		{
			p.SetState(516)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(517)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(518)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(519)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VectorcontrolContext).vect = datoscompuestos.NewArregloDeclaracionSinLista((func() int {
			if localctx.(*VectorcontrolContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VectorcontrolContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VectorcontrolContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VectorcontrolContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VectorcontrolContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*VectorcontrolContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*VectorcontrolContext).Get_tipodato().GetTipo())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(522)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VectorcontrolContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(523)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*VectorcontrolContext).prin = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(524)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(525)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(526)

			var _x = p.Tipodato()

			localctx.(*VectorcontrolContext)._tipodato = _x
		}
		{
			p.SetState(527)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(528)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(529)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*VectorcontrolContext).secu = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VectorcontrolContext).vect = datoscompuestos.NewArregloDeclaracionId((func() int {
			if localctx.(*VectorcontrolContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VectorcontrolContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VectorcontrolContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VectorcontrolContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VectorcontrolContext).GetPrin() == nil {
				return ""
			} else {
				return localctx.(*VectorcontrolContext).GetPrin().GetText()
			}
		}()), localctx.(*VectorcontrolContext).Get_tipodato().GetTipo(), (func() string {
			if localctx.(*VectorcontrolContext).GetSecu() == nil {
				return ""
			} else {
				return localctx.(*VectorcontrolContext).GetSecu().GetText()
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

// IBlockparamsContext is an interface to support dynamic dispatch.
type IBlockparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_bloqueparams returns the _bloqueparams rule contexts.
	Get_bloqueparams() IBloqueparamsContext

	// Set_bloqueparams sets the _bloqueparams rule contexts.
	Set_bloqueparams(IBloqueparamsContext)

	// GetBlopas returns the blopas rule context list.
	GetBlopas() []IBloqueparamsContext

	// SetBlopas sets the blopas rule context list.
	SetBlopas([]IBloqueparamsContext)

	// GetBlkpar returns the blkpar attribute.
	GetBlkpar() []interface{}

	// SetBlkpar sets the blkpar attribute.
	SetBlkpar([]interface{})

	// Getter signatures
	AllBloqueparams() []IBloqueparamsContext
	Bloqueparams(i int) IBloqueparamsContext

	// IsBlockparamsContext differentiates from other interfaces.
	IsBlockparamsContext()
}

type BlockparamsContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	blkpar        []interface{}
	_bloqueparams IBloqueparamsContext
	blopas        []IBloqueparamsContext
}

func NewEmptyBlockparamsContext() *BlockparamsContext {
	var p = new(BlockparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_blockparams
	return p
}

func InitEmptyBlockparamsContext(p *BlockparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_blockparams
}

func (*BlockparamsContext) IsBlockparamsContext() {}

func NewBlockparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockparamsContext {
	var p = new(BlockparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_blockparams

	return p
}

func (s *BlockparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockparamsContext) Get_bloqueparams() IBloqueparamsContext { return s._bloqueparams }

func (s *BlockparamsContext) Set_bloqueparams(v IBloqueparamsContext) { s._bloqueparams = v }

func (s *BlockparamsContext) GetBlopas() []IBloqueparamsContext { return s.blopas }

func (s *BlockparamsContext) SetBlopas(v []IBloqueparamsContext) { s.blopas = v }

func (s *BlockparamsContext) GetBlkpar() []interface{} { return s.blkpar }

func (s *BlockparamsContext) SetBlkpar(v []interface{}) { s.blkpar = v }

func (s *BlockparamsContext) AllBloqueparams() []IBloqueparamsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBloqueparamsContext); ok {
			len++
		}
	}

	tst := make([]IBloqueparamsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBloqueparamsContext); ok {
			tst[i] = t.(IBloqueparamsContext)
			i++
		}
	}

	return tst
}

func (s *BlockparamsContext) Bloqueparams(i int) IBloqueparamsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueparamsContext); ok {
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

	return t.(IBloqueparamsContext)
}

func (s *BlockparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlockparams(s)
	}
}

func (s *BlockparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlockparams(s)
	}
}

func (p *SwiftGrammarParser) Blockparams() (localctx IBlockparamsContext) {
	localctx = NewBlockparamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SwiftGrammarParserRULE_blockparams)

	localctx.(*BlockparamsContext).blkpar = []interface{}{}
	var listInt []IBloqueparamsContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-6)) & ^0x3f) == 0 && ((int64(1)<<(_la-6))&2341890530142584851) != 0) {
		{
			p.SetState(534)

			var _x = p.Bloqueparams()

			localctx.(*BlockparamsContext)._bloqueparams = _x
		}
		localctx.(*BlockparamsContext).blopas = append(localctx.(*BlockparamsContext).blopas, localctx.(*BlockparamsContext)._bloqueparams)

		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt = localctx.(*BlockparamsContext).GetBlopas()
	for _, e := range listInt {
		localctx.(*BlockparamsContext).blkpar = append(localctx.(*BlockparamsContext).blkpar, e.GetBlopas())
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

// IBloqueparamsContext is an interface to support dynamic dispatch.
type IBloqueparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_COMA returns the _COMA token.
	Get_COMA() antlr.Token

	// Set_COMA sets the _COMA token.
	Set_COMA(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetBlopas returns the blopas attribute.
	GetBlopas() interfaces.Expression

	// SetBlopas sets the blopas attribute.
	SetBlopas(interfaces.Expression)

	// Getter signatures
	COMA() antlr.TerminalNode
	Expr() IExprContext

	// IsBloqueparamsContext differentiates from other interfaces.
	IsBloqueparamsContext()
}

type BloqueparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	blopas interfaces.Expression
	_COMA  antlr.Token
	_expr  IExprContext
}

func NewEmptyBloqueparamsContext() *BloqueparamsContext {
	var p = new(BloqueparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_bloqueparams
	return p
}

func InitEmptyBloqueparamsContext(p *BloqueparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_bloqueparams
}

func (*BloqueparamsContext) IsBloqueparamsContext() {}

func NewBloqueparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueparamsContext {
	var p = new(BloqueparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_bloqueparams

	return p
}

func (s *BloqueparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueparamsContext) Get_COMA() antlr.Token { return s._COMA }

func (s *BloqueparamsContext) Set_COMA(v antlr.Token) { s._COMA = v }

func (s *BloqueparamsContext) Get_expr() IExprContext { return s._expr }

func (s *BloqueparamsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *BloqueparamsContext) GetBlopas() interfaces.Expression { return s.blopas }

func (s *BloqueparamsContext) SetBlopas(v interfaces.Expression) { s.blopas = v }

func (s *BloqueparamsContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *BloqueparamsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BloqueparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBloqueparams(s)
	}
}

func (s *BloqueparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBloqueparams(s)
	}
}

func (p *SwiftGrammarParser) Bloqueparams() (localctx IBloqueparamsContext) {
	localctx = NewBloqueparamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftGrammarParserRULE_bloqueparams)
	p.SetState(548)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserCOMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(541)

			var _m = p.Match(SwiftGrammarParserCOMA)

			localctx.(*BloqueparamsContext)._COMA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(542)

			var _x = p.expr(0)

			localctx.(*BloqueparamsContext)._expr = _x
		}

		localctx.(*BloqueparamsContext).blopas = datoscompuestos.NewArregloParametros((func() int {
			if localctx.(*BloqueparamsContext).Get_COMA() == nil {
				return 0
			} else {
				return localctx.(*BloqueparamsContext).Get_COMA().GetLine()
			}
		}()), (func() int {
			if localctx.(*BloqueparamsContext).Get_COMA() == nil {
				return 0
			} else {
				return localctx.(*BloqueparamsContext).Get_COMA().GetColumn()
			}
		}()), localctx.(*BloqueparamsContext).Get_expr().GetE())

	case SwiftGrammarParserTRU, SwiftGrammarParserFAL, SwiftGrammarParserNULO, SwiftGrammarParserNUMBER, SwiftGrammarParserCADENA, SwiftGrammarParserID_VALIDO, SwiftGrammarParserCHARACTER, SwiftGrammarParserPARIZQ, SwiftGrammarParserNOT, SwiftGrammarParserSUB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(545)

			var _x = p.expr(0)

			localctx.(*BloqueparamsContext)._expr = _x
		}

		localctx.(*BloqueparamsContext).blopas = datoscompuestos.NewArregloParametro(localctx.(*BloqueparamsContext).Get_expr().GetE())

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

// IVectoragregarContext is an interface to support dynamic dispatch.
type IVectoragregarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// GetPrin returns the prin token.
	GetPrin() antlr.Token

	// GetSecu returns the secu token.
	GetSecu() antlr.Token

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// SetPrin sets the prin token.
	SetPrin(antlr.Token)

	// SetSecu sets the secu token.
	SetSecu(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetPop returns the pop rule contexts.
	GetPop() IExprContext

	// GetSop returns the sop rule contexts.
	GetSop() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetPop sets the pop rule contexts.
	SetPop(IExprContext)

	// SetSop sets the sop rule contexts.
	SetSop(IExprContext)

	// GetVeadct returns the veadct attribute.
	GetVeadct() interfaces.Instruction

	// SetVeadct sets the veadct attribute.
	SetVeadct(interfaces.Instruction)

	// Getter signatures
	AllID_VALIDO() []antlr.TerminalNode
	ID_VALIDO(i int) antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	APPEND() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PARDER() antlr.TerminalNode
	AllCORCHIZQ() []antlr.TerminalNode
	CORCHIZQ(i int) antlr.TerminalNode
	AllCORCHDER() []antlr.TerminalNode
	CORCHDER(i int) antlr.TerminalNode
	IG() antlr.TerminalNode

	// IsVectoragregarContext differentiates from other interfaces.
	IsVectoragregarContext()
}

type VectoragregarContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	veadct     interfaces.Instruction
	_ID_VALIDO antlr.Token
	_expr      IExprContext
	prin       antlr.Token
	pop        IExprContext
	secu       antlr.Token
	sop        IExprContext
}

func NewEmptyVectoragregarContext() *VectoragregarContext {
	var p = new(VectoragregarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectoragregar
	return p
}

func InitEmptyVectoragregarContext(p *VectoragregarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectoragregar
}

func (*VectoragregarContext) IsVectoragregarContext() {}

func NewVectoragregarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectoragregarContext {
	var p = new(VectoragregarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vectoragregar

	return p
}

func (s *VectoragregarContext) GetParser() antlr.Parser { return s.parser }

func (s *VectoragregarContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *VectoragregarContext) GetPrin() antlr.Token { return s.prin }

func (s *VectoragregarContext) GetSecu() antlr.Token { return s.secu }

func (s *VectoragregarContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *VectoragregarContext) SetPrin(v antlr.Token) { s.prin = v }

func (s *VectoragregarContext) SetSecu(v antlr.Token) { s.secu = v }

func (s *VectoragregarContext) Get_expr() IExprContext { return s._expr }

func (s *VectoragregarContext) GetPop() IExprContext { return s.pop }

func (s *VectoragregarContext) GetSop() IExprContext { return s.sop }

func (s *VectoragregarContext) Set_expr(v IExprContext) { s._expr = v }

func (s *VectoragregarContext) SetPop(v IExprContext) { s.pop = v }

func (s *VectoragregarContext) SetSop(v IExprContext) { s.sop = v }

func (s *VectoragregarContext) GetVeadct() interfaces.Instruction { return s.veadct }

func (s *VectoragregarContext) SetVeadct(v interfaces.Instruction) { s.veadct = v }

func (s *VectoragregarContext) AllID_VALIDO() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID_VALIDO)
}

func (s *VectoragregarContext) ID_VALIDO(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, i)
}

func (s *VectoragregarContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *VectoragregarContext) APPEND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAPPEND, 0)
}

func (s *VectoragregarContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *VectoragregarContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *VectoragregarContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *VectoragregarContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *VectoragregarContext) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHIZQ)
}

func (s *VectoragregarContext) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, i)
}

func (s *VectoragregarContext) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHDER)
}

func (s *VectoragregarContext) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, i)
}

func (s *VectoragregarContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *VectoragregarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectoragregarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectoragregarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVectoragregar(s)
	}
}

func (s *VectoragregarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVectoragregar(s)
	}
}

func (p *SwiftGrammarParser) Vectoragregar() (localctx IVectoragregarContext) {
	localctx = NewVectoragregarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SwiftGrammarParserRULE_vectoragregar)
	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(550)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*VectoragregarContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(552)
			p.Match(SwiftGrammarParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(553)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(554)

			var _x = p.expr(0)

			localctx.(*VectoragregarContext)._expr = _x
		}
		{
			p.SetState(555)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VectoragregarContext).veadct = datoscompuestos.NewArregloAppend((func() string {
			if localctx.(*VectoragregarContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*VectoragregarContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*VectoragregarContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(558)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*VectoragregarContext).prin = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(559)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(560)

			var _x = p.expr(0)

			localctx.(*VectoragregarContext).pop = _x
		}
		{
			p.SetState(561)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(562)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(563)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*VectoragregarContext).secu = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(564)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(565)

			var _x = p.expr(0)

			localctx.(*VectoragregarContext).sop = _x
		}
		{
			p.SetState(566)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VectoragregarContext).veadct = datoscompuestos.NewArregloAppendArreglo((func() string {
			if localctx.(*VectoragregarContext).GetPrin() == nil {
				return ""
			} else {
				return localctx.(*VectoragregarContext).GetPrin().GetText()
			}
		}()), localctx.(*VectoragregarContext).GetPop().GetE(), (func() string {
			if localctx.(*VectoragregarContext).GetSecu() == nil {
				return ""
			} else {
				return localctx.(*VectoragregarContext).GetSecu().GetText()
			}
		}()), localctx.(*VectoragregarContext).GetSop().GetE())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(569)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*VectoragregarContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(570)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(571)

			var _x = p.expr(0)

			localctx.(*VectoragregarContext).pop = _x
		}
		{
			p.SetState(572)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(573)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(574)

			var _x = p.expr(0)

			localctx.(*VectoragregarContext).sop = _x
		}
		localctx.(*VectoragregarContext).veadct = datoscompuestos.NewArregloAppendExp((func() string {
			if localctx.(*VectoragregarContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*VectoragregarContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*VectoragregarContext).GetPop().GetE(), localctx.(*VectoragregarContext).GetSop().GetE())

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

// IVectorremoverContext is an interface to support dynamic dispatch.
type IVectorremoverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Get_PUNTO returns the _PUNTO token.
	Get_PUNTO() antlr.Token

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// Set_PUNTO sets the _PUNTO token.
	Set_PUNTO(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetVermct returns the vermct attribute.
	GetVermct() interfaces.Instruction

	// SetVermct sets the vermct attribute.
	SetVermct(interfaces.Instruction)

	// Getter signatures
	ID_VALIDO() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	REMOVELAST() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	REMOVE() antlr.TerminalNode
	AT() antlr.TerminalNode
	DOS_PUNTOS() antlr.TerminalNode
	Expr() IExprContext

	// IsVectorremoverContext differentiates from other interfaces.
	IsVectorremoverContext()
}

type VectorremoverContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	vermct     interfaces.Instruction
	_ID_VALIDO antlr.Token
	_PUNTO     antlr.Token
	_expr      IExprContext
}

func NewEmptyVectorremoverContext() *VectorremoverContext {
	var p = new(VectorremoverContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectorremover
	return p
}

func InitEmptyVectorremoverContext(p *VectorremoverContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectorremover
}

func (*VectorremoverContext) IsVectorremoverContext() {}

func NewVectorremoverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorremoverContext {
	var p = new(VectorremoverContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vectorremover

	return p
}

func (s *VectorremoverContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorremoverContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *VectorremoverContext) Get_PUNTO() antlr.Token { return s._PUNTO }

func (s *VectorremoverContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *VectorremoverContext) Set_PUNTO(v antlr.Token) { s._PUNTO = v }

func (s *VectorremoverContext) Get_expr() IExprContext { return s._expr }

func (s *VectorremoverContext) Set_expr(v IExprContext) { s._expr = v }

func (s *VectorremoverContext) GetVermct() interfaces.Instruction { return s.vermct }

func (s *VectorremoverContext) SetVermct(v interfaces.Instruction) { s.vermct = v }

func (s *VectorremoverContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, 0)
}

func (s *VectorremoverContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *VectorremoverContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVELAST, 0)
}

func (s *VectorremoverContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *VectorremoverContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *VectorremoverContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVE, 0)
}

func (s *VectorremoverContext) AT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAT, 0)
}

func (s *VectorremoverContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOS_PUNTOS, 0)
}

func (s *VectorremoverContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VectorremoverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorremoverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorremoverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVectorremover(s)
	}
}

func (s *VectorremoverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVectorremover(s)
	}
}

func (p *SwiftGrammarParser) Vectorremover() (localctx IVectorremoverContext) {
	localctx = NewVectorremoverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftGrammarParserRULE_vectorremover)
	p.SetState(595)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(579)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*VectorremoverContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(580)

			var _m = p.Match(SwiftGrammarParserPUNTO)

			localctx.(*VectorremoverContext)._PUNTO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(581)
			p.Match(SwiftGrammarParserREMOVELAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(582)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(583)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VectorremoverContext).vermct = datoscompuestos.NewArregloRemoveLast((func() int {
			if localctx.(*VectorremoverContext).Get_PUNTO() == nil {
				return 0
			} else {
				return localctx.(*VectorremoverContext).Get_PUNTO().GetLine()
			}
		}()), (func() int {
			if localctx.(*VectorremoverContext).Get_PUNTO() == nil {
				return 0
			} else {
				return localctx.(*VectorremoverContext).Get_PUNTO().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VectorremoverContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*VectorremoverContext).Get_ID_VALIDO().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(585)

			var _m = p.Match(SwiftGrammarParserID_VALIDO)

			localctx.(*VectorremoverContext)._ID_VALIDO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(586)

			var _m = p.Match(SwiftGrammarParserPUNTO)

			localctx.(*VectorremoverContext)._PUNTO = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(587)
			p.Match(SwiftGrammarParserREMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(588)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(589)
			p.Match(SwiftGrammarParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(590)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(591)

			var _x = p.expr(0)

			localctx.(*VectorremoverContext)._expr = _x
		}
		{
			p.SetState(592)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VectorremoverContext).vermct = datoscompuestos.NewArregloRemovePos((func() int {
			if localctx.(*VectorremoverContext).Get_PUNTO() == nil {
				return 0
			} else {
				return localctx.(*VectorremoverContext).Get_PUNTO().GetLine()
			}
		}()), (func() int {
			if localctx.(*VectorremoverContext).Get_PUNTO() == nil {
				return 0
			} else {
				return localctx.(*VectorremoverContext).Get_PUNTO().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VectorremoverContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*VectorremoverContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*VectorremoverContext).Get_expr().GetE())

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

// IVectorvacioContext is an interface to support dynamic dispatch.
type IVectorvacioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Get_PUNTO returns the _PUNTO token.
	Get_PUNTO() antlr.Token

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// Set_PUNTO sets the _PUNTO token.
	Set_PUNTO(antlr.Token)

	// GetVeemct returns the veemct attribute.
	GetVeemct() interfaces.Expression

	// SetVeemct sets the veemct attribute.
	SetVeemct(interfaces.Expression)

	// Getter signatures
	ID_VALIDO() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	ISEMPTY() antlr.TerminalNode

	// IsVectorvacioContext differentiates from other interfaces.
	IsVectorvacioContext()
}

type VectorvacioContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	veemct     interfaces.Expression
	_ID_VALIDO antlr.Token
	_PUNTO     antlr.Token
}

func NewEmptyVectorvacioContext() *VectorvacioContext {
	var p = new(VectorvacioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectorvacio
	return p
}

func InitEmptyVectorvacioContext(p *VectorvacioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectorvacio
}

func (*VectorvacioContext) IsVectorvacioContext() {}

func NewVectorvacioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorvacioContext {
	var p = new(VectorvacioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vectorvacio

	return p
}

func (s *VectorvacioContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorvacioContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *VectorvacioContext) Get_PUNTO() antlr.Token { return s._PUNTO }

func (s *VectorvacioContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *VectorvacioContext) Set_PUNTO(v antlr.Token) { s._PUNTO = v }

func (s *VectorvacioContext) GetVeemct() interfaces.Expression { return s.veemct }

func (s *VectorvacioContext) SetVeemct(v interfaces.Expression) { s.veemct = v }

func (s *VectorvacioContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, 0)
}

func (s *VectorvacioContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *VectorvacioContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserISEMPTY, 0)
}

func (s *VectorvacioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorvacioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorvacioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVectorvacio(s)
	}
}

func (s *VectorvacioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVectorvacio(s)
	}
}

func (p *SwiftGrammarParser) Vectorvacio() (localctx IVectorvacioContext) {
	localctx = NewVectorvacioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SwiftGrammarParserRULE_vectorvacio)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(597)

		var _m = p.Match(SwiftGrammarParserID_VALIDO)

		localctx.(*VectorvacioContext)._ID_VALIDO = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(598)

		var _m = p.Match(SwiftGrammarParserPUNTO)

		localctx.(*VectorvacioContext)._PUNTO = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(599)
		p.Match(SwiftGrammarParserISEMPTY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*VectorvacioContext).veemct = datoscompuestos.NewArregloIsEmpty((func() int {
		if localctx.(*VectorvacioContext).Get_PUNTO() == nil {
			return 0
		} else {
			return localctx.(*VectorvacioContext).Get_PUNTO().GetLine()
		}
	}()), (func() int {
		if localctx.(*VectorvacioContext).Get_PUNTO() == nil {
			return 0
		} else {
			return localctx.(*VectorvacioContext).Get_PUNTO().GetColumn()
		}
	}()), (func() string {
		if localctx.(*VectorvacioContext).Get_ID_VALIDO() == nil {
			return ""
		} else {
			return localctx.(*VectorvacioContext).Get_ID_VALIDO().GetText()
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

// IVectorcountContext is an interface to support dynamic dispatch.
type IVectorcountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Get_PUNTO returns the _PUNTO token.
	Get_PUNTO() antlr.Token

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// Set_PUNTO sets the _PUNTO token.
	Set_PUNTO(antlr.Token)

	// GetVecnct returns the vecnct attribute.
	GetVecnct() interfaces.Expression

	// SetVecnct sets the vecnct attribute.
	SetVecnct(interfaces.Expression)

	// Getter signatures
	ID_VALIDO() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	COUNT() antlr.TerminalNode

	// IsVectorcountContext differentiates from other interfaces.
	IsVectorcountContext()
}

type VectorcountContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	vecnct     interfaces.Expression
	_ID_VALIDO antlr.Token
	_PUNTO     antlr.Token
}

func NewEmptyVectorcountContext() *VectorcountContext {
	var p = new(VectorcountContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectorcount
	return p
}

func InitEmptyVectorcountContext(p *VectorcountContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectorcount
}

func (*VectorcountContext) IsVectorcountContext() {}

func NewVectorcountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorcountContext {
	var p = new(VectorcountContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vectorcount

	return p
}

func (s *VectorcountContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorcountContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *VectorcountContext) Get_PUNTO() antlr.Token { return s._PUNTO }

func (s *VectorcountContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *VectorcountContext) Set_PUNTO(v antlr.Token) { s._PUNTO = v }

func (s *VectorcountContext) GetVecnct() interfaces.Expression { return s.vecnct }

func (s *VectorcountContext) SetVecnct(v interfaces.Expression) { s.vecnct = v }

func (s *VectorcountContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, 0)
}

func (s *VectorcountContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *VectorcountContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOUNT, 0)
}

func (s *VectorcountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorcountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorcountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVectorcount(s)
	}
}

func (s *VectorcountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVectorcount(s)
	}
}

func (p *SwiftGrammarParser) Vectorcount() (localctx IVectorcountContext) {
	localctx = NewVectorcountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SwiftGrammarParserRULE_vectorcount)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(602)

		var _m = p.Match(SwiftGrammarParserID_VALIDO)

		localctx.(*VectorcountContext)._ID_VALIDO = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(603)

		var _m = p.Match(SwiftGrammarParserPUNTO)

		localctx.(*VectorcountContext)._PUNTO = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(604)
		p.Match(SwiftGrammarParserCOUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*VectorcountContext).vecnct = datoscompuestos.NewArregloCount((func() int {
		if localctx.(*VectorcountContext).Get_PUNTO() == nil {
			return 0
		} else {
			return localctx.(*VectorcountContext).Get_PUNTO().GetLine()
		}
	}()), (func() int {
		if localctx.(*VectorcountContext).Get_PUNTO() == nil {
			return 0
		} else {
			return localctx.(*VectorcountContext).Get_PUNTO().GetColumn()
		}
	}()), (func() string {
		if localctx.(*VectorcountContext).Get_ID_VALIDO() == nil {
			return ""
		} else {
			return localctx.(*VectorcountContext).Get_ID_VALIDO().GetText()
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

// IVectoraccessContext is an interface to support dynamic dispatch.
type IVectoraccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Get_CORCHDER returns the _CORCHDER token.
	Get_CORCHDER() antlr.Token

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// Set_CORCHDER sets the _CORCHDER token.
	Set_CORCHDER(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetVepposct returns the vepposct attribute.
	GetVepposct() interfaces.Expression

	// SetVepposct sets the vepposct attribute.
	SetVepposct(interfaces.Expression)

	// Getter signatures
	ID_VALIDO() antlr.TerminalNode
	CORCHIZQ() antlr.TerminalNode
	Expr() IExprContext
	CORCHDER() antlr.TerminalNode

	// IsVectoraccessContext differentiates from other interfaces.
	IsVectoraccessContext()
}

type VectoraccessContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	vepposct   interfaces.Expression
	_ID_VALIDO antlr.Token
	_expr      IExprContext
	_CORCHDER  antlr.Token
}

func NewEmptyVectoraccessContext() *VectoraccessContext {
	var p = new(VectoraccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectoraccess
	return p
}

func InitEmptyVectoraccessContext(p *VectoraccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectoraccess
}

func (*VectoraccessContext) IsVectoraccessContext() {}

func NewVectoraccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectoraccessContext {
	var p = new(VectoraccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vectoraccess

	return p
}

func (s *VectoraccessContext) GetParser() antlr.Parser { return s.parser }

func (s *VectoraccessContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *VectoraccessContext) Get_CORCHDER() antlr.Token { return s._CORCHDER }

func (s *VectoraccessContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *VectoraccessContext) Set_CORCHDER(v antlr.Token) { s._CORCHDER = v }

func (s *VectoraccessContext) Get_expr() IExprContext { return s._expr }

func (s *VectoraccessContext) Set_expr(v IExprContext) { s._expr = v }

func (s *VectoraccessContext) GetVepposct() interfaces.Expression { return s.vepposct }

func (s *VectoraccessContext) SetVepposct(v interfaces.Expression) { s.vepposct = v }

func (s *VectoraccessContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, 0)
}

func (s *VectoraccessContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *VectoraccessContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VectoraccessContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *VectoraccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectoraccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectoraccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVectoraccess(s)
	}
}

func (s *VectoraccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVectoraccess(s)
	}
}

func (p *SwiftGrammarParser) Vectoraccess() (localctx IVectoraccessContext) {
	localctx = NewVectoraccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SwiftGrammarParserRULE_vectoraccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(607)

		var _m = p.Match(SwiftGrammarParserID_VALIDO)

		localctx.(*VectoraccessContext)._ID_VALIDO = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(608)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(609)

		var _x = p.expr(0)

		localctx.(*VectoraccessContext)._expr = _x
	}
	{
		p.SetState(610)

		var _m = p.Match(SwiftGrammarParserCORCHDER)

		localctx.(*VectoraccessContext)._CORCHDER = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*VectoraccessContext).vepposct = datoscompuestos.NewArregloAccess((func() int {
		if localctx.(*VectoraccessContext).Get_CORCHDER() == nil {
			return 0
		} else {
			return localctx.(*VectoraccessContext).Get_CORCHDER().GetLine()
		}
	}()), (func() int {
		if localctx.(*VectoraccessContext).Get_CORCHDER() == nil {
			return 0
		} else {
			return localctx.(*VectoraccessContext).Get_CORCHDER().GetColumn()
		}
	}()), (func() string {
		if localctx.(*VectoraccessContext).Get_ID_VALIDO() == nil {
			return ""
		} else {
			return localctx.(*VectoraccessContext).Get_ID_VALIDO().GetText()
		}
	}()), localctx.(*VectoraccessContext).Get_expr().GetE())

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

// IMatrizcontrolContext is an interface to support dynamic dispatch.
type IMatrizcontrolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID_VALIDO returns the _ID_VALIDO token.
	Get_ID_VALIDO() antlr.Token

	// Get_DOS_PUNTOS returns the _DOS_PUNTOS token.
	Get_DOS_PUNTOS() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID_VALIDO sets the _ID_VALIDO token.
	Set_ID_VALIDO(antlr.Token)

	// Set_DOS_PUNTOS sets the _DOS_PUNTOS token.
	Set_DOS_PUNTOS(antlr.Token)

	// Get_tipomatriz returns the _tipomatriz rule contexts.
	Get_tipomatriz() ITipomatrizContext

	// Get_defmatriz returns the _defmatriz rule contexts.
	Get_defmatriz() IDefmatrizContext

	// Set_tipomatriz sets the _tipomatriz rule contexts.
	Set_tipomatriz(ITipomatrizContext)

	// Set_defmatriz sets the _defmatriz rule contexts.
	Set_defmatriz(IDefmatrizContext)

	// GetMatct returns the matct attribute.
	GetMatct() interfaces.Instruction

	// SetMatct sets the matct attribute.
	SetMatct(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID_VALIDO() antlr.TerminalNode
	IG() antlr.TerminalNode
	Defmatriz() IDefmatrizContext
	DOS_PUNTOS() antlr.TerminalNode
	Tipomatriz() ITipomatrizContext

	// IsMatrizcontrolContext differentiates from other interfaces.
	IsMatrizcontrolContext()
}

type MatrizcontrolContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	matct       interfaces.Instruction
	_VAR        antlr.Token
	_ID_VALIDO  antlr.Token
	_DOS_PUNTOS antlr.Token
	_tipomatriz ITipomatrizContext
	_defmatriz  IDefmatrizContext
}

func NewEmptyMatrizcontrolContext() *MatrizcontrolContext {
	var p = new(MatrizcontrolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_matrizcontrol
	return p
}

func InitEmptyMatrizcontrolContext(p *MatrizcontrolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_matrizcontrol
}

func (*MatrizcontrolContext) IsMatrizcontrolContext() {}

func NewMatrizcontrolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrizcontrolContext {
	var p = new(MatrizcontrolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_matrizcontrol

	return p
}

func (s *MatrizcontrolContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrizcontrolContext) Get_VAR() antlr.Token { return s._VAR }

func (s *MatrizcontrolContext) Get_ID_VALIDO() antlr.Token { return s._ID_VALIDO }

func (s *MatrizcontrolContext) Get_DOS_PUNTOS() antlr.Token { return s._DOS_PUNTOS }

func (s *MatrizcontrolContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *MatrizcontrolContext) Set_ID_VALIDO(v antlr.Token) { s._ID_VALIDO = v }

func (s *MatrizcontrolContext) Set_DOS_PUNTOS(v antlr.Token) { s._DOS_PUNTOS = v }

func (s *MatrizcontrolContext) Get_tipomatriz() ITipomatrizContext { return s._tipomatriz }

func (s *MatrizcontrolContext) Get_defmatriz() IDefmatrizContext { return s._defmatriz }

func (s *MatrizcontrolContext) Set_tipomatriz(v ITipomatrizContext) { s._tipomatriz = v }

func (s *MatrizcontrolContext) Set_defmatriz(v IDefmatrizContext) { s._defmatriz = v }

func (s *MatrizcontrolContext) GetMatct() interfaces.Instruction { return s.matct }

func (s *MatrizcontrolContext) SetMatct(v interfaces.Instruction) { s.matct = v }

func (s *MatrizcontrolContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *MatrizcontrolContext) ID_VALIDO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID_VALIDO, 0)
}

func (s *MatrizcontrolContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *MatrizcontrolContext) Defmatriz() IDefmatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefmatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefmatrizContext)
}

func (s *MatrizcontrolContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOS_PUNTOS, 0)
}

func (s *MatrizcontrolContext) Tipomatriz() ITipomatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipomatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipomatrizContext)
}

func (s *MatrizcontrolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrizcontrolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrizcontrolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterMatrizcontrol(s)
	}
}

func (s *MatrizcontrolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitMatrizcontrol(s)
	}
}

func (p *SwiftGrammarParser) Matrizcontrol() (localctx IMatrizcontrolContext) {
	localctx = NewMatrizcontrolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SwiftGrammarParserRULE_matrizcontrol)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(613)

		var _m = p.Match(SwiftGrammarParserVAR)

		localctx.(*MatrizcontrolContext)._VAR = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(614)

		var _m = p.Match(SwiftGrammarParserID_VALIDO)

		localctx.(*MatrizcontrolContext)._ID_VALIDO = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(617)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserDOS_PUNTOS {
		{
			p.SetState(615)

			var _m = p.Match(SwiftGrammarParserDOS_PUNTOS)

			localctx.(*MatrizcontrolContext)._DOS_PUNTOS = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(616)

			var _x = p.Tipomatriz()

			localctx.(*MatrizcontrolContext)._tipomatriz = _x
		}

	}
	{
		p.SetState(619)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(620)

		var _x = p.Defmatriz()

		localctx.(*MatrizcontrolContext)._defmatriz = _x
	}

	if localctx.(*MatrizcontrolContext).Get_DOS_PUNTOS() != nil {
		localctx.(*MatrizcontrolContext).matct = datoscompuestos.NewMatrizDeclaracion((func() int {
			if localctx.(*MatrizcontrolContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*MatrizcontrolContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*MatrizcontrolContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*MatrizcontrolContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*MatrizcontrolContext).Get_ID_VALIDO() == nil {
				return ""
			} else {
				return localctx.(*MatrizcontrolContext).Get_ID_VALIDO().GetText()
			}
		}()), localctx.(*MatrizcontrolContext).Get_tipomatriz().GetTipomat(), localctx.(*MatrizcontrolContext).Get_defmatriz().GetDefmat())
	} else {
		fmt.Println("Nada")
		//localctx.(*MatrizcontrolContext).matct = datoscompuestos.NewMatrizDeclaracionSinTipo((func() int { if localctx.(*MatrizcontrolContext).Get_VAR() == nil { return 0 } else { return localctx.(*MatrizcontrolContext).Get_VAR().GetLine() }}()), (func() int { if localctx.(*MatrizcontrolContext).Get_VAR() == nil { return 0 } else { return localctx.(*MatrizcontrolContext).Get_VAR().GetColumn() }}()), (func() string { if localctx.(*MatrizcontrolContext).Get_ID_VALIDO() == nil { return "" } else { return localctx.(*MatrizcontrolContext).Get_ID_VALIDO().GetText() }}()) , localctx.(*MatrizcontrolContext).Get_defmatriz().GetDefmat())
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

// ITipomatrizContext is an interface to support dynamic dispatch.
type ITipomatrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CORCHIZQ returns the _CORCHIZQ token.
	Get_CORCHIZQ() antlr.Token

	// Set_CORCHIZQ sets the _CORCHIZQ token.
	Set_CORCHIZQ(antlr.Token)

	// Get_tipomatriz returns the _tipomatriz rule contexts.
	Get_tipomatriz() ITipomatrizContext

	// Get_tipodato returns the _tipodato rule contexts.
	Get_tipodato() ITipodatoContext

	// Set_tipomatriz sets the _tipomatriz rule contexts.
	Set_tipomatriz(ITipomatrizContext)

	// Set_tipodato sets the _tipodato rule contexts.
	Set_tipodato(ITipodatoContext)

	// GetTipomat returns the tipomat attribute.
	GetTipomat() interfaces.Expression

	// SetTipomat sets the tipomat attribute.
	SetTipomat(interfaces.Expression)

	// Getter signatures
	CORCHIZQ() antlr.TerminalNode
	Tipomatriz() ITipomatrizContext
	CORCHDER() antlr.TerminalNode
	Tipodato() ITipodatoContext

	// IsTipomatrizContext differentiates from other interfaces.
	IsTipomatrizContext()
}

type TipomatrizContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	tipomat     interfaces.Expression
	_CORCHIZQ   antlr.Token
	_tipomatriz ITipomatrizContext
	_tipodato   ITipodatoContext
}

func NewEmptyTipomatrizContext() *TipomatrizContext {
	var p = new(TipomatrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_tipomatriz
	return p
}

func InitEmptyTipomatrizContext(p *TipomatrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_tipomatriz
}

func (*TipomatrizContext) IsTipomatrizContext() {}

func NewTipomatrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipomatrizContext {
	var p = new(TipomatrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_tipomatriz

	return p
}

func (s *TipomatrizContext) GetParser() antlr.Parser { return s.parser }

func (s *TipomatrizContext) Get_CORCHIZQ() antlr.Token { return s._CORCHIZQ }

func (s *TipomatrizContext) Set_CORCHIZQ(v antlr.Token) { s._CORCHIZQ = v }

func (s *TipomatrizContext) Get_tipomatriz() ITipomatrizContext { return s._tipomatriz }

func (s *TipomatrizContext) Get_tipodato() ITipodatoContext { return s._tipodato }

func (s *TipomatrizContext) Set_tipomatriz(v ITipomatrizContext) { s._tipomatriz = v }

func (s *TipomatrizContext) Set_tipodato(v ITipodatoContext) { s._tipodato = v }

func (s *TipomatrizContext) GetTipomat() interfaces.Expression { return s.tipomat }

func (s *TipomatrizContext) SetTipomat(v interfaces.Expression) { s.tipomat = v }

func (s *TipomatrizContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *TipomatrizContext) Tipomatriz() ITipomatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipomatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipomatrizContext)
}

func (s *TipomatrizContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *TipomatrizContext) Tipodato() ITipodatoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipodatoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipodatoContext)
}

func (s *TipomatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipomatrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipomatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTipomatriz(s)
	}
}

func (s *TipomatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTipomatriz(s)
	}
}

func (p *SwiftGrammarParser) Tipomatriz() (localctx ITipomatrizContext) {
	localctx = NewTipomatrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SwiftGrammarParserRULE_tipomatriz)
	p.SetState(633)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(623)

			var _m = p.Match(SwiftGrammarParserCORCHIZQ)

			localctx.(*TipomatrizContext)._CORCHIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(624)

			var _x = p.Tipomatriz()

			localctx.(*TipomatrizContext)._tipomatriz = _x
		}
		{
			p.SetState(625)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*TipomatrizContext).tipomat = datoscompuestos.NewMatrizDimension((func() int {
			if localctx.(*TipomatrizContext).Get_CORCHIZQ() == nil {
				return 0
			} else {
				return localctx.(*TipomatrizContext).Get_CORCHIZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*TipomatrizContext).Get_CORCHIZQ() == nil {
				return 0
			} else {
				return localctx.(*TipomatrizContext).Get_CORCHIZQ().GetColumn()
			}
		}()), localctx.(*TipomatrizContext).Get_tipomatriz().GetTipomat())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(628)

			var _m = p.Match(SwiftGrammarParserCORCHIZQ)

			localctx.(*TipomatrizContext)._CORCHIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(629)

			var _x = p.Tipodato()

			localctx.(*TipomatrizContext)._tipodato = _x
		}
		{
			p.SetState(630)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*TipomatrizContext).tipomat = datoscompuestos.NewMatrizTipo((func() int {
			if localctx.(*TipomatrizContext).Get_CORCHIZQ() == nil {
				return 0
			} else {
				return localctx.(*TipomatrizContext).Get_CORCHIZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*TipomatrizContext).Get_CORCHIZQ() == nil {
				return 0
			} else {
				return localctx.(*TipomatrizContext).Get_CORCHIZQ().GetColumn()
			}
		}()), localctx.(*TipomatrizContext).Get_tipodato().GetTipo())

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

// IDefmatrizContext is an interface to support dynamic dispatch.
type IDefmatrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listavaloresmat returns the _listavaloresmat rule contexts.
	Get_listavaloresmat() IListavaloresmatContext

	// Set_listavaloresmat sets the _listavaloresmat rule contexts.
	Set_listavaloresmat(IListavaloresmatContext)

	// GetDefmat returns the defmat attribute.
	GetDefmat() interfaces.Instruction

	// SetDefmat sets the defmat attribute.
	SetDefmat(interfaces.Instruction)

	// Getter signatures
	Listavaloresmat() IListavaloresmatContext

	// IsDefmatrizContext differentiates from other interfaces.
	IsDefmatrizContext()
}

type DefmatrizContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	defmat           interfaces.Instruction
	_listavaloresmat IListavaloresmatContext
}

func NewEmptyDefmatrizContext() *DefmatrizContext {
	var p = new(DefmatrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defmatriz
	return p
}

func InitEmptyDefmatrizContext(p *DefmatrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defmatriz
}

func (*DefmatrizContext) IsDefmatrizContext() {}

func NewDefmatrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefmatrizContext {
	var p = new(DefmatrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_defmatriz

	return p
}

func (s *DefmatrizContext) GetParser() antlr.Parser { return s.parser }

func (s *DefmatrizContext) Get_listavaloresmat() IListavaloresmatContext { return s._listavaloresmat }

func (s *DefmatrizContext) Set_listavaloresmat(v IListavaloresmatContext) { s._listavaloresmat = v }

func (s *DefmatrizContext) GetDefmat() interfaces.Instruction { return s.defmat }

func (s *DefmatrizContext) SetDefmat(v interfaces.Instruction) { s.defmat = v }

func (s *DefmatrizContext) Listavaloresmat() IListavaloresmatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavaloresmatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListavaloresmatContext)
}

func (s *DefmatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefmatrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefmatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDefmatriz(s)
	}
}

func (s *DefmatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDefmatriz(s)
	}
}

func (p *SwiftGrammarParser) Defmatriz() (localctx IDefmatrizContext) {
	localctx = NewDefmatrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SwiftGrammarParserRULE_defmatriz)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(635)

		var _x = p.Listavaloresmat()

		localctx.(*DefmatrizContext)._listavaloresmat = _x
	}
	localctx.(*DefmatrizContext).defmat = localctx.(*DefmatrizContext).Get_listavaloresmat().GetListvlamat()

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

// IListavaloresmatContext is an interface to support dynamic dispatch.
type IListavaloresmatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listavaloresmat2 returns the _listavaloresmat2 rule contexts.
	Get_listavaloresmat2() IListavaloresmat2Context

	// Get_simplematriz returns the _simplematriz rule contexts.
	Get_simplematriz() ISimplematrizContext

	// Set_listavaloresmat2 sets the _listavaloresmat2 rule contexts.
	Set_listavaloresmat2(IListavaloresmat2Context)

	// Set_simplematriz sets the _simplematriz rule contexts.
	Set_simplematriz(ISimplematrizContext)

	// GetListvlamat returns the listvlamat attribute.
	GetListvlamat() interfaces.Instruction

	// SetListvlamat sets the listvlamat attribute.
	SetListvlamat(interfaces.Instruction)

	// Getter signatures
	CORCHIZQ() antlr.TerminalNode
	Listavaloresmat2() IListavaloresmat2Context
	CORCHDER() antlr.TerminalNode
	Simplematriz() ISimplematrizContext

	// IsListavaloresmatContext differentiates from other interfaces.
	IsListavaloresmatContext()
}

type ListavaloresmatContext struct {
	antlr.BaseParserRuleContext
	parser            antlr.Parser
	listvlamat        interfaces.Instruction
	_listavaloresmat2 IListavaloresmat2Context
	_simplematriz     ISimplematrizContext
}

func NewEmptyListavaloresmatContext() *ListavaloresmatContext {
	var p = new(ListavaloresmatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmat
	return p
}

func InitEmptyListavaloresmatContext(p *ListavaloresmatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmat
}

func (*ListavaloresmatContext) IsListavaloresmatContext() {}

func NewListavaloresmatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListavaloresmatContext {
	var p = new(ListavaloresmatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmat

	return p
}

func (s *ListavaloresmatContext) GetParser() antlr.Parser { return s.parser }

func (s *ListavaloresmatContext) Get_listavaloresmat2() IListavaloresmat2Context {
	return s._listavaloresmat2
}

func (s *ListavaloresmatContext) Get_simplematriz() ISimplematrizContext { return s._simplematriz }

func (s *ListavaloresmatContext) Set_listavaloresmat2(v IListavaloresmat2Context) {
	s._listavaloresmat2 = v
}

func (s *ListavaloresmatContext) Set_simplematriz(v ISimplematrizContext) { s._simplematriz = v }

func (s *ListavaloresmatContext) GetListvlamat() interfaces.Instruction { return s.listvlamat }

func (s *ListavaloresmatContext) SetListvlamat(v interfaces.Instruction) { s.listvlamat = v }

func (s *ListavaloresmatContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *ListavaloresmatContext) Listavaloresmat2() IListavaloresmat2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavaloresmat2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListavaloresmat2Context)
}

func (s *ListavaloresmatContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *ListavaloresmatContext) Simplematriz() ISimplematrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimplematrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimplematrizContext)
}

func (s *ListavaloresmatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListavaloresmatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListavaloresmatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListavaloresmat(s)
	}
}

func (s *ListavaloresmatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListavaloresmat(s)
	}
}

func (p *SwiftGrammarParser) Listavaloresmat() (localctx IListavaloresmatContext) {
	localctx = NewListavaloresmatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SwiftGrammarParserRULE_listavaloresmat)
	p.SetState(646)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(638)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(639)

			var _x = p.listavaloresmat2(0)

			localctx.(*ListavaloresmatContext)._listavaloresmat2 = _x
		}
		{
			p.SetState(640)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ListavaloresmatContext).listvlamat = localctx.(*ListavaloresmatContext).Get_listavaloresmat2().GetMylisttmatt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(643)

			var _x = p.Simplematriz()

			localctx.(*ListavaloresmatContext)._simplematriz = _x
		}
		localctx.(*ListavaloresmatContext).listvlamat = localctx.(*ListavaloresmatContext).Get_simplematriz().GetSimmat()

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

// IListavaloresmat2Context is an interface to support dynamic dispatch.
type IListavaloresmat2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op rule contexts.
	GetOp() IListavaloresmat2Context

	// Get_listavaloresmat returns the _listavaloresmat rule contexts.
	Get_listavaloresmat() IListavaloresmatContext

	// Get_listaexpresions returns the _listaexpresions rule contexts.
	Get_listaexpresions() IListaexpresionsContext

	// SetOp sets the op rule contexts.
	SetOp(IListavaloresmat2Context)

	// Set_listavaloresmat sets the _listavaloresmat rule contexts.
	Set_listavaloresmat(IListavaloresmatContext)

	// Set_listaexpresions sets the _listaexpresions rule contexts.
	Set_listaexpresions(IListaexpresionsContext)

	// GetMylisttmatt returns the mylisttmatt attribute.
	GetMylisttmatt() interfaces.Instruction

	// SetMylisttmatt sets the mylisttmatt attribute.
	SetMylisttmatt(interfaces.Instruction)

	// Getter signatures
	Listavaloresmat() IListavaloresmatContext
	Listaexpresions() IListaexpresionsContext
	COMA() antlr.TerminalNode
	Listavaloresmat2() IListavaloresmat2Context

	// IsListavaloresmat2Context differentiates from other interfaces.
	IsListavaloresmat2Context()
}

type Listavaloresmat2Context struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	mylisttmatt      interfaces.Instruction
	op               IListavaloresmat2Context
	_listavaloresmat IListavaloresmatContext
	_listaexpresions IListaexpresionsContext
}

func NewEmptyListavaloresmat2Context() *Listavaloresmat2Context {
	var p = new(Listavaloresmat2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmat2
	return p
}

func InitEmptyListavaloresmat2Context(p *Listavaloresmat2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmat2
}

func (*Listavaloresmat2Context) IsListavaloresmat2Context() {}

func NewListavaloresmat2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Listavaloresmat2Context {
	var p = new(Listavaloresmat2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmat2

	return p
}

func (s *Listavaloresmat2Context) GetParser() antlr.Parser { return s.parser }

func (s *Listavaloresmat2Context) GetOp() IListavaloresmat2Context { return s.op }

func (s *Listavaloresmat2Context) Get_listavaloresmat() IListavaloresmatContext {
	return s._listavaloresmat
}

func (s *Listavaloresmat2Context) Get_listaexpresions() IListaexpresionsContext {
	return s._listaexpresions
}

func (s *Listavaloresmat2Context) SetOp(v IListavaloresmat2Context) { s.op = v }

func (s *Listavaloresmat2Context) Set_listavaloresmat(v IListavaloresmatContext) {
	s._listavaloresmat = v
}

func (s *Listavaloresmat2Context) Set_listaexpresions(v IListaexpresionsContext) {
	s._listaexpresions = v
}

func (s *Listavaloresmat2Context) GetMylisttmatt() interfaces.Instruction { return s.mylisttmatt }

func (s *Listavaloresmat2Context) SetMylisttmatt(v interfaces.Instruction) { s.mylisttmatt = v }

func (s *Listavaloresmat2Context) Listavaloresmat() IListavaloresmatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavaloresmatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListavaloresmatContext)
}

func (s *Listavaloresmat2Context) Listaexpresions() IListaexpresionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaexpresionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaexpresionsContext)
}

func (s *Listavaloresmat2Context) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *Listavaloresmat2Context) Listavaloresmat2() IListavaloresmat2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavaloresmat2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListavaloresmat2Context)
}

func (s *Listavaloresmat2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Listavaloresmat2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Listavaloresmat2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListavaloresmat2(s)
	}
}

func (s *Listavaloresmat2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListavaloresmat2(s)
	}
}

func (p *SwiftGrammarParser) Listavaloresmat2() (localctx IListavaloresmat2Context) {
	return p.listavaloresmat2(0)
}

func (p *SwiftGrammarParser) listavaloresmat2(_p int) (localctx IListavaloresmat2Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListavaloresmat2Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListavaloresmat2Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 64
	p.EnterRecursionRule(localctx, 64, SwiftGrammarParserRULE_listavaloresmat2, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(655)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserCORCHIZQ:
		{
			p.SetState(649)

			var _x = p.Listavaloresmat()

			localctx.(*Listavaloresmat2Context)._listavaloresmat = _x
		}
		localctx.(*Listavaloresmat2Context).mylisttmatt = datoscompuestos.NewMatrizListaNivel(localctx.(*Listavaloresmat2Context).Get_listavaloresmat().GetListvlamat())

	case SwiftGrammarParserTRU, SwiftGrammarParserFAL, SwiftGrammarParserNULO, SwiftGrammarParserNUMBER, SwiftGrammarParserCADENA, SwiftGrammarParserID_VALIDO, SwiftGrammarParserCHARACTER, SwiftGrammarParserPARIZQ, SwiftGrammarParserNOT, SwiftGrammarParserSUB, SwiftGrammarParserCOMA:
		{
			p.SetState(652)

			var _x = p.Listaexpresions()

			localctx.(*Listavaloresmat2Context)._listaexpresions = _x
		}
		localctx.(*Listavaloresmat2Context).mylisttmatt = datoscompuestos.NewMatrizListaExpresion(localctx.(*Listavaloresmat2Context).Get_listaexpresions().GetBlkparf())

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(664)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListavaloresmat2Context(p, _parentctx, _parentState)
			localctx.(*Listavaloresmat2Context).op = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listavaloresmat2)
			p.SetState(657)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(658)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(659)

				var _x = p.Listavaloresmat()

				localctx.(*Listavaloresmat2Context)._listavaloresmat = _x
			}
			localctx.(*Listavaloresmat2Context).mylisttmatt = datoscompuestos.NewMatrizListaExpresionList(localctx.(*Listavaloresmat2Context).GetOp().GetMylisttmatt(), localctx.(*Listavaloresmat2Context).Get_listavaloresmat().GetListvlamat())

		}
		p.SetState(666)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
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

// IListaexpresionsContext is an interface to support dynamic dispatch.
type IListaexpresionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listaexpresion returns the _listaexpresion rule contexts.
	Get_listaexpresion() IListaexpresionContext

	// Set_listaexpresion sets the _listaexpresion rule contexts.
	Set_listaexpresion(IListaexpresionContext)

	// GetFunpar returns the funpar rule context list.
	GetFunpar() []IListaexpresionContext

	// SetFunpar sets the funpar rule context list.
	SetFunpar([]IListaexpresionContext)

	// GetBlkparf returns the blkparf attribute.
	GetBlkparf() []interface{}

	// SetBlkparf sets the blkparf attribute.
	SetBlkparf([]interface{})

	// Getter signatures
	AllListaexpresion() []IListaexpresionContext
	Listaexpresion(i int) IListaexpresionContext

	// IsListaexpresionsContext differentiates from other interfaces.
	IsListaexpresionsContext()
}

type ListaexpresionsContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	blkparf         []interface{}
	_listaexpresion IListaexpresionContext
	funpar          []IListaexpresionContext
}

func NewEmptyListaexpresionsContext() *ListaexpresionsContext {
	var p = new(ListaexpresionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listaexpresions
	return p
}

func InitEmptyListaexpresionsContext(p *ListaexpresionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listaexpresions
}

func (*ListaexpresionsContext) IsListaexpresionsContext() {}

func NewListaexpresionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaexpresionsContext {
	var p = new(ListaexpresionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listaexpresions

	return p
}

func (s *ListaexpresionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaexpresionsContext) Get_listaexpresion() IListaexpresionContext {
	return s._listaexpresion
}

func (s *ListaexpresionsContext) Set_listaexpresion(v IListaexpresionContext) { s._listaexpresion = v }

func (s *ListaexpresionsContext) GetFunpar() []IListaexpresionContext { return s.funpar }

func (s *ListaexpresionsContext) SetFunpar(v []IListaexpresionContext) { s.funpar = v }

func (s *ListaexpresionsContext) GetBlkparf() []interface{} { return s.blkparf }

func (s *ListaexpresionsContext) SetBlkparf(v []interface{}) { s.blkparf = v }

func (s *ListaexpresionsContext) AllListaexpresion() []IListaexpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListaexpresionContext); ok {
			len++
		}
	}

	tst := make([]IListaexpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListaexpresionContext); ok {
			tst[i] = t.(IListaexpresionContext)
			i++
		}
	}

	return tst
}

func (s *ListaexpresionsContext) Listaexpresion(i int) IListaexpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaexpresionContext); ok {
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

	return t.(IListaexpresionContext)
}

func (s *ListaexpresionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaexpresionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaexpresionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListaexpresions(s)
	}
}

func (s *ListaexpresionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListaexpresions(s)
	}
}

func (p *SwiftGrammarParser) Listaexpresions() (localctx IListaexpresionsContext) {
	localctx = NewListaexpresionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SwiftGrammarParserRULE_listaexpresions)

	localctx.(*ListaexpresionsContext).blkparf = []interface{}{}
	var listInt []IListaexpresionContext

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(668)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(667)

				var _x = p.Listaexpresion()

				localctx.(*ListaexpresionsContext)._listaexpresion = _x
			}
			localctx.(*ListaexpresionsContext).funpar = append(localctx.(*ListaexpresionsContext).funpar, localctx.(*ListaexpresionsContext)._listaexpresion)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(670)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	listInt = localctx.(*ListaexpresionsContext).GetFunpar()
	for _, e := range listInt {
		localctx.(*ListaexpresionsContext).blkparf = append(localctx.(*ListaexpresionsContext).blkparf, e.GetFunpar())
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

// IListaexpresionContext is an interface to support dynamic dispatch.
type IListaexpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_COMA returns the _COMA token.
	Get_COMA() antlr.Token

	// Set_COMA sets the _COMA token.
	Set_COMA(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetFunpar returns the funpar attribute.
	GetFunpar() interfaces.Expression

	// SetFunpar sets the funpar attribute.
	SetFunpar(interfaces.Expression)

	// Getter signatures
	COMA() antlr.TerminalNode
	Expr() IExprContext

	// IsListaexpresionContext differentiates from other interfaces.
	IsListaexpresionContext()
}

type ListaexpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	funpar interfaces.Expression
	_COMA  antlr.Token
	_expr  IExprContext
}

func NewEmptyListaexpresionContext() *ListaexpresionContext {
	var p = new(ListaexpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listaexpresion
	return p
}

func InitEmptyListaexpresionContext(p *ListaexpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listaexpresion
}

func (*ListaexpresionContext) IsListaexpresionContext() {}

func NewListaexpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaexpresionContext {
	var p = new(ListaexpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listaexpresion

	return p
}

func (s *ListaexpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaexpresionContext) Get_COMA() antlr.Token { return s._COMA }

func (s *ListaexpresionContext) Set_COMA(v antlr.Token) { s._COMA = v }

func (s *ListaexpresionContext) Get_expr() IExprContext { return s._expr }

func (s *ListaexpresionContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListaexpresionContext) GetFunpar() interfaces.Expression { return s.funpar }

func (s *ListaexpresionContext) SetFunpar(v interfaces.Expression) { s.funpar = v }

func (s *ListaexpresionContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListaexpresionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListaexpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaexpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaexpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListaexpresion(s)
	}
}

func (s *ListaexpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListaexpresion(s)
	}
}

func (p *SwiftGrammarParser) Listaexpresion() (localctx IListaexpresionContext) {
	localctx = NewListaexpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SwiftGrammarParserRULE_listaexpresion)
	p.SetState(681)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserCOMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(674)

			var _m = p.Match(SwiftGrammarParserCOMA)

			localctx.(*ListaexpresionContext)._COMA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(675)

			var _x = p.expr(0)

			localctx.(*ListaexpresionContext)._expr = _x
		}

		localctx.(*ListaexpresionContext).funpar = datoscompuestos.NewArregloParametros((func() int {
			if localctx.(*ListaexpresionContext).Get_COMA() == nil {
				return 0
			} else {
				return localctx.(*ListaexpresionContext).Get_COMA().GetLine()
			}
		}()), (func() int {
			if localctx.(*ListaexpresionContext).Get_COMA() == nil {
				return 0
			} else {
				return localctx.(*ListaexpresionContext).Get_COMA().GetColumn()
			}
		}()), localctx.(*ListaexpresionContext).Get_expr().GetE())

	case SwiftGrammarParserTRU, SwiftGrammarParserFAL, SwiftGrammarParserNULO, SwiftGrammarParserNUMBER, SwiftGrammarParserCADENA, SwiftGrammarParserID_VALIDO, SwiftGrammarParserCHARACTER, SwiftGrammarParserPARIZQ, SwiftGrammarParserNOT, SwiftGrammarParserSUB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(678)

			var _x = p.expr(0)

			localctx.(*ListaexpresionContext)._expr = _x
		}

		localctx.(*ListaexpresionContext).funpar = datoscompuestos.NewArregloParametro(localctx.(*ListaexpresionContext).Get_expr().GetE())

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

// ISimplematrizContext is an interface to support dynamic dispatch.
type ISimplematrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Get_tipomatriz returns the _tipomatriz rule contexts.
	Get_tipomatriz() ITipomatrizContext

	// GetOp returns the op rule contexts.
	GetOp() ISimplematrizContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_tipomatriz sets the _tipomatriz rule contexts.
	Set_tipomatriz(ITipomatrizContext)

	// SetOp sets the op rule contexts.
	SetOp(ISimplematrizContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetSimmat returns the simmat attribute.
	GetSimmat() interfaces.Instruction

	// SetSimmat sets the simmat attribute.
	SetSimmat(interfaces.Instruction)

	// Getter signatures
	Tipomatriz() ITipomatrizContext
	PARIZQ() antlr.TerminalNode
	REPEATING() antlr.TerminalNode
	AllDOS_PUNTOS() []antlr.TerminalNode
	DOS_PUNTOS(i int) antlr.TerminalNode
	COMA() antlr.TerminalNode
	COUNT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	Simplematriz() ISimplematrizContext
	Expr() IExprContext

	// IsSimplematrizContext differentiates from other interfaces.
	IsSimplematrizContext()
}

type SimplematrizContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	simmat      interfaces.Instruction
	_tipomatriz ITipomatrizContext
	op          ISimplematrizContext
	_NUMBER     antlr.Token
	_expr       IExprContext
}

func NewEmptySimplematrizContext() *SimplematrizContext {
	var p = new(SimplematrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_simplematriz
	return p
}

func InitEmptySimplematrizContext(p *SimplematrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_simplematriz
}

func (*SimplematrizContext) IsSimplematrizContext() {}

func NewSimplematrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimplematrizContext {
	var p = new(SimplematrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_simplematriz

	return p
}

func (s *SimplematrizContext) GetParser() antlr.Parser { return s.parser }

func (s *SimplematrizContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *SimplematrizContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *SimplematrizContext) Get_tipomatriz() ITipomatrizContext { return s._tipomatriz }

func (s *SimplematrizContext) GetOp() ISimplematrizContext { return s.op }

func (s *SimplematrizContext) Get_expr() IExprContext { return s._expr }

func (s *SimplematrizContext) Set_tipomatriz(v ITipomatrizContext) { s._tipomatriz = v }

func (s *SimplematrizContext) SetOp(v ISimplematrizContext) { s.op = v }

func (s *SimplematrizContext) Set_expr(v IExprContext) { s._expr = v }

func (s *SimplematrizContext) GetSimmat() interfaces.Instruction { return s.simmat }

func (s *SimplematrizContext) SetSimmat(v interfaces.Instruction) { s.simmat = v }

func (s *SimplematrizContext) Tipomatriz() ITipomatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipomatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipomatrizContext)
}

func (s *SimplematrizContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *SimplematrizContext) REPEATING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREPEATING, 0)
}

func (s *SimplematrizContext) AllDOS_PUNTOS() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserDOS_PUNTOS)
}

func (s *SimplematrizContext) DOS_PUNTOS(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOS_PUNTOS, i)
}

func (s *SimplematrizContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *SimplematrizContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOUNT, 0)
}

func (s *SimplematrizContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
}

func (s *SimplematrizContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *SimplematrizContext) Simplematriz() ISimplematrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimplematrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimplematrizContext)
}

func (s *SimplematrizContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SimplematrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimplematrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimplematrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterSimplematriz(s)
	}
}

func (s *SimplematrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitSimplematriz(s)
	}
}

func (p *SwiftGrammarParser) Simplematriz() (localctx ISimplematrizContext) {
	localctx = NewSimplematrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SwiftGrammarParserRULE_simplematriz)
	p.SetState(707)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(683)

			var _x = p.Tipomatriz()

			localctx.(*SimplematrizContext)._tipomatriz = _x
		}
		{
			p.SetState(684)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(685)
			p.Match(SwiftGrammarParserREPEATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(686)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(687)

			var _x = p.Simplematriz()

			localctx.(*SimplematrizContext).op = _x
		}
		{
			p.SetState(688)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(689)
			p.Match(SwiftGrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(690)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(691)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*SimplematrizContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(692)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SimplematrizContext).simmat = datoscompuestos.NewMatrizSimpleUno(localctx.(*SimplematrizContext).Get_tipomatriz().GetTipomat(), localctx.(*SimplematrizContext).GetOp().GetSimmat(), (func() string {
			if localctx.(*SimplematrizContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*SimplematrizContext).Get_NUMBER().GetText()
			}
		}()), (func() int {
			if localctx.(*SimplematrizContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*SimplematrizContext).Get_NUMBER().GetLine()
			}
		}()), (func() int {
			if localctx.(*SimplematrizContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*SimplematrizContext).Get_NUMBER().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(695)

			var _x = p.Tipomatriz()

			localctx.(*SimplematrizContext)._tipomatriz = _x
		}
		{
			p.SetState(696)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(697)
			p.Match(SwiftGrammarParserREPEATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(698)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(699)

			var _x = p.expr(0)

			localctx.(*SimplematrizContext)._expr = _x
		}
		{
			p.SetState(700)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(701)
			p.Match(SwiftGrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(702)
			p.Match(SwiftGrammarParserDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(703)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*SimplematrizContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(704)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SimplematrizContext).simmat = datoscompuestos.NewMatrizSimpleDos(localctx.(*SimplematrizContext).Get_tipomatriz().GetTipomat(), localctx.(*SimplematrizContext).Get_expr().GetE(), (func() string {
			if localctx.(*SimplematrizContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*SimplematrizContext).Get_NUMBER().GetText()
			}
		}()), (func() int {
			if localctx.(*SimplematrizContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*SimplematrizContext).Get_NUMBER().GetLine()
			}
		}()), (func() int {
			if localctx.(*SimplematrizContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*SimplematrizContext).Get_NUMBER().GetColumn()
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

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Get_listaexpresions returns the _listaexpresions rule contexts.
	Get_listaexpresions() IListaexpresionsContext

	// Set_listaexpresions sets the _listaexpresions rule contexts.
	Set_listaexpresions(IListaexpresionsContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Listaexpresions() IListaexpresionsContext
	PARDER() antlr.TerminalNode

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	prnt             interfaces.Instruction
	_PRINT           antlr.Token
	_listaexpresions IListaexpresionsContext
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *PrintstmtContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *PrintstmtContext) Get_listaexpresions() IListaexpresionsContext { return s._listaexpresions }

func (s *PrintstmtContext) Set_listaexpresions(v IListaexpresionsContext) { s._listaexpresions = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) Listaexpresions() IListaexpresionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaexpresionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaexpresionsContext)
}

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (p *SwiftGrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(709)

		var _m = p.Match(SwiftGrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(710)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(711)

		var _x = p.Listaexpresions()

		localctx.(*PrintstmtContext)._listaexpresions = _x
	}
	{
		p.SetState(712)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*PrintstmtContext).prnt = funciones.NewPrint((func() int {
		if localctx.(*PrintstmtContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
		}
	}()), (func() int {
		if localctx.(*PrintstmtContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
		}
	}()), localctx.(*PrintstmtContext).Get_listaexpresions().GetBlkparf())

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

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 32:
		var t *Listavaloresmat2Context = nil
		if localctx != nil {
			t = localctx.(*Listavaloresmat2Context)
		}
		return p.Listavaloresmat2_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 13)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Listavaloresmat2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
