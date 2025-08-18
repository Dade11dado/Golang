package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	for i := range 10 {
		go func() {
			for j := range 10 {
				c <- fmt.Sprintf("This is loop %d in func %d", j, i)
			}
		}()
	}

	for k := range 100 {
		fmt.Println(k, <-c)
	}
}
