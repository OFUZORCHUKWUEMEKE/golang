package main

import "fmt"

// Useful to run operations after functions complete
// The defer keyword can be used to execute code after a functions runs
func One() {
	fmt.Println("1")
}
func Two() {
	fmt.Println("2")
}
func Three() {
	fmt.Println("3")
}

func main() {
	fmt.Println("Begin")

	defer One()

	defer Two()

	defer Three()

	fmt.Println("Ending")
}
