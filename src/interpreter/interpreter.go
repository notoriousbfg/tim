package interpreter

import (
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"runtime/debug"
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
		Globals: &env.Environment{
			Enclosing: nil,
			Values:    make(map[string]interface{}),
		},
	}
	interpreter.defineGlobals()
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
	Globals     *env.Environment
	stdErr      io.Writer
}

func (i *Interpreter) printToStdErr() {
	// hide stacktrace
	debug.SetTraceback("none")
	if err := recover(); err != nil {
		if e, ok := err.(errors.RuntimeError); ok {
			_, _ = i.stdErr.Write([]byte(e.Error() + "\n"))
			os.Exit(70)
		} else {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

func (i *Interpreter) defineGlobals() {
	i.Globals.Define("print", Print{})
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
	return i.lookupVariable(expr.Name)
}

func (i *Interpreter) lookupVariable(name token.Token) interface{} {
	global, err := i.Globals.Get(name)
	if err != nil {
		if _, ok := err.(*errors.RuntimeError); !ok {
			panic(err)
		}
	}
	if global != nil {
		return global
	}

	val, err := i.Environment.Get(name)
	if err != nil {
		panic(err)
	}
	return val
}

func (i *Interpreter) VisitCallStmt(stmt tree.CallStmt) interface{} {
	// do we need to reexecute a statement if it's a pointer to an already executed statement?
	callee := i.evaluate(stmt.Callee)
	var arguments []interface{}
	for _, arg := range stmt.Arguments {
		arguments = append(arguments, i.evaluate(arg))
	}
	return callee.(Callable).Call(i, stmt.Initialiser, arguments)
}

func (i *Interpreter) VisitExpressionStmt(stmt tree.ExpressionStmt) interface{} {
	return i.evaluate(stmt.Expr)
}

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

func (i *Interpreter) Print(stmt tree.Stmt) {
	fmt.Println(i.printValue(i.execute(stmt)))
}

func (i *Interpreter) printValue(value interface{}) string {
	var output string
	switch t := value.(type) {
	case []interface{}:
		output = "("
		for index, item := range t {
			output += i.printValue(item)
			if index < len(t)-1 {
				output += ", "
			}
		}
		output += ")"
	case string:
		output = fmt.Sprintf("\"%s\"", value.(string))
	default:
		output = fmt.Sprintf("%v", value)
	}
	return output
}

func (i *Interpreter) evaluate(expr tree.Expr) interface{} {
	return expr.Accept(i)
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

// func (i *Interpreter) Print(initialiser tree.ListStmt) {
// 	var values []interface{}
// 	for _, item := range initialiser.Items {
// 		values = append(values, i.execute(item))
// 	}

// }

func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}
