package tree

type Callable interface {
	Call(interpreter *Interpreter, arguments []interface{}) interface{}
}
