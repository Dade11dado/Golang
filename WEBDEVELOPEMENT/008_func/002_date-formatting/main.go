package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fc).ParseFiles("tpl.html"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fc = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
