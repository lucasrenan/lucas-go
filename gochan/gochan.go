package main

import (
	"fmt"
	"runtime"
	"time"
)

func printHi(name string) {
	fmt.Printf("Hi %s\n", name)
}

type Job struct {
	ID       string
	Producer int
}

func consumer(id int, ch chan *Job) {
	for job := range ch {
		fmt.Printf("consumer: %d <- %+v\n", id, job)
		time.Sleep(100 * time.Millisecond) // simulate work
	}
}

func producer(id int, ch chan *Job) {
	msgID := 0
	for {
		msgID++
		job := &Job{
			ID:       fmt.Sprintf("%d", msgID),
			Producer: id,
		}
		ch <- job
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Printf("Using threads: %d\n", runtime.GOMAXPROCS(0))

	// var wg sync.WaitGroup
	// for _, name := range []string{"Rick", "Summer", "Morty"} {
	// 	wg.Add(1)
	// 	go func(n string) {
	// 		defer wg.Done()
	// 		printHi(n)
	// 	}(name)
	// }

	// wg.Wait()

	msgCh := make(chan *Job)
	go producer(1, msgCh)
	go consumer(1, msgCh)
	go consumer(2, msgCh)

	time.Sleep(time.Second)
}
