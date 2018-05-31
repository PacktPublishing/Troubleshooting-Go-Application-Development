package main

import "fmt"

func main() {
	primes := []int{2, 3, 7, 11, 13, 17}
	for _, p := range primes {
		fmt.Println(p)
	}
}
