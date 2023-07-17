package interpreter

import (
	"fmt"
	"reflect"
	"tim/errors"
)

func subtract(left, right interface{}) interface{} {
	if isNaN(left) || isNaN(right) {
		panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
	}

	if isInt(left) && isInt(right) {
		return left.(int) - right.(int)
	}

	// if either are floats
	if isFloat(left) || isFloat(right) {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat - rightFloat
	}

	return nil
}

func add(left, right interface{}) interface{} {
	if isString(left) || isString(right) {
		return fmt.Sprint(left, right)
	}

	// if both are ints
	if isInt(left, right) {
		return left.(int) + right.(int)
	}

	leftFloat, _ := toFloat(left)
	rightFloat, _ := toFloat(right)

	return leftFloat + rightFloat
}

func divide(left, right interface{}) interface{} {
	if isZero(left) || isZero(right) {
		panic(errors.NewRuntimeError(errors.DivisionByZero))
	}

	if isNaN(left) || isNaN(right) {
		panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
	}

	if isInt(left, right) {
		return left.(int) / right.(int)
	}

	if isFloat(left) || isFloat(right) {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat / rightFloat
	}

	return nil
}

func multiply(left, right interface{}) interface{} {
	if isNaN(left) || isNaN(right) {
		panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
	}

	if isInt(left, right) {
		return left.(int) * right.(int)
	}

	if isFloat(left) || isFloat(right) {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat * rightFloat
	}

	return nil
}

func greaterThan(left, right interface{}) bool {
	if isNaN(left) || isNaN(right) {
		panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
	}

	if isInt(left, right) {
		return left.(int) > right.(int)
	}

	if isFloat(left) || isFloat(right) {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat > rightFloat
	}

	return false
}

func greaterThanOrEqual(left, right interface{}) bool {
	if isNaN(left) || isNaN(right) {
		panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
	}

	if isInt(left, right) {
		return left.(int) >= right.(int)
	}

	if isFloat(left) || isFloat(right) {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat >= rightFloat
	}

	return false
}

func lessThan(left, right interface{}) bool {
	if isNaN(left) || isNaN(right) {
		panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
	}

	if isInt(left, right) {
		return left.(int) < right.(int)
	}

	if isFloat(left) || isFloat(right) {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat < rightFloat
	}

	return false
}

func lessThanOrEqual(left, right interface{}) bool {
	if isNaN(left) || isNaN(right) {
		panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
	}

	if isInt(left, right) {
		return left.(int) <= right.(int)
	}

	if isFloat(left) || isFloat(right) {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat <= rightFloat
	}

	return false
}

func equal(left, right interface{}) bool {
	if isNumber(left, right) {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat == rightFloat
	}

	return left == right
}

func notEqual(left, right interface{}) bool {
	return !equal(left, right)
}

func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}

func toFloat(val interface{}) (float64, error) {
	switch t := val.(type) {
	case int:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return float64(t), nil
	default:
		return 0, fmt.Errorf("could not convert value of type '%t' to float64", val)
	}
}

func isInt(args ...interface{}) bool {
	if len(args) == 1 {
		if _, ok := args[0].(int); !ok {
			return false
		}
	}
	for _, arg := range args {
		if _, ok := arg.(int); !ok {
			return false
		}
	}
	return true
}

func isString(args ...interface{}) bool {
	if len(args) == 1 {
		if _, ok := args[0].(string); !ok {
			return false
		}
	}
	for _, arg := range args {
		if _, ok := arg.(string); !ok {
			return false
		}
	}
	return true
}

func isFloat(args ...interface{}) bool {
	if len(args) == 1 {
		if _, ok := args[0].(float64); !ok {
			return false
		}
	}
	for _, arg := range args {
		if _, ok := arg.(float64); !ok {
			return false
		}
	}
	return true
}

func isNumber(args ...interface{}) bool {
	if len(args) == 1 {
		if isFloat(args[0]) || isInt(args[0]) {
			return true
		}
	}
	for _, arg := range args {
		if !isFloat(arg) && !isInt(arg) {
			return false
		}
	}
	return true
}

func isNaN(args ...interface{}) bool {
	if len(args) == 1 {
		if !isFloat(args[0]) && !isInt(args[0]) {
			return true
		}
	}
	for _, arg := range args {
		if !isFloat(arg) && !isInt(arg) {
			return true
		}
	}
	return false
}
