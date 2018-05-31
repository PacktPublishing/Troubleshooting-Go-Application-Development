package main

import "time"

func main() {
	for {
		t := time.NewTicker(time.Nanosecond)
		<-t.C
		t.Stop()

	}

}
