// Code generated from Aim8.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Aim8

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 91, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 3, 2, 3, 2, 3, 2, 5, 2, 30, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 36,
	10, 3, 12, 3, 14, 3, 39, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 7, 4, 48, 10, 4, 12, 4, 14, 4, 51, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 5, 5, 58, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 64, 10, 6, 12, 6, 14,
	6, 67, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 77,
	10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 2, 5, 4, 2, 12, 13, 16, 16, 4, 2, 11, 11, 14, 14, 4, 2, 9, 9, 15, 15,
	2, 86, 2, 29, 3, 2, 2, 2, 4, 31, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8, 57,
	3, 2, 2, 2, 10, 59, 3, 2, 2, 2, 12, 70, 3, 2, 2, 2, 14, 76, 3, 2, 2, 2,
	16, 78, 3, 2, 2, 2, 18, 82, 3, 2, 2, 2, 20, 84, 3, 2, 2, 2, 22, 86, 3,
	2, 2, 2, 24, 88, 3, 2, 2, 2, 26, 30, 5, 22, 12, 2, 27, 30, 5, 24, 13, 2,
	28, 30, 5, 4, 3, 2, 29, 26, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 28, 3,
	2, 2, 2, 30, 3, 3, 2, 2, 2, 31, 32, 7, 3, 2, 2, 32, 37, 5, 2, 2, 2, 33,
	34, 7, 7, 2, 2, 34, 36, 5, 2, 2, 2, 35, 33, 3, 2, 2, 2, 36, 39, 3, 2, 2,
	2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 40, 3, 2, 2, 2, 39, 37,
	3, 2, 2, 2, 40, 41, 7, 4, 2, 2, 41, 5, 3, 2, 2, 2, 42, 43, 7, 20, 2, 2,
	43, 44, 7, 5, 2, 2, 44, 49, 5, 8, 5, 2, 45, 46, 7, 8, 2, 2, 46, 48, 5,
	8, 5, 2, 47, 45, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49,
	50, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 7, 6, 2,
	2, 53, 7, 3, 2, 2, 2, 54, 58, 5, 2, 2, 2, 55, 58, 5, 6, 4, 2, 56, 58, 5,
	10, 6, 2, 57, 54, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58,
	9, 3, 2, 2, 2, 59, 60, 7, 5, 2, 2, 60, 65, 5, 12, 7, 2, 61, 62, 7, 8, 2,
	2, 62, 64, 5, 12, 7, 2, 63, 61, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2, 65, 63,
	3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2,
	68, 69, 7, 6, 2, 2, 69, 11, 3, 2, 2, 2, 70, 71, 5, 14, 8, 2, 71, 72, 5,
	20, 11, 2, 72, 73, 5, 8, 5, 2, 73, 13, 3, 2, 2, 2, 74, 77, 5, 6, 4, 2,
	75, 77, 5, 16, 9, 2, 76, 74, 3, 2, 2, 2, 76, 75, 3, 2, 2, 2, 77, 15, 3,
	2, 2, 2, 78, 79, 5, 6, 4, 2, 79, 80, 5, 18, 10, 2, 80, 81, 5, 6, 4, 2,
	81, 17, 3, 2, 2, 2, 82, 83, 9, 2, 2, 2, 83, 19, 3, 2, 2, 2, 84, 85, 9,
	3, 2, 2, 85, 21, 3, 2, 2, 2, 86, 87, 9, 4, 2, 2, 87, 23, 3, 2, 2, 2, 88,
	89, 7, 19, 2, 2, 89, 25, 3, 2, 2, 2, 8, 29, 37, 49, 57, 65, 76,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'['", "']'", "','", "';'", "'\u039B'", "'\u03BB'", "'\u27F6'",
	"'\u2227'", "'\u2228'", "'->'", "'NIL'", "'~'", "'1'", "'0'",
}
var symbolicNames = []string{
	"", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "COMMA", "SEMICOLON", "NULL",
	"LAMBDA", "ARROW", "AND", "OR", "ALT_ARROW", "ALT_NULL", "NOT", "TRUE",
	"FALSE", "ATOM", "NAME",
}

