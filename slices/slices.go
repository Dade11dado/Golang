package main

import "fmt"

func main() {

	//decalred like a array
	xs := []string{"Hello", "World"}

	fmt.Println(xs)

	//append to a slice
	x1 := []int{42, 43, 44}
	fmt.Println(x1)
	fmt.Println("-------------")
	x1 = append(x1, 3, 6, 7)
	fmt.Println(x1)

	//slicing a slice
	xslice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("No slicing: %v \n", xslice)

	//[inclusive:exclusive]
	fmt.Printf("Slicing inclusive exclusive: %v \n", xslice[0:4])

	//[:exclusive]
	fmt.Printf("Slicing :exclusive: %v \n", xslice[:4])

	//[inclusive:]
	fmt.Printf("Slicing inclusive: %v \n", xslice[2:])

	//[:]
	fmt.Printf("Slicing : %v \n", xslice[:])

	//How to delete from a slice
	xslice = append(xslice[:4], xslice[5:]...)
	fmt.Printf("Deleted slice is: %v \n", xslice)

	//create a slice with the make function
	xmake := make([]int, 2, 10)
	fmt.Printf("Array has length of %v and capacity of %v... there it is %v\n", len(xmake), cap(xmake), xmake)

	//multidimensional slices
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Moneypenny", "jenny", "Guiness", "Wolverine"}
	fmt.Println(jb)
	fmt.Println(jm)
	xxs := [][]string{jb, jm}
	fmt.Printf("Subarray has pointer at %p \n", &xxs)
	//exercise
	icecream := []string{
		"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
		"Browned ButterCookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
		"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar",
		"Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy ChocolatecPeanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough",
		"RaspberrySorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	fmt.Printf("The length of the slice is %v \n", len(icecream))

	// for i, v := range icecream {
	//
	// fmt.Printf("The value at index %v is %v \n", i, v)
	// }
}
