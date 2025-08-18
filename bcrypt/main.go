package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "passwrod123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	fmt.Println(bs)
}
