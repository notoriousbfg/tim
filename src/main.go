package main

import (
	"encoding/json"
	"fmt"
	"tim/lexer"
	"tim/parser"
	"tim/tree"
)

func main() {
	l := lexer.New("\"hello\" + \"tim\"")
	p := parser.New(l.Tokens)
	parsed := p.Parse()
	fmt.Println(tree.Print(parsed))
	json, _ := json.Marshal(tree.Interpret(parsed, true))
	fmt.Println(string(json))

}
