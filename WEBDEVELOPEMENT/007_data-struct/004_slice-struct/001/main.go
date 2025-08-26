package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no belief",
	}
	mlk := sage{
		Name:  "MLK",
		Motto: "Hatred never ceased with heatred but with love",
	}
	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}
	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	muhamme := sage{
		Name:  "Muhammed",
		Motto: "To overcome evil with good",
	}

	sages := []sage{jesus, mlk, muhamme, gandhi, buddha}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}
}
