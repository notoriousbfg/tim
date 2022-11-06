package lexer

import (
	"fmt"
	"tim/token"
)

func New(input string) {
	lexer := Lexer{
		Input: input,
	}
	err := lexer.ReadInput()
	if err != nil {
		panic(err)
	}
}

type Lexer struct {
	Input    string
	Tokens   []token.Token
	Position int
}

func (l Lexer) ReadInput() error {
	for _, char := range l.Input {
		err := l.ReadChar(string(char))
		if err != nil {
			return err
		}
		l.NextChar()
	}
	l.AddToken(token.EOF, "", l.Position)
	fmt.Printf("%+v", l.Tokens)
	return nil
}

func (l *Lexer) ReadChar(char string) error {
	switch char {
	case "(":
		l.AddToken(token.LEFT_PAREN, "(", l.Position)
	case ")":
		l.AddToken(token.RIGHT_PAREN, ")", l.Position)
		// default:
		// 	return errors.New("unsupported type")
	}
	return nil
}

func (l *Lexer) NextChar() {
	l.Position++
}

func (l *Lexer) AddToken(tokenType token.TokenType, text string, position int) {
	l.Tokens = append(l.Tokens, token.Token{
		Type:     tokenType,
		Text:     text,
		Position: position,
	})
}
