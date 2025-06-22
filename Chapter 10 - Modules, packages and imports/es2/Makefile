.DEFAULT_GOAL := build

.PHONY: tidy fmt vet build
tidy:
	go mod tidy

fmt: tidy
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean