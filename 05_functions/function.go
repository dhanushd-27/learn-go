package main

import (
	"fmt"
)

// Function to add two integers
func add(a int, b int) int {
	return a + b
}

// Anonymous function to print a message
func addAno(a, b int) (sub int, sum int) {
	sub = a - b
	sum = a + b
	return
}

func main() {
	// Call the add function
	result := add(5, 3)
	fmt.Println("The sum is:", result)

	// Call the add function with different arguments
	result = add(10, 20)
	fmt.Println("The sum is:", result)

	// Call the anonymous function
	sub, sum := addAno(10, 5)
	fmt.Println("The sum is:", sum)
	fmt.Println("The sub is:", sub)
}