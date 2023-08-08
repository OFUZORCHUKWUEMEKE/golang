package main

import "fmt"

// Channels offer one-way commmunication
// Conceptually the same as a two-ended pipe
// -Write data in one end and read out the other
//   This is also called sending ande= recieving
// Utilizing channels enables goroutines to communicate
//  - Can send / recieve messages or computational results
//  Channels can be buffered or unbuffered

func main() {
	channel := make(chan int)
	go func() { channel <- 1 }()
	go func() { channel <- 2 }()
	go func() { channel <- 3 }()

	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)

}
