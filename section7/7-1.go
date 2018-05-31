package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func writeMessage(output io.Writer) {
	msg := "The best message you'll ever see!\n"
	_, err := output.Write([]byte(msg))
	if err != nil {
		log.Printf("Couldn't write message: %s", err)
	}
	c, ok := output.(io.Closer)
	if !ok {
		fmt.Println("not a closer; quitting")
		return
	}
	fmt.Println("closing output")
	c.Close()
}

func main() {
	f, err := ioutil.TempFile("", "gopher")
	if err != nil {
		log.Printf("Couldn't get temp file: %s", err)
		return
	}
	writeMessage(f)
	fmt.Printf("Wrote to %s\n", f.Name())
	http.HandleFunc("/", index)
	fmt.Println("starting server...")
	http.ListenAndServe("localhost:8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	writeMessage(w)
}
