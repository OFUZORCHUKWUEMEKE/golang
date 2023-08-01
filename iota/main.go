package main

import "fmt"

const (
	Add = iota
	Substract
	Multiply
	Divide
)

type Operation int

func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Substract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	}
	panic("unhandled operation")
}

func main() {
	add := Operation(Add)
	sum := Operation(Substract)
	mult := Operation(Multiply)
	fmt.Println(add.calculate(2, 3))
	fmt.Println(sum.calculate(2, 3))
	fmt.Println(mult.calculate(2, 4))

}
