package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 0
	for x < 100 {
		x++
		if x := rand.Intn(5); x == 3 {
			fmt.Println("X is 3")
		}
	}

}
