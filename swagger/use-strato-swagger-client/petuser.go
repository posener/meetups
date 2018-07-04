package main

import (
	"fmt"
	"github.com/posener/meetups/swagger/strato-swagger/client/pet"
	"context"
)

type PetUser struct {
	Pet pet.API
}

func (pu *PetUser) Duplicate(ctx context.Context) error {
	pets, err := pu.Pet.List(ctx, nil)
	if err != nil {
		return fmt.Errorf("listing pets: %v", err)
	}
	for _, p := range pets.Payload {
		id := p.ID
		p.ID = 0
		_, err := pu.Pet.Create(ctx, &pet.CreateParams{Pet: p})
		if err != nil {
			return fmt.Errorf("duplicating pet %d: %v", id, err)
		}
	}
	return nil
}
