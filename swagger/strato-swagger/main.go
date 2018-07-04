package main

import (
	"net/http"
	"log"

	"github.com/posener/meetups/swagger/strato-swagger/restapi"
	"github.com/posener/meetups/swagger/strato-swagger/internal"
)

func main() {
	h, err := restapi.Handler(restapi.Config{
		PetAPI: &internal.PetImplemented{},
		Logger: log.Printf,
	})
	if err != nil {
		panic(err)
	}
	log.Println("Serving, see http://127.0.0.1:8080")
	err = http.ListenAndServe(":8080", h)
	if err != nil {
		panic(err)
	}
}