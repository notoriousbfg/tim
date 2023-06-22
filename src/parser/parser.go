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
		statements = append(statements, p.statement())
	}
	return statements
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
	if p.match(token.LEFT_PAREN) {
		expr := p.Expression()
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return tree.Grouping{Expression: expr}
	}
	panic(p.error(p.peek(), "expect expression."))
}

func (p *Parser) match(tokenTypes ...token.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tokenType token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
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

func (p *Parser) peek() token.Token {
	return p.Tokens[p.Current]
}

func (p *Parser) previous() token.Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) consume(tokenType token.TokenType, message string) token.Token {
	if p.check(tokenType) {
		return p.advance()
	}

	panic(p.error(p.peek(), message))
}

func (p *Parser) statement() tree.Stmt {
	// this will not be a feature of the language but i intend to removed it when i get to function declarations later
	if p.match(token.PRINT) {
		return p.printStatement()
	}

	return p.expressionStatement()
}

func (p *Parser) printStatement() tree.Stmt {
	value := p.Expression()
	// i'm not sure if this is correct. i don't want semicolons
	p.consume(token.NEWLINE, "Expect '\\n' after value.")
	printStmt := &tree.PrintStmt{
		Expr: value,
	}
	// printStmt.Print(value)
	return printStmt
}

func (p *Parser) expressionStatement() tree.Stmt {
	value := p.Expression()
	// i'm not sure if this is correct. i don't want semicolons
	p.consume(token.NEWLINE, "Expect '\\n' after value.")
	exprStmt := &tree.ExpressionStmt{
		Expr: value,
	}
	// exprStmt.Expression(value)
	return exprStmt
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
		Message: fmt.Sprintf("[line %d] Error%s: %s\n", thisToken.Line+1, where, message),
	}
	return err
}

type ParseError struct {
	Message string
}

func (pe *ParseError) Error() string {
	return pe.Message
}
