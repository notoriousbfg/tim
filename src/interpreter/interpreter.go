package interpreter

import (
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"tim/env"
	"tim/errors"
	"tim/token"
	"tim/tree"
)

func Interpret(statements []tree.Stmt, printPanics bool) (result []interface{}) {
	interpreter := &Interpreter{
		Level: 0,
		Environment: &env.Environment{
			Enclosing: nil,
			Values:    make(map[string]interface{}),
		},
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
	Level       int
	Environment *env.Environment
	stdErr      io.Writer
}

func (i *Interpreter) printToStdErr() {
	if err := recover(); err != nil {
		if e, ok := err.(errors.RuntimeError); ok {
			_, _ = i.stdErr.Write([]byte(e.Error() + "\n"))
			os.Exit(70)
		} else {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

func (i *Interpreter) VisitBinaryExpr(expr tree.Binary) interface{} {
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
			panic(errors.NewRuntimeError(errors.DivisionByZero))
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

func (i *Interpreter) VisitLiteralExpr(expr tree.Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitGroupingExpr(expr tree.Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitUnaryExpr(expr tree.Unary) interface{} {
	right := i.evaluate(expr)
	switch expr.Operator.Type {
	case token.BANG:
		return !i.IsTruthy(right)
	case token.MINUS:
		return right.(float64) * -1 // ? zeros are going to be a PITA
	}
	return nil
}

func (i *Interpreter) VisitVariableExpr(expr tree.Variable) interface{} {
	return i.Environment.Get(expr.Name)
}

func (i *Interpreter) VisitCallStmt(stmt tree.CallStmt) interface{} {
	// do we need to reexecute a statement if it's a pointer to an already executed statement?
	callee := i.execute(*stmt.Callee)
	var arguments []interface{}
	for _, arg := range stmt.Arguments {
		arguments = append(arguments, i.evaluate(arg))
	}
	// var listArguments []interface{}
	// for _, arg := range expr.ListArguments {
	// 	listArguments = append(listArguments, i.evaluate(arg))
	// }
	return callee.(Callable).Call(i, arguments)
}

func (i *Interpreter) VisitExpressionStmt(stmt tree.ExpressionStmt) interface{} {
	return i.evaluate(stmt.Expr)
}

// func (i *Interpreter) VisitPrintStmt(stmt PrintStmt) interface{} {
// 	value := i.evaluate(stmt.Expr)
// 	// _, _ = i.stdOut.Write([]byte(i.stringify(value) + "\n"))
// 	fmt.Println(i.stringify(value))
// 	return value
// }

func (i *Interpreter) VisitVariableStmt(stmt tree.VariableStmt) interface{} {
	var value interface{}
	if stmt.Initializer != nil {
		value = i.evaluate(stmt.Initializer)
	}
	i.Environment.Define(stmt.Name.Text, value)
	return nil
}

func (i *Interpreter) VisitListStmt(stmt tree.ListStmt) interface{} {
	i.Level++
	defer func() {
		i.Level--
	}()
	// does this suck?
	environment := i.Environment
	if i.Level > 1 {
		environment = env.NewEnvironment(i.Environment)
	}
	return i.executeList(stmt.Items, environment)
}

func (i *Interpreter) executeList(items []tree.Stmt, environment *env.Environment) []interface{} {
	previous := i.Environment
	i.Environment = environment
	var values []interface{}
	for _, item := range items {
		values = append(values, i.execute(item))
	}
	i.Environment = previous
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

func (i *Interpreter) evaluate(expr tree.Expr) interface{} {
	expression := expr.Accept(i)
	return expression
}

// func (i *Interpreter) checkStringOperand(val interface{}) bool {
// 	if _, ok := val.(string); ok {
// 		return true
// 	}
// 	return false
// }

func (i *Interpreter) checkNumberOperands(left interface{}, right interface{}) {
	if _, ok := left.(float64); ok {
		if _, ok = right.(float64); ok {
			return
		}
	}
	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
}

func (i *Interpreter) execute(stmt tree.Stmt) interface{} {
	return stmt.Accept(i)
}

// func (i *Interpreter) stringify(value interface{}) string {
// 	if value == nil {
// 		return "nil"
// 	}
// 	return fmt.Sprint(value)
// }

func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}
