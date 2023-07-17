package parser

import (
	"fmt"
	"tim/token"
	"tim/tree"
)

func New(tokens []token.Token) *Parser {
	return &Parser{
		Tokens:  tokens,
		Current: 0,
	}
}

type Parser struct {
	Tokens            []token.Token
	Current           int
	PreviousStatement tree.Stmt
}

func (p *Parser) Parse() []tree.Stmt {
	statements := make([]tree.Stmt, 0)
	for !p.isAtEnd() {
		statements = append(statements, p.Declaration())
	}
	return statements
}

func (p *Parser) Declaration() tree.Stmt {
	if p.match(token.LEFT_PAREN) {
		return p.Iterable()
	}

	if p.checkSequence(token.IDENTIFIER, token.COLON) {
		identifier := p.peek()

		// advance for identifier and colon
		// p.advance()
		// p.advance()
		p.advanceBy(2)

		return p.VarDeclaration(identifier)
	}

	// if p.checkSequence(token.DOUBLE_ARROW, token.RIGHT_BRACE) {
	// 	return p.FuncDeclaration()
	// }

	return p.Statement()

	// todo: error handling
}

func (p *Parser) Iterable() tree.Stmt {
	var items []tree.Stmt
	for !p.check(token.RIGHT_PAREN) && !p.isAtEnd() {
		if p.check(token.COMMA) {
			p.advance()
		}

		items = append(items, p.Declaration())
	}

	p.consume(token.RIGHT_PAREN, "expected ')' after expression")

	// if the next two chars are double arrow and right parent, this is a func declaration
	// items becomes the function args
	if p.checkSequence(token.DOUBLE_ARROW, token.LEFT_BRACE) {
		p.advanceBy(2)

		return p.FunctionDeclaration(items)
	}

	// otherwise this is a list
	var listFunctions []tree.CallStmt
	for p.checkSequence(token.DOT, token.IDENTIFIER) {
		p.advance()

		listFunctions = append(listFunctions, p.Call())
	}

	if !p.check(token.RIGHT_PAREN) && !p.check(token.COMMA) && !p.check(token.DOT) && !p.check(token.RIGHT_BRACE) {
		p.consume(token.SEMICOLON, "expected ')'")
	}

	p.expectSemicolon()

	return tree.ListStmt{
		Items:     items,
		Functions: listFunctions,
	}
}

func (p *Parser) Call() tree.CallStmt {
	// name of function
	callee := p.Primary()

	p.consume(token.LEFT_PAREN, "expect '(' after function declaration")

	var arguments []tree.Expr
	for !p.check(token.RIGHT_PAREN) {
		if p.check(token.COMMA) {
			p.advance()
		}

		arguments = append(arguments, p.Expression())
	}

	closingParen := p.consume(token.RIGHT_PAREN, "expected ')'")

	return tree.CallStmt{
		Callee:       callee,
		ClosingParen: closingParen,
		Arguments:    arguments,
	}
}

func (p *Parser) VarDeclaration(identifier token.Token) tree.Stmt {
	initializer := p.Declaration()

	return tree.VariableStmt{
		Name:        identifier,
		Initializer: initializer,
	}
}

func (p *Parser) FunctionDeclaration(arguments []tree.Stmt) tree.Stmt {
	body := p.Block()

	return tree.FuncStmt{
		Body:      body,
		Arguments: arguments,
	}
}

func (p *Parser) Block() []tree.Stmt {
	statements := make([]tree.Stmt, 0)

	for !p.check(token.RIGHT_BRACE) && !p.isAtEnd() {
		statements = append(statements, p.Declaration())
	}

	p.consume(token.RIGHT_BRACE, "expect '}' after block")

	return statements
}

func (p *Parser) Statement() tree.Stmt {
	if p.match(token.RETURN) {
		return p.ReturnStatement()
	}

	return p.ExpressionStatement()
}

func (p *Parser) ReturnStatement() tree.Stmt {
	returnToken := p.previous()

	var value tree.Stmt
	if !p.check(token.SEMICOLON) {
		value = p.Declaration()
	}

	p.expectSemicolon()

	return tree.ReturnStmt{
		Token: returnToken,
		Value: value,
	}
}

func (p *Parser) ExpressionStatement() tree.Stmt {
	value := p.Expression()
	p.expectSemicolon()
	exprStmt := tree.ExpressionStmt{
		Expr: value,
	}
	return exprStmt
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

		// we're using a variable or function, not declaring one
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

// func (p *Parser) matchSequence(tokenTypes ...token.TokenType) bool {
// 	if p.isAtEnd() {
// 		return false
// 	}
// 	startIndex := p.Current
// 	for index, tokenType := range tokenTypes {
// 		thisToken := p.peekAt(startIndex + index)
// 		if thisToken.Type != tokenType {
// 			return false
// 		}
// 	}
// 	for range tokenTypes {
// 		p.advance()
// 	}
// 	return true
// }

// check that the current token is of a type
func (p *Parser) check(tokenType token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) checkSequence(tokenTypes ...token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	startIndex := p.Current
	for index, tokenType := range tokenTypes {
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

func (p *Parser) advanceBy(count int) token.Token {
	for i := 0; i < count; i++ {
		if !p.isAtEnd() {
			p.Current++
		}
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
