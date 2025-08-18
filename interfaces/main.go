package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am ", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent ", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	sa := secretAgent{
		person: person{
			first: "Davide"},
		ltk: true,
	}

	p2 := person{
		first: "Maurizio"}

	saySomething(sa)
	saySomething(p2)

}
