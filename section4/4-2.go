package main

import "fmt"

func getNumber() int {
	return 42
}

func main() {
	num := getNumber()
	fmt.Println(num)
}
