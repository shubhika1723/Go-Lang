package main

import (
	"fmt"
	"sync"
)

func square(n int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Step: mark this goroutine as finished at the end

	// square the number
	sq := n * n

	// send the result into the channel
	results <- sq
}

func main() {
	// Step 1: input numbers
	numbers := []int{2, 3, 4, 5, 6}

	// Step 2: make a channel to receive squared results
	results := make(chan int)

	// Step 3: WaitGroup to track goroutines
	var wg sync.WaitGroup

	// Step 4: start a goroutine for each number
	for _, n := range numbers {
		wg.Add(1) // one more goroutine is starting
		go square(n, results, &wg)
	}

	// Step 5: close channel after all goroutines finish
	go func() {
		wg.Wait()      // wait until wg counter becomes 0
		close(results) // tell receiver: no more values coming
	}()

	// Step 6: read and print results from the channel
	fmt.Println("Squared results:")
	for r := range results {
		fmt.Println(r)
	}
}
