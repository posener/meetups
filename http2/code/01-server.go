package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a server on port 8000
	// Exactly how you would run an HTTP1.1 server
	srv := &http.Server{Addr: ":8000", Handler: http.HandlerFunc(handle)}

	// Start the server with TLS, since we are running HTTP2 it must be run with TLS.
	// Exactly how you would run an HTTP1.1 server with TLS connection.
	log.Printf("Go to https://localhost:8000")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Log the request protocol
	log.Printf("Got connection (%s): %s", r.Proto, r.URL)

	// Send response body
	w.Write([]byte("Hello"))
}
