package tree_test

import (
	"fmt"
	"testing"
	"tim/lexer"
	"tim/parser"
	"tim/tree"
)

type InterpretedCase struct {
	InputString string
	Expected    interface{}
}

func TestInterpreter(t *testing.T) {
	cases := map[string]InterpretedCase{
		"basic addition": {
			InputString: "200 + 200",
			Expected:    400.00,
		},
		"basic addition with decimals": {
			InputString: "200.23 + 200",
			Expected:    400.23,
		},
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			p := parser.New(l.Tokens)
			parsed := p.Parse()
			fmt.Println(tree.Print(parsed))
			actual := tree.Interpret(parsed)
			if testcase.Expected != actual {
				fmt.Printf("%t, %t", testcase.Expected, actual)
				t.Fatal("expressions do not match", testcase.Expected, actual)
			}
		})
	}
}
