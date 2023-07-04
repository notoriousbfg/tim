package tree

import (
	"tim/token"
)

type Expr interface {
	Accept(visitor ExprVisitor) interface{}
	Print(visitor PrintVisitor) string
}

type ExprVisitor interface {
	VisitBinaryExpr(expr Binary) interface{}
	VisitGroupingExpr(expr Grouping) interface{}
	VisitLiteralExpr(expr Literal) interface{}
	VisitUnaryExpr(expr Unary) interface{}
	VisitVariableExpr(expr Variable) interface{}
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (b Binary) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitBinaryExpr(b)
}

func (b Binary) Print(visitor PrintVisitor) string {
	return "binary expression"
}

type Grouping struct {
	Expression Expr
}

func (g Grouping) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(g)
}

func (g Grouping) Print(visitor PrintVisitor) string {
	return "grouping expression"
}

type Literal struct {
	Value interface{}
}

func (l Literal) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(l)
}

func (l Literal) Print(visitor PrintVisitor) string {
	return "literal expression"
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (u Unary) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitUnaryExpr(u)
}

func (u Unary) Print(visitor PrintVisitor) string {
	return "unary expression"
}

type Variable struct {
	Name token.Token
}

func (v Variable) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitVariableExpr(v)
}

func (v Variable) Print(visitor PrintVisitor) string {
	return ""
}
