package tree_test

import (
	"bytes"
	"fmt"
	"io"
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
	StdOut      string
}

func TestInterpreter(t *testing.T) {
	cases := map[string]InterpretedCase{
		"basic expression": {
			InputString: "(\"hello world\")",
			Expected:    "hello world",
		},
		"basic addition: 2 integers": {
			InputString: "(200 + 200)",
			Expected:    400,
		},
		"basic addition: 1 integer, 1 float": {
			InputString: "(200 + 200.45)",
			Expected:    400.45,
		},
		"basic subtraction: 2 integers": {
			InputString: "(300 - 200)",
			Expected:    100,
		},
		"subtraction: string and number": {
			InputString: "(\"hello\" - 13)",
			Err:         tree.NewRuntimeError(tree.OperandsMustBeNumber),
		},
		"concatenation: 2 strings": {
			InputString: "(\"hello \" + \"world\")",
			Expected:    "hello world",
		},
		"concatenation: 1 string and 1 number": {
			InputString: "(\"hello \" + 123)",
			Expected:    "hello 123",
		},
		"division by zero panics": {
			InputString: "(10 / 0)",
			Err:         tree.NewRuntimeError(tree.DivisionByZero),
		},
		"print expression": {
			InputString: "print \"hello world\"",
			StdOut:      "hello world",
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

			if len(testcase.StdOut) > 0 {
				actual = captureStdOut(func() {
					tree.Interpret(parsed, false)
				})
				assert.Equal(t, testcase.StdOut, actual)
				return
			}

			assert.Equal(t, testcase.Expected, actual, "expressions do not match", fmt.Sprintf("%t", testcase.Expected), fmt.Sprintf("%t", actual))
		})
	}
}

func captureStdOut(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	// reading our temp stdout
	return out
}
