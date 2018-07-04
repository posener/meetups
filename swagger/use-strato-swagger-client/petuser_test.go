package main

import (
	"context"
	"testing"

	"github.com/go-openapi/swag"
	"github.com/posener/meetups/swagger/strato-swagger/client/pet"
	"github.com/posener/meetups/swagger/strato-swagger/models"
	"github.com/stretchr/testify/assert"
)

var petList = []*models.Pet{
	{ID: 1, Kind: "dog", Name: swag.String("Bonni")},
	{ID: 2, Kind: "cat", Name: swag.String("Mitzi")},
}

func TestPetUser_Duplicate(t *testing.T) {
	var (
		m   pet.MockAPI
		pu  = PetUser{Pet: &m}
		ctx = context.Background()
	)

	m.On("List", ctx, (*pet.ListParams)(nil)).
		Return(&pet.ListOK{Payload: petList}, nil).
		Once()

	m.On("Create", ctx, &pet.CreateParams{Pet: &models.Pet{Kind: "dog", Name: swag.String("Bonni")}}).
		Return(&pet.CreateCreated{Payload: petList[0]}, nil).
		Once()

	m.On("Create", ctx, &pet.CreateParams{Pet: &models.Pet{Kind: "cat", Name: swag.String("Mitzi")}}).
		Return(&pet.CreateCreated{Payload: petList[1]}, nil).
		Once()

	err := pu.Duplicate(ctx)

	assert.Nil(t, err)
	m.AssertExpectations(t)
}
