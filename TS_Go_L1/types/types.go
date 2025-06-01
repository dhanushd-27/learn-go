package types

import "fmt"

func Types() {
	var a int = 10
	var b int = 20
	var c int = a + b
	fmt.Println(c)

	var d float32 = 10.0
	var e float32 = 20.0
	var f float32 = d + e
	fmt.Println(f)

	var g string = "Hello"
	var h string = "World"
	var i string = g + h
	fmt.Println(i)

	var j bool = true
	var k bool = false
	var l bool = j && k
	fmt.Println(l)
}
