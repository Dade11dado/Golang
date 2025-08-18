package main

import "fmt"

func main() {

	m := make(map[string][]string)

	m["bond_james"] = []string{"shaken", "Not_stirred", "martinis", "fast_cars"}
	m["money_penny"] = []string{"intelligence", "literature", "computer", "science"}
	m["no_dr"] = []string{"cats", "icecream", "sunsets"}
	m["flming_ian"] = []string{"steak", "cigars", "espionage"}

	for _, v := range m {
		for i, p := range v {
			fmt.Printf("Value %v has index %d \n", p, i)
		}
	}

	fmt.Println("----------deleting element -----------")

	delete(m, "flming_ian")

}
