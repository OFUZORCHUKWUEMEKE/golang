package main

import "errors"

// GO has no exceptions
// Errors are returned as the last return value from a function
//   -Encodes failure as part of the function signature
// Simple to determine if a function can fail
// Return nil if no error occurred
// Error implements the error interface from stf
//  -One function to implement:Error string

type error interface {
	Error() string
}

func divide(lhs, rhs int) (int, error) {
	if rhs == 0 {
		return 0, errors.New("Error")
	} else {
		return rhs / lhs, nil
	}
	// return rhs / lhs, nil
}

func main() {
	errors.New("Cannot divide by Zero")
}
