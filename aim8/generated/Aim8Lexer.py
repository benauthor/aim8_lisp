# Generated from Aim8.g4 by ANTLR 4.7.1
from antlr4 import *
from io import StringIO
from typing.io import TextIO
import sys


def serializedATN():
    with StringIO() as buf:
        buf.write("\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\24")
        buf.write("T\b\1\4\2\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7")
        buf.write("\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4\f\t\f\4\r\t\r\4\16")
        buf.write("\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22\4\23\t\23")
        buf.write("\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3")
        buf.write("\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\r\3\16")
        buf.write("\3\16\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\6\22")
        buf.write("L\n\22\r\22\16\22M\3\23\6\23Q\n\23\r\23\16\23R\2\2\24")
        buf.write("\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31")
        buf.write("\16\33\17\35\20\37\21!\22#\23%\24\3\2\4\3\2C\\\3\2c|\2")
        buf.write("U\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13")
        buf.write("\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3")
        buf.write("\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2")
        buf.write("\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2")
        buf.write("%\3\2\2\2\3\'\3\2\2\2\5)\3\2\2\2\7+\3\2\2\2\t-\3\2\2\2")
        buf.write("\13/\3\2\2\2\r\61\3\2\2\2\17\63\3\2\2\2\21\65\3\2\2\2")
        buf.write("\23\67\3\2\2\2\259\3\2\2\2\27;\3\2\2\2\31=\3\2\2\2\33")
        buf.write("@\3\2\2\2\35D\3\2\2\2\37F\3\2\2\2!H\3\2\2\2#K\3\2\2\2")
        buf.write("%P\3\2\2\2\'(\7*\2\2(\4\3\2\2\2)*\7+\2\2*\6\3\2\2\2+,")
        buf.write("\7]\2\2,\b\3\2\2\2-.\7_\2\2.\n\3\2\2\2/\60\7.\2\2\60\f")
        buf.write("\3\2\2\2\61\62\7=\2\2\62\16\3\2\2\2\63\64\7\u039d\2\2")
        buf.write("\64\20\3\2\2\2\65\66\7\u03bd\2\2\66\22\3\2\2\2\678\7\u27f8")
        buf.write("\2\28\24\3\2\2\29:\7\u2229\2\2:\26\3\2\2\2;<\7\u222a\2")
        buf.write("\2<\30\3\2\2\2=>\7/\2\2>?\7@\2\2?\32\3\2\2\2@A\7P\2\2")
        buf.write("AB\7K\2\2BC\7N\2\2C\34\3\2\2\2DE\7\u0080\2\2E\36\3\2\2")
        buf.write("\2FG\7\63\2\2G \3\2\2\2HI\7\62\2\2I\"\3\2\2\2JL\t\2\2")
        buf.write("\2KJ\3\2\2\2LM\3\2\2\2MK\3\2\2\2MN\3\2\2\2N$\3\2\2\2O")
        buf.write("Q\t\3\2\2PO\3\2\2\2QR\3\2\2\2RP\3\2\2\2RS\3\2\2\2S&\3")
        buf.write("\2\2\2\5\2MR\2")
        return buf.getvalue()


class Aim8Lexer(Lexer):

    atn = ATNDeserializer().deserialize(serializedATN())

    decisionsToDFA = [ DFA(ds, i) for i, ds in enumerate(atn.decisionToState) ]

    LPAREN = 1
    RPAREN = 2
    LBRACKET = 3
    RBRACKET = 4
    COMMA = 5
    SEMICOLON = 6
    NULL = 7
    LAMBDA = 8
    ARROW = 9
    AND = 10
    OR = 11
    ALT_ARROW = 12
    ALT_NULL = 13
    NOT = 14
    TRUE = 15
    FALSE = 16
    ATOM = 17
    NAME = 18

    channelNames = [ u"DEFAULT_TOKEN_CHANNEL", u"HIDDEN" ]

    modeNames = [ "DEFAULT_MODE" ]

    literalNames = [ "<INVALID>",
            "'('", "')'", "'['", "']'", "','", "';'", "'\u039B'", "'\u03BB'", 
            "'\u27F6'", "'\u2227'", "'\u2228'", "'->'", "'NIL'", "'~'", 
            "'1'", "'0'" ]

    symbolicNames = [ "<INVALID>",
            "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "COMMA", "SEMICOLON", 
            "NULL", "LAMBDA", "ARROW", "AND", "OR", "ALT_ARROW", "ALT_NULL", 
            "NOT", "TRUE", "FALSE", "ATOM", "NAME" ]

    ruleNames = [ "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "COMMA", "SEMICOLON", 
                  "NULL", "LAMBDA", "ARROW", "AND", "OR", "ALT_ARROW", "ALT_NULL", 
                  "NOT", "TRUE", "FALSE", "ATOM", "NAME" ]

    grammarFileName = "Aim8.g4"

    def __init__(self, input=None, output:TextIO = sys.stdout):
        super().__init__(input, output)
        self.checkVersion("4.7.1")
        self._interp = LexerATNSimulator(self, self.atn, self.decisionsToDFA, PredictionContextCache())
        self._actions = None
        self._predicates = None


