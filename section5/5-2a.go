package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	msg := "Message 42: Don't panic!"
	conn, err := net.DialTimeout("tcp", "localhost:23", time.Second)
	if err != nil {
		fmt.Printf("Failed to create connection: %s\n", err)
		return
	}

	conn.Write([]byte(msg))
	fmt.Println("All done!")
}
