package main

import "fmt"

type Product struct {
	price int
	name  string
}

func exercise(list [4]Product) {
	var cost, totalItem int
	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItem += 1
		}
	}

	fmt.Println("Last item on the list is", list[totalItem-1])
}
