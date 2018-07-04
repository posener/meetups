package internal

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/posener/meetups/swagger/strato-swagger/restapi/operations/pet"
	"github.com/go-openapi/swag"
	"github.com/posener/meetups/swagger/strato-swagger/models"
)

type Pet struct {}

func (*Pet) Create(ctx context.Context, params pet.CreateParams) middleware.Responder {
	return middleware.NotImplemented("Not implemented")
}

func (*Pet) Get(ctx context.Context, params pet.GetParams) middleware.Responder {
	return middleware.NotImplemented("Not implemented")
}

func (*Pet) List(ctx context.Context, params pet.ListParams) middleware.Responder {
	return middleware.NotImplemented("Not implemented")
}

var petList = []*models.Pet{
	{ID: 0, Name: swag.String("Bobby"), Kind: "dog"},
	{ID: 1, Name: swag.String("Lola"), Kind: "cat"},
	{ID: 2, Name: swag.String("Bella"), Kind: "dog"},
	{ID: 3, Name: swag.String("Maggie"), Kind: "cat"},
}

type PetImplemented struct {}

func (*PetImplemented) Create(ctx context.Context, params pet.CreateParams) middleware.Responder {
	return middleware.NotImplemented("Not implemented")
}

func (*PetImplemented) Get(ctx context.Context, params pet.GetParams) middleware.Responder {
	return middleware.NotImplemented("Not implemented")
}

func (*PetImplemented) List(ctx context.Context, params pet.ListParams) middleware.Responder {
	if params.Kind == nil {
		return pet.NewListOK().WithPayload(petList)
	}
	var pets []*models.Pet
	for _, pet := range petList {
		if *params.Kind == pet.Kind {
			pets = append(pets, pet)
		}
	}
	return pet.NewListOK().WithPayload(pets)
}
