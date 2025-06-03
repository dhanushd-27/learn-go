package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("January 2, 2006"))

	date := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(date)

	fmt.Println(date.Format("2006-01-02"))
}