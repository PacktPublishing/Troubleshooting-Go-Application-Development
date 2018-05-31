package main

import "fmt"

type animal struct {
	name  string
	sound string
}

func (a animal) makeNoise() {
	fmt.Printf("The %s goes %q.\n", a.name, a.sound)
}

func main() {
	d := animal{name: "dog", sound: "bark"}
	d.makeNoise()
	print(d)
	print("monkey")

}

func print(thing interface{}) {
	fmt.Printf("thing is a %T with a value of %#v\n", thing, thing)
	x, ok := thing.(animal)
	if ok {
		fmt.Printf("x is a %T with a value of %#v\n", x, x)
		x.makeNoise()
	} else {
		fmt.Printf("thing %v is not an animal\n", thing)
	}

}
