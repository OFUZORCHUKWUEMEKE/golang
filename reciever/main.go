package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}
func (lot *ParkingLot) occupyeSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	// make
	lot := ParkingLot{spaces: make([]Space, 6)}

	fmt.Println("Initial:", lot)
	lot.occupyeSpace(1)
	occupySpace(&lot, 2)

	lot.vacateSpace(3)

	fmt.Println("AFter occupied", lot)
}
