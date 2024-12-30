package main

import (
	"errors"
	"fmt"
)

type Operation string

const (
	Add      Operation = "+"
	Subtract Operation = "-"
	Multiply Operation = "*"
	Divide   Operation = "/"
)

func main() {
	var a, b float64 = 1, 2
	var op Operation = Add

	result, err := calculate(a, b, op)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("The result is", result)
}

func calculate(a, b float64, operation Operation) (float64, error) {
	switch operation {
	case Add:
		return add(a, b), nil
	case Subtract:
		return subtract(a, b), nil
	case Multiply:
		return multiply(a, b), nil
	case Divide:
		return divide(a, b)
	default:
		return 0, errors.New("unsupported operation")
	}
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}
