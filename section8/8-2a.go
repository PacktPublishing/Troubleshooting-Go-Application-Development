package main

import "fmt"

func main() {
	i := 23
	fmt.Printf("i is %d\n", i)
	colors := []string{"purple", "gray", "green", "black"}
	for i, color := range colors {
		fmt.Printf("Color %d is %s\n", i, color)
		j := 5 * i
		fmt.Printf("j: %d\n", j)
	}
	fmt.Printf("i is %d\n", i)
}
