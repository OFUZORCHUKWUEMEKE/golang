package main

import "fmt"

// type Emeke struct {
// 	name string
// 	age  int
// }

func main() {
	// fmt.print("Processing request again")
	dataName := []string{"name", "age", "dob", "account_number", "bank_name"}
	name := make(map[string]int)

	fmt.Println(dataName)

	num := 0

	name["Emeke"] = 10

	name["Ekenem"] = 10

	name["Ifeanyi"] = 10

	name["Chidinma"] = 10

	for _, amount := range name {
		num += amount
		fmt.Println(num)
	}
	fmt.Println("Total Number", num)

	fmt.Println(name)
	// fmt.Println(name)
}
