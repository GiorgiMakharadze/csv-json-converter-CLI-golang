.DEFAULT_GOAL := build

fmt:
	go fmt ./...

lint: fmt
	golint ./...

vet: fmt
	go vet ./...
	shadow ./...
build: vet
	go build -o main

.PHONY: fmt lint vet build