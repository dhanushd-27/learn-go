package main

import (
	"fmt"
)

func main() {
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}

	// While loop (using for)
	j := 0
	for j < 5 {
		fmt.Println("While loop iteration:", j)
		j++
	}

	// Infinite loop (break with condition)
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println("Infinite loop iteration:", k)
		k++
	}

	var gender = "male"

	switch gender {
	case "male":
		println("You are a male")
	case "female":
		println("You are a female")
	default:
		println("You are neither male nor female")

	}
}