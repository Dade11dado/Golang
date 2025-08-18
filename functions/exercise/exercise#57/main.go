package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := person{
		first: "Davide",
		age:   28,
	}
	b := []int{1, 2, 3, 4, 5, 6, 7}
	foo1(b...)
	defer fmt.Println(foo())
	fmt.Println(bar())
	p1.speak()

	sq := square{
		length: 5,
		width:  4,
	}

	circ := circle{
		radius: 4,
	}

	info(sq)
	info(circ)

}

type person struct {
	first string
	age   int
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return s.length * s.width
}

func (p person) speak() {
	fmt.Printf("My name is %s and i am %d years old \n", p.first, p.age)
}

func foo() string {
	return "I am the second"
}

func foo1(a ...int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

func bar() (int, string) {
	return 1, "I am the third"
}

func bar1(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
