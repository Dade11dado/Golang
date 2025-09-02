package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Mcleod-Key", "This is from mcleod")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<h1>Anycode you want</h1>")
}
func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
