package token

const (
	// single char tokens
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	COLON
	PLUS
	MINUS

	// two-char tokens
	FUNC_BODY
	EQUAL
	GREATER
	LESS
	GREATER_EQUAL
	LESS_EQUAL

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	CALL

	EOF
)

type TokenType int

func (tt TokenType) String() string {
	switch tt {
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case EOF:
		return "EOF"
	default:
		return ""
	}
}

type Token struct {
	Type     TokenType
	Text     string
	Position int
}
