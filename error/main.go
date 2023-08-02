package main

import (
	"errors"
	"fmt"
)

// GO has no exceptions
// Errors are returned as the last return value from a function
//   -Encodes failure as part of the function signature
// Simple to determine if a function can fail
// Return nil if no error occurred
// Error implements the error interface from stf
//  -One function to impl ement:Error string

// Use errors.New() as the last return value from a function
// use errors.As() to retrieve an error , or errors.Is() to check the error type
// implement the errror interface for custom errore
// Always implement the interface as a reciever function

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no element found"))
	} else {
		return s.values[index], nil
	}
}

func main() {
	stuff := Stuff{}
	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}

}
