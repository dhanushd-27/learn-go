package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	fileVal := "Hello, I am Learning about handling files in golang"

	file, _ := os.Create("./hi.txt")

	len, _ := io.WriteString(file, fileVal)

	defer file.Close()

	fmt.Println(len)

	// ioutil is used to handle files
	readFile()
}

func readFile() {
	content, _ := os.ReadFile("./hi.txt")
	fmt.Println(string(content))
}