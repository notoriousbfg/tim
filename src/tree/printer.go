package tree

import "fmt"

type Printer struct{}

func (p Printer) VisitBinaryExpr(expr Binary) interface{} {
	return p.parenthesize(expr.Operator.Text, expr.Left, expr.Right)
}

func (p Printer) VisitGroupingExpr(expr Grouping) interface{} {
	return p.parenthesize("group", expr.Expression)
}

func (p Printer) VisitLiteralExpr(expr Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}

	return fmt.Sprint(expr.Value)
}

func (p Printer) VisitUnaryExpr(expr Unary) interface{} {
	return p.parenthesize(expr.Operator.Text, expr.Right)
}

func (p Printer) print(expr Expr) string {
	return expr.Accept(p).(string)
}

func (p Printer) parenthesize(name string, exprs ...Expr) string {
	var str string
	str += "(" + name
	for _, expr := range exprs {
		str += " " + p.print(expr)
	}
	str += ")"
	return str
}
