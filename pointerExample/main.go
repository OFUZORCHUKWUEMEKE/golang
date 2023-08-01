package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}
func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	fmt.Println("Checking Out")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main() {
	shirt := Item{"Shirt", Active}
	pant := Item{"Shirt", Active}
	purse := Item{"Shirt", Active}
	watch := Item{"Shirt", Active}

	items := []Item{shirt, pant, purse, watch}
	fmt.Println(items)

	deactivate(&items[0].tag)

	fmt.Println(items)

	checkout(items)

	fmt.Println(items)

}
