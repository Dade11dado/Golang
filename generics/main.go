package main

import "fmt"

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(add(2.4, 3.7))

}

type myNumbers interface {
	int | float64
}

// This is typed sets
func add[T myNumbers](a, b T) T {
	return a + b
}

//This is typed constraint
// func add[T float64 | int](a, b T) T {
// 	return a + b
// }
