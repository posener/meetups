package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>I love %s!</h1>", r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}
