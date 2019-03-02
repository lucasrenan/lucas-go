package main

import "fmt"

func hello(ch chan int) {
	fmt.Println("hello outside")
	ch <- 2
}

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("hello inline")
		ch <- 1
	}()

	go hello(ch)

	fmt.Println("hello main")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
