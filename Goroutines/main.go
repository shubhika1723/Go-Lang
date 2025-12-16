package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()

	fmt.Println("Main function done")
	time.Sleep(time.Second) // make sure goroutine gets time to run
}

func sayHello() {
	fmt.Println("Hello from goroutine!")
}
