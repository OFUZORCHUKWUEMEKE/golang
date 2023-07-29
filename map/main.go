package main

import "fmt"

// myMap1 := make(map[string]int)

// Maps store data in key-value pairs

func main() {
	shoppingList := make(map[string]int)
	// len(shoppingList)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted", shoppingList)

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal")
	if !found {
		fmt.Println(found)

	} else {
		fmt.Println("yup", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println(totalItems)

}
