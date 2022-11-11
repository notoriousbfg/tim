package main

import (
	"tim/lexer"
	"tim/parser"
	"tim/tree"
)

func main() {
	l := lexer.New("200 + 200 == 400")
	p := parser.New(l.Tokens)
	parsed := p.Parse()
	tree.Interpret(parsed)
}
