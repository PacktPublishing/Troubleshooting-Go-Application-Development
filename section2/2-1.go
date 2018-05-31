package main

import "fmt"

func main() {

	nums := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			nums <- i
		}
		close(nums)

	}()

	for i := range nums {
		fmt.Println(i)
	}

}
