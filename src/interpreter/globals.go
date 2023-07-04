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

func (p Print) Call(i *Interpreter, initialiser tree.ListStmt, _ []interface{}) interface{} {
	fmt.Println(printValue(i.Execute(initialiser)))
	return nil
}

func (p Print) String() string {
	return "<native fn>"
}

type Join struct {
}

func (p Join) Arity() int {
	return 0
}

func (p Join) Call(i *Interpreter, initialiser tree.ListStmt, _ []interface{}) interface{} {
	return ""
}

func (p Print) Join() string {
	return "<native fn>"
}

func printValue(value interface{}) string {
	var output string
	switch t := value.(type) {
	case []interface{}:
		output = "("
		for index, item := range t {
			output += printValue(item)
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
