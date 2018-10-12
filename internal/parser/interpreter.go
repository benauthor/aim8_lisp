package parser

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*BaseAim8Listener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func PrintTree(rawInput string) {
	fmt.Println(rawInput)
	input := antlr.NewInputStream(rawInput)
	lexer := NewAim8Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewAim8Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Sexpr()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
