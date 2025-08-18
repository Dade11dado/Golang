package main

import "fmt"

func main() {
	p1 := person{"Maurizio"}
	fmt.Println(p1)
	fmt.Println(valueSemantic(p1))
	pointerSemantic(&p1)
	fmt.Println(p1)
}

type person struct {
	first string
}

func valueSemantic(p person) person {
	p.first = "Davide"
	return p
}

func pointerSemantic(p *person) {
	p.first = "Davide"
}
