# Generated from Aim8.g4 by ANTLR 4.7.1
# encoding: utf-8
from antlr4 import *
from io import StringIO
from typing.io import TextIO
import sys

def serializedATN():
    with StringIO() as buf:
        buf.write("\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\24")
        buf.write("[\4\2\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b")
        buf.write("\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4\f\t\f\4\r\t\r\3\2\3\2")
        buf.write("\3\2\5\2\36\n\2\3\3\3\3\3\3\3\3\7\3$\n\3\f\3\16\3\'\13")
        buf.write("\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\7\4\60\n\4\f\4\16\4\63")
        buf.write("\13\4\3\4\3\4\3\5\3\5\3\5\5\5:\n\5\3\6\3\6\3\6\3\6\7\6")
        buf.write("@\n\6\f\6\16\6C\13\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\5")
        buf.write("\bM\n\b\3\t\3\t\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r")
        buf.write("\3\r\3\r\2\2\16\2\4\6\b\n\f\16\20\22\24\26\30\2\5\4\2")
        buf.write("\f\r\20\20\4\2\13\13\16\16\4\2\t\t\17\17\2V\2\35\3\2\2")
        buf.write("\2\4\37\3\2\2\2\6*\3\2\2\2\b9\3\2\2\2\n;\3\2\2\2\fF\3")
        buf.write("\2\2\2\16L\3\2\2\2\20N\3\2\2\2\22R\3\2\2\2\24T\3\2\2\2")
        buf.write("\26V\3\2\2\2\30X\3\2\2\2\32\36\5\26\f\2\33\36\5\30\r\2")
        buf.write("\34\36\5\4\3\2\35\32\3\2\2\2\35\33\3\2\2\2\35\34\3\2\2")
        buf.write("\2\36\3\3\2\2\2\37 \7\3\2\2 %\5\2\2\2!\"\7\7\2\2\"$\5")
        buf.write("\2\2\2#!\3\2\2\2$\'\3\2\2\2%#\3\2\2\2%&\3\2\2\2&(\3\2")
        buf.write("\2\2\'%\3\2\2\2()\7\4\2\2)\5\3\2\2\2*+\7\24\2\2+,\7\5")
        buf.write("\2\2,\61\5\b\5\2-.\7\b\2\2.\60\5\b\5\2/-\3\2\2\2\60\63")
        buf.write("\3\2\2\2\61/\3\2\2\2\61\62\3\2\2\2\62\64\3\2\2\2\63\61")
        buf.write("\3\2\2\2\64\65\7\6\2\2\65\7\3\2\2\2\66:\5\2\2\2\67:\5")
        buf.write("\6\4\28:\5\n\6\29\66\3\2\2\29\67\3\2\2\298\3\2\2\2:\t")
        buf.write("\3\2\2\2;<\7\5\2\2<A\5\f\7\2=>\7\b\2\2>@\5\f\7\2?=\3\2")
        buf.write("\2\2@C\3\2\2\2A?\3\2\2\2AB\3\2\2\2BD\3\2\2\2CA\3\2\2\2")
        buf.write("DE\7\6\2\2E\13\3\2\2\2FG\5\16\b\2GH\5\24\13\2HI\5\b\5")
        buf.write("\2I\r\3\2\2\2JM\5\6\4\2KM\5\20\t\2LJ\3\2\2\2LK\3\2\2\2")
        buf.write("M\17\3\2\2\2NO\5\6\4\2OP\5\22\n\2PQ\5\6\4\2Q\21\3\2\2")
        buf.write("\2RS\t\2\2\2S\23\3\2\2\2TU\t\3\2\2U\25\3\2\2\2VW\t\4\2")
        buf.write("\2W\27\3\2\2\2XY\7\23\2\2Y\31\3\2\2\2\b\35%\619AL")
        return buf.getvalue()


