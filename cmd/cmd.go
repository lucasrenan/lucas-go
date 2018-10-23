package main

import (
	"flag"
	"fmt"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()

	fmt.Printf("port = %d\n", port)
	// fmt.Printf("args = %")
}
