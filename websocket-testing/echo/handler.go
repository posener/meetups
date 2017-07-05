package echo

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func Handler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	var message string
	for {
		err = c.ReadJSON(&message)
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
		err = c.WriteJSON(message)
		if err != nil {
			log.Println("write:", err)
			return
		}
	}
}
