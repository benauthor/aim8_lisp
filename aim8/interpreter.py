#!/usr/bin/env python3
import sys

from antlr4 import (
    InputStream,
    ParseTreeWalker,
    CommonTokenStream,
)
from antlr4.tree import Tree

from .generated.Aim8Lexer import Aim8Lexer
from .generated.Aim8Parser import Aim8Parser
from .generated.Aim8Listener import Aim8Listener


def aim8EvalString(input):
    lexer = Aim8Lexer(InputStream(input))
    stream = CommonTokenStream(lexer)
    parser = Aim8Parser(stream)
    tree = parser.expr()
    # TMP
    # import ipdb;ipdb.set_trace()
    printer = TreePrinter()
    walker = ParseTreeWalker()
    walker.walk(printer, tree)
    # /TMP
    return aim8eval(tree)


def main():
    input = sys.argv[1]
    print(aim8EvalString(input))


class _TRUE(object):
    def __repr__(self):
        return "#1"


class _FALSE(object):
    def __repr__(self):
        return "#0"


class _NIL(object):
    def __repr__(self):
        return "Λ"


TRUE = _TRUE()
FALSE = _FALSE()
NIL = _NIL()


def null(node):
    result = aim8eval(node)
    if result == NIL:
        return TRUE
    return FALSE


funcs = {
    "null": null
}


def apply(funcname, *args):
    if funcname not in funcs:
        raise NameError("%s is not defined" % funcname)
    return funcs[funcname](*args)


def aim8eval(node):
    if isinstance(node, Aim8Parser.ExprContext):
        return aim8eval(node.children[0])
    elif isinstance(node, Aim8Parser.SexprContext):
        return aim8eval(node.children[0])
    elif isinstance(node, Aim8Parser.Enclosd_sexprContext):
        return list(filter(lambda n: n is not None,
                           [aim8eval(c) for c in node.children]))
    elif isinstance(node, Aim8Parser.MexprContext):
        name = node.children[0].getText()
        rest = [c for c in node.children[1:]
                if not isinstance(c, Tree.TerminalNodeImpl)]
        return apply(name, *rest)
    elif isinstance(node, Aim8Parser.NullContext):
        return NIL
    elif isinstance(node, Aim8Parser.AtomContext):
        return node.getText()
    elif isinstance(node, Tree.TerminalNodeImpl):
        return None
    raise RuntimeError("not yet implemented %s" % node.__class__)


class TreePrinter(Aim8Listener):
    def enterEveryRule(self, ctx):
        print(ctx.__class__)
        print(ctx.getText())


if __name__ == '__main__':
    main()
