package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog")
}

type hotcat int

func (d hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Cat")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
