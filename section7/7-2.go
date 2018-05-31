package main

import (
	"fmt"
	"log"
	"os"
)

type adder interface {
	add()
}

type nums struct {
	a int
	b int
}

func (n nums) add() {
	fmt.Printf("%d + %d = %d\n", n.a, n.b, n.a+n.b)
}

type words struct {
	first  string
	second string
}

func (w words) add() {
	fmt.Printf("%s %s\n", w.first, w.second)
}

func main() {
	n := nums{a: 3, b: 4}
	worker(n)
	w := words{first: "maximum", second: "effort"}
	worker(w)
	log.SetOutput(os.Stdout)
	log.Println("done")
}

func worker(n adder) {
	n.add()
}
