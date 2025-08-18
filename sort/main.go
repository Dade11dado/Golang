package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type BAge []person

type BName []person

func (p person) String() string {
	return fmt.Sprintf("%s: %d \n", p.first, p.age)
}

func (p BAge) Len() int {
	return len(p)
}

func (p BAge) Less(i, j int) bool {
	return p[i].age < p[j].age
}

func (p BAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p BName) Len() int {
	return len(p)
}

func (p BName) Less(i, j int) bool {
	return p[i].first < p[j].first
}

func (p BName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(BAge(people))
	fmt.Println(people)
	sort.Sort(BName(people))
	fmt.Println(people)

}
