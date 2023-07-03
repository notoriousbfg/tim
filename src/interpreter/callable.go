package interpreter

import "tim/tree"

type Callable interface {
	Call(interpreter *Interpreter, initialiser tree.Stmt, arguments []interface{}) interface{}
}