class Aim8Parser ( Parser ):

    grammarFileName = "Aim8.g4"

    atn = ATNDeserializer().deserialize(serializedATN())

    decisionsToDFA = [ DFA(ds, i) for i, ds in enumerate(atn.decisionToState) ]

    sharedContextCache = PredictionContextCache()

    literalNames = [ "<INVALID>", "'('", "')'", "'['", "']'", "','", "';'", 
                     "'\u039B'", "'\u03BB'", "'\u27F6'", "'\u2227'", "'\u2228'", 
                     "'->'", "'NIL'", "'~'", "'1'", "'0'" ]

    symbolicNames = [ "<INVALID>", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", 
                      "COMMA", "SEMICOLON", "NULL", "LAMBDA", "ARROW", "AND", 
                      "OR", "ALT_ARROW", "ALT_NULL", "NOT", "TRUE", "FALSE", 
                      "ATOM", "NAME" ]

    RULE_sexpr = 0
    RULE_enclosd_sexpr = 1
    RULE_mexpr = 2
    RULE_expr = 3
    RULE_conditional_expr = 4
    RULE_conditional_pair = 5
    RULE_predicate = 6
    RULE_proposition = 7
    RULE_connective = 8
    RULE_arrow = 9
    RULE_null = 10
    RULE_atom = 11

    ruleNames =  [ "sexpr", "enclosd_sexpr", "mexpr", "expr", "conditional_expr", 
                   "conditional_pair", "predicate", "proposition", "connective", 
                   "arrow", "null", "atom" ]

    EOF = Token.EOF
    LPAREN=1
    RPAREN=2
    LBRACKET=3
    RBRACKET=4
    COMMA=5
    SEMICOLON=6
    NULL=7
    LAMBDA=8
    ARROW=9
    AND=10
    OR=11
    ALT_ARROW=12
    ALT_NULL=13
    NOT=14
    TRUE=15
    FALSE=16
    ATOM=17
    NAME=18

    def __init__(self, input:TokenStream, output:TextIO = sys.stdout):
        super().__init__(input, output)
        self.checkVersion("4.7.1")
        self._interp = ParserATNSimulator(self, self.atn, self.decisionsToDFA, self.sharedContextCache)
        self._predicates = None



    class SexprContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def null(self):
            return self.getTypedRuleContext(Aim8Parser.NullContext,0)


        def atom(self):
            return self.getTypedRuleContext(Aim8Parser.AtomContext,0)


        def enclosd_sexpr(self):
            return self.getTypedRuleContext(Aim8Parser.Enclosd_sexprContext,0)


        def getRuleIndex(self):
            return Aim8Parser.RULE_sexpr

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterSexpr" ):
                listener.enterSexpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitSexpr" ):
                listener.exitSexpr(self)




    def sexpr(self):

        localctx = Aim8Parser.SexprContext(self, self._ctx, self.state)
        self.enterRule(localctx, 0, self.RULE_sexpr)
        try:
            self.state = 27
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [Aim8Parser.NULL, Aim8Parser.ALT_NULL]:
                self.enterOuterAlt(localctx, 1)
                self.state = 24
                self.null()
                pass
            elif token in [Aim8Parser.ATOM]:
                self.enterOuterAlt(localctx, 2)
                self.state = 25
                self.atom()
                pass
            elif token in [Aim8Parser.LPAREN]:
                self.enterOuterAlt(localctx, 3)
                self.state = 26
                self.enclosd_sexpr()
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class Enclosd_sexprContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def LPAREN(self):
            return self.getToken(Aim8Parser.LPAREN, 0)

        def sexpr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(Aim8Parser.SexprContext)
            else:
                return self.getTypedRuleContext(Aim8Parser.SexprContext,i)


        def RPAREN(self):
            return self.getToken(Aim8Parser.RPAREN, 0)

        def COMMA(self, i:int=None):
            if i is None:
                return self.getTokens(Aim8Parser.COMMA)
            else:
                return self.getToken(Aim8Parser.COMMA, i)

        def getRuleIndex(self):
            return Aim8Parser.RULE_enclosd_sexpr

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterEnclosd_sexpr" ):
                listener.enterEnclosd_sexpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitEnclosd_sexpr" ):
                listener.exitEnclosd_sexpr(self)




    def enclosd_sexpr(self):

        localctx = Aim8Parser.Enclosd_sexprContext(self, self._ctx, self.state)
        self.enterRule(localctx, 2, self.RULE_enclosd_sexpr)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 29
            self.match(Aim8Parser.LPAREN)
            self.state = 30
            self.sexpr()
            self.state = 35
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while _la==Aim8Parser.COMMA:
                self.state = 31
                self.match(Aim8Parser.COMMA)
                self.state = 32
                self.sexpr()
                self.state = 37
                self._errHandler.sync(self)
                _la = self._input.LA(1)

            self.state = 38
            self.match(Aim8Parser.RPAREN)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class MexprContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def NAME(self):
            return self.getToken(Aim8Parser.NAME, 0)

        def LBRACKET(self):
            return self.getToken(Aim8Parser.LBRACKET, 0)

        def expr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(Aim8Parser.ExprContext)
            else:
                return self.getTypedRuleContext(Aim8Parser.ExprContext,i)


        def RBRACKET(self):
            return self.getToken(Aim8Parser.RBRACKET, 0)

        def SEMICOLON(self, i:int=None):
            if i is None:
                return self.getTokens(Aim8Parser.SEMICOLON)
            else:
                return self.getToken(Aim8Parser.SEMICOLON, i)

        def getRuleIndex(self):
            return Aim8Parser.RULE_mexpr

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterMexpr" ):
                listener.enterMexpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitMexpr" ):
                listener.exitMexpr(self)




    def mexpr(self):

        localctx = Aim8Parser.MexprContext(self, self._ctx, self.state)
        self.enterRule(localctx, 4, self.RULE_mexpr)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 40
            self.match(Aim8Parser.NAME)
            self.state = 41
            self.match(Aim8Parser.LBRACKET)
            self.state = 42
            self.expr()
            self.state = 47
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while _la==Aim8Parser.SEMICOLON:
                self.state = 43
                self.match(Aim8Parser.SEMICOLON)
                self.state = 44
                self.expr()
                self.state = 49
                self._errHandler.sync(self)
                _la = self._input.LA(1)

            self.state = 50
            self.match(Aim8Parser.RBRACKET)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class ExprContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def sexpr(self):
            return self.getTypedRuleContext(Aim8Parser.SexprContext,0)


        def mexpr(self):
            return self.getTypedRuleContext(Aim8Parser.MexprContext,0)


        def conditional_expr(self):
            return self.getTypedRuleContext(Aim8Parser.Conditional_exprContext,0)


        def getRuleIndex(self):
            return Aim8Parser.RULE_expr

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterExpr" ):
                listener.enterExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitExpr" ):
                listener.exitExpr(self)




    def expr(self):

        localctx = Aim8Parser.ExprContext(self, self._ctx, self.state)
        self.enterRule(localctx, 6, self.RULE_expr)
        try:
            self.state = 55
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [Aim8Parser.LPAREN, Aim8Parser.NULL, Aim8Parser.ALT_NULL, Aim8Parser.ATOM]:
                self.enterOuterAlt(localctx, 1)
                self.state = 52
                self.sexpr()
                pass
            elif token in [Aim8Parser.NAME]:
                self.enterOuterAlt(localctx, 2)
                self.state = 53
                self.mexpr()
                pass
            elif token in [Aim8Parser.LBRACKET]:
                self.enterOuterAlt(localctx, 3)
                self.state = 54
                self.conditional_expr()
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class Conditional_exprContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def LBRACKET(self):
            return self.getToken(Aim8Parser.LBRACKET, 0)

        def conditional_pair(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(Aim8Parser.Conditional_pairContext)
            else:
                return self.getTypedRuleContext(Aim8Parser.Conditional_pairContext,i)


        def RBRACKET(self):
            return self.getToken(Aim8Parser.RBRACKET, 0)

        def SEMICOLON(self, i:int=None):
            if i is None:
                return self.getTokens(Aim8Parser.SEMICOLON)
            else:
                return self.getToken(Aim8Parser.SEMICOLON, i)

        def getRuleIndex(self):
            return Aim8Parser.RULE_conditional_expr

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterConditional_expr" ):
                listener.enterConditional_expr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitConditional_expr" ):
                listener.exitConditional_expr(self)




    def conditional_expr(self):

        localctx = Aim8Parser.Conditional_exprContext(self, self._ctx, self.state)
        self.enterRule(localctx, 8, self.RULE_conditional_expr)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 57
            self.match(Aim8Parser.LBRACKET)
            self.state = 58
            self.conditional_pair()
            self.state = 63
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while _la==Aim8Parser.SEMICOLON:
                self.state = 59
                self.match(Aim8Parser.SEMICOLON)
                self.state = 60
                self.conditional_pair()
                self.state = 65
                self._errHandler.sync(self)
                _la = self._input.LA(1)

            self.state = 66
            self.match(Aim8Parser.RBRACKET)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class Conditional_pairContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def predicate(self):
            return self.getTypedRuleContext(Aim8Parser.PredicateContext,0)


        def arrow(self):
            return self.getTypedRuleContext(Aim8Parser.ArrowContext,0)


        def expr(self):
            return self.getTypedRuleContext(Aim8Parser.ExprContext,0)


        def getRuleIndex(self):
            return Aim8Parser.RULE_conditional_pair

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterConditional_pair" ):
                listener.enterConditional_pair(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitConditional_pair" ):
                listener.exitConditional_pair(self)




    def conditional_pair(self):

        localctx = Aim8Parser.Conditional_pairContext(self, self._ctx, self.state)
        self.enterRule(localctx, 10, self.RULE_conditional_pair)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 68
            self.predicate()
            self.state = 69
            self.arrow()
            self.state = 70
            self.expr()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class PredicateContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def mexpr(self):
            return self.getTypedRuleContext(Aim8Parser.MexprContext,0)


        def proposition(self):
            return self.getTypedRuleContext(Aim8Parser.PropositionContext,0)


        def getRuleIndex(self):
            return Aim8Parser.RULE_predicate

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterPredicate" ):
                listener.enterPredicate(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitPredicate" ):
                listener.exitPredicate(self)




    def predicate(self):

        localctx = Aim8Parser.PredicateContext(self, self._ctx, self.state)
        self.enterRule(localctx, 12, self.RULE_predicate)
        try:
            self.state = 74
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,5,self._ctx)
            if la_ == 1:
                self.enterOuterAlt(localctx, 1)
                self.state = 72
                self.mexpr()
                pass

            elif la_ == 2:
                self.enterOuterAlt(localctx, 2)
                self.state = 73
                self.proposition()
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class PropositionContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def mexpr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(Aim8Parser.MexprContext)
            else:
                return self.getTypedRuleContext(Aim8Parser.MexprContext,i)


        def connective(self):
            return self.getTypedRuleContext(Aim8Parser.ConnectiveContext,0)


        def getRuleIndex(self):
            return Aim8Parser.RULE_proposition

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterProposition" ):
                listener.enterProposition(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitProposition" ):
                listener.exitProposition(self)




    def proposition(self):

        localctx = Aim8Parser.PropositionContext(self, self._ctx, self.state)
        self.enterRule(localctx, 14, self.RULE_proposition)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 76
            self.mexpr()
            self.state = 77
            self.connective()
            self.state = 78
            self.mexpr()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class ConnectiveContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def AND(self):
            return self.getToken(Aim8Parser.AND, 0)

        def OR(self):
            return self.getToken(Aim8Parser.OR, 0)

        def NOT(self):
            return self.getToken(Aim8Parser.NOT, 0)

        def getRuleIndex(self):
            return Aim8Parser.RULE_connective

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterConnective" ):
                listener.enterConnective(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitConnective" ):
                listener.exitConnective(self)




    def connective(self):

        localctx = Aim8Parser.ConnectiveContext(self, self._ctx, self.state)
        self.enterRule(localctx, 16, self.RULE_connective)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 80
            _la = self._input.LA(1)
            if not((((_la) & ~0x3f) == 0 and ((1 << _la) & ((1 << Aim8Parser.AND) | (1 << Aim8Parser.OR) | (1 << Aim8Parser.NOT))) != 0)):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class ArrowContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def ARROW(self):
            return self.getToken(Aim8Parser.ARROW, 0)

        def ALT_ARROW(self):
            return self.getToken(Aim8Parser.ALT_ARROW, 0)

        def getRuleIndex(self):
            return Aim8Parser.RULE_arrow

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterArrow" ):
                listener.enterArrow(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitArrow" ):
                listener.exitArrow(self)




    def arrow(self):

        localctx = Aim8Parser.ArrowContext(self, self._ctx, self.state)
        self.enterRule(localctx, 18, self.RULE_arrow)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 82
            _la = self._input.LA(1)
            if not(_la==Aim8Parser.ARROW or _la==Aim8Parser.ALT_ARROW):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class NullContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def NULL(self):
            return self.getToken(Aim8Parser.NULL, 0)

        def ALT_NULL(self):
            return self.getToken(Aim8Parser.ALT_NULL, 0)

        def getRuleIndex(self):
            return Aim8Parser.RULE_null

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterNull" ):
                listener.enterNull(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitNull" ):
                listener.exitNull(self)




    def null(self):

        localctx = Aim8Parser.NullContext(self, self._ctx, self.state)
        self.enterRule(localctx, 20, self.RULE_null)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 84
            _la = self._input.LA(1)
            if not(_la==Aim8Parser.NULL or _la==Aim8Parser.ALT_NULL):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx

    class AtomContext(ParserRuleContext):

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def ATOM(self):
            return self.getToken(Aim8Parser.ATOM, 0)

        def getRuleIndex(self):
            return Aim8Parser.RULE_atom

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAtom" ):
                listener.enterAtom(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAtom" ):
                listener.exitAtom(self)




    def atom(self):

        localctx = Aim8Parser.AtomContext(self, self._ctx, self.state)
        self.enterRule(localctx, 22, self.RULE_atom)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 86
            self.match(Aim8Parser.ATOM)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx





