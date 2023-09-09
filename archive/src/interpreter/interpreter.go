package interpreter

import (
	"fmt"
	"io"
	"os"
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
		result = append(result, interpreter.Execute(statement))
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
	if err := recover(); err != nil {
		// hide stacktrace
		// debug.SetTraceback("none")
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
	// i.Globals.Define("join", Join{})
	i.Globals.Define("range", Range{})
	i.Globals.Define("get", Get{})
}

func (i *Interpreter) VisitBinaryExpr(expr tree.Binary) interface{} {
	left := i.Evaluate(expr.Left)
	right := i.Evaluate(expr.Right)

	var returnValue interface{}
	switch expr.Operator.Type {
	case token.PLUS:
		returnValue = add(left, right)
	case token.MINUS:
		returnValue = subtract(left, right)
	case token.SLASH:
		returnValue = divide(left, right)
	case token.STAR:
		returnValue = multiply(left, right)
	case token.GREATER:
		returnValue = greaterThan(left, right)
	case token.GREATER_EQUAL:
		returnValue = greaterThanOrEqual(left, right)
	case token.LESS:
		returnValue = lessThan(left, right)
	case token.LESS_EQUAL:
		returnValue = lessThanOrEqual(left, right)
	case token.BANG_EQUAL:
		returnValue = notEqual(left, right)
	case token.DOUBLE_EQUAL:
		returnValue = equal(left, right)
	}
	return returnValue
}

func (i *Interpreter) VisitLiteralExpr(expr tree.Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitGroupingExpr(expr tree.Grouping) interface{} {
	return i.Evaluate(expr.Expression)
}

func (i *Interpreter) VisitUnaryExpr(expr tree.Unary) interface{} {
	right := i.Evaluate(expr)
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

// func (i *Interpreter) VisitCallStmt(stmt tree.CallStmt) interface{} {
// 	// do we need to reexecute a statement if it's a pointer to an already executed statement?
// 	callee := i.Evaluate(stmt.Callee)
// 	var arguments []interface{}
// 	for _, arg := range stmt.Arguments {
// 		arguments = append(arguments, i.Evaluate(arg))
// 	}
// 	return callee.(Callable).Call(i, , arguments)
// }

func (i *Interpreter) callFunction(stmt tree.CallStmt, caller interface{}) interface{} {
	callee := i.Evaluate(stmt.Callee)
	var arguments []interface{}
	for _, arg := range stmt.Arguments {
		arguments = append(arguments, i.Evaluate(arg))
	}
	return callee.(Callable).Call(i, caller, arguments)
}

func (i *Interpreter) VisitReturnStmt(stmt tree.ReturnStmt) interface{} {
	var value interface{}
	if stmt.Value != nil {
		value = i.Execute(stmt.Value)
	}
	return Return{
		Value: value,
	}
}

func (i *Interpreter) VisitExpressionStmt(stmt tree.ExpressionStmt) interface{} {
	return i.Evaluate(stmt.Expr)
}

func (i *Interpreter) VisitVariableStmt(stmt tree.VariableStmt) interface{} {
	var value interface{}
	if stmt.Initializer != nil {
		value = i.Execute(stmt.Initializer)
	}
	i.Environment.Define(stmt.Name.Text, value)
	if _, ok := stmt.Initializer.(tree.FuncStmt); ok {
		return "<closure>"
	}
	if value == nil {
		return "<variable>"
	}
	return value
}

func (i *Interpreter) VisitFunctionStmt(stmt tree.FuncStmt) interface{} {
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
	return i.executeList(stmt.Items, stmt.Functions, environment)
}

func (i *Interpreter) executeList(items []tree.Stmt, functions []tree.CallStmt, environment *env.Environment) interface{} {
	previous := i.Environment
	i.Environment = environment
	values := NewOrderedMap()
	for index, item := range items {
		value := i.Execute(item)

		if variableStmt, ok := item.(tree.VariableStmt); ok {
			values.Set(variableStmt.Name.Literal, value)
		} else {
			values.Set(index, value)
		}
	}

	var returnVal interface{}
	if len(functions) == 0 {
		returnVal = values
	} else {
		// potential performance bottleneck
		for index, function := range functions {
			// pipe values into first function call
			if index == 0 {
				returnVal = i.callFunction(function, values)
			} else {
				returnVal = i.callFunction(function, returnVal)
			}
		}
	}
	i.Environment = previous
	return returnVal
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

func (i *Interpreter) Evaluate(expr tree.Expr) interface{} {
	return expr.Accept(i)
}

func (i *Interpreter) Execute(stmt tree.Stmt) interface{} {
	return stmt.Accept(i)
}
