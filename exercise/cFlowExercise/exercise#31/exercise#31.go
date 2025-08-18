package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for v, _ := range m {
		fmt.Printf("The value associated with the key %v is %v \n", v, m[v])
	}
}
