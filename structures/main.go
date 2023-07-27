package main

import "fmt"

type Emeke struct {
	name        string
	age         int
	bio         string
	description string
}

type Passenger struct {
	name        string
	age         int
	availabilty bool
}

type Bus struct {
	Frontseat Passenger
}

var sample struct {
	field string
	a, b  int
}

func main() {
	data := Emeke{
		name:        "Ofuzor Chukwuemeke",
		age:         22,
		bio:         "I love Building Softwares with Golang and Javascript",
		description: "I love writing code",
	}
	fmt.Println(data)

	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{name: "Emeke", age: 22, availabilty: true}
	)

	fmt.Println(bill)

	var emeke Passenger

	emeke.age = 22
	emeke.availabilty = true
	emeke.name = "Ofuzor Chukwuemeke"

	fmt.Println(emeke)

}
