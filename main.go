package main

import (
	"fmt"
)

func sum(start int, end int) int {
	total := 0
	for i := total; i <= end; i++ {
		total += 1
	}

	return total
}

func fib(n int) int {
	a, b := 1, 1
	for n > 0 {
		a, b = b, a+b
		n--
	}

	return a
}

func main() {
	fmt.Println("hello world")
	fmt.Println("sum(1, 10) = %d\n", sum(1, 10))
	fmt.Println("fib(10) = %d\n", fib(10))

	s := "Hello"
	for i := range s {
		fmt.Println(i)
	}

	for i, c := range s {
		fmt.Println("%d: %c\n", i, c)
	}

	for i, c := range s[1:3] {
		fmt.Println("%d: %c\n", i, c)
	}

	for _, c := range s {
		fmt.Println(c)
	}
}
