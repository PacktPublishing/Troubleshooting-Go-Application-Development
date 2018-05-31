package main

import "fmt"

func main() {
	list := []string{"milk", "bread", "eggs"}
	fmt.Println(list)

	backup := make([]string, len(list))
	i := copy(backup, list)
	fmt.Printf("Copied %d items\n", i)
	fmt.Println(backup)
}
