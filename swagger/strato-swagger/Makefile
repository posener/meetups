swagger=docker run --rm -e GOPATH=$(GOPATH):/go -v $(PWD):$(PWD) -w $(PWD) -u $(shell id -u):$(shell id -u) stratoscale/swagger:v1.0.14

clean:
	rm -rf restapi client models

generate-server:
	$(swagger) generate server

generate-client:
	$(swagger) generate client
