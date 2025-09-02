package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Doggy")
	case "/cat":
		io.WriteString(w, "Cat")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
