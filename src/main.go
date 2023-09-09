package main

import (
	"fmt"
	"os"
	"tim/lexer"
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
}
