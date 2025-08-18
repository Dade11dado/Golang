package main

import "fmt"

func main() {

	func(s string) {
		fmt.Printf("%s this is me", s)
	}("Davide")

	//function expressions

	x := func() {
		fmt.Println("Anonymous")
	}

	x()
}

func foo() {
	fmt.Println("Hooray")
}
