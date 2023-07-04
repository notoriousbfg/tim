package interpreter

import (
	"tim/tree"
)

type Print struct {
}

func (p Print) Arity() int {
	return 0
}

func (p Print) Call(i *Interpreter, initialiser tree.ListStmt, _ []interface{}) interface{} {
	i.Print(initialiser)
	return nil
}

func (p Print) String() string {
	return "<native fn>"
}
