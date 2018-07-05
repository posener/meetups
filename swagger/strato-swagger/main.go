package main

import (
	"log"
	"net/http"

	"github.com/posener/meetups/swagger/strato-swagger/internal"
	"github.com/posener/meetups/swagger/strato-swagger/restapi"
)

func main() {
	h, err := restapi.Handler(restapi.Config{
		// Injecting the PetAPI business logic implementer here.
		// change to internal.PetImplemented in order to get the pet list working
		PetAPI: &internal.Pet{},
		//PetAPI: &internal.PetImplemented{},
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
