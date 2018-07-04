package main

import (
	"context"

	"github.com/Stratoscale/golib/log"
	"github.com/posener/meetups/swagger/strato-swagger/client"
)

func main() {
	var (
		cl   = client.New(client.Config{})
		user = PetUser{Pet: cl.Pet}
		ctx  = context.Background()
	)

	err := user.Duplicate(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
