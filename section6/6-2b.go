package main

import "fmt"

func main() {
	eNum := getEvenNum(8)
	fmt.Println(eNum.half())

}

type evenNum struct {
	val int
}

func (en evenNum) half() int {
	return en.val / 2
}

func getEvenNum(i int) *evenNum {
	if i%2 == 1 {
		return nil
	}

	return &evenNum{val: i}
}
