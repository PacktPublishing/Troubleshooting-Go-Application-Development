package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	s := "42"
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to convert %q to number: %s\n", s, err)
	}
	fmt.Printf("The number is %d\n", i)

	s = "23"

	i, err = strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to convert %q to number: %s\n", s, err)
	}
	fmt.Printf("The number is %d\n", i)

}
