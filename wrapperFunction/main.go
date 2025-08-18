package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("readfile in maain %s \n", err)
	}

	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(fn string) ([]byte, error) {
	xb, err := os.ReadFile(fn)
	if err != nil {
		return nil, fmt.Errorf(" \t error in read file %s \n", err)
	}
	return xb, nil
}
