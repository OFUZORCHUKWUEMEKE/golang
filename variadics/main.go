package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	a := []int{2, 3, 4}
	b := []int{4, 5, 6}

	all := append(a, b...)

	answer := sum(all...)

	fmt.Println(answer)
}
