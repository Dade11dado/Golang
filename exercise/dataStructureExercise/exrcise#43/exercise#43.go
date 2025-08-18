package main

import "fmt"

func main() {

	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range a {
		fmt.Printf("Value %v is of type %T \n", v, v)
	}

	b := append(a[:5])
	fmt.Println(b)
	c := append(a[5:])
	fmt.Println(c)
	d := append(a[2:7])
	fmt.Println(d)
	e := append(a[1:6])
	fmt.Println(e)

}