var ruleNames = []string{
	"sexpr", "enclosd_sexpr", "mexpr", "expr", "conditional_expr", "conditional_pair",
	"predicate", "proposition", "connective", "arrow", "null", "atom",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Aim8Parser struct {
	*antlr.BaseParser
}

func NewAim8Parser(input antlr.TokenStream) *Aim8Parser {
	this := new(Aim8Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Aim8.g4"

	return this
}

// Aim8Parser tokens.
const (
	Aim8ParserEOF       = antlr.TokenEOF
	Aim8ParserLPAREN    = 1
	Aim8ParserRPAREN    = 2
	Aim8ParserLBRACKET  = 3
	Aim8ParserRBRACKET  = 4
	Aim8ParserCOMMA     = 5
	Aim8ParserSEMICOLON = 6
	Aim8ParserNULL      = 7
	Aim8ParserLAMBDA    = 8
	Aim8ParserARROW     = 9
	Aim8ParserAND       = 10
	Aim8ParserOR        = 11
	Aim8ParserALT_ARROW = 12
	Aim8ParserALT_NULL  = 13
	Aim8ParserNOT       = 14
	Aim8ParserTRUE      = 15
	Aim8ParserFALSE     = 16
	Aim8ParserATOM      = 17
	Aim8ParserNAME      = 18
)

// Aim8Parser rules.
const (
	Aim8ParserRULE_sexpr            = 0
	Aim8ParserRULE_enclosd_sexpr    = 1
	Aim8ParserRULE_mexpr            = 2
	Aim8ParserRULE_expr             = 3
	Aim8ParserRULE_conditional_expr = 4
	Aim8ParserRULE_conditional_pair = 5
	Aim8ParserRULE_predicate        = 6
	Aim8ParserRULE_proposition      = 7
	Aim8ParserRULE_connective       = 8
	Aim8ParserRULE_arrow            = 9
	Aim8ParserRULE_null             = 10
	Aim8ParserRULE_atom             = 11
)

// ISexprContext is an interface to support dynamic dispatch.
type ISexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSexprContext differentiates from other interfaces.
	IsSexprContext()
}

type SexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySexprContext() *SexprContext {
	var p = new(SexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_sexpr
	return p
}

func (*SexprContext) IsSexprContext() {}

func NewSexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SexprContext {
	var p = new(SexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_sexpr

	return p
}

func (s *SexprContext) GetParser() antlr.Parser { return s.parser }

func (s *SexprContext) Null() INullContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullContext)
}

func (s *SexprContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *SexprContext) Enclosd_sexpr() IEnclosd_sexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnclosd_sexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnclosd_sexprContext)
}

func (s *SexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterSexpr(s)
	}
}

func (s *SexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitSexpr(s)
	}
}

func (p *Aim8Parser) Sexpr() (localctx ISexprContext) {
	localctx = NewSexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Aim8ParserRULE_sexpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Aim8ParserNULL, Aim8ParserALT_NULL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Null()
		}

	case Aim8ParserATOM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Atom()
		}

	case Aim8ParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Enclosd_sexpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnclosd_sexprContext is an interface to support dynamic dispatch.
type IEnclosd_sexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnclosd_sexprContext differentiates from other interfaces.
	IsEnclosd_sexprContext()
}

type Enclosd_sexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnclosd_sexprContext() *Enclosd_sexprContext {
	var p = new(Enclosd_sexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_enclosd_sexpr
	return p
}

func (*Enclosd_sexprContext) IsEnclosd_sexprContext() {}

func NewEnclosd_sexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enclosd_sexprContext {
	var p = new(Enclosd_sexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_enclosd_sexpr

	return p
}

func (s *Enclosd_sexprContext) GetParser() antlr.Parser { return s.parser }

func (s *Enclosd_sexprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(Aim8ParserLPAREN, 0)
}

func (s *Enclosd_sexprContext) AllSexpr() []ISexprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISexprContext)(nil)).Elem())
	var tst = make([]ISexprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISexprContext)
		}
	}

	return tst
}

func (s *Enclosd_sexprContext) Sexpr(i int) ISexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISexprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISexprContext)
}

func (s *Enclosd_sexprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(Aim8ParserRPAREN, 0)
}

func (s *Enclosd_sexprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Aim8ParserCOMMA)
}

func (s *Enclosd_sexprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Aim8ParserCOMMA, i)
}

func (s *Enclosd_sexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enclosd_sexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enclosd_sexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterEnclosd_sexpr(s)
	}
}

func (s *Enclosd_sexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitEnclosd_sexpr(s)
	}
}

func (p *Aim8Parser) Enclosd_sexpr() (localctx IEnclosd_sexprContext) {
	localctx = NewEnclosd_sexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Aim8ParserRULE_enclosd_sexpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(Aim8ParserLPAREN)
	}
	{
		p.SetState(30)
		p.Sexpr()
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Aim8ParserCOMMA {
		{
			p.SetState(31)
			p.Match(Aim8ParserCOMMA)
		}
		{
			p.SetState(32)
			p.Sexpr()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(Aim8ParserRPAREN)
	}

	return localctx
}

// IMexprContext is an interface to support dynamic dispatch.
type IMexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMexprContext differentiates from other interfaces.
	IsMexprContext()
}

type MexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMexprContext() *MexprContext {
	var p = new(MexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_mexpr
	return p
}

func (*MexprContext) IsMexprContext() {}

func NewMexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MexprContext {
	var p = new(MexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_mexpr

	return p
}

func (s *MexprContext) GetParser() antlr.Parser { return s.parser }

func (s *MexprContext) NAME() antlr.TerminalNode {
	return s.GetToken(Aim8ParserNAME, 0)
}

func (s *MexprContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(Aim8ParserLBRACKET, 0)
}

func (s *MexprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MexprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MexprContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(Aim8ParserRBRACKET, 0)
}

func (s *MexprContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(Aim8ParserSEMICOLON)
}

func (s *MexprContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(Aim8ParserSEMICOLON, i)
}

func (s *MexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterMexpr(s)
	}
}

func (s *MexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitMexpr(s)
	}
}

func (p *Aim8Parser) Mexpr() (localctx IMexprContext) {
	localctx = NewMexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Aim8ParserRULE_mexpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(Aim8ParserNAME)
	}
	{
		p.SetState(41)
		p.Match(Aim8ParserLBRACKET)
	}
	{
		p.SetState(42)
		p.Expr()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Aim8ParserSEMICOLON {
		{
			p.SetState(43)
			p.Match(Aim8ParserSEMICOLON)
		}
		{
			p.SetState(44)
			p.Expr()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(Aim8ParserRBRACKET)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Sexpr() ISexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISexprContext)
}

func (s *ExprContext) Mexpr() IMexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMexprContext)
}

func (s *ExprContext) Conditional_expr() IConditional_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditional_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditional_exprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *Aim8Parser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Aim8ParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Aim8ParserLPAREN, Aim8ParserNULL, Aim8ParserALT_NULL, Aim8ParserATOM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Sexpr()
		}

	case Aim8ParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Mexpr()
		}

	case Aim8ParserLBRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Conditional_expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditional_exprContext is an interface to support dynamic dispatch.
type IConditional_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditional_exprContext differentiates from other interfaces.
	IsConditional_exprContext()
}

type Conditional_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditional_exprContext() *Conditional_exprContext {
	var p = new(Conditional_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_conditional_expr
	return p
}

func (*Conditional_exprContext) IsConditional_exprContext() {}

func NewConditional_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional_exprContext {
	var p = new(Conditional_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_conditional_expr

	return p
}

func (s *Conditional_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Conditional_exprContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(Aim8ParserLBRACKET, 0)
}

func (s *Conditional_exprContext) AllConditional_pair() []IConditional_pairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditional_pairContext)(nil)).Elem())
	var tst = make([]IConditional_pairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditional_pairContext)
		}
	}

	return tst
}

func (s *Conditional_exprContext) Conditional_pair(i int) IConditional_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditional_pairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditional_pairContext)
}

func (s *Conditional_exprContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(Aim8ParserRBRACKET, 0)
}

func (s *Conditional_exprContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(Aim8ParserSEMICOLON)
}

func (s *Conditional_exprContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(Aim8ParserSEMICOLON, i)
}

func (s *Conditional_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Conditional_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Conditional_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterConditional_expr(s)
	}
}

func (s *Conditional_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitConditional_expr(s)
	}
}

func (p *Aim8Parser) Conditional_expr() (localctx IConditional_exprContext) {
	localctx = NewConditional_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Aim8ParserRULE_conditional_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(Aim8ParserLBRACKET)
	}
	{
		p.SetState(58)
		p.Conditional_pair()
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Aim8ParserSEMICOLON {
		{
			p.SetState(59)
			p.Match(Aim8ParserSEMICOLON)
		}
		{
			p.SetState(60)
			p.Conditional_pair()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(Aim8ParserRBRACKET)
	}

	return localctx
}

// IConditional_pairContext is an interface to support dynamic dispatch.
type IConditional_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditional_pairContext differentiates from other interfaces.
	IsConditional_pairContext()
}

type Conditional_pairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditional_pairContext() *Conditional_pairContext {
	var p = new(Conditional_pairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_conditional_pair
	return p
}

func (*Conditional_pairContext) IsConditional_pairContext() {}

func NewConditional_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional_pairContext {
	var p = new(Conditional_pairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_conditional_pair

	return p
}

func (s *Conditional_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Conditional_pairContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *Conditional_pairContext) Arrow() IArrowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrowContext)
}

func (s *Conditional_pairContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Conditional_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Conditional_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Conditional_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterConditional_pair(s)
	}
}

func (s *Conditional_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitConditional_pair(s)
	}
}

func (p *Aim8Parser) Conditional_pair() (localctx IConditional_pairContext) {
	localctx = NewConditional_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Aim8ParserRULE_conditional_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Predicate()
	}
	{
		p.SetState(69)
		p.Arrow()
	}
	{
		p.SetState(70)
		p.Expr()
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) Mexpr() IMexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMexprContext)
}

