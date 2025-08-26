package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	State string
	Name  string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	b := sage{
		Name:  "Buddha",
		State: "India",
	}

	g := sage{
		Name:  "Gandhi",
		State: "India",
	}
	m := sage{
		Name:  "MLK",
		State: "America",
	}
	f := car{
		Manufacturer: "Mercedez",
		Model:        "Gla",
		Doors:        3,
	}

	c := car{
		Manufacturer: "Fiat",
		Model:        "Panda",
		Doors:        3,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
