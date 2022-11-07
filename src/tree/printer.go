package tree

import "fmt"

type Printer struct{}

func (p Printer) VisitBinaryExpr(expr Binary) interface{} {
	return p.Parenthesize(expr.Operator.Text, expr.Left, expr.Right)
}

func (p Printer) VisitGroupingExpr(expr Grouping) interface{} {
	return p.Parenthesize("group", expr.Expression)
}

func (p Printer) VisitLiteralExpr(expr Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}

	return fmt.Sprint(expr.Value)
}

func (p Printer) VisitUnaryExpr(expr Unary) interface{} {
	return p.Parenthesize(expr.Operator.Text, expr.Right)
}

func (p Printer) Print(expr Expr) string {
	return expr.Accept(p).(string)
}

func (p Printer) Parenthesize(name string, exprs ...Expr) string {
	var str string
	str += "(" + name
	for _, expr := range exprs {
		str += " " + p.Print(expr)
	}
	str += ")"
	return str
}
