package main

import (
	"fmt"
	"sync"
)

type person struct {
	first string
}

type human interface {
	Speak()
}

func (p *person) Speak() {
	fmt.Println("Hello, my name is ", p.first)
}

func saySomething(h human) {
	h.Speak()
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("First go routine")
		wg.Done()
	}()
	go func() {
		fmt.Println("Second go routine")
		wg.Done()
	}()
	wg.Wait()

	//p1 := person{}

	p2 := person{
		first: "MAurizo",
	}

	//saySomething(p1) //doesn't work since it requires a pointer
	saySomething(&p2)
}
