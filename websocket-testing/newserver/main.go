package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
})

func main() {
	server := httptest.NewServer(handler)
	defer server.Close()
	response, _ := http.Get(server.URL + "/Go")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
