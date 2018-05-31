package main

import "fmt"

func main() {
	x := 3
	y := 4
	fmt.Printf("before x: %d, y: %d\n", x, y)
	x, y = 5000, 7000
	fmt.Printf("x: %d, y: %d\n", x, y)
	fmt.Printf("after x: %d, y: %d\n", x, y)
}
