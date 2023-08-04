package main

import (
	"fmt"
	"time"
)

// Goroutines allow functions to run concurrently
// can also run functions literals/closures
// Go will automatically select parallel or asynchronous execution
// new goroutines can be created with the go keyword

func count(amount int) {
	counter := 0
	for i := 1; i <= amount; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
	wait := func(ms time.Duration) {
		time.Sleep(ms * time.Millisecond)
		counter += 1
	}
	fmt.Println("Launching goroutines")
	go wait(100)
	go wait(900)
	go wait(1000)
	fmt.Println("Launched.  Counter=", counter)
	time.Sleep(1100 * time.Millisecond)
	fmt.Println("Waited 1100ms. Counter =", counter)
}

func main() {
	go count(5)
	fmt.Println("Wait for Goroutines")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("end program")
}
