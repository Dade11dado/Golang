package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("this is the book ", b.title)
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM", s.String())
}

type count int

func (c count) String() string {
	return fmt.Sprint("The numver is ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "West with the night",
	}

	var c count = 42

	logInfo(b)
	logInfo(c)
}
