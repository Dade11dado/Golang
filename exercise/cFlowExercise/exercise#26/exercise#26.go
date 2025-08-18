package main

import "fmt"

func main() {
	x := 0
	for x < 10 {
		fmt.Printf("X is now %v \n", x)
		x++
		if x == 6 {
			break
		}
	}
}
