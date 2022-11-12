package tree

import (
	"fmt"
	"strconv"
	"tim/token"
)

func Interpret(expression Expr) interface{} {
	interpreter := &Interpreter{}
	value, err := interpreter.evaluate(expression)
	if err != nil {
		panic(err) // todo - how to report runtime errors?
	}
	return value
}

type Interpreter struct{}

func (i *Interpreter) VisitBinaryExpr(expr Binary) (interface{}, error) {
	left, _ := i.evaluate(expr.Left)
	right, _ := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.MINUS:
		i.CheckNumberOperands(left, right)
		return left.(float64) - right.(float64), nil
	case token.PLUS:
		i.CheckNumberOperands(left, right)
		return left.(float64) + right.(float64), nil
	case token.STAR:
		i.CheckNumberOperands(left, right)
		return left.(float64) * right.(float64), nil
	case token.GREATER:
		i.CheckNumberOperands(left, right)
		return left.(float64) > right.(float64), nil
	case token.GREATER_EQUAL:
		i.CheckNumberOperands(left, right)
		return left.(float64) >= right.(float64), nil
	case token.LESS:
		i.CheckNumberOperands(left, right)
		return left.(float64) < right.(float64), nil
	case token.LESS_EQUAL:
		i.CheckNumberOperands(left, right)
		return left.(float64) <= right.(float64), nil
	case token.BANG_EQUAL:
		return left == right, nil
	case token.DOUBLE_EQUAL:
		return left != right, nil
	}
	return nil, nil
}

func (i *Interpreter) VisitLiteralExpr(expr Literal) (interface{}, error) {
	return expr.Value, nil
}

func (i *Interpreter) VisitGroupingExpr(expr Grouping) (interface{}, error) {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitUnaryExpr(expr Unary) (interface{}, error) {
	right, _ := i.evaluate(expr)
	switch expr.Operator.Type {
	case token.BANG:
		return !i.IsTruthy(right), nil
	case token.MINUS:
		return right.(float64) * -1, nil // ? zeros are going to be a PITA
	}
	return nil, nil
}

func (i *Interpreter) IsTruthy(val interface{}) bool {
	if val == nil {
		return false
	}
	switch i := val.(type) {
	case float64:
	case float32:
	case int64:
	case int32:
	case int:
	case uint64:
	case uint32:
	case uint:
		return i != 0
	}
	return true
}

func (i *Interpreter) CheckNumberOperands(left interface{}, right interface{}) {
	if _, ok := left.(float64); ok {
		if _, ok = right.(float64); ok {
			return
		}
	}
	panic(NewRuntimeError("operands must be a number"))
}

func (i *Interpreter) evaluate(expr Expr) (interface{}, error) {
	expression, err := expr.Accept(i)
	if err != nil {
		return nil, err
	}
	return expression, nil
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
		converted, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return 0, NewRuntimeError(fmt.Sprintf("%v is not a number", i))
		}
		return converted, nil
	default:
		return 0, NewRuntimeError(fmt.Sprintf("%v is not a number", i))
	}
	return 0, NewRuntimeError(fmt.Sprintf("%v is not a number", val)) // never
}
