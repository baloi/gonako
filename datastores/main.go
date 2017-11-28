// datastores/main.go
// shows a redis-like datastore in go
// Just using parts taken from https://github.com/felixge/go-redis/blob/master/main.go
package main

import (
    "fmt"
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

func main() {
    fmt.Printf("Server started\n")
    var id int64
    store := &store{
        data: make(map[string]string),
        // lock: &sync.RWMutex{},
    }

    id = 0
    client := &client{id: id, conn: conn, store: store}
    // go client.serve()

}
