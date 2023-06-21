package tree

import (
	"fmt"
	"io"
	"math"
	"os"
	"tim/token"
)

func Interpret(expression Expr) interface{} {
	interpreter := &Interpreter{}

	defer func() {
		if err := recover(); err != nil {
			if e, ok := err.(RuntimeError); ok {
				_, _ = interpreter.stdErr.Write([]byte(e.Error() + "\n"))
				os.Exit(70)
			} else {
				fmt.Printf("Error: %s\n", err)
			}
		}
	}()

	return interpreter.evaluate(expression)
}

type Interpreter struct {
	stdErr io.Writer
}

func (i *Interpreter) VisitBinaryExpr(expr Binary) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	// will timlang inherit the same floating-point arithmetic hell as go?
	var returnValue interface{}
	switch expr.Operator.Type {
	case token.MINUS:
		i.checkNumberOperands(left, right)
		returnValue = left.(float64) - right.(float64)
	case token.PLUS:
		if i.checkStringOperand(left) || i.checkStringOperand(right) {
			returnValue = fmt.Sprint(left, right)
		} else {
			i.checkNumberOperands(left, right)
			returnValue = left.(float64) + right.(float64)
		}
	case token.SLASH:
		i.checkNumberOperands(left, right)
		returnValue = left.(float64) / right.(float64)
	case token.STAR:
		i.checkNumberOperands(left, right)
		returnValue = left.(float64) * right.(float64)
	case token.GREATER:
		i.checkNumberOperands(left, right)
		returnValue = left.(float64) > right.(float64)
	case token.GREATER_EQUAL:
		i.checkNumberOperands(left, right)
		returnValue = left.(float64) >= right.(float64)
	case token.LESS:
		i.checkNumberOperands(left, right)
		returnValue = left.(float64) < right.(float64)
	case token.LESS_EQUAL:
		i.checkNumberOperands(left, right)
		returnValue = left.(float64) <= right.(float64)
	case token.BANG_EQUAL:
		returnValue = left != right
	case token.DOUBLE_EQUAL:
		returnValue = left == right
	}
	// if returnValue can be expressed as an int, return it as such
	if returnFloat, isFloat := returnValue.(float64); isFloat {
		if wholeFloat := math.Trunc(returnFloat); returnFloat == wholeFloat {
			return int(wholeFloat)
		}
		return returnFloat
	}
	return returnValue
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

func (i *Interpreter) evaluate(expr Expr) interface{} {
	expression := expr.Accept(i)
	return expression
}

func (i *Interpreter) checkStringOperand(val interface{}) bool {
	if _, ok := val.(string); ok {
		return true
	}
	return false
}

func (i *Interpreter) checkNumberOperands(left interface{}, right interface{}) {
	if _, ok := left.(float64); ok {
		if _, ok = right.(float64); ok {
			return
		}
	}
	panic(NewRuntimeError("operands must be number"))
}
