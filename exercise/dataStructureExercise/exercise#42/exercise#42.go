package main

import "fmt"

func main() {

	a := [5]int{}
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4

	for i, v := range a {
		fmt.Printf("Value at index %v is %v \n", i, v)
	}

}
