package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

type home string
type dog string
type me string

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", "Home")
	if err != nil {
		log.Fatal(err)
	}
}
func (h dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", "Dog")
	if err != nil {
		log.Fatal(err)
	}
}
func (h me) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", "Me")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var home home
	var dog dog
	var me me

	http.Handle("/", home)
	http.Handle("/dog/", dog)
	http.Handle("/me/", me)

	http.ListenAndServe(":8080", nil)

}
