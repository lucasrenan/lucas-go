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

	// bug
	for _, name := range []string{"1", "2", "3"} {
		defer func() {
			fmt.Printf("defer1 of %s\n", name)
		}()
	}

	// fix
	for _, name := range []string{"1", "2", "3"} {
		defer func(n string) {
			fmt.Printf("defer2 of %s\n", n)
		}(name)
	}
}
