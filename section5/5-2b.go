package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	msg := "Go can read and write files!\n"
	filename := "/tmp/message.txt"

	err := ioutil.WriteFile(filename, []byte(msg), 0644)
	if err != nil {
		fmt.Printf("Failed to write file: %s\n", err)
	}

	fmt.Println("All done!")

}
