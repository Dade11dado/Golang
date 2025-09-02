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

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(w, "index.html", r.PostForm)

}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
