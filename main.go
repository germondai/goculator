package main

import "fmt"

func main() {
	var a, b int = 1, 2
	var result = add(a, b)
	fmt.Println("The result is", result)
}

func add(a, b int) int {
	return a + b
}
