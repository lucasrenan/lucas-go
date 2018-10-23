package main

import (
	"fmt"
	"time"
)

func worker(val int, out chan int) {
	time.Sleep(time.Duration(val) * time.Millisecond) // simulate work
	out <- val * 2
}

func timeout() {
	ch := make(chan int)

	go worker(100, ch)
	select {
	case val, ok := <-ch:
		fmt.Printf("value: %d - ok: %v\n", val, ok)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout")
	}
}

func main() {
	timeout()
}
