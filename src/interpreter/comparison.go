package interpreter

import (
	"fmt"
	"reflect"
	"tim/errors"
)

func subtract(left, right interface{}) interface{} {
	if areInts, ints := isInt(left, right); areInts {
		return ints[0] - ints[1]
	}

	if areFloats, floats := isFloat(left, right); areFloats {
		return floats[0] - floats[1]
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
}

func add(left, right interface{}) interface{} {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if leftType.String() == "string" || rightType.String() == "string" {
		return fmt.Sprint(left, right)
	}

	if (leftType == rightType) && leftType.String() == "int" {
		return left.(int) + right.(int)
	}

	leftFloat, _ := interfaceToFloat(left)
	rightFloat, _ := interfaceToFloat(right)

	return leftFloat + rightFloat
}

func divide(left, right interface{}) interface{} {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if (leftType == rightType) && leftType.String() == "int" {
		return left.(int) / right.(int)
	}

	if leftType.String() == "float64" && rightType.String() == "float64" {
		leftFloat, _ := interfaceToFloat(left)
		rightFloat, _ := interfaceToFloat(right)

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
		leftFloat, _ := interfaceToFloat(left)
		rightFloat, _ := interfaceToFloat(right)

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
		leftFloat, _ := interfaceToFloat(left)
		rightFloat, _ := interfaceToFloat(right)

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
		leftFloat, _ := interfaceToFloat(left)
		rightFloat, _ := interfaceToFloat(right)

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
		leftFloat, _ := interfaceToFloat(left)
		rightFloat, _ := interfaceToFloat(right)

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
		leftFloat, _ := interfaceToFloat(left)
		rightFloat, _ := interfaceToFloat(right)

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
		leftFloat, _ := interfaceToFloat(left)
		rightFloat, _ := interfaceToFloat(right)

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
		leftFloat, _ := interfaceToFloat(left)
		rightFloat, _ := interfaceToFloat(right)

		return leftFloat != rightFloat
	}

	panic(errors.NewRuntimeError(errors.OperandsMustBeNumber))
}

func interfaceToFloat(val interface{}) (float64, error) {
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

func isInt(args ...interface{}) (bool, []int) {
	ints := make([]int, 0)

	for _, arg := range args {
		if thisInt, ok := arg.(int); ok {
			ints = append(ints, thisInt)
		} else {
			return false, []int{}
		}
	}

	return true, ints
}

func isFloat(args ...interface{}) (bool, []float64) {
	floats := make([]float64, 0)

	for _, arg := range args {
		if thisFloat, ok := arg.(float64); ok {
			floats = append(floats, thisFloat)
		} else {
			return false, []float64{}
		}
	}

	return true, floats
}
