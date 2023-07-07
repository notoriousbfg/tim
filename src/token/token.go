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
	RETURN    // >>

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	TRUE
	FALSE
	NIL

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
	case RETURN:
		return "RETURN"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case IDENTIFIER:
		return "IDENTIFIER"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
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

// temporary
func ListFunctions() []string {
	return []string{
		"print",
	}
}

func IsListFunction(name string) bool {
	for _, function := range ListFunctions() {
		if name == function {
			return true
		}
	}
	return false
}

func Keywords() map[string]TokenType {
	return map[string]TokenType{
		"true":  TRUE,
		"false": FALSE,
		"nil":   NIL,
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
