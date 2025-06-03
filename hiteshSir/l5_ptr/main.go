package main

import "fmt"

func main() {
	// pointers are used to store the address of a variable
	// & is used to get the address of a variable
	// * is used to get the value of a variable

	var a int = 10
	var b *int = &a

	fmt.Println(a, b, *b)

	*b = *b + 2

	fmt.Println(a, b, *b)
}