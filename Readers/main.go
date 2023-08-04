package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Reader and writers are interfaces that allow reading from  and writing to i/o sources
// Network sockets ,files , arbitrary arrays
// Multiple implementations in standard library

func main() {
	r := bufio.NewReader(os.Stdin)
	sum := 0

	for {
		input, inputErr := r.ReadString(' ')

		n := strings.TrimSpace(input)

		if n == "" {
			continue
		}
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}
		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
		}
	}
	fmt.Printf("sum : %v\n", sum)
}
