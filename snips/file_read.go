package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	dat, err := ioutil.ReadFile("test.txt")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("test.txt")
	check(err)

	b1 := make([]byte, 5) // space is made for b1 which will be a byte array of 5
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
}
