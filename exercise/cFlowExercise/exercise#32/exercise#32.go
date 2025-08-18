package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	fmt.Println("--------------------")

	for v, _ := range m {
		fmt.Printf("The value associated with the key %v is %v \n", v, m[v])
	}
	fmt.Println("--------------------")

	age := m["James"]
	fmt.Println(age)

	fmt.Println("--------------------")
	//check for q in map
	q := m["Q"]
	fmt.Println(q)

	fmt.Println("--------------------")
	//using comma ok to check for q in the map
	if v, ok := m["q"]; ok {
		fmt.Println(v)
	} else {
		fmt.Printf("Q is not present and v is %v", v)
	}

}
