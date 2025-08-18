package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := range 42 {
		fmt.Printf("Iteration number %v \n", i)
		x := rand.Intn(5)
		switch {
		case x == 0:
			fmt.Printf("Value X is 0 \n")
		case x == 1:
			fmt.Printf("Value X is 1 \n")
		case x == 2:
			fmt.Printf("Value X is 2 \n")
		case x == 3:
			fmt.Printf("Value X is 3 \n")
		case x == 4:
			fmt.Printf("Value X is 4 \n")
		}
	}
}
