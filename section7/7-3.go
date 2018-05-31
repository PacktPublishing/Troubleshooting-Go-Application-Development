package main

import "fmt"

type incrementer interface {
	incr()
}

type counter struct {
	count int
}

func (c *counter) incr() {
	c.count++
}

func main() {
	c := &counter{}
	fmt.Printf("Counter before: %d\n", c.count)
	tick(c)
	fmt.Printf("Counter after: %d\n", c.count)
}

func tick(i incrementer) {
	i.incr()
}
