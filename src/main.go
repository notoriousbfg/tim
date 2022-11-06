package main

import "tim/lexer"

func main() {
	l := lexer.New("(five: 5 >= 4)")
	l.PrintTokens()
}
