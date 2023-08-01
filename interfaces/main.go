package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook that thing")
}

func (s salad) PrepareDish() {
	fmt.Println("Cook salad well well")
	fmt.Println("Add Dressing")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare Dishes")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Println("--dISH", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken wings"), salad("Salad tujnkwkjnjfnn")}
	prepareDishes(dishes)
}
