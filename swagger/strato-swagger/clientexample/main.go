package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/posener/meetups/swagger/strato-swagger/client"
	"github.com/posener/meetups/swagger/strato-swagger/client/pet"
)

var (
	kind   = flag.String("kind", "", "Filter by kind")
	urlStr = flag.String("url", "http://localhost:8080/api", "Custom url")
)

func main() {
	flag.Parse()

	u, err := url.Parse(*urlStr)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go catchSignal(cancel)

	c := client.New(client.Config{URL: u})
	params := &pet.ListParams{Context: context.Background()}
	if *kind != "" {
		params.Kind = kind
	}

	//log.Println("Sleeping")
	//time.Sleep(time.Second * 3)
	//log.Println("Continueing")

	pets, err := c.Pet.List(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range pets.Payload {
		fmt.Printf("\t%d Kind=%v Name=%v\n", p.ID, p.Kind, *p.Name)
	}
}
func catchSignal(cancel context.CancelFunc) {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig
	log.Println("Cancelling due to interrupt")
	cancel()
}
