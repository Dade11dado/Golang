package main

import "fmt"

func main() {
	f := incrementor()
	fmt.Println(f()) //it increments the value by one
	fmt.Println(f()) //it increments the value by one
	fmt.Println(f()) //it increments the value by one
	fmt.Println(f()) //it increments the value by one
	fmt.Println(f()) //it increments the value by one
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
