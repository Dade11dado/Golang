package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type sage struct {
	Name  string
	State string
}

type data struct {
	Wisdom []sage
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.html"))
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	b := sage{
		Name:  "Buddha",
		State: "India",
	}

	c := sage{
		Name:  "MLK",
		State: "America",
	}

	d := data{
		Wisdom: []sage{b, c},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", d)
	if err != nil {
		log.Fatal(err)
	}
}
