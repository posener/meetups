package main

import (
	"io"
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

type flushWriter struct {
	w io.Writer
}

func (fw flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	fw.w.Write([]byte("wow"))
	// Flush - send the buffered written data to the client
	if f, ok := fw.w.(http.Flusher); ok {
		f.Flush()
	}
	return
}

func handle(w http.ResponseWriter, r *http.Request) {
	// First flash response headers
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
	// Copy from the request body to the response writer and flush
	// (send to client)
	io.Copy(flushWriter{w: w}, r.Body)
}
