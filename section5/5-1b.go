package main

import "net/http"

func main() {
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
