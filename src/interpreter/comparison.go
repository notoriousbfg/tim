package interpreter

import (
	"fmt"
	"reflect"
	"tim/errors"
)

func subtract(left, right interface{}) interface{} {
	// if both ints
	// if areInts, ints := isInt(left, right); areInts {
	// 	return ints[0] - ints[1]
	// }

	if isInt(left) && isInt(right) {
		return left.(int) - right.(int)
	}

	// if either are floats
	if isFloat(left) || isFloat(right) {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat - rightFloat
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
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
	if isInt(left, right) {
		return left.(int) / right.(int)
	}

	if isFloat(left) || isFloat(right) {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat / rightFloat
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
}

func multiply(left, right interface{}) interface{} {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if (leftType == rightType) && leftType.String() == "int" {
		return left.(int) * right.(int)
	}

	if leftType.String() == "float64" && rightType.String() == "float64" {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat * rightFloat
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
}

func greaterThan(left, right interface{}) interface{} {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if (leftType == rightType) && leftType.String() == "int" {
		return left.(int) > right.(int)
	}

	if leftType.String() == "float64" && rightType.String() == "float64" {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat > rightFloat
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
}

func greaterThanOrEqual(left, right interface{}) interface{} {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if (leftType == rightType) && leftType.String() == "int" {
		return left.(int) >= right.(int)
	}

	if leftType.String() == "float64" && rightType.String() == "float64" {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat >= rightFloat
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
}

func lessThan(left, right interface{}) interface{} {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if (leftType == rightType) && leftType.String() == "int" {
		return left.(int) < right.(int)
	}

	if leftType.String() == "float64" && rightType.String() == "float64" {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat < rightFloat
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
}

func lessThanOrEqual(left, right interface{}) interface{} {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if (leftType == rightType) && leftType.String() == "int" {
		return left.(int) <= right.(int)
	}

	if leftType.String() == "float64" && rightType.String() == "float64" {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat <= rightFloat
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
}

func equal(left, right interface{}) interface{} {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if leftType.String() == "string" || rightType.String() == "string" {
		return left.(string) == right.(string)
	}

	if (leftType == rightType) && leftType.String() == "int" {
		return left.(int) == right.(int)
	}

	if leftType.String() == "float64" && rightType.String() == "float64" {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat == rightFloat
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
}

func notEqual(left, right interface{}) interface{} {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if leftType.String() == "string" || rightType.String() == "string" {
		return left.(string) != right.(string)
	}

	if (leftType == rightType) && leftType.String() == "int" {
		return left.(int) != right.(int)
	}

	if leftType.String() == "float64" && rightType.String() == "float64" {
		leftFloat, _ := toFloat(left)
		rightFloat, _ := toFloat(right)

		return leftFloat != rightFloat
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
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
	for _, arg := range args {
		if _, ok := arg.(int); !ok {
			return false
		}
	}

	return true
}

func isString(args ...interface{}) bool {
	for _, arg := range args {
		if _, ok := arg.(string); !ok {
			return false
		}
	}

	return true
}

func isFloat(args ...interface{}) bool {
	for _, arg := range args {
		if _, ok := arg.(float64); !ok {
			return false
		}
	}

	return true
}
