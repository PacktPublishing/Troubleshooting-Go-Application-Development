package main

import "fmt"

func main() {
	things := []string{"screwdriver", "torch", "phone booth"}
	for i := 0; i < 5; i++ {
		if i > len(things)-1 {
			fmt.Printf("No index %d\n", i)
			continue
		}

		fmt.Printf("Item %d is %s\n", i, things[i])
	}
}
