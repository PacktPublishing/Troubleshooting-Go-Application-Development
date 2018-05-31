package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("sending %d\n", i)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("done")
}
