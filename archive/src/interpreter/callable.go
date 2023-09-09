package interpreter

type Callable interface {
	Call(interpreter *Interpreter, caller interface{}, arguments []interface{}) interface{}
}
