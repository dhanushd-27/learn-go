package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a number: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	num2 := num * 2

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("You entered: ", num2)
	}
}