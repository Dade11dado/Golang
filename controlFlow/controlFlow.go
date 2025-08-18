package main

import (
	"fmt"
	"math/rand"
	"time"
)

// FUNCITON INIT RUNS BEFORE EVERYTHING
func init() {
	fmt.Println("Run before everythin")
}

func main() {
	//SEQUENTIAL CODE
	fmt.Println("This is the first statement to run")
	fmt.Println("This is the second statement to run")
	x := 40
	y := 5
	fmt.Printf("X = %v and Y = %v", x, y)

	//CONDITIONAL CODE
	//If statement and switch statement

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("Equal to or greater than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("greater than the meaning of life")
	}

	if x > 30 && y > 42 {
		fmt.Println("They are bot evaluated")
	}

	//STATEMENT STATEMENT

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and is greater or equal to x which is %v", z, x)
	} else {
		fmt.Printf("z is %v and is less than x that is %v", z, x)
	}

	//COMMA OK
	timeZone := make(map[string]interface{})
	timeZone["House"] = 12
	tz := "Home"
	//if timeZOne[tz] exist second is set to value and ok to true,
	//otherwise second is set to 0 and ok to false
	if seconds, ok := timeZone[tz]; ok {
		fmt.Println(seconds)
	}

	//SWITCH STATEMENT
	//switch can be used as if if the case equals true
	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL TO 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("this is the default case for x")
	}

	//classic switch
	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("this is the default case for x")
	}

	// no default fallthrough
	//If case is reached switch exits, with fallthrough
	//it continues anyway with all the others
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}

	// no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of ALL OF THE fallthrough statements: x is 43")
		fallthrough
	default:
		fmt.Println("printed because of ALL OF THE fallthrough statements: this is the default case for x")
	}

	//creating channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// get an int64, convert it to type time.Duration, then use it in time.Sleep
	// Int63n returns an int64
	// type duration's underlying type is int64
	// time.Sleep takes type duration
	// time.Millisecond is a constant in the time package
	// https://pkg.go.dev/time#pkg-constants
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))
	// fmt.Printf("%v \t %T\n", d1, d1)
	// time.Sleep(d1 * time.Millisecond)
	// fmt.Println("Done")

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// A "select" statement chooses which of a set of possible send or receive operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations.
	// https://go.dev/ref/spec#Select_statements
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

	// LOOPS / INTERATIVE
	// for statements

	/*
		The Go for loop is similar to—but not the same as—C's.
		It unifies for and while and there is no do-while.
		There are three forms, only one of which has semicolons.

		// Like a C for
		for init; condition; post { }

		// Like a C while
		for condition { }

		// Like a C for(;;)
		for { }

	*/
	// https://go.dev/doc/effective_go#for

	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v \t first  for loop\n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	// break
	// takes you out of the loop
	for {
		fmt.Printf("y is %v \t\t third  for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue
	// takes to next iteration
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}

	// nested loops
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}

	// for range loop
	// data structures - slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for range loop
	// data structures - map
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}

}
