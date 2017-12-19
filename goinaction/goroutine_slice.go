package main

import (
    "fmt"
    "runtime"
    "sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {
    // Allocate 1 logical processors for the schedule rot use
    runtime.GOMAXPROCS(1)
    // add a count of two, one for each goroutine
    wg.Add(2)

    // create two goroutines
    fmt.Println("Create Goroutines")

    go printPrime("A")
    go printPrime("B")

    fmt.Println("Waiting to finish")
    wg.Wait()
    fmt.Println("Terminating program")
}

func printPrime(prefix string) {
    defer wg.Done()
next:
    for outer := 2; outer < 5000; outer ++ {
        for inner := 2; inner < outer; inner++ {
            if outer % inner == 0 {
                continue next
            }
        }
        fmt.Printf("%s:%d\n", prefix, outer)
    }
    fmt.Println("Completed", prefix)
}
