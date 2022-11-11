package tree

type RuntimeError struct {
	message string
}

func (r *RuntimeError) Error() string {
	return r.message
}

func NewRuntimeError(msg string) *RuntimeError { // does this suck?
	return &RuntimeError{message: msg}
}
