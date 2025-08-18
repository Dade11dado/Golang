package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func main() {

	p1 := person{
		first: "Davide"}

	p2 := person{
		first: "Maurizo"}

	p1.speak()
	p2.speak()
	x := []int{1, 2, 3, 4, 5, 6, 7, 89, 9}
	foo()
	foo1("Hello")
	foo2("Davide")
	foo3("Daisy", 9)
	// ... opens up a slice
	fmt.Println(sum(x...))

}

// "func" + [receiver] + identifier + ([parameters]) + [(returns)] + {}
// no params no return
func foo() {
	fmt.Println("I am foo")
}

// one parameter
func foo1(s string) {
	fmt.Println(s)
}

// one paramters one return
func foo2(s string) string {
	return fmt.Sprintf("They call me mister %v", s)
}

// one parametrs multiple returns
func foo3(s string, a int) (string, int) {
	a *= 7
	return fmt.Sprintf("Todd %v is this old", s), a
}

// variadic parameters
func sum(ii ...int) int {
	fmt.Print(ii)
	fmt.Printf("%T	\n", ii)

	value := 0
	for _, v := range ii {
		value += v
	}
	return value
}
