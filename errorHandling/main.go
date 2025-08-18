package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()
	//ECAMPLE 1
	// n, err := fmt.Println("Hello")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(n)

	//EXAMPLE 2
	// var answer1, answer2, answer3 string

	// fmt.Print("Name: ")
	// _, err := fmt.Scan(&answer1)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Print("Fav Food: ")
	// _, err = fmt.Scan(&answer2)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Print("Fav Sport: ")
	// _, err = fmt.Scan(&answer3)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(answer1, answer2, answer3)

	//We can use fmt/log - log also writes in a file
	//wwe can use .panic or .fatal. .,fatal is the worst scenario
	//Set up to set the file where log will write
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	_, err = os.Open("no.file.txt")
	if err != nil {
		//This just prints to the console the error
		//fmt.Println("Err happened", err)
		//This print with hours and date
		//log.Println("err happened", err)
		//Fatalln is equal to Println followed by a call to os.Exit()
		//after writing to log
		//log.Fatalln(err)
		//stop normal executions, can call a callback
		panic(err)

	}
}
