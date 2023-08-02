package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("SAMPLE")

	var newString strings.Builder

	buffer := make([]byte, 4)

	for {
		numBytes, err := reader.Read(buffer)
		chunk := buffer[:numBytes]

		newString.Write(chunk)
		fmt.Printf("Read %v bytes: %c\n", numBytes)

		if err == io.EOF {
			break
		}

	}
	fmt.Printf("%v\n", newString.String())
}
