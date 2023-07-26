package main

import "fmt"

func price() int {
	return 100
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p > 2:
		fmt.Println("Cheap Item")
	case p < 10:
		fmt.Println("Cheap But Expensive")
	case p > 5:
		fmt.Println("Thats whats up")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Ticket is An Economy Ticket")
	case Business:
		fmt.Println("Ticket is a Business Ticket")
	case FirstClass:
		fmt.Print("Ticket is a FirstClass Ticket")
	default:
		fmt.Println("Ticket is a ", ticket)
	}
}
