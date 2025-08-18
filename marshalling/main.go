package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	p1 := person{
		First: "Davide",
		Last:  "Vavassori",
		Age:   28,
	}

	p2 := person{
		First: "Maurizio",
		Last:  "Ferrari",
		Age:   35,
	}

	people := []person{p1, p2}

	p3 := person{}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	abs := []byte(bs)

	err1 := json.Unmarshal(abs, &p3)

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(p3)
}
