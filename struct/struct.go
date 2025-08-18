package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	as := struct {
		name    string
		surname string
	}{
		name:    "Davide",
		surname: "vavasori"}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   47},
		ltk: true}

	ps := person{
		first: "jenny",
		last:  "Monepenny",
		age:   27}

	fmt.Println(sa1, ps)
	fmt.Printf("The type of sa1 is %T \n", as)
}
