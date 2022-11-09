package tree

import (
	"math"
	"tim/token"
)

type Interpreter struct{}

// func (i *Interpreter) VisitBinaryExpr(expr Binary) interface{} {

// }

func (i *Interpreter) VisitLiteralExpr(expr Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitGroupingExpr(expr Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitUnaryExpr(expr Unary) interface{} {
	right := i.evaluate(expr)
	switch expr.Operator.Type {
	case token.MINUS:
		return math.Copysign(right.(float64), -1) // ? zeros are going to be a PITA
	}
	return nil
}

func (i *Interpreter) evaluate(expr Expr) interface{} {
	return expr.Accept(i)
}
