package lexer

import (
	"fmt"
	"strconv"
	"strings"
	"tim/token"
	"unicode"

	"golang.org/x/exp/slices"
)

func New(input string) Lexer {
	lexer := Lexer{
		Input:   input,
		Line:    1,
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
	Input      string
	Tokens     []token.Token
	Start      int
	Current    int
	Line       int
	insertSemi bool
}

func (l *Lexer) ReadInput() error {
	for !l.isAtEnd() {
		l.Start = l.Current
		err := l.ReadChar()
		if err != nil {
			return err
		}
		if l.insertSemi {
			l.insertSemi = false
			l.AddToken(token.SEMICOLON, ";", "\\n")
		}
	}
	l.Start++
	l.AddToken(token.EOF, "", "")
	return nil
}

func (l *Lexer) ReadChar() error {
	char := l.NextChar()
	canInsertSemi := false
	switch char {
	case "(":
		l.AddToken(token.LEFT_PAREN, char, char)
	case ")":
		l.AddToken(token.RIGHT_PAREN, char, char)
		if !slices.Contains([]string{".", ")", ",", "}", "=>"}, l.peek()) {
			canInsertSemi = true
		}
	case "{":
		canInsertSemi = false
		l.AddToken(token.LEFT_BRACE, char, char)
	case "}":
		canInsertSemi = false
		l.AddToken(token.RIGHT_BRACE, char, char)
	case ",":
		canInsertSemi = false
		l.AddToken(token.COMMA, char, char)
	// case ";":
	// 	l.AddToken(token.SEMICOLON, char, char)
	case ".":
		canInsertSemi = false
		l.AddToken(token.DOT, char, char)
	case "+":
		canInsertSemi = false
		if l.matchNext("+") {
			// canInsertSemi = true
			l.AddToken(token.INCREMENT, "++", "++")
		} else {
			l.AddToken(token.PLUS, char, char)
		}
	case "-":
		canInsertSemi = false
		if l.matchNext("-") {
			l.AddToken(token.DECREMENT, "--", "--")
		} else {
			l.AddToken(token.MINUS, char, char)
		}
	case "*":
		canInsertSemi = false
		l.AddToken(token.STAR, char, char)
	case "/":
		canInsertSemi = false
		l.AddToken(token.SLASH, char, char)
	case "?":
		canInsertSemi = false
		l.AddToken(token.QUESTION, char, char)
	case "!":
		canInsertSemi = false
		if l.matchNext("=") {
			l.AddToken(token.BANG_EQUAL, "!=", "!=")
		} else {
			l.AddToken(token.BANG, char, char)
		}
	case "=":
		canInsertSemi = false
		if l.matchNext(">") {
			l.AddToken(token.DOUBLE_ARROW, "=>", "=>")
		} else if l.matchNext("=") {
			l.AddToken(token.DOUBLE_EQUAL, "==", "==")
		} else {
			l.AddToken(token.EQUAL, char, char)
		}
	case "<":
		canInsertSemi = false
		if l.matchNext("=") {
			l.AddToken(token.LESS_EQUAL, "<=", "<=")
		} else {
			l.AddToken(token.LESS, char, char)
		}
	case ">":
		canInsertSemi = false
		if l.matchNext("=") {
			l.AddToken(token.GREATER_EQUAL, ">=", ">=")
		} else if l.matchNext(">") {
			l.AddToken(token.RETURN, ">>", ">>")
		} else {
			l.AddToken(token.GREATER, char, char)
		}
	case ":":
		canInsertSemi = false
		l.AddToken(token.COLON, char, char)
	case "\"":
		canInsertSemi = false
		l.matchString()
		// canInsertSemi = true
	case "\n":
		canInsertSemi = false
		// canInsertSemi = false
		// l.AddToken(token.SEMICOLON, ";", "\\n")
		l.Line++
	case " ", "\r", "\t":
		break
	default:
		canInsertSemi = false
		if isDigit(char) {
			l.matchNumber()
		} else if isLetter(char) {
			l.matchIdentifier()
		} else {
			return fmt.Errorf("unsupported type: %s", char)
		}
	}
	l.insertSemi = canInsertSemi
	return nil
}

func (l *Lexer) NextChar() string {
	char := string(l.Input[l.Current])
	l.Current++
	return char
}

func (l *Lexer) AddToken(tokenType token.TokenType, text string, literal interface{}) {
	l.Tokens = append(l.Tokens, token.Token{
		Type:     tokenType,
		Text:     text,
		Literal:  literal,
		Position: l.Start,
		Line:     l.Line,
	})
}

func (l *Lexer) TokenTypes() []token.TokenType {
	var types []token.TokenType
	for _, token := range l.Tokens {
		types = append(types, token.Type)
	}
	return types
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

func (l *Lexer) peekNext() string {
	if l.Current+1 >= len(l.Input) {
		return ""
	}
	return string(l.Input[l.Current+1])
}

func (l *Lexer) matchNumber() {
	for isDigit(l.peek()) {
		l.NextChar()
	}

	if l.peek() == "." && isDigit(l.peekNext()) {
		l.NextChar()

		for isDigit(l.peek()) {
			l.NextChar()
		}
	}

	text := l.Input[l.Start:l.Current]

	var val interface{}
	if strings.Contains(text, ".") {
		val, _ = strconv.ParseFloat(text, 64)
	} else {
		intVal, _ := strconv.ParseInt(text, 10, 0)
		val = int(intVal)
	}

	l.AddToken(token.NUMBER, text, val)
}

func (l *Lexer) matchString() {
	for l.peek() != "\"" && !l.isAtEnd() {
		l.NextChar()
	}

	l.NextChar()
	text := l.Input[l.Start+1 : l.Current-1]
	l.AddToken(token.STRING, text, text)
}

func (l *Lexer) matchIdentifier() {
	for isAlphaNumeric(l.peek()) && !l.isAtEnd() {
		l.NextChar()
	}

	text := l.Input[l.Start:l.Current]
	if tokenType, ok := token.Keywords()[text]; ok {
		l.AddToken(tokenType, text, text)
	} else {
		l.AddToken(token.IDENTIFIER, text, text)
	}
}

func (l *Lexer) matchNext(expected string) bool {
	if string(l.Input[l.Current]) != expected {
		return false
	}
	l.NextChar()
	return true
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
