
clean:
	rm -rf restapi client models cmd

swagger:
	curl -L https://github.com/go-swagger/go-swagger/releases/download/0.14.0/swagger_linux_amd64 -o $@
	chmod +x $@

generate: swagger
	./swagger generate server
	./swagger generate client