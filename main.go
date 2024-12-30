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
	fmt.Println("Welcome to the Enhanced Go Calculator!")
	fmt.Println("Supported operations: +, -, *, /")
	fmt.Println("----------------------------------")

	num1, err := getInput("Enter the first number: ")
	handleError(err)

	op, err := getOperation()
	handleError(err)

	num2, err := getInput("Enter the second number: ")
	handleError(err)

	result, err := calculate(num1, num2, op)
	handleError(err)

	fmt.Printf("Result: %.2f\n", result)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}

func getInput(prompt string) (float64, error) {
	var input float64
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0, errors.New("invalid input, please enter a number")
	}
	return input, nil
}

func getOperation() (Operation, error) {
	var op string
	fmt.Print("Enter an operator (+, -, *, /): ")
	_, err := fmt.Scanln(&op)
	if err != nil {
		return "", errors.New("invalid input, please enter a valid operator")
	}
	return Operation(op), nil
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
