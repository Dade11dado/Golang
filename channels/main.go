package main

import "fmt"

func main() {
	//create a channel c of type int, with numbrer you specify that this is
	//a buffeer channel
	//a buffer channel allow variable to sit in there even if not
	//ready
	c := make(chan int, 1)

	//If i declare the channel with an arrow between chan and the type
	//I specify the direction the data will only be able to have
	d := make(chan<- int) //only able to send
	e := make(<-chan int) // only able to receive

	go func() {
		//like this we put values in channel
		c <- 42
		d <- 43
		fmt.Println(<-e)

	}()

	//To get data use the arrow
	fmt.Println(<-c)
	fmt.Println("--------")
	fmt.Printf("%T\n", c)
}
