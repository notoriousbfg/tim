package interpreter_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"tim/errors"
	"tim/interpreter"
	"tim/lexer"
	"tim/parser"

	"github.com/stretchr/testify/assert"
)

type InterpretedCase struct {
	InputString string
	Expected    interface{}
	Err         error
	StdOut      string
}

func TestExpressions(t *testing.T) {
	cases := map[string]InterpretedCase{
		"basic expression": {
			InputString: "(\"hello world\").print()",
			StdOut:      "(\"hello world\")",
		},
		"basic addition: 2 integers": {
			InputString: "(200 + 200).print()",
			StdOut:      "(400)",
		},
		"basic addition: 1 integer, 1 float": {
			InputString: "(200 + 200.45).print()",
			StdOut:      "(400.45)",
		},
		"basic subtraction: 2 integers": {
			InputString: "(300 - 200).print()",
			StdOut:      "(100)",
		},
		"subtraction: string and number": {
			InputString: "(\"hello\" - 13).print()",
			Err:         errors.NewRuntimeError(errors.OperandsMustBeNumber),
		},
		"concatenation: 2 strings": {
			InputString: "(\"hello \" + \"world\").print()",
			StdOut:      "(\"hello world\")",
		},
		"concatenation: 1 string and 1 number": {
			InputString: "(\"hello \" + 123).print()",
			StdOut:      "(\"hello 123\")",
		},
		"division by zero panics": {
			InputString: "(10 / 0).print()",
			Err:         errors.NewRuntimeError(errors.DivisionByZero),
		},
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			p := parser.New(l.Tokens)
			parsed := p.Parse()

			if testcase.Err != nil {
				assert.PanicsWithError(t, testcase.Err.Error(), func() {
					interpreter.Interpret(parsed, false)
				}, "did not panic with '%s'", testcase.Err.Error())
			} else if testcase.StdOut != "" {
				stdOut := captureStdOut(func() {
					interpreter.Interpret(parsed, true)
				})
				assert.Equal(t, testcase.StdOut, stdOut, "expressions do not match", testcase.StdOut, stdOut)
			} else {
				actual := interpreter.Interpret(parsed, true)
				assert.Equal(t, testcase.Expected, actual, "expressions do not match", fmt.Sprintf("%t", testcase.Expected), fmt.Sprintf("%t", actual))
			}

			// if len(testcase.StdOut) > 0 {
			// 	actual = []interface{}{
			// 		captureStdOut(func() {
			// 			interpreter.Interpret(parsed, false)
			// 		}),
			// 	}
			// 	assert.Equal(t, testcase.StdOut, actual)
			// 	return
			// }
		})
	}
}

func TestStatements(t *testing.T) {
	cases := map[string]InterpretedCase{
		"list within list": {
			InputString: "(1, 2, 3, (hello: \"world\")).print()",
			StdOut:      "(1, 2, 3, (\"world\"))",
		},
		"variable declaration": {
			InputString: "(myVar: \"sup\").print()",
			StdOut:      "(\"sup\")",
		},
		"function declaration": {
			InputString: "(myFunc: () => { >> \"hello\"}).print()",
			StdOut:      "(\"<closure>\")",
		},
		"get list offset": {
			InputString: "(1, 2, 3).get(1).print()",
			StdOut:      "2",
		},
		"get list variable": {
			InputString: "((one: 1, two: 2, three: 3).get(\"one\")).print()",
			StdOut:      "1",
		},
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			p := parser.New(l.Tokens)
			parsed := p.Parse()

			if testcase.Err != nil {
				assert.PanicsWithError(t, testcase.Err.Error(), func() {
					interpreter.Interpret(parsed, false)
				}, "did not panic with '%s'", testcase.Err.Error())
			} else if testcase.StdOut != "" {
				stdOut := captureStdOut(func() {
					interpreter.Interpret(parsed, true)
				})
				assert.Equal(t, testcase.StdOut, stdOut, "expressions do not match", testcase.StdOut, stdOut)
			} else {
				actual := interpreter.Interpret(parsed, true)
				assert.Equal(t, testcase.Expected, actual, "expressions do not match", fmt.Sprintf("%t", testcase.Expected), fmt.Sprintf("%t", actual))
			}
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
