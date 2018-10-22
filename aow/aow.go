package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("aow.txt")
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	counts := map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			counts[strings.ToLower(word)]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error: can't scan - %s", err)
	}

	maxWord, maxCount := "", 0
	for word, count := range counts {
		if count > maxCount {
			maxWord, maxCount = word, count
		}
	}

	fmt.Printf("%s (%d)\n", maxWord, maxCount)
}
