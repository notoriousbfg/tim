package lexer

import (
	"errors"
	"fmt"
	"strconv"
	"tim/token"
	"unicode"
)

func New(input string) Lexer {
	lexer := Lexer{
		Input:   input,
		Start:   0,
		Current: 0,
	}
	err := lexer.ReadInput()
	if err != nil {
		panic(err)
	}
	return lexer
}

type Lexer struct {
	Input   string
	Tokens  []token.Token
	Start   int
	Current int
}

func (l *Lexer) ReadInput() error {
	for !l.isAtEnd() {
		l.Start = l.Current
		err := l.ReadChar()
		if err != nil {
			return err
		}
	}
	l.Start++
	l.AddToken(token.EOF, "")
	return nil
}

func (l *Lexer) ReadChar() error {
	char := l.NextChar()
	switch char {
	case "(":
		l.AddToken(token.LEFT_PAREN, char)
	case ")":
		l.AddToken(token.RIGHT_PAREN, char)
	case "{":
		l.AddToken(token.LEFT_BRACE, char)
	case "}":
		l.AddToken(token.RIGHT_BRACE, char)
	case ",":
		l.AddToken(token.COMMA, char)
	case ".":
		l.AddToken(token.DOT, char)
	case "+":
		l.AddToken(token.PLUS, char)
	case "-":
		l.AddToken(token.MINUS, char)
	case ":":
		l.AddToken(token.COLON, char)
	case "\"":
		l.matchString()
	case " ", "\n", "\t":
		break
	default:
		if isDigit(char) {
			l.matchNumber()
		} else if isLetter(char) {
			l.matchIdentifier()
		} else {
			return errors.New("unsupported type")
		}
	}
	return nil
}

func (l *Lexer) NextChar() string {
	defer func() {
		if !l.isAtEnd() {
			l.Current++
		}
	}()

	return string(l.Input[l.Current])
}

func (l *Lexer) AddToken(tokenType token.TokenType, text string) {
	l.Tokens = append(l.Tokens, token.Token{
		Type:     tokenType,
		Text:     text,
		Position: l.Start,
	})
}

func (l *Lexer) TokenTypeStrings() []string {
	var strings []string
	for _, token := range l.Tokens {
		strings = append(strings, token.Type.String())
	}
	return strings
}

func (l *Lexer) isAtEnd() bool {
	return l.Current >= len(l.Input)
}

func (l *Lexer) peek() string {
	if l.isAtEnd() {
		return ""
	}

	return string(l.Input[l.Current])
}

func (l *Lexer) matchNumber() {
	if l.isAtEnd() {
		return
	}

	for isDigit(l.peek()) {
		l.NextChar()
	}

	l.AddToken(token.NUMBER, l.Input[l.Start:l.Current])
}

func (l *Lexer) matchString() {
	for l.peek() != "\"" && !l.isAtEnd() {
		l.NextChar()
	}

	l.NextChar()

	l.AddToken(token.STRING, l.Input[l.Start+1:l.Current-1])
}

func (l *Lexer) matchIdentifier() {
	for isAlphaNumeric(l.peek()) {
		l.NextChar()
	}

	l.AddToken(token.IDENTIFIER, l.Input[l.Start:l.Current])
}

func (l *Lexer) PrintTokens() {
	for _, token := range l.Tokens {
		fmt.Printf("%+v \n", token)
	}
}

func isDigit(ch string) bool {
	_, err := strconv.Atoi(ch)
	return err == nil
}

func isLetter(ch string) bool {
	for _, r := range ch {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isAlphaNumeric(ch string) bool {
	return isDigit(ch) || isLetter(ch)
}
