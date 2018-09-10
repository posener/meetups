package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
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
	pngName := fmt.Sprintf("/gopher-%d.png", rand.Int())

	// ERROR: always push
	if p, ok := w.(http.Pusher); ok {
		log.Printf("Pushing...")
		err := p.Push(pngName, nil)
		if err != nil {
			log.Printf("Failed push: %s", err)
		}
	}

	// Server according to URL path
	switch {
	case r.URL.Path == "/":
		// Serve index
		time.Sleep(5 * time.Second)
		w.Write([]byte(fmt.Sprintf(`<html><body><img src="%s"></body></html>`, pngName)))

	case strings.HasPrefix(r.URL.Path, "/gopher-"):
		// Serve image
		w.Header().Set("Content-Type", "image/x-icon")
		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, r, "gopher.png")
	}
}
