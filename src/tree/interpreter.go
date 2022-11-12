package tree

import (
	"math"
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

func NewInterpreter() {}

type Interpreter struct{}

func (i *Interpreter) VisitBinaryExpr(expr Binary) (interface{}, error) {
	left, _ := i.evaluate(expr.Left)
	right, _ := i.evaluate(expr.Right)

	// will timlang inherit the same floating-point arithmetic hell as go?
	var returnValue interface{}
	switch expr.Operator.Type {
	case token.MINUS:
		if err := i.checkNumberOperands(left, right); err != nil {
			return nil, err
		}
		returnValue = left.(float64) - right.(float64)
	case token.PLUS:
		if err := i.checkNumberOperands(left, right); err != nil {
			return nil, err
		}
		returnValue = left.(float64) + right.(float64)
	case token.STAR:
		if err := i.checkNumberOperands(left, right); err != nil {
			return nil, err
		}
		returnValue = left.(float64) * right.(float64)
	case token.GREATER:
		if err := i.checkNumberOperands(left, right); err != nil {
			return nil, err
		}
		returnValue = left.(float64) > right.(float64)
	case token.GREATER_EQUAL:
		if err := i.checkNumberOperands(left, right); err != nil {
			return nil, err
		}
		returnValue = left.(float64) >= right.(float64)
	case token.LESS:
		if err := i.checkNumberOperands(left, right); err != nil {
			return nil, err
		}
		returnValue = left.(float64) < right.(float64)
	case token.LESS_EQUAL:
		if err := i.checkNumberOperands(left, right); err != nil {
			return nil, err
		}
		returnValue = left.(float64) <= right.(float64)
	case token.BANG_EQUAL:
		returnValue = left == right
	case token.DOUBLE_EQUAL:
		returnValue = left != right
	}
	// if returnValue can be expressed as an int, return it as such
	if returnFloat, isFloat := returnValue.(float64); isFloat {
		if wholeFloat := math.Trunc(returnFloat); returnFloat == wholeFloat {
			return int(wholeFloat), nil
		}
		return returnFloat, nil
	}
	return returnValue, nil
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

func (i *Interpreter) evaluate(expr Expr) (interface{}, error) {
	expression, err := expr.Accept(i)
	if err != nil {
		return nil, err
	}
	return expression, nil
}

func (i *Interpreter) checkNumberOperands(left interface{}, right interface{}) error {
	if _, ok := left.(float64); ok {
		if _, ok = right.(float64); ok {
			return nil
		}
	}
	return NewRuntimeError("Operands must be number")
}
