// readfile - Exercise 3
// Write a program that uses a single synchronous filesystem operation to read a file
// and print the number of newlines it contains to the console (stdout), similar to running cat file | wc -l.
// The full path to the file to read will be provided as the first command-line argument.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1]
	buffer, _ := ioutil.ReadFile(fileName)
	count := 0
	for i := range buffer {
		if buffer[i] == 10 {
			count++
		}
	}
	fmt.Println(count)
}
