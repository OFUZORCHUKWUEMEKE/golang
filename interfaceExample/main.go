package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Motorcyle string

type Car string

type Truck string

func (m Motorcyle) String() string {
	return fmt.Sprintf("Motorcyle: %v", string(m))
}
func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}
func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (m Motorcyle) PickLift() Lift {
	return SmallLift
}

func (m Car) PickLift() Lift {
	return StandardLift
}

func (m Truck) PickLift() Lift {
	return LargeLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift")
	case StandardLift:
		fmt.Printf("send %v to standard lift")
	case LargeLift:
		fmt.Printf("send %v to standard lift")
	}
}

func main() {
	car := Car("Sporty")
	truck := Truck("Mountain Cruiser")
	motorcycle := Motorcyle("Motorcyle")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)
}
