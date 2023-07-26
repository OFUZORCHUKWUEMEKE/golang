package main

import "fmt"

func main() {
	var myName = "Ofuzor Chukwuemeke"

	const (
		Economy    = 1
		Business   = 2
		FirstClass = 3
	)

	username := "admin"
	fmt.Println("My name is", myName)

	fmt.Println(username)

	// type name string

	name := "Ofuzor Chukwuemeke"
	part1, part2 := 1, 2

	fmt.Println(part1, part2, name)
}
