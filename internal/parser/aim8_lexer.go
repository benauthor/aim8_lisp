// Code generated from Aim8.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 84, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18,
	9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 6, 18, 76, 10, 18, 13, 18, 14, 18,
	77, 3, 19, 6, 19, 81, 10, 19, 13, 19, 14, 19, 82, 2, 2, 20, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14,
	27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 3, 2, 4, 3, 2, 67, 92,
	3, 2, 99, 124, 2, 85, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31,
	3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3,
	39, 3, 2, 2, 2, 5, 41, 3, 2, 2, 2, 7, 43, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2,
	11, 47, 3, 2, 2, 2, 13, 49, 3, 2, 2, 2, 15, 51, 3, 2, 2, 2, 17, 53, 3,
	2, 2, 2, 19, 55, 3, 2, 2, 2, 21, 57, 3, 2, 2, 2, 23, 59, 3, 2, 2, 2, 25,
	61, 3, 2, 2, 2, 27, 64, 3, 2, 2, 2, 29, 68, 3, 2, 2, 2, 31, 70, 3, 2, 2,
	2, 33, 72, 3, 2, 2, 2, 35, 75, 3, 2, 2, 2, 37, 80, 3, 2, 2, 2, 39, 40,
	7, 42, 2, 2, 40, 4, 3, 2, 2, 2, 41, 42, 7, 43, 2, 2, 42, 6, 3, 2, 2, 2,
	43, 44, 7, 93, 2, 2, 44, 8, 3, 2, 2, 2, 45, 46, 7, 95, 2, 2, 46, 10, 3,
	2, 2, 2, 47, 48, 7, 46, 2, 2, 48, 12, 3, 2, 2, 2, 49, 50, 7, 61, 2, 2,
	50, 14, 3, 2, 2, 2, 51, 52, 7, 925, 2, 2, 52, 16, 3, 2, 2, 2, 53, 54, 7,
	957, 2, 2, 54, 18, 3, 2, 2, 2, 55, 56, 7, 10232, 2, 2, 56, 20, 3, 2, 2,
	2, 57, 58, 7, 8745, 2, 2, 58, 22, 3, 2, 2, 2, 59, 60, 7, 8746, 2, 2, 60,
	24, 3, 2, 2, 2, 61, 62, 7, 47, 2, 2, 62, 63, 7, 64, 2, 2, 63, 26, 3, 2,
	2, 2, 64, 65, 7, 80, 2, 2, 65, 66, 7, 75, 2, 2, 66, 67, 7, 78, 2, 2, 67,
	28, 3, 2, 2, 2, 68, 69, 7, 128, 2, 2, 69, 30, 3, 2, 2, 2, 70, 71, 7, 51,
	2, 2, 71, 32, 3, 2, 2, 2, 72, 73, 7, 50, 2, 2, 73, 34, 3, 2, 2, 2, 74,
	76, 9, 2, 2, 2, 75, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 75, 3, 2, 2,
	2, 77, 78, 3, 2, 2, 2, 78, 36, 3, 2, 2, 2, 79, 81, 9, 3, 2, 2, 80, 79,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2,
	83, 38, 3, 2, 2, 2, 5, 2, 77, 82, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'['", "']'", "','", "';'", "'\u039B'", "'\u03BB'", "'\u27F6'",
	"'\u2227'", "'\u2228'", "'->'", "'NIL'", "'~'", "'1'", "'0'",
}

var lexerSymbolicNames = []string{
	"", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "COMMA", "SEMICOLON", "NULL",
	"LAMBDA", "ARROW", "AND", "OR", "ALT_ARROW", "ALT_NULL", "NOT", "TRUE",
	"FALSE", "ATOM", "NAME",
}

var lexerRuleNames = []string{
	"LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "COMMA", "SEMICOLON", "NULL",
	"LAMBDA", "ARROW", "AND", "OR", "ALT_ARROW", "ALT_NULL", "NOT", "TRUE",
	"FALSE", "ATOM", "NAME",
}

type Aim8Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewAim8Lexer(input antlr.CharStream) *Aim8Lexer {

	l := new(Aim8Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Aim8.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Aim8Lexer tokens.
const (
	Aim8LexerLPAREN    = 1
	Aim8LexerRPAREN    = 2
	Aim8LexerLBRACKET  = 3
	Aim8LexerRBRACKET  = 4
	Aim8LexerCOMMA     = 5
	Aim8LexerSEMICOLON = 6
	Aim8LexerNULL      = 7
	Aim8LexerLAMBDA    = 8
	Aim8LexerARROW     = 9
	Aim8LexerAND       = 10
	Aim8LexerOR        = 11
	Aim8LexerALT_ARROW = 12
	Aim8LexerALT_NULL  = 13
	Aim8LexerNOT       = 14
	Aim8LexerTRUE      = 15
	Aim8LexerFALSE     = 16
	Aim8LexerATOM      = 17
	Aim8LexerNAME      = 18
)
