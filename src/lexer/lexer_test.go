package lexer_test

import (
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
				token.SEMICOLON,
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
				token.SEMICOLON,
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
				token.SEMICOLON,
				token.EOF,
			},
		},
		"function": {
			InputString: "(addOne: (myNumber) => { >> myNumber + 1 })",
			Types: []token.TokenType{
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.COLON,
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.RIGHT_PAREN,
				token.DOUBLE_ARROW,
				token.LEFT_BRACE,
				token.RETURN,
				token.IDENTIFIER,
				token.PLUS,
				token.NUMBER,
				token.RIGHT_BRACE,
				token.RIGHT_PAREN,
				token.SEMICOLON,
				token.EOF,
			},
		},
		"boolean": {
			InputString: "(trueValue: true)",
			Types: []token.TokenType{
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.COLON,
				token.TRUE,
				token.RIGHT_PAREN,
				token.SEMICOLON,
				token.EOF,
			},
		},
		"double equal and star": {
			InputString: "(isTrue: (5 * 10) == 50)",
			Types: []token.TokenType{
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.COLON,
				token.LEFT_PAREN,
				token.NUMBER,
				token.STAR,
				token.NUMBER,
				token.RIGHT_PAREN,
				token.DOUBLE_EQUAL,
				token.NUMBER,
				token.RIGHT_PAREN,
				token.SEMICOLON,
				token.EOF,
			},
		},
		"decimal number": {
			InputString: "(five: 200.32)",
			Types: []token.TokenType{
				token.LEFT_PAREN,
				token.IDENTIFIER,
				token.COLON,
				token.NUMBER,
				token.RIGHT_PAREN,
				token.SEMICOLON,
				token.EOF,
			},
		},
	}

	for name, testcase := range cases {
		t.Run(name, func(t *testing.T) {
			l := lexer.New(testcase.InputString)
			if !slicesMatch(l.TokenTypes(), testcase.Types) {
				t.Fatal("types do not match", testcase.Types, l.TokenTypes())
			}
		})
	}
}

func slicesMatch(a []token.TokenType, b []token.TokenType) bool {
	if len(a) != len(b) {
		return false
	}

	for index, aType := range a {
		bType := b[index]
		if bType != aType {
			return false
		}
	}
	return true
}
