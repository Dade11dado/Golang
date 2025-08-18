package main

import "fmt"

func main() {
	x := 100
	for x > 0 {
		x--
		if x%2 != 0 {
			continue
		}
		fmt.Printf("X is %v \n", x)
	}
}
