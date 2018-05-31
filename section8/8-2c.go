package main

import "fmt"

type shirt struct {
	size int
}

func main() {
	s := shirt{}
	var a int
	s.size, a = 5, 3
	fmt.Printf("size: %d, a: %d\n", s.size, a)
}
