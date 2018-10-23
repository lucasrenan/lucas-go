package main

import (
	"expvar"
	"fmt"
)

var (
	nCalls  = expvar.NewInt("num_calls")
	nErrors = expvar.NewInt("num_errors")
)

func handler(i int) {
	nCalls.Add(1) // this is goroutine safe
	if i > 7 {
		fmt.Printf("error: i (%d) > 7 \n", i)
		nErrors.Add(1)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		handler(i)
	}

	// expvar.Do(func())
}
