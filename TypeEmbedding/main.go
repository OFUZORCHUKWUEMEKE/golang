package main

import "fmt"

// Embedding intefaces allow you to "embed" an interface into another interface
// Implementing the interface requires all embedded functions to be implemented
// Reduces the need to write duplicate  inteface declarations
// changes in embedded interfaces automatically propagate

const (
	small = iota
	medium
	large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "large"}[b]
}
func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}
type Shipper interface {
	Ship() Shipping
}
type WareHouse interface {
	Shipper
	Conveyor
}

type spamMail struct {
	amount int
}

func (s spamMail) String() string {
	return "Spam Mail"
}

func (s *spamMail) Ship() Shipping {
	return Air
}
func (s *spamMail) Convey() BeltSize {
	return small
}

func automate(item WareHouse) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	mail := spamMail{amount: 40000}
	automate(&mail)

	// waste := ToxicWaste{300}
	// // automate(&waste)
}
