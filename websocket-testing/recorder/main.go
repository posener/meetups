package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
})

func main() {
	rec := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "http://example.com/Stratoscale", nil)
	handler.ServeHTTP(rec, request)
	fmt.Println(rec.Body.String())
}
