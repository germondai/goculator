package main

import (
	"errors"
	"fmt"

	"./operations"
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
	operations := map[Operation]func(float64, float64) (float64, error){
		Add:      operations.Add,
		Subtract: operations.Subtract,
		Multiply: operations.Multiply,
		Divide:   operations.Divide,
	}

	opFn, exists := operations[operation]

	if exists {
		return opFn(a, b)
	}

	return 0, errors.New("unsupported operation")
}
