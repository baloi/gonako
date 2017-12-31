// datastores/main.go
// shows a redis-like datastore in go
// Just using parts taken from https://github.com/felixge/go-redis/blob/master/main.go
package main

import (
    "fmt"
    "bufio"
    "net"
    "log"
)

type store struct {
    data map[string]string
    //lock *sync.RWMutex
}

type client struct {
    id int64
    conn net.Conn
    reader *bufio.Reader
    store *store
}

func (store *store) Set(key string, val string) {
    //store.lock.Lock()
    //defer store.lock.Unlock()
    store.data[key] = val
}

func (client *client) serve() {

}

func main() {
    fmt.Printf("Server started\n")
    addr := "8080"
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Printf("Error: listen(): %s", err)
        os.Exit(1)
    }

    log.Printf("Accepting connections at: %s", addr)
    store := &store{
        data: make(map[string]string),
        // lock: &sync.RWMutex{},
    }

    var id int64
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Error: Accepting(): %s", err)
            continue
        }
        id++
        client := &client{id: id, conn: conn, store: store}
        go client.serve()
    }

}
