package main

import (
	"tim/lexer"
	"tim/parser"
)

func main() {
	l := lexer.New(`
		(five: 5)
		(ten: 10)

		(add: (x, y) => {
			return x + y
		})

		(result: (five, ten).call(add))
	`)

	p := parser.New(l.Tokens)

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
