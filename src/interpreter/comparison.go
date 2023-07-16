package interpreter

import (
	"fmt"
	"reflect"
	"tim/errors"
)

func subtract(left, right interface{}) interface{} {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if (leftType == rightType) && leftType.String() == "int" {
		return left.(int) - right.(int)
	}

	if leftType.String() == "float64" && rightType.String() == "float64" {
		leftFloat, _ := interfaceToFloat(left)
		rightFloat, _ := interfaceToFloat(right)

		return leftFloat - rightFloat
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
