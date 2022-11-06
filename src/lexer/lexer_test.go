package lexer_test

import (
	"reflect"
	"testing"
	"tim/lexer"
	"tim/token"
)

type TokenCase struct {
	InputString string
	Types       []token.TokenType
}

func TestLexer(t *testing.T) {
	cases := map[string]TokenCase{
		"basic string variable": {
			InputString: "(hello: \"hello\")",
			Types: []token.TokenType{
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.COLON,
				token.STRING,
				token.RIGHT_PAREN,
				token.EOF,
			},
		},
		"basic integer variable": {
			InputString: "(five: 5)",
			Types: []token.TokenType{
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.COLON,
				token.NUMBER,
				token.RIGHT_PAREN,
				token.EOF,
			},
		},
		"greater than or equal operator": {
			InputString: "(five: 5 >= 4)",
			Types: []token.TokenType{
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.COLON,
				token.NUMBER,
				token.GREATER_EQUAL,
				token.NUMBER,
				token.RIGHT_PAREN,
				token.EOF,
			},
		},
		"function": {
			InputString: "(addOne: (myNumber) => { return myNumber + 1 })",
			Types: []token.TokenType{
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.COLON,
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.RIGHT_PAREN,
				token.DOUBLE_ARROW,
				token.LEFT_BRACE,
				token.IDENTIFIER,
				token.IDENTIFIER,
				token.PLUS,
				token.NUMBER,
				token.RIGHT_BRACE,
				token.RIGHT_PAREN,
				token.EOF,
			},
		},
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			if !reflect.DeepEqual(l.TokenTypes(), testcase.Types) {
				t.Fatal("types do not match", testcase.Types, l.TokenTypes())
			}
		})
	}
}
