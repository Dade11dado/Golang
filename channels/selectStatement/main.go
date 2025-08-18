package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)

	fmt.Println("About to exit")
}

func send(e, o, q chan<- int) {
	for i := range 100 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel: ", v)
		case v := <-o:
			fmt.Println("From the odd channel: ", v)
		case v := <-q:
			fmt.Println("Quitting: ", v)
			return
		}
	}
}
