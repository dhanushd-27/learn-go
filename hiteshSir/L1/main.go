package main

import (
	"fmt"
	"l1/foo"
)

func initt() {
	fmt.Println("init")
	const a int = 10

	fmt.Println(a);
	fmt.Printf("Var a is of type %T\n", a)
}

func main() {
	foo.Foo()
	initt()
}
