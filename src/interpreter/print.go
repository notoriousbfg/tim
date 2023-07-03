package interpreter

import (
	"fmt"
	"tim/tree"
)

type Print struct {
}

func (p Print) Arity() int {
	return 0
}

func (p Print) Call(_ *Interpreter, initialiser tree.Stmt, _ []interface{}) interface{} {
	var printVal interface{}
	switch stmt := initialiser.(type) {
	case tree.ListStmt:
		printVal = stmt.Items // todo
	default:
		printVal = nil
	}
	fmt.Print(printVal)
	return nil
}

func (p Print) String() string {
	return "<native fn>"
}
