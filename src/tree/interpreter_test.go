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
		"basic addition: 2 integers": {
			InputString: "200 + 200",
			Expected:    400,
		},
		"basic addition: 1 integer, 1 float": {
			InputString: "200 + 200.45",
			Expected:    400.45,
		},
		"basic subtraction: 2 integers": {
			InputString: "300 - 200",
			Expected:    100,
		},
		"concatenation: 2 strings": {
			InputString: "\"hello \" + \"world\"",
			Expected:    "hello world",
		},
		"concatenation: 1 string and 1 number": {
			InputString: "\"hello \" + 123",
			Expected:    "hello 123",
		},
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			p := parser.New(l.Tokens)
			parsed := p.Parse()
			actual := tree.Interpret(parsed)
			if testcase.Expected != actual {
				t.Fatal("expressions do not match", fmt.Sprintf("%t", testcase.Expected), fmt.Sprintf("%t", actual))
			}
		})
	}
}
