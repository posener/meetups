package main

import (
	"net/http"

	"fmt"
	"net/http/httptest"

	"github.com/gorilla/websocket"
	"github.com/posener/wstest-meetup/echo"
)

func main() {
	server := httptest.NewServer(http.HandlerFunc(echo.Handler))
	defer server.Close()

	d := websocket.Dialer{}
	c, resp, err := d.Dial("ws://"+server.Listener.Addr().String()+"/ws", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	fmt.Println("resp status code:", resp.StatusCode)
	c.WriteJSON("test")
	var ret string
	c.ReadJSON(&ret)
	fmt.Println("read:", ret)
}
