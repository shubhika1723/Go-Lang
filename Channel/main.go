package main

import "fmt"

func main() {
	ch := make(chan string) // create channel

	go func() {
		ch <- "Message from goroutine" // send to channel
	}()

	msg := <-ch // receive from channel
	fmt.Println(msg)
}
