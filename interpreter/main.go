// main.go
package main

import (
	"fmt"
	"os"
	"interpreter/repl"
)

func main() {
	fmt.Println("L & M interpreter system. Welcome!")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}