func (s *PredicateContext) Proposition() IPropositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropositionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropositionContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *Aim8Parser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Aim8ParserRULE_predicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Mexpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Proposition()
		}

	}

	return localctx
}

// IPropositionContext is an interface to support dynamic dispatch.
type IPropositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropositionContext differentiates from other interfaces.
	IsPropositionContext()
}

type PropositionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropositionContext() *PropositionContext {
	var p = new(PropositionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_proposition
	return p
}

func (*PropositionContext) IsPropositionContext() {}

func NewPropositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropositionContext {
	var p = new(PropositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_proposition

	return p
}

func (s *PropositionContext) GetParser() antlr.Parser { return s.parser }

func (s *PropositionContext) AllMexpr() []IMexprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMexprContext)(nil)).Elem())
	var tst = make([]IMexprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMexprContext)
		}
	}

	return tst
}

func (s *PropositionContext) Mexpr(i int) IMexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMexprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMexprContext)
}

func (s *PropositionContext) Connective() IConnectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectiveContext)
}

func (s *PropositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterProposition(s)
	}
}

func (s *PropositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitProposition(s)
	}
}

func (p *Aim8Parser) Proposition() (localctx IPropositionContext) {
	localctx = NewPropositionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Aim8ParserRULE_proposition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Mexpr()
	}
	{
		p.SetState(77)
		p.Connective()
	}
	{
		p.SetState(78)
		p.Mexpr()
	}

	return localctx
}

// IConnectiveContext is an interface to support dynamic dispatch.
type IConnectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnectiveContext differentiates from other interfaces.
	IsConnectiveContext()
}

type ConnectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnectiveContext() *ConnectiveContext {
	var p = new(ConnectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_connective
	return p
}

func (*ConnectiveContext) IsConnectiveContext() {}

func NewConnectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectiveContext {
	var p = new(ConnectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_connective

	return p
}

func (s *ConnectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectiveContext) AND() antlr.TerminalNode {
	return s.GetToken(Aim8ParserAND, 0)
}

func (s *ConnectiveContext) OR() antlr.TerminalNode {
	return s.GetToken(Aim8ParserOR, 0)
}

func (s *ConnectiveContext) NOT() antlr.TerminalNode {
	return s.GetToken(Aim8ParserNOT, 0)
}

func (s *ConnectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConnectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterConnective(s)
	}
}

func (s *ConnectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitConnective(s)
	}
}

func (p *Aim8Parser) Connective() (localctx IConnectiveContext) {
	localctx = NewConnectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Aim8ParserRULE_connective)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Aim8ParserAND)|(1<<Aim8ParserOR)|(1<<Aim8ParserNOT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArrowContext is an interface to support dynamic dispatch.
type IArrowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrowContext differentiates from other interfaces.
	IsArrowContext()
}

type ArrowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrowContext() *ArrowContext {
	var p = new(ArrowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_arrow
	return p
}

func (*ArrowContext) IsArrowContext() {}

func NewArrowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrowContext {
	var p = new(ArrowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_arrow

	return p
}

func (s *ArrowContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrowContext) ARROW() antlr.TerminalNode {
	return s.GetToken(Aim8ParserARROW, 0)
}

func (s *ArrowContext) ALT_ARROW() antlr.TerminalNode {
	return s.GetToken(Aim8ParserALT_ARROW, 0)
}

func (s *ArrowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterArrow(s)
	}
}

func (s *ArrowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitArrow(s)
	}
}

func (p *Aim8Parser) Arrow() (localctx IArrowContext) {
	localctx = NewArrowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Aim8ParserRULE_arrow)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Aim8ParserARROW || _la == Aim8ParserALT_ARROW) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INullContext is an interface to support dynamic dispatch.
type INullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullContext differentiates from other interfaces.
	IsNullContext()
}

type NullContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullContext() *NullContext {
	var p = new(NullContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_null
	return p
}

func (*NullContext) IsNullContext() {}

func NewNullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullContext {
	var p = new(NullContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_null

	return p
}

func (s *NullContext) GetParser() antlr.Parser { return s.parser }

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(Aim8ParserNULL, 0)
}

func (s *NullContext) ALT_NULL() antlr.TerminalNode {
	return s.GetToken(Aim8ParserALT_NULL, 0)
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitNull(s)
	}
}

func (p *Aim8Parser) Null() (localctx INullContext) {
	localctx = NewNullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Aim8ParserRULE_null)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Aim8ParserNULL || _la == Aim8ParserALT_NULL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Aim8ParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Aim8ParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) ATOM() antlr.TerminalNode {
	return s.GetToken(Aim8ParserATOM, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Aim8Listener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *Aim8Parser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Aim8ParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(Aim8ParserATOM)
	}

	return localctx
}
