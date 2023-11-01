// Code generated from C3DGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parserc3d

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// import "Optimizar/interfacesc3d"
// import "Optimizar/instructionsc3d"

// import "Optimizar/environmentc3d"
// import "Optimizar/expressionsc3d"
// import "strings"

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type C3DGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var C3DGrammarLexerLexerStaticData struct {
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

func c3dgrammarlexerLexerInit() {
	staticData := &C3DGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'0'",
	}
	staticData.RuleNames = []string{
		"T__0",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 1, 5, 6, -1, 2, 0, 7, 0, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 4, 0,
		1, 1, 0, 0, 0, 1, 3, 1, 0, 0, 0, 3, 4, 5, 48, 0, 0, 4, 2, 1, 0, 0, 0, 1,
		0, 0,
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

// C3DGrammarLexerInit initializes any static state used to implement C3DGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewC3DGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func C3DGrammarLexerInit() {
	staticData := &C3DGrammarLexerLexerStaticData
	staticData.once.Do(c3dgrammarlexerLexerInit)
}

// NewC3DGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewC3DGrammarLexer(input antlr.CharStream) *C3DGrammarLexer {
	C3DGrammarLexerInit()
	l := new(C3DGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &C3DGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "C3DGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// C3DGrammarLexerT__0 is the C3DGrammarLexer token.
const C3DGrammarLexerT__0 = 1
