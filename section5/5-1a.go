package main

import "fmt"

func main() {
	print("hello")
}

func print(msg string) {
	fmt.Println(msg)
	panic("broke!")
}
