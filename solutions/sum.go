// sum - Exercise 2
// Write a program that accepts one or more numbers as command-line arguments and
// prints the sum of those numbers to the console (stdout).

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	total := 0
	params := os.Args[1:]
	for i := range params {
		param, _ := strconv.Atoi(params[i])
		total += param
	}
	fmt.Println(total)
}
