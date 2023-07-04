package tree

import "fmt"

type PrintVisitor interface {
	StmtVisitor
	ExprVisitor
}

func Print(stmt Stmt) {
	printer := &Printer{}
	printer.Print(stmt)
}

// should implement every method in PrintVisitor
type Printer struct{}

func (p *Printer) VisitExpressionStmt(stmt ExpressionStmt) interface{} {
	return stmt.Expr.Accept(p)
}

func (p *Printer) VisitVariableStmt(stmt VariableStmt) interface{} {
	return "VisitVariableStmt"
}

func (p *Printer) VisitListStmt(stmt ListStmt) interface{} {
	return "VisitListStmt"
}

func (p *Printer) VisitCallStmt(stmt CallStmt) interface{} {
	return "VisitCallStmt"
}

func (p *Printer) VisitBinaryExpr(expr Binary) interface{} {
	return "VisitBinaryExpr"
}

func (p *Printer) VisitGroupingExpr(expr Grouping) interface{} {
	return "VisitGroupingExpr"
}

func (p *Printer) VisitLiteralExpr(expr Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}

	return fmt.Sprint(expr.Value)
}

func (p *Printer) VisitUnaryExpr(expr Unary) interface{} {
	return "VisitUnaryExpr"
}

func (p *Printer) VisitVariableExpr(expr Variable) interface{} {
	return expr.Name
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
