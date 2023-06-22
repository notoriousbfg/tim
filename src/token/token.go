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
	STAR
	SLASH
	QUESTION
	SEMICOLON

	// one or two-char tokens
	DOUBLE_ARROW
	DOUBLE_EQUAL
	BANG
	BANG_EQUAL
	EQUAL
	GREATER
	LESS
	GREATER_EQUAL
	LESS_EQUAL
	INCREMENT // ++
	DECREMENT // --

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	CALL
	RETURN
	TRUE
	FALSE
	NIL
	EACH
	FILTER
	MAP
	PRINT

	NEWLINE
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
	case QUESTION:
		return "QUESTION"
	case SEMICOLON:
		return "SEMICOLON"
	case EQUAL:
		return "EQUAL"
	case DOUBLE_ARROW:
		return "DOUBLE_ARROW"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case INCREMENT:
		return "INCREMENT"
	case DECREMENT:
		return "DECREMENT"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case IDENTIFIER:
		return "IDENTIFIER"
	case CALL:
		return "CALL"
	case RETURN:
		return "RETURN"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case EACH:
		return "EACH"
	case FILTER:
		return "FILTER"
	case MAP:
		return "MAP"
	case PRINT:
		return "PRINT"
	case NEWLINE:
		return "NEWLINE"
	case EOF:
		return "EOF"
	default:
		return ""
	}
}

type Token struct {
	Type     TokenType
	Text     string
	Literal  interface{}
	Position int
	Line     int
}

func Keywords() map[string]TokenType {
	return map[string]TokenType{
		"call":   CALL,
		"return": RETURN,
		"true":   TRUE,
		"false":  FALSE,
		"nil":    NIL,
		"each":   EACH,
		"filter": FILTER,
		"map":    MAP,
		"print":  PRINT,
	}
}

func IsKeyword(tt TokenType) bool {
	for _, tokenType := range Keywords() {
		if tt == tokenType {
			return true
		}
	}
	return false
}
