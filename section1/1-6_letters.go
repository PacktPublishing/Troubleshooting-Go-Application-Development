package main

import "fmt"

func main() {
	for _, c := range []rune("Gopher!") {
		fmt.Println(string(c))
	}
}
