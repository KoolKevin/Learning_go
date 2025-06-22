.DEFAULT_GOAL := build

.PHONY:fmt vet build
tidy:
	go mod tidy

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean