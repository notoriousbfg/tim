package parser

import (
	"fmt"
	"tim/token"
	"tim/tree"
)

func New(tokens []token.Token) Parser {
	return Parser{
		Tokens:  tokens,
		Current: 0,
	}
}

type Parser struct {
	Tokens  []token.Token
	Current int
}

func (p *Parser) Parse() []tree.Stmt {
	statements := make([]tree.Stmt, 0)
	for !p.isAtEnd() {
		statements = append(statements, p.Declaration())
	}
	return statements
}

func (p *Parser) Declaration() tree.Stmt {
	// if p.match(token.IDENTIFIER) {
	// 	// if p.check(token.COLON) {
	// 	// 	// name := p.consume(token.IDENTIFIER, "expect variable name")
	// 	// 	// fmt.Printf("var: %s, name: %+v", name)
	// 	// 	return p.varDeclaration(p.previous())
	// 	// }
	// }

	// variable declaration
	// we must use check in case it is not a declaration but an expression
	// if p.match(token.IDENTIFIER) {
	// 	identifier := p.previous()

	// 	if p.match(token.COLON) {
	// 		return p.VarDeclaration(identifier)
	// 	} else {
	// 		// this is REALLY problematic
	// 		return tree.Variable{Name: p.Expression()}
	// 	}
	// }

	if p.checkSequence(token.IDENTIFIER, token.COLON) {
		identifier := p.peek()
		p.advance()
		p.advance()
		return p.VarDeclaration(identifier)
	}

	if p.match(token.LEFT_PAREN) {
		return p.List()
	}

	return p.Statement()

	// todo: error handling
}

func (p *Parser) List() tree.Stmt {
	var items []tree.Stmt

	items = append(items, p.Declaration())

	for !p.check(token.RIGHT_PAREN) && !p.isAtEnd() {
		if p.match(token.COMMA) {
			items = append(items, p.Declaration())
		}
	}

	p.consume(token.RIGHT_PAREN, "expect ')' after list")
	p.consume(token.SEMICOLON, "expect ';' after list")

	// // there could be multiple functions
	// var functions []tree.Expr
	// if p.check(token.DOT) {
	// 	p.advance()

	// 	// this would only work for a single call
	// 	p.consume(token.IDENTIFIER, "expect identifier after '.'")
	// 	callable := p.Call()
	// 	functions = append(functions, callable)

	// 	// for !p.check(token.SEMICOLON) && !p.isAtEnd() {
	// 	// 	// callable := p.Call()
	// 	// 	// fmt.Printf("%+v", callable)
	// 	// 	// functions = append(functions, callable)
	// 	// 	// p.advance()
	// 	// }
	// }

	p.expectSemicolon()

	return tree.ListStmt{
		Items: items,
		// Functions: functions,
	}
}

// func (p *Parser) ListItem() tree.Stmt {

// }

func (p *Parser) Statement() tree.Stmt {
	// if p.match(token.PRINT) {
	// 	return p.printStatement()
	// }

	return p.ExpressionStatement()
}

// func (p *Parser) printStatement() tree.Stmt {
// 	value := p.Expression()
// 	p.consume(token.SEMICOLON, "Expect ';' after value.")
// 	printStmt := &tree.PrintStmt{
// 		Expr: value,
// 	}
// 	// printStmt.Print(value)
// 	return printStmt
// }

func (p *Parser) ExpressionStatement() tree.Stmt {
	value := p.Expression()
	p.expectSemicolon()
	exprStmt := tree.ExpressionStmt{
		Expr: value,
	}
	// exprStmt.Expression(value)
	return exprStmt
}

// todo: allow a variable to be another list
func (p *Parser) VarDeclaration(identifier token.Token) tree.Stmt {
	initializer := p.Expression()

	return &tree.VariableStmt{
		Name:        identifier,
		Initializer: initializer,
	}
}

func (p *Parser) Expression() tree.Expr {
	return p.Equality()
}

func (p *Parser) Equality() tree.Expr {
	expr := p.Comparison()

	for p.match(token.DOUBLE_EQUAL, token.BANG_EQUAL) {
		expr = tree.Binary{
			Left:     expr,
			Operator: p.previous(),
			Right:    p.Comparison(),
		}
	}

	return expr
}

