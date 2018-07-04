# go-swagger Demo

## Generate Server

```bash
./swagger generate server
go run ./cmd/minimal-pet-store-example-server/main.go --port 8080
curl -i http://127.0.0.1:8080/
curl -i http://127.0.0.1:8080/swagger.json
curl -i http://127.0.0.1:8080/api/pets
```

Edit `restapi/configure_minimal_pet_store_example.go`

```go
var petList = []*models.Pet{
	{ID: 0, Name: swag.String("Bobby"), Kind: "dog"},
	{ID: 1, Name: swag.String("Lola"), Kind: "cat"},
	{ID: 2, Name: swag.String("Bella"), Kind: "dog"},
	{ID: 3, Name: swag.String("Maggie"), Kind: "cat"},
}

[...]

func configureAPI(api *operations.MinimalPetStoreExampleAPI) http.Handler {
	[...]
	api.PetListHandler = pet.ListHandlerFunc(func(params pet.ListParams) middleware.Responder {
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
	})
	[...]
}
```

```bash
go run ./cmd/minimal-pet-store-example-server/main.go --port 8080
curl -s  http://127.0.0.1:8080/api/pets | jq
curl -s  http://127.0.0.1:8080/api/pets?kind=dog | jq
```

## Generate Client

```bash
./swagger generate client
go run ./clientexample/main.go
go run ./clientexample/main.go -kind dog
```