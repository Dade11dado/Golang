package main

import "fmt"

func main() {

	a := [3]int{42, 43, 44}
	fmt.Println(a)

	b := [...]string{"Hello", "Goophers"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	//declare a vaiable of type [7]int
	var ni [7]int
	//assign a value to index position 0
	ni[0] = 42
	fmt.Printf("%v \t\t %T \n", ni, ni)

	//array literal
	ni2 := [4]int{55, 56, 57, 58}
	fmt.Printf("%v \t\t %T \n", ni2, ni2)

	// array literals have compiler count elements
	ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
	fmt.Printf("%v \t\t %T \n", ns, ns)
}
