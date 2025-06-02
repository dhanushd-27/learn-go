package main

import (
	"fmt"
	"l1/foo"
)

func initt() {
	fmt.Println("init")
	const a int = 10

	fmt.Println(a);
}

func main() {
	foo.Foo()
	initt()
}
