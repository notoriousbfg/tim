package lexer

import (
	"fmt"
	"strconv"
	"tim/token"
)

func New(input string) {
	lexer := Lexer{
		Input:   input,
		Start:   0,
		Current: 0,
	}
	err := lexer.ReadInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", lexer.Tokens)
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
		fmt.Printf("%d:%d \n", l.Start, l.Current)
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
	case " ", "\n", "\t":
		break
	case "\"":
		l.matchString()
	default:
		if isDigit(char) {
			l.matchNumber()
		}
		// } else {
		// 	return errors.New("unsupported type")
		// }
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

func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
