package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		log.Printf("Failed to open connection: %s", err)
		log.Printf("The conn is type %T", conn)

		return
	}

	conn.Write([]byte("hello"))
}
