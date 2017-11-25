package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    //var c chan int
    c := make(chan string)

    go boring("boring!", c)
    for i := 0; i < 5; i++ {
        fmt.Printf("You say: %q\n", <-c) // Receive expressing is
                                         // just a value.
    }
    fmt.Println("You're boring; I'm leaving.")

    /*
    // Sending on a channel
    c <- 1
    // Receiving from a channel
    // The "arrow" indicates the direction of the data flow
    value = <-c
    fmt.Println("value = %d", value)
    */
}

func boring(msg string, c chan string) {
    for i := 0; ; i++ {
        c <- fmt.Sprintf("%s %d", msg, i) // Expressing to be sent
                                        // can be any suitable value
        time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
}
