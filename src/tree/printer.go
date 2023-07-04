package tree

import "fmt"

type PrintVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) string
	VisitVariableStmt(stmt VariableStmt) string
	VisitListStmt(stmt ListStmt) string
	VisitCallStmt(stmt CallStmt) string
	VisitBinaryExpr(expr Binary) string
	VisitGroupingExpr(expr Grouping) string
	VisitLiteralExpr(expr Literal) string
	VisitUnaryExpr(expr Unary) string
	VisitVariableExpr(expr Variable) string
}

func Print(stmt Stmt) {
	printer := &Printer{}
	printer.Print(stmt)
}

// should implement every method in PrintVisitor
type Printer struct{}

func (p *Printer) VisitExpressionStmt(stmt ExpressionStmt) string {
	// return stmt.Expr.Accept(p)
	return stmt.Print(p)
}

func (p *Printer) VisitVariableStmt(stmt VariableStmt) string {
	return stmt.Print(p)
}

func (p *Printer) VisitListStmt(stmt ListStmt) string {
	return stmt.Print(p)
}

func (p *Printer) VisitCallStmt(stmt CallStmt) string {
	return stmt.Print(p)
}

func (p *Printer) VisitBinaryExpr(expr Binary) string {
	return expr.Print(p)
}

func (p *Printer) VisitGroupingExpr(expr Grouping) string {
	return expr.Print(p)
}

func (p *Printer) VisitLiteralExpr(expr Literal) string {
	// if expr.Value == nil {
	// 	return "nil"
	// }

	// return fmt.Sprint(expr.Value)
	return expr.Print(p)
}

func (p *Printer) VisitUnaryExpr(expr Unary) string {
	return expr.Print(p)
}

func (p *Printer) VisitVariableExpr(expr Variable) string {
	return expr.Print(p)
}

func (p *Printer) Print(stmt Stmt) {
	fmt.Println(stmt.Print(p))
}

// func (p *Printer) parenthesise(name string, stmts ...Stmt) string {
// 	var str string
// 	str += "(" + name
// 	for _, stmt := range stmts {
// 		str += " " + p.Print(stmt)
// 	}
// 	str += ")"
// 	return str
// }
