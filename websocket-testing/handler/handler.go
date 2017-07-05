package main

import (
	"fmt"
	"net/http"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>I love %s!</h1>", r.URL.Path[1:])
	})
	http.HandleFunc("/", h)
	http.ListenAndServe(":8080", nil)
}
