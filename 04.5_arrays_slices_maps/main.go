package main

import (
	"fmt"
)

func arrayEx() {
	// Declare and initialize an array
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// Accessing elements
	fmt.Println("First element:", arr[0])
	fmt.Println("Second element:", arr[1])
}

func sliceEx() {
	// Declare and initialize a slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)

	// Accessing elements
	fmt.Println("First element:", slice[0])
	fmt.Println("Second element:", slice[1])

	// Append to a slice
	slice = append(slice, 6)
	fmt.Println("Slice after append:", slice)

	// Slicing a slice
	subSlice := slice[1:4]
	fmt.Println("Sub-slice:", subSlice)

	// Copy a slice by value
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	fmt.Println("Copied slice:", copiedSlice)

	// Copy a slice by reference
	referenceSlice := slice
	fmt.Println("Reference slice:", referenceSlice)

	// Modify the original slice
	referenceSlice[0] = 100
	fmt.Println("Modified reference slice:", referenceSlice)
	fmt.Println("Original slice after modification:", slice)
}

func mapEx() {
	// Declare and initialize a map
	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("Map:", myMap)

	// Accessing elements
	fmt.Println("Value for key 'one':", myMap["one"])

	// Adding a new key-value pair
	myMap["four"] = 4
	fmt.Println("Map after adding 'four':", myMap)

	// Deleting a key-value pair
	delete(myMap, "two")
	fmt.Println("Map after deleting 'two':", myMap)
}

func main() {
	arrayEx()
	sliceEx()
	mapEx()
}