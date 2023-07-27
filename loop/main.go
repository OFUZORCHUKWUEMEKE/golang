package main

import "fmt"

func main() {
	sum := 0
	fmt.Println(sum)
	for i := 1; i <= 10; i++ {
		sum += 1
		fmt.Println("sum is", sum)
	}
	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement. Sum is", sum)
	}

	ForLoop()
}

func ForLoop() {
	sum := 0
	for i := sum; i < 50; i++ {
		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		if divisibleBy3 && divisibleBy5 {
			fmt.Println("FizzBuzz")
		} else if divisibleBy3 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}
