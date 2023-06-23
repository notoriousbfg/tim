package tree

import "tim/token"

type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) interface{}
	// VisitPrintStmt(stmt PrintStmt) interface{}
	VisitVarStmt(stmt VariableStmt) interface{}
	VisitListStmt(stmt ListStmt) interface{}
}

type ExpressionStmt struct {
	Expr Expr
}

func (es ExpressionStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitExpressionStmt(es)
}

// type PrintStmt struct {
// 	Expr Expr
// }

// func (ps *PrintStmt) Print(expr Expr) {
// 	ps.Expr = expr
// }

// func (ps PrintStmt) Accept(visitor StmtVisitor) interface{} {
// 	return visitor.VisitPrintStmt(ps)
// }

type VariableStmt struct {
	Name        token.Token
	Initializer Expr
}

func (vs VariableStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitVarStmt(vs)
}

type ListStmt struct {
	Statements []Stmt
}

func (ls ListStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitListStmt(ls)
}
