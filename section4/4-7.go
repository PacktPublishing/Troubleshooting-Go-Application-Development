package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			m[i] = i
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Printf("map: %+v\n", m)
}
