// list_try.go
package main

import "fmt"

func deleteElement(s []string, elem string) []string {
	var index int
	index = 0
	for x :=  0; x < len(s) && s[x] != elem; x++ {
		index++;
	}
	s = append(s[:index], s[index+1:]...)
	return s
}

func main() {
	s := make([]string, 3) // will be an empty list of type string of length 3
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	s = append(s, "x")
	s = append(s, "z")
	fmt.Println("s is now", s)

	// deleting an element
	s = deleteElement(s, "b")

	fmt.Println("s is now ", s)
}
