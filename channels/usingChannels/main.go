package main

import "fmt"

func main() {

	c := make(chan int)
	//Send
	go foo(c)

	//Receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

// send
func foo(c chan<- int) {
	for i := range 100 {
		c <- i
	}
	close(c)
}