func (p *Parser) Comparison() tree.Expr {
	expr := p.Term()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		expr = tree.Binary{
			Left:     expr,
			Operator: p.previous(),
			Right:    p.Term(),
		}
	}
	return expr
}

func (p *Parser) Term() tree.Expr {
	expr := p.Factor()
	for p.match(token.MINUS, token.PLUS) {
		expr = tree.Binary{
			Left:     expr,
			Operator: p.previous(),
			Right:    p.Factor(),
		}
	}
	return expr
}

func (p *Parser) Factor() tree.Expr {
	expr := p.Unary()
	for p.match(token.STAR, token.SLASH) {
		expr = tree.Binary{
			Left:     expr,
			Operator: p.previous(),
			Right:    p.Unary(),
		}
	}
	return expr
}

func (p *Parser) Unary() tree.Expr {
	if p.match(token.MINUS) {
		return tree.Unary{
			Operator: p.previous(),
			Right:    p.Unary(),
		}
	}
	return p.Primary()
}

func (p *Parser) Call() tree.Expr {
	var expr tree.Expr
	for {
		if p.match(token.LEFT_PAREN) {
			expr = p.finishCall(expr)
		} else {
			break
		}
	}

	return expr
}

func (p *Parser) finishCall(callee tree.Expr) tree.Expr {
	var arguments []tree.Expr
	if !p.check(token.RIGHT_PAREN) {
		for p.match(token.COMMA) {
			arguments = append(arguments, p.Expression())
		}
	}
	paren := p.consume(token.RIGHT_PAREN, "expect ')' after arguments")
	return tree.Call{
		Callee:    callee,
		Paren:     paren,
		Arguments: arguments,
	}
}

func (p *Parser) Primary() tree.Expr {
	if p.match(token.FALSE) {
		return tree.Literal{Value: false}
	}
	if p.match(token.TRUE) {
		return tree.Literal{Value: true}
	}
	if p.match(token.NIL) {
		return tree.Literal{Value: nil}
	}
	if p.match(token.NUMBER, token.STRING) {
		return tree.Literal{Value: p.previous().Literal}
	}
	if p.match(token.IDENTIFIER) {
		identifier := p.previous()
		// we're not declaring a variable
		if !p.check(token.COLON) {
			return tree.Variable{Name: identifier}
		}
	}
	panic(p.error(p.peek(), "expect expression."))
}

// check that the current token is any of the types and advance if so
func (p *Parser) match(tokenTypes ...token.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

// check that the current token is of a type
func (p *Parser) check(tokenType token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) checkSequence(tokens ...token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	startIndex := p.Current
	for index, tokenType := range tokens {
		thisToken := p.peekAt(startIndex + index)
		if thisToken.Type != tokenType {
			return false
		}
	}
	return true
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.Current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

// get the token at the current index
func (p *Parser) peek() token.Token {
	return p.Tokens[p.Current]
}

// get the token at the specified index
func (p *Parser) peekAt(index int) token.Token {
	return p.Tokens[index]
}

// get the previous token
func (p *Parser) previous() token.Token {
	return p.Tokens[p.Current-1]
}

// if the token is of the specified type advance, otherwise panic
func (p *Parser) consume(tokenType token.TokenType, message string) token.Token {
	if p.check(tokenType) {
		return p.advance()
	}

	panic(p.error(p.peek(), message))
}

func (p *Parser) expectSemicolon() {
	nextToken := p.peek()
	if nextToken.Type == token.SEMICOLON && nextToken.Literal == "\\n" {
		p.advance()
	}
}

// func (p *Parser) synchronise() {
// 	p.advance()

// 	for !p.isAtEnd() {
// 		if p.previous().Type == token.NEWLINE {
// 			return
// 		}

// 		switch p.peek().Type {
// 		case token.LEFT_BRACE: // almost all new statements in timlang begin with a left brace
// 			return
// 		}

// 		p.advance()
// 	}
// }

func (p *Parser) error(thisToken token.Token, message string) *ParseError {
	var where string
	if thisToken.Type == token.EOF {
		where = " at end"
	} else {
		where = " at '" + thisToken.Text + "'"
	}
	err := &ParseError{
		Message: fmt.Sprintf("[line %d] Error%s: %s\n", thisToken.Line, where, message),
	}
	return err
}

type ParseError struct {
	Message string
}

func (pe *ParseError) Error() string {
	return pe.Message
}
