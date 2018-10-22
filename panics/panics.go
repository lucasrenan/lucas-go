package main

import (
	"fmt"
)

func div(a, b int) int {
	fmt.Printf("div: %d/%d\n", a, b)
	return a / b
}

func safeDiv(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("safeDiv: %s\n", err)
		}
	}()

	return div(a, b)
}

func main() {
	r1 := div(2, 3)
	fmt.Println(r1)
	// r2 := div(7, 0)
	// fmt.Println(r2)

	safeDiv(7, 0)
}
