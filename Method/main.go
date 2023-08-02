package main

import (
	"fmt"
	"math"
)

// Go does not have classes. Howerver , you can define methods on Types
// A method is a function with a special reciever argument

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.X*v.Y)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	v := Vertex{2, 3}
	answer := v.Abs()
	answer2 := f.Abs()
	fmt.Println(answer)
	fmt.Println(answer2)
}
