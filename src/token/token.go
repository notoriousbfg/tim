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
	DOUBLE_ARROW
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
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"
	case COLON:
		return "COLON"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case IDENTIFIER:
		return "IDENTIFIER"
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
