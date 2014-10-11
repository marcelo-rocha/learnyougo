// readfilego - Exercise 4
// Write a program that read a file and print the number of newlines it contains to
// the console (stdout), similar to running cat file | wc -l.
// The full path to the file to read will be provided as the first command-line argument.
// Hints:
// The solution to this problem is almost the same as the previous problem except you
// must now checking for errors at each operation using "os" package

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	afile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := afile.Close(); err != nil {
			panic(err)
		}
	}()

	count := 0
	buffer := make([]byte, 1000)
	for {
		len, err := afile.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if len == 0 {
			break
		}
		for i := 0; i < len; i++ {
			if buffer[i] == 10 {
				count++
			}
		}
	}
	fmt.Println(count)
}
