.DEFAULT_GOAL := build

.PHONY: tidy generate fmt vet build clean
tidy:
	go mod tidy

generate: tidy 
	go generate ./...

fmt: generate
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean