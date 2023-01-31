generate:
	go generate ./...

install:
	go get ./..
	go mod vendor 

test:
	go test -v ./...
