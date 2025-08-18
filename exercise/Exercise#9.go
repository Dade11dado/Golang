package main

import "fmt"

var myName string = "Davide"

const year int = 1997

func main() {
	place := "Iseo"

	fmt.Printf("My name is %v, I was born in %d and I live in %v", myName, year, place)
}
