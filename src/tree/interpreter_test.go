package tree_test

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
	"tim/lexer"
	"tim/parser"
	"tim/tree"

	"github.com/stretchr/testify/assert"
)

type InterpretedCase struct {
	InputString string
	Expected    interface{}
	Err         error
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
		"subtraction: string and number": {
			InputString: "\"hello\" - 13",
			Err:         tree.NewRuntimeError(tree.OperandsMustBeNumber),
		},
		"concatenation: 2 strings": {
			InputString: "\"hello \" + \"world\"",
			Expected:    "hello world",
		},
		"concatenation: 1 string and 1 number": {
			InputString: "\"hello \" + 123",
			Expected:    "hello 123",
		},
		"division by zero panics": {
			InputString: "10 / 0",
			Err:         tree.NewRuntimeError(tree.DivisionByZero),
		},
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			p := parser.New(l.Tokens)
			parsed := p.Parse()
			actual := tree.Interpret(parsed, true)

			if testcase.Err != nil {
				assert.PanicsWithError(t, testcase.Err.Error(), func() {
					tree.Interpret(parsed, false)
				})
			}

			assert.Equal(t, testcase.Expected, actual, "expressions do not match", fmt.Sprintf("%t", testcase.Expected), fmt.Sprintf("%t", actual))
		})
	}
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
