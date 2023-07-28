package main

import "fmt"

func main() {
	slices := []string{"hello", "world", "!"}
	for i, element := range slices {
		fmt.Println(i, element, ":")
		for _, ch := range element {
			fmt.Println(" %q\n", ch)
		}
	}
}
