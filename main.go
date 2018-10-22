package main

import (
	"fmt"
	"time"
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
	fmt.Printf("sum(1, 10) = %d\n", sum(1, 10))
	fmt.Printf("fib(10) = %d\n", fib(10))

	s := "Hello"
	for i := range s {
		fmt.Println(i)
	}

	for i, c := range s {
		fmt.Printf("%d: %c\n", i, c)
	}

	for i, c := range s[1:3] {
		fmt.Printf("%d: %c\n", i, c)
	}

	for _, c := range s {
		fmt.Println(c)
	}

	heroes := map[string]string{
		"Superman":     "Clark Kent",
		"Batman":       "Bruce Wayne",
		"Wonder Woman": "Diana Prince",
	}

	for h := range heroes {
		fmt.Println(h)
	}

	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	start := time.Now()
	end := start.Add(3 * time.Second)
	for range time.Tick(time.Second) {
		fmt.Println("tick")
		if time.Now().After(end) {
			break
		}
	}
}
