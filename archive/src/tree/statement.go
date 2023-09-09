package tree

import (
	"tim/token"
)

type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) interface{}
	VisitVariableStmt(stmt VariableStmt) interface{}
	VisitListStmt(stmt ListStmt) interface{}
	VisitFunctionStmt(stmt FuncStmt) interface{}
	VisitReturnStmt(stmt ReturnStmt) interface{}
}

type ExpressionStmt struct {
	Expr Expr
}

func (es ExpressionStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitExpressionStmt(es)
}

type VariableStmt struct {
	Name        token.Token
	Initializer Stmt
}

func (vs VariableStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitVariableStmt(vs)
}

type ListStmt struct {
	Items     []Stmt
	Functions []CallStmt
}

func (ls ListStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitListStmt(ls)
}

type CallStmt struct {
	Callee       Expr
	ClosingParen token.Token
	Arguments    []Expr
}

// i don't know what to do with this
func (cs CallStmt) Accept(visitor StmtVisitor) interface{} {
	return "<not supported>"
}

type FuncStmt struct {
	Body      []Stmt
	Arguments []Stmt
}

func (fs FuncStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitFunctionStmt(fs)
}

type ReturnStmt struct {
	Token token.Token
	Value Stmt
}

func (rs ReturnStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitReturnStmt(rs)
}
