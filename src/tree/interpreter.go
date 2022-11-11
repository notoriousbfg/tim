package tree

import (
	"fmt"
	"math"
	"strconv"
	"tim/token"
)

type Interpreter struct{}

func (i *Interpreter) VisitBinaryExpr(expr Binary) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	leftFloat, err := convertInterfaceToFloat(left)
	if err != nil {
		panic(err)
	}
	rightFloat, err := convertInterfaceToFloat(right)
	if err != nil {
		panic(err)
	}

	switch expr.Operator.Type {
	case token.MINUS:
		return leftFloat - rightFloat
	case token.PLUS:
		return leftFloat + rightFloat
	case token.STAR:
		return leftFloat * rightFloat
	case token.GREATER:
		return leftFloat > rightFloat
	case token.GREATER_EQUAL:
		return leftFloat >= rightFloat
	case token.LESS:
		return leftFloat < rightFloat
	case token.LESS_EQUAL:
		return leftFloat <= rightFloat
	case token.BANG_EQUAL:
		return !i.IsEqual(left, right)
	case token.DOUBLE_EQUAL:
		return i.IsEqual(left, right)
	}
	return nil
}

func (i *Interpreter) VisitLiteralExpr(expr Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitGroupingExpr(expr Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitUnaryExpr(expr Unary) interface{} {
	right := i.evaluate(expr)
	switch expr.Operator.Type {
	case token.BANG:
		return !i.IsTruthy(right)
	case token.MINUS:
		return right.(float64) * -1 // ? zeros are going to be a PITA
	}
	return nil
}

func (i *Interpreter) IsTruthy(val interface{}) bool {
	if val == nil {
		return false
	}
	floatVal, err := convertInterfaceToFloat(val)
	if err != nil {
		return false
	}
	return floatVal != 0
}

func (i *Interpreter) IsEqual(a interface{}, b interface{}) bool {
	if a == nil && b == nil {
		return false
	}
	if a == nil {
		return false
	}
	return a == b
}

func (i *Interpreter) evaluate(expr Expr) interface{} {
	return expr.Accept(i)
}

func convertInterfaceToFloat(val interface{}) (float64, error) {
	switch i := val.(type) {
	case float64:
		return i, nil
	case float32:
	case int64:
	case int32:
	case int:
	case uint64:
	case uint32:
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	}
	return math.NaN(), fmt.Errorf("can't convert %v to float64", val)
}
