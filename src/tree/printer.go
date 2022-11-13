package tree

import "fmt"

func Print(expr Expr) string {
	printer := &Printer{}
	return printer.Print(expr)
}

type Printer struct{}

func (p *Printer) VisitBinaryExpr(expr Binary) interface{} {
	return p.parenthesise(expr.Operator.Text, expr.Left, expr.Right)
}

func (p *Printer) VisitGroupingExpr(expr Grouping) interface{} {
	return p.parenthesise("group", expr.Expression)
}

func (p *Printer) VisitLiteralExpr(expr Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}

	return fmt.Sprint(expr.Value)
}

func (p *Printer) VisitUnaryExpr(expr Unary) interface{} {
	return p.parenthesise(expr.Operator.Text, expr.Right)
}

func (p *Printer) Print(expr Expr) string {
	expression := expr.Accept(p)
	return expression.(string)
}

func (p *Printer) parenthesise(name string, exprs ...Expr) string {
	var str string
	str += "(" + name
	for _, expr := range exprs {
		str += " " + p.Print(expr)
	}
	str += ")"
	return str
}
