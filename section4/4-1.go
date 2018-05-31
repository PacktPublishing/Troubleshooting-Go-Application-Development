package main

import "fmt"

func main() {
	sizes := map[string]int{
		"shoe": 11,
	}

	for _, item := range []string{"shoe", "sock"} {
		size, found := sizes[item]
		if !found {
			fmt.Printf("%s not found\n", item)
			continue
		}
		fmt.Printf("The %s is size %d\n", item, size)
	}
}
