package main

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httptest"

	"fmt"

	"github.com/gorilla/websocket"
	"github.com/posener/wstest-meetup/echo"
)

func NewTestDialer(h http.Handler) *websocket.Dialer {
	client, server := net.Pipe()

	go func() {
		req, err := http.ReadRequest(bufio.NewReader(server))
		if err != nil {
			return
		}
		rec := &recorder{conn: server}
		h.ServeHTTP(rec, req)
	}()

	return &websocket.Dialer{
		NetDial: func(network, addr string) (net.Conn, error) {
			return client, nil
		},
	}
}

// RECORDER START OMIT
// recorder implements http.ResponseWriter and http.Hijacker interface
type recorder struct {
	httptest.ResponseRecorder
	conn net.Conn
}

// Hijack implements the Hijacker interface
func (r *recorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	rw := bufio.NewReadWriter(bufio.NewReader(r.conn), bufio.NewWriter(r.conn))
	return r.conn, rw, nil
}

// WriteHeader is part of the ResponseWriter interface
func (r *recorder) WriteHeader(code int) {
	resp := http.Response{StatusCode: code}
	resp.Write(r.conn)
}

// RECORDER END OMIT

func main() {
	d := NewTestDialer(http.HandlerFunc(echo.Handler))
	c, resp, err := d.Dial("ws://whatever/ws", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp status code:", resp.StatusCode)
	err = c.WriteJSON("test")
	if err != nil {
		panic(err)
	}
	var ret string
	c.ReadJSON(&ret)
	fmt.Println("read:", ret)
}
