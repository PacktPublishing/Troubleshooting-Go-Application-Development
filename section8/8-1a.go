package main

import "fmt"

const a = 1

var b = 2
var c int = 3

func main() {
	d := 4
	var e int32 = 5
	for i := 1; i < 2; i++ {
		fmt.Printf("%d + %d = %d\n", i, a, i+a)
		fmt.Printf("%d + %d = %d\n", i, b, i+b)
		fmt.Printf("%d + %d = %d\n", i, c, i+c)
		fmt.Printf("%d + %d = %d\n", i, d, i+d)
		fmt.Printf("%d + %d = %d\n", i, e, i+int(e))
	}
}
