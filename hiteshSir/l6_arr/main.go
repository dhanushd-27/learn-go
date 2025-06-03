package main

import (
	"fmt"
)

func main() {
	var users []string = make([]string, 3)
	println(len(users))
	println(cap(users))
	fmt.Print(users[0] == "")

	var users2 = []string{"a", "b", "c"}
	users2 = append(users2, "d")

	users2 = append(users2[:1], users2[2:]...)
	fmt.Println(users2)
}
