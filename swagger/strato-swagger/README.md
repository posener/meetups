# Stratoscale/swagger Demo

```bash
alias swagger='docker run --rm -e GOPATH=${GOPATH}:/go -v $(pwd):$(pwd) -w $(pwd) -u $(id -u):$(id -u) stratoscale/swagger:v1.0.14'
swagger generate server
go run ./main.go
curl -i http://127.0.0.1:8080/
curl -i http://127.0.0.1:8080/swagger.json
curl -i http://127.0.0.1:8080/api/pets
```

Set in main.go Pet -> PetImplemented

```bash
go run ./main.go
curl -s  http://127.0.0.1:8080/api/pets | jq
curl -s  http://127.0.0.1:8080/api/pets?kind=dog | jq
```

## Generate Client

```bash
swagger generate client
go run ./clientexample/main.go
go run ./clientexample/main.go -kind dog
```