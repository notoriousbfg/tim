package main

import (
	"tim/lexer"
	"tim/parser"
	"tim/tree"
)

func main() {
	l := lexer.New("200 + 200 != 400")
	p := parser.New(l.Tokens)
	tree.Print(p.Parse())

	// expression := tree.Binary{
	// 	Left: tree.Unary{
	// 		Operator: token.Token{
	// 			Type:     token.MINUS,
	// 			Text:     "-",
	// 			Position: 0,
	// 		},
	// 		Right: tree.Literal{
	// 			Value: 123,
	// 		},
	// 	},
	// 	Operator: token.Token{
	// 		Type:     token.PLUS,
	// 		Text:     "+",
	// 		Position: 5,
	// 	},
	// 	Right: tree.Grouping{
	// 		Expression: tree.Literal{
	// 			Value: 456,
	// 		},
	// 	},
	// }

	// printer := tree.Printer{}
	// fmt.Println(printer.Print(expression))
}
