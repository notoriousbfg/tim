package interpreter

import (
	"fmt"
	"tim/errors"
)

type Print struct {
}

func (p Print) Arity() int {
	return 0
}

func (p Print) Call(i *Interpreter, _ []interface{}) interface{} {
	fmt.Println(printValue(i.PrevValue))
	return nil
}

func (p Print) String() string {
	return "<native fn>"
}

type Join struct {
}

func (j Join) Arity() int {
	return 0
}

func (j Join) Call(i *Interpreter, arguments []interface{}) interface{} {
	if len(arguments) > 1 {
		panic(errors.NewRuntimeError("maximum of 1 arguments allowed for method 'join'"))
	}
	var delimiter string
	if len(arguments) == 1 {
		delimiter = arguments[0].(string)
	}
	var args []interface{}
	var format string
	if values, ok := i.PrevValue.([]interface{}); ok {
		for index, item := range values {
			format += "%v"
			args = append(args, item)

			if len(delimiter) > 0 && index < len(values)-1 {
				format += "%s"
				args = append(args, delimiter)
			}
		}
	}
	return fmt.Sprintf(format, args...)
}

func (j Join) String() string {
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
