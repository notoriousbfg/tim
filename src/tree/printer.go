package tree

import "fmt"

func Print(expr Expr) string {
	printer := &Printer{}
	return printer.Print(expr)
}

type Printer struct{}

func (p *Printer) VisitBinaryExpr(expr Binary) (interface{}, error) {
	return p.parenthesise(expr.Operator.Text, expr.Left, expr.Right), nil
}

func (p *Printer) VisitGroupingExpr(expr Grouping) (interface{}, error) {
	return p.parenthesise("group", expr.Expression), nil
}

func (p *Printer) VisitLiteralExpr(expr Literal) (interface{}, error) {
	if expr.Value == nil {
		return "nil", nil
	}

	return fmt.Sprint(expr.Value), nil
}

func (p *Printer) VisitUnaryExpr(expr Unary) (interface{}, error) {
	return p.parenthesise(expr.Operator.Text, expr.Right), nil
}

func (p *Printer) Print(expr Expr) string {
	expression, err := expr.Accept(p)
	if err != nil {
		return fmt.Sprintf("there was an error")
	}
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
