package main

import (
	"fmt"

	"github.com/shawnmilo/girya"
)

func main() {
	kg := 24.0
	fmt.Printf("%.1f kg is %.1f lbs.\n", kg, girya.KgToLbs(kg))
}
