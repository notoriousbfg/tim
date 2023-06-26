package errors

const (
	OperandsMustBeNumber = "operands must be number"
	DivisionByZero       = "division by zero"
)

type RuntimeError struct {
	message string
}

func (r RuntimeError) Error() string {
	return r.message
}

func NewRuntimeError(msg string) *RuntimeError {
	return &RuntimeError{message: msg}
}
