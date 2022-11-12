package main

import (
	"encoding/json"
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
	json, _ := json.Marshal(tree.Interpret(parsed))
	fmt.Println(string(json))

}
