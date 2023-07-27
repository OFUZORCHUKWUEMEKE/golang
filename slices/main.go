package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	slice1 := numbers[:]
	slice2 := numbers[1:]
	slice3 := numbers[1:3]
	slice4 := numbers[1:2]
	fmt.Println(slice1, slice2, slice3, slice4)

	// Append

	numbers = append(numbers, 4, 5, 6)
	fmt.Println(numbers)

	slice := make([]int, 10)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	// slice[1] = 10
	fmt.Println(slice)
}
