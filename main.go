package main

import (
	"errors"
	"fmt"

	"./operations"
	"./utils"
)

func main() {
	fmt.Println("Welcome to the Enhanced Go Calculator!")
	fmt.Println("Supported operations: +, -, *, /")
	fmt.Println("----------------------------------")

	num1, err := getInput("Enter the first number: ")
	utils.HandleError(err)

	op, err := getOperation()
	utils.HandleError(err)

	num2, err := getInput("Enter the second number: ")
	utils.HandleError(err)

	result, err := calculate(num1, num2, op)
	utils.HandleError(err)

	fmt.Printf("Result: %.2f\n", result)
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

func getOperation() (operations.Operation, error) {
	var op string
	fmt.Print("Enter an operator (+, -, *, /): ")
	_, err := fmt.Scanln(&op)
	if err != nil {
		return "", errors.New("invalid input, please enter a valid operator")
	}
	return operations.Operation(op), nil
}

func calculate(a, b float64, operation operations.Operation) (float64, error) {
	operations := map[operations.Operation]func(float64, float64) (float64, error){
		operations.AddOp:      operations.Add,
		operations.SubtractOp: operations.Subtract,
		operations.MultiplyOp: operations.Multiply,
		operations.DivideOp:   operations.Divide,
	}

	opFn, exists := operations[operation]

	if exists {
		return opFn(a, b)
	}

	return 0, errors.New("unsupported operation")
}
