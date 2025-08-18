package main

import "fmt"

type customError struct {
	err string
}

func (c customError) Error() string {
	return fmt.Sprintf("There was an error: %v
	", c.err)
}

func main() {
	c := customError{
		"OOOOOOH",
	}
	foo(c)
}

func foo(er error) {
	fmt.Println(er)
}
