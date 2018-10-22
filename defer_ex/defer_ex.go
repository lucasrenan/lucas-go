package main

import (
	"fmt"
)

func cleanup(name string) {
	fmt.Printf("cleaning up %s\n", name)
}

func work() {
	resource := "A"
	defer cleanup(resource)
	defer cleanup("B")

	fmt.Println("working")
}

func main() {
	work()
}
