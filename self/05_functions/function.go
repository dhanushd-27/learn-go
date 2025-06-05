package main

import (
	"fmt"
)

func main() {
	var feb func(int) int

	feb = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		return feb(n-1) + feb(n-2)
	}
	fmt.Println(feb(10))
}