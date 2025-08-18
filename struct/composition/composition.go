package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type foo int

func main() {
	var a foo = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
