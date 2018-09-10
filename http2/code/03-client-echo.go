package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
)

const url = "https://localhost:8000"

func main() {
	// Create a pool with the server certificate since it is not signed by a known CA
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{RootCAs: caCertPool}

	// Use the proper transport in the client
	client := &http.Client{Transport: &http2.Transport{TLSClientConfig: tlsConfig}}

	// Create a pipe - an object that implements `io.Reader` and `io.Writer`. // Whatever is written to the writer part will be read by the reader part.
	pr, pw := io.Pipe()

	// Create an `http.Request` and set its body as the reader part of the
	// pipe - after sending the request, whatever will be written to the pipe,
	// will be sent as the request body.
	// This makes the request content dynamic, so we don't need to define it
	// before sending the request.
	req, err := http.NewRequest(http.MethodPut, url, ioutil.NopCloser(pr))
	if err != nil {
		log.Fatal(err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got: %d", resp.StatusCode)

	// Run a loop which writes every second to the writer part of the pipe
	// the current time.
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Fprintf(pw, "It is now %v\n", time.Now())
		}
	}()

	// Copy the server's response to stdout.
	_, err = io.Copy(os.Stdout, resp.Body)
	log.Fatal(err)
}
