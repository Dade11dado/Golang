package main

import (
	"log"
	"os"
	"text/template"
)

var tp *template.Template

func init() {
	tp = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	err := tp.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Release self focus; embrace other focus")
	if err != nil {
		log.Fatal(err)
	}

}
