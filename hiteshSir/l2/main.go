package main

import (
	"bufio"
	"fmt"
	"l2/struct"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// The comma-ok or comma-error pattern is common in Go
	// Here reader.ReadString returns two values:
	// 1. input: the string that was read
	// 2. err: any error that occurred during reading
	// The _ means we're ignoring the error value
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", input)

	r := structs.Rect{Width: 10, Height: 20}
	fmt.Println(r.Width, r.Height)

	structs.Val(r)

	fmt.Println(structs.Area(&r))
}
