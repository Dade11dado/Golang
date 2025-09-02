package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type Page struct {
	Title   string
	Heading string
	Input   string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}
func main() {
	home := Page{
		Title:   "Nothing escaped",
		Heading: "Nothing is escaped with text/template",
		Input:   `<script>alert("Yo bitch");</script>`,
	}

	err := tpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatal(err)
	}
}
