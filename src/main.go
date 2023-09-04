package main

import (
	"fmt"
	"os"
	"tim/interpreter"
	"tim/lexer"
	"tim/parser"
)

func main() {
	b, err := os.ReadFile("./basic.tim")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)

	l := lexer.New(input)

	fmt.Printf("%+v", l.Tokens)
	fmt.Println()

	p := parser.New(l.Tokens)
	parsed := p.Parse()

	fmt.Println()

	for _, p := range parsed {
		fmt.Printf("%+v \n\n", p)
	}

	// doesn't work
	// for _, stmt := range parsed {
	// 	fmt.Println(tree.Print(stmt))
	// }

	// json, _ := json.Marshal(interpreter.Interpret(parsed, true))
	// fmt.Println(string(json))

	interpreter.Interpret(parsed, true)
}
