package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("budget application...")
	dat, err := ioutil.ReadFile("budget.txt")
	check(err)
	fmt.Println(string(dat))
}
