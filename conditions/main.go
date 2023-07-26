package main

import "fmt"

const (
	monday    = 0
	tuesday   = 1
	wednesday = 2
	thursday  = 3
	friday    = 4
	saturday  = 5
)

const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func weekday(day int) bool {
	return day <= 4
}

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func main() {
	today, role := tuesday, Guest

	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !weekday(today) {
		accessGranted()
	} else if role == Member && weekday(today) {
		accessGranted()
	} else {
		accessDenied()
	}

	accessGranted()
}
