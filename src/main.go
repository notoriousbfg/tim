package main

import (
	"fmt"
	"tim/lexer"
	"tim/parser"
	"tim/tree"
)

func main() {
	l := lexer.New("200 + 200 == 400")
	p := parser.New(l.Tokens)
	parsed := p.Parse()
	fmt.Println(tree.Print(parsed))
	tree.Interpret(parsed)
}
