package main

import (
	"net"
	"log"
	"io"
	"time"
	"fmt"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // client disconnect
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	fmt.Println("Listening on port 8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Accepting connections")
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}
