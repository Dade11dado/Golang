package main

import "fmt"

func main() {
	x := func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Iteration number %d \n", i)
		}
	}

	x()

	y := first()
	fmt.Println(y())
}

func first() func() int {
	x := 10
	return func() int {
		x += 10
		return x
	}
}
