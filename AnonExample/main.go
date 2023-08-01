package main

import "fmt"

const (
	Form   = "fit"
	INJURY = "injured"
)

type Player struct {
	number   int
	position string
	health   string
	starter  bool
	star     bool
}

type GoalKeeper struct {
	number int
	health string
}

type Coach struct {
	name string
}

type Squad struct {
	players    [10]Player
	coach      Coach
	goalKeeper GoalKeeper
}

func (play *Squad) starter(Player) {

}

func main() {
	fmt.Println("data")
}
