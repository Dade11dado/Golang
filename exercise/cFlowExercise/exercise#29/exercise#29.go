package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("Inner loop number %v, outer loop number %v \n", i, j)
		}
	}
}
