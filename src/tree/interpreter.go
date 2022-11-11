package tree

import (
	"encoding/json"
	"fmt"
	"strconv"
	"tim/token"
)

func Interpret(expression Expr) {
	interpreter := &Interpreter{}
	value, err := interpreter.evaluate(expression)
	if err != nil {
		panic(err) // todo - how to report runtime errors?
	}
	json, _ := json.Marshal(value)
	fmt.Println(string(json))
}

type Interpreter struct{}

func (i *Interpreter) VisitBinaryExpr(expr Binary) (interface{}, error) {
	left, _ := i.evaluate(expr.Left)
	right, _ := i.evaluate(expr.Right)

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
		return leftFloat - rightFloat, nil
	case token.PLUS:
		return leftFloat + rightFloat, nil
	case token.STAR:
		return leftFloat * rightFloat, nil
	case token.GREATER:
		return leftFloat > rightFloat, nil
	case token.GREATER_EQUAL:
		return leftFloat >= rightFloat, nil
	case token.LESS:
		return leftFloat < rightFloat, nil
	case token.LESS_EQUAL:
		return leftFloat <= rightFloat, nil
	case token.BANG_EQUAL:
		return !i.IsEqual(left, right), nil
	case token.DOUBLE_EQUAL:
		return i.IsEqual(left, right), nil
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
	case string:
		converted, _ := strconv.Atoi(i)
		return converted != 0
		// default:
		// return 0, NewRuntimeError(fmt.Sprintf("%v is not a number", val))
	}
	return true
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
	default:
		// return 0, NewRuntimeError(fmt.Sprintf("%v is not a number", val))
	}
	return 0, nil // never
}
