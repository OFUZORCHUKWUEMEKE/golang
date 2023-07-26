package main

import "fmt"

func main() {
	result := addTwo(2, 4)
	sub := subTwo(100, 10)
	fmt.Println(result)
	fmt.Println(sub)
}

// Addition
func addTwo(a, b int) int {
	return a + b
}

// Substraction
func subTwo(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
