package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	go func() {
		defer close(ch)
		for i := range 10 {
			fmt.Println("sent data ", i)
			ch <- i
		}
	}()

	time.Sleep(time.Second * 5)
	val, ok := <- ch
	fmt.Println(val, ok)
	for value := range ch {
		println(value)
	}
}
