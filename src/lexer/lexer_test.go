package lexer_test

import (
	"reflect"
	"testing"
	"tim/lexer"
)

type TokenCase struct {
	InputString string
	Types       []string
}

func TestLexer(t *testing.T) {
	cases := map[string]TokenCase{
		"basic string variable": {
			InputString: "(hello: \"hello\")",
			Types: []string{
				"LEFT_PAREN",
				"IDENTIFIER",
				"COLON",
				"STRING",
				"RIGHT_PAREN",
				"EOF",
			},
		},
		"basic integer variable": {
			InputString: "(five: 5)",
			Types: []string{
				"LEFT_PAREN",
				"IDENTIFIER",
				"COLON",
				"NUMBER",
				"RIGHT_PAREN",
				"EOF",
			},
		},
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			if !reflect.DeepEqual(l.TokenTypeStrings(), testcase.Types) {
				t.Fatal("types do not match", testcase.Types, l.TokenTypeStrings())
			}
		})
	}
}
