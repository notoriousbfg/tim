package interpreter

type Callable interface {
	Call(interpreter *Interpreter, arguments []interface{}) interface{}
}
