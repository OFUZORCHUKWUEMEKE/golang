package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func sumFile(rd bufio.Reader) int {
	sum := 0
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error", err)
		}
		num, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			fmt.Println("Error", err)
		}
		sum += num
	}
	return sum
}

func main() {

}
