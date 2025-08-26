package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("/Users/davidevavassori/Desktop/Golang/WEBDEVELOPEMENT/011_composition/004_method/tpl.html"))
}

func main() {
	p := person{
		Name: "Davide",
		Age:  56,
	}

	err := tpl.Execute(os.Stdout, p)

	if err != nil {
		log.Fatal(err)
	}
}
