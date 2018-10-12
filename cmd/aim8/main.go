package main

import (
	"os"

	"github.com/benauthor/aim8_lisp/internal/parser"
)

func main() {
	parser.PrintTree(os.Args[1])
}
