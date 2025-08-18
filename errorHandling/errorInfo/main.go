package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("Norgate math redux, square root of negative number")
		return 0, norgateMathError{
			"50.22",
			"59.3",
			nme}
	}
	return 42, nil
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("A norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}
