/*
 * sync_try.go
 * Concurrent programming using Mutex.
 * This is what can be used if we do not need
 * to pass messages around in goroutines (through channels)
 * but want to lock a resource which can be done with
 * sync.Mutex. A more simple technique than channels.
 */

package main

import (
    "fmt"
    "sync"
    "time"
)
// SafeCounter is safe to use concurrently
type SafeCounter struct {
    val map[string]int
    mux sync.Mutex
}

// Inc increments the counter for the given key
func (counter *SafeCounter) Inc(key string) {
    counter.mux.Lock()
    // Lock so only one goroutine at a time can access the map counter.val
    counter.val[key]++
    counter.mux.Unlock()
}

// Value returns the current value for the given key
func (counter *SafeCounter) Value(key string) int {
    counter.mux.Lock()
    defer counter.mux.Unlock()
    return counter.val[key]
}

func main() {
    fmt.Println("initializing")
    c := SafeCounter { val: make(map[string]int) }
    for i := 0; i < 1000; i++ {
        go c.Inc("somekey")
    }

    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))
}
