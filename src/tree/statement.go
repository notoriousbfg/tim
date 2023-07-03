package tree

import (
	"tim/token"
)

type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) interface{}
	// VisitPrintStmt(stmt PrintStmt) interface{}
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

type VariableStmt struct {
	Name        token.Token
	Initializer Expr
}

func (vs VariableStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitVariableStmt(vs)
}

type ListStmt struct {
	Items     []Stmt
	Functions []Expr
}

func (ls ListStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitListStmt(ls)
}

func (ls ListStmt) Length() int {
	return len(ls.Items)
}

type CallStmt struct {
	Initialiser  Stmt
	Callee       Expr
	ClosingParen token.Token
	Arguments    []Expr
}

func (cs CallStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitCallStmt(cs)
}
