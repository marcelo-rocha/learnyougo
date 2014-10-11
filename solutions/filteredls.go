// FILTERED LS - Exercise 5
// Create a program that prints a list of files in a given directory, filtered by the extension of the files.
// You will be provided a directory name as the first argument to your program
// (e.g. '/path/to/dir/') and a file extension to filter by as the second argument.
// For example, if you get 'txt' as the second argument then you will need to filter
// the list to only files that end with .txt.
// The list of files should be printed to the console, one file per line

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	pattern := path.Join(os.Args[1], "*."+os.Args[2])
	files, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}
	for i := range files {
		fmt.Println(files[i])
	}
}
