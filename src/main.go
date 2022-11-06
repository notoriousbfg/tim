package main

import "tim/lexer"

func main() {
	l := lexer.New(`
		(five: 5)
		(ten: 10)

		(add: (x, y) => {
			return x + y
		})

		(result: (five, ten).call(add))
	`)
	l.PrintTokens()
}
