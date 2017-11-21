// dup3.go version of program reads a file as a while, not per line into
// memory in one gulp.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// counts is a hash map with a string key and int value
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// split the data read from file(s) into lines delimited by "\n"
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	// Now print the duplicate lines and their respective counts
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
