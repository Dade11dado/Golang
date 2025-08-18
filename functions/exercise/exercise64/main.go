//We will be doing now some functions to do basic math calculation

package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)
}

// doMath will acept 3 parametrs, two ints, which are the numbers
// and a function, whic is the function that will do the calculations
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

// add will acept two int as parametrs and do the adding on them
func add(a int, b int) int {
	return a + b
}

// subtract will acept two int and will do the sibtration to them
func subtract(a int, b int) int {
	return a - b
}
