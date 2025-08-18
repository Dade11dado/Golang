package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 2))
}

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, i int) string {
	x := f(i)
	y := fmt.Sprintf("The result is %d", x)
	return y
}
