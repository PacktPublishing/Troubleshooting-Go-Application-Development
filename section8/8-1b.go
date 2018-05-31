package main

import (
	"fmt"
	"time"
)

const a = 1

var b = 2

func main() {
	dur := time.Duration(b)
	time.Sleep(time.Second * dur)
	fmt.Println("done")
}
