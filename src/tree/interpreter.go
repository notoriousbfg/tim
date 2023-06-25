package tree

import (
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"tim/token"
)

func Interpret(statements []Stmt, printPanics bool) (result []interface{}) {
	interpreter := &Interpreter{
		Environment: NewEnvironment(),
	}
	if printPanics {
		defer interpreter.printToStdErr()
	}
	for _, statement := range statements {
		result = append(result, interpreter.execute(statement))
	}
	return
}

type Interpreter struct {
	stdErr      io.Writer
	Environment *Environment
}

func (i *Interpreter) printToStdErr() {
	if err := recover(); err != nil {
		if e, ok := err.(RuntimeError); ok {
			_, _ = i.stdErr.Write([]byte(e.Error() + "\n"))
			os.Exit(70)
		} else {
			fmt.Printf("Error: %s\n", err)
		}
	}
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
		// if i.checkStringOperand(left) || i.checkStringOperand(right) {
		// 	returnValue = fmt.Sprint(left, right)
		// } else {
		i.checkNumberOperands(left, right)
		returnValue = left.(float64) + right.(float64)
		// }
	case token.SLASH:
		i.checkNumberOperands(left, right)
		if isZero(left) || isZero(right) {
			panic(NewRuntimeError(DivisionByZero))
		}
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

func (i *Interpreter) VisitCallExpr(expr Call) interface{} {
	callee := i.evaluate(expr.Callee)
	var arguments []interface{}
	for _, arg := range expr.Arguments {
		arguments = append(arguments, i.evaluate(arg))
	}
	// var listArguments []interface{}
	// for _, arg := range expr.ListArguments {
	// 	listArguments = append(listArguments, i.evaluate(arg))
	// }
	return callee.(Callable).Call(i, arguments)
}

func (i *Interpreter) VisitVariableExpr(expr Variable) interface{} {
	return i.Environment.Get(expr.Name)
}

func (i *Interpreter) VisitExpressionStmt(stmt ExpressionStmt) interface{} {
	return i.evaluate(stmt.Expr)
}

// func (i *Interpreter) VisitPrintStmt(stmt PrintStmt) interface{} {
// 	value := i.evaluate(stmt.Expr)
// 	// _, _ = i.stdOut.Write([]byte(i.stringify(value) + "\n"))
// 	fmt.Println(i.stringify(value))
// 	return value
// }

func (i *Interpreter) VisitVarStmt(stmt VariableStmt) interface{} {
	var value interface{}
	if stmt.Initializer != nil {
		value = i.evaluate(stmt.Initializer)
	}
	i.Environment.Define(stmt.Name.Text, value)
	return nil
}

func (i *Interpreter) VisitListStmt(stmt ListStmt) interface{} {
	var values []interface{}
	for _, item := range stmt.Items {
		values = append(values, i.execute(item))
	}
	// should we evaluate list here?
	// for _, function := range stmt.Functions {

	// }
	return values
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
	panic(NewRuntimeError(OperandsMustBeNumber))
}

func (i *Interpreter) execute(stmt Stmt) interface{} {
	return stmt.Accept(i)
}

func (i *Interpreter) stringify(value interface{}) string {
	if value == nil {
		return "nil"
	}
	return fmt.Sprint(value)
}

func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}
