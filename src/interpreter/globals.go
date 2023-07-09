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

func (p Print) Call(i *Interpreter, caller interface{}, _ []interface{}) interface{} {
	fmt.Print(printValue(caller))
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

func (j Join) Call(i *Interpreter, caller interface{}, arguments []interface{}) interface{} {
	if len(arguments) > 1 {
		panic(errors.NewRuntimeError("maximum of 1 arguments allowed for method 'join'"))
	}
	var delimiter string
	if len(arguments) == 1 {
		delimiter = arguments[0].(string)
	}
	return joinValues(caller, delimiter)
}

func (j Join) String() string {
	return "<native fn>"
}

type Range struct {
}

func (r Range) Arity() int {
	return 0
}

func (r Range) Call(i *Interpreter, caller interface{}, arguments []interface{}) interface{} {
	if len(arguments) > 2 {
		panic(errors.NewRuntimeError("maximum of 2 arguments allowed for method 'range'"))
	}
	return makeRange(arguments[0].(float64), arguments[1].(float64))
}

type Get struct {
}

func (g Get) Arity() int {
	return 0
}

func (g Get) Call(i *Interpreter, caller interface{}, arguments []interface{}) interface{} {
	if len(arguments) > 1 {
		panic(errors.NewRuntimeError("maximum of 1 argument allowed for method 'get'"))
	}

	selector := arguments[0]

	switch t := caller.(type) {
	case []interface{}:
		if intSelector, ok := selector.(float64); ok {
			return t[int(intSelector)]
		}

		// if strSelector, ok := selector.(string); ok {

		// }
	}

	return nil
}

func makeRange(min, max float64) []interface{} {
	a := make([]interface{}, int(max-min+1))
	for i := range a {
		a[i] = min + float64(i)
	}
	return a
}

func (r Range) String() string {
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

func joinValues(values interface{}, delimiter string) string {
	var args []interface{}
	var format string
	if values, ok := values.([]interface{}); ok {
		for index, item := range values {
			if itemValues, ok := item.([]interface{}); ok {
				format += "%v"
				args = append(args, joinValues(itemValues, delimiter))

				if len(delimiter) > 0 && index < len(values)-1 {
					format += "%s"
					args = append(args, delimiter)
				}
			} else {
				format += "%v"
				args = append(args, item)

				if len(delimiter) > 0 && index < len(values)-1 {
					format += "%s"
					args = append(args, delimiter)
				}
			}
		}
	}
	return fmt.Sprintf(format, args...)
}
