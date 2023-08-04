package main

import "fmt"

// Function literals provide a way to define a function within a function
// Possible to assign function literals to variables
// They can be passed to a function as parameters
// More dynamic code
// Also known as closures or anonymous functions
// closures allow data to be encapsulated within

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(lhs, rhs int) int) int {
	fmt.Printf("Running a computation with %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 2 =", compute(2, 2, add))
}
