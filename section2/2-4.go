package main

import (
	"fmt"
	"sync"
)

func msg(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			msg(num, &wg)
		}(i)
	}
	wg.Wait()
	fmt.Println("done")
}
