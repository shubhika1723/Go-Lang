package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // expecting 2 goroutines

	go func() {
		fmt.Println("Task 1 done")
		wg.Done()
	}()

	go func() {
		fmt.Println("Task 2 done")
		wg.Done()
	}()

	wg.Wait() // wait for both goroutines
	fmt.Println("All goroutines finished")
}
