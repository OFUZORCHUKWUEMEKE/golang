package main

import "fmt"

func main() {
	two()
	fmt.Println("Another One")
	two()
	addthree(1, 2, 3)
	fmt.Printf("whats popping")
	addUp(1, 19)
}

func two() (int, int) {
	return 5, 5
}

func addthree(a, b, c int) int {
	return a + b + c
}

func addUp(a, b int) int {
	return a + b
}
