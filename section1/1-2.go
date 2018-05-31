package main

import "fmt"

var colors = []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}

func main() {
	for _, c := range colors {
		fmt.Println(c)
	}
}
