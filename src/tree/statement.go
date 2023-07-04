package tree

import (
	"tim/token"
)

type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
	Print(visitor PrintVisitor) string
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) interface{}
	VisitVariableStmt(stmt VariableStmt) interface{}
	VisitListStmt(stmt ListStmt) interface{}
	VisitCallStmt(stmt CallStmt) interface{}
}

type ExpressionStmt struct {
	Expr Expr
}

func (es ExpressionStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitExpressionStmt(es)
}

func (es ExpressionStmt) Print(visitor PrintVisitor) string {
	return "expression statement"
}

type VariableStmt struct {
	Name        token.Token
	Initializer Expr
}

func (vs VariableStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitVariableStmt(vs)
}

func (vs VariableStmt) Print(visitor PrintVisitor) string {
	return "variable statement"
}

type ListStmt struct {
	Items     []Stmt
	Functions []Expr
}

func (ls ListStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitListStmt(ls)
}

func (ls ListStmt) Print(visitor PrintVisitor) string {
	return "list statement"
}

func (ls ListStmt) Length() int {
	return len(ls.Items)
}

type CallStmt struct {
	Initialiser  ListStmt
	Callee       Expr
	ClosingParen token.Token
	Arguments    []Expr
}

func (cs CallStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitCallStmt(cs)
}

func (cs CallStmt) Print(visitor PrintVisitor) string {
	return "call statement"
}
