build:
	go build

generate:
	go generate ./...

install:
	go get ./..
	go mod vendor 

test:
	go test -v ./...

install_dev_dep:
	go install github.com/golang/mock/mockgen@v1.6.0
