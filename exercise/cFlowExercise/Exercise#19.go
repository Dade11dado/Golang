package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	x := rand.Intn(251)
	fmt.Printf("The value name is x and the number %v is stored in it \n", x)

	switch {
	case x < 100:
		fmt.Printf("Between 0 and 100")
	case x < 200:
		fmt.Printf("Between 101 and 200")
	case x < 250:
		fmt.Printf("Between 201 and 250")
	}
}
