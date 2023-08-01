package main

import (
	"example.com/mymodule/packages/display"
	"example.com/mymodule/packages/msg"
)

func main() {
	msg.Existing("Thats whats up")
	msg.Hi()
	display.Display("I love Golang")
}
