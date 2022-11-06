package main

import "tim/lexer"

func main() {
	l := lexer.New("(hello: \"hello\" + \"tim\")")
	l.PrintTokens()
}
