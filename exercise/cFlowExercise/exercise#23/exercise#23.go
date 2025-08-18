package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x, y := rand.Intn(10), rand.Intn(10)
	fmt.Printf("X value is %v and Y value is %v", x, y)

	//with if statements

	if x < 4 && y < 4 {
		fmt.Println("they are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Printf("They are both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("X is between 4 and 6 inclusive")
	} else if y != 5 {
		fmt.Println("Y is not 5")
	} else {
		fmt.Println("None of the cases were met")
	}

	//with switch statements
	switch {
	case x < 4 && y < 4:
		fmt.Println("they are both less than 4")
	case x > 6 && y > 6:
		fmt.Printf("They are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("X is between 4 and 6 inclusive")
	case y != 5:
		fmt.Println("Y is not 5")
	default:
		fmt.Println("None of the cases were met")
	}
}
