package main

import "fmt"

func main() {
	c := make(chan int)

	go feed(c)

	receive(c)
}

func feed(c chan int) {
	go func() {
		for i := range 100 {
			c <- i
		}
		close(c)
	}()
}

func receive(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
