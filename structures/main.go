package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
}

func Area(data rectangle) int {
	return data.breadth * data.length
}

func Perimeter(data rectangle) int {
	return 2*data.breadth + 2*data.length
}

func main() {
	data := rectangle{length: 50, breadth: 4}
	area := Area(data)
	fmt.Println(area)
	perimeter := Perimeter(data)
	fmt.Println(perimeter)

	//  double(area)
}
