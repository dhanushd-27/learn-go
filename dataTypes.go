package main

import (
	"log"
)

func dataTypesDemo() {
	var number int = 42;
	var name string = "John Doe";
	var isActive bool = true;
	var pi float64 = 3.14;
	var numbers []int = []int{1, 2, 3, 4, 5};

	log.Println("Integer:", number)
	log.Println("String:", name)
	log.Println("Boolean:", isActive)
	log.Println("Float:", pi)
	log.Println("Slice:", numbers)
}

func main() {
	dataTypesDemo()
	log.Println("Data types demo completed.")
}