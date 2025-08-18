package main

import "fmt"

type person struct {
	first            string
	last             string
	favoriteIcecream []string
}

func main() {
	p1 := person{
		first:            "Davide",
		last:             "Vavasori",
		favoriteIcecream: []string{"Pistacchio", "Cream"}}

	for _, v := range p1.favoriteIcecream {
		fmt.Printf("%v favourite icream is %v \n", p1.first, v)
	}

}
