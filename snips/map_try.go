package main

import "fmt"

func main() {
	celebs := map[string]int{
		"Nicolas Cage":      50,
		"Selena Gomez":      21,
		"Jude Law":          41,
		"Scarlet Johansson": 29,
	}

	fmt.Printf("%#v", celebs)
}
