package main
// Demonstrates how to create goroutines and how
// the scheduler behaves

import (
    "fmt"
    "runtime"
    "sync"
)

// main is the entry point for all Go programs
func main() {
    // Allocate 1 logical processor
    // for the scheduler to use
    runtime.GOMAXPROCS(1)

    // wg us used to wait for the program to
    // to finish. Add a count of two, one for
    // each goroutine
    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("start Goroutines")

    // Declare and anonymous function and create
    // a goroutine.
    go func() {
        // Schedule the call to Done to tell main
        // we are done.
        defer wg.Done()
        // Display the alphabed three times
        for count := 0; count < 3; count ++ {
            for char := 'a'; char < 'a'+26; char++ {
                fmt.Printf("%c ", char)
            }
        }
    } ()

    go func() {
        // Schedule the call to Done to tell main
        // we are done
        defer wg.Done()
        // Display the alphabet 2 times
        for count := 0; count < 2; count ++ {
            for char := 'A'; char < 'A'+ 26; char++ {
                fmt.Printf("%c ", char)
            }
        }
    } ()

    fmt.Println("Waiting to finish")
    wg.Wait()

    fmt.Println("\nTernimating Program")

}


