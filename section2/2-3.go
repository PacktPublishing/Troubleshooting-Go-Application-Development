package main

import "fmt"

func status(msg *string) {
	fmt.Println(*msg)
}

func main() {
	msg := "hello"
	defer status(&msg)

	fmt.Println(msg)
	msg = "goodbye"
}
