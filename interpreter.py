#!/usr/bin/env python3
import sys

from antlr4 import (
    InputStream,
    ParseTreeWalker,
    CommonTokenStream,
)

from Aim8Lexer import Aim8Lexer
from Aim8Parser import Aim8Parser
from Aim8Listener import Aim8Listener


def main(argv):
    input = InputStream(argv[1])
    lexer = Aim8Lexer(input)
    stream = CommonTokenStream(lexer)
    parser = Aim8Parser(stream)
    tree = parser.expr()
    printer = TreePrinter()
    walker = ParseTreeWalker()
    walker.walk(printer, tree)


class TreePrinter(Aim8Listener):
    def enterEveryRule(self, ctx):
        print(ctx.__class__)
        print(ctx.getText())


if __name__ == '__main__':
    main(sys.argv)
