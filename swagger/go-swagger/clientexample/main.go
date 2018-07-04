package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"github.com/posener/meetups/swagger/go-swagger/client"
	"github.com/posener/meetups/swagger/go-swagger/client/pet"
)

var kind = flag.String("kind", "", "filter by kind")

func main() {
	flag.Parse()
	c := client.Default
	params := &pet.ListParams{Context: context.Background()}
	if *kind != "" {
		params.Kind = kind
	}
	pets, err := c.Pet.List(params)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range pets.Payload {
		fmt.Printf("\t%d Kind=%v Name=%v\n", p.ID, p.Kind, *p.Name)
	}
}