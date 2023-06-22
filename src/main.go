package main

import (
	"encoding/json"
	"fmt"
	"os"
	"tim/lexer"
	"tim/parser"
	"tim/tree"
)

func main() {
	b, err := os.ReadFile("./script.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)

	// l := lexer.New("(\"hello\" + \" tim\")\n")
	l := lexer.New(input)

	// fmt.Printf("%+v", l.Tokens)

	p := parser.New(l.Tokens)
	parsed := p.Parse()

	fmt.Println()

	// for _, p := range parsed {
	// 	fmt.Println(p)
	// }

	// doesn't work
	// for _, stmt := range parsed {
	// 	fmt.Println(tree.Print(stmt))
	// }

	json, _ := json.Marshal(tree.Interpret(parsed, true))
	fmt.Println(string(json))

}
