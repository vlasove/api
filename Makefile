.PHONY: build, lint, test

build:
	go build -v ./cmd/apiserver

lint:
	golint ./... && golangci-lint run

test:
	go test -v -race ./...

.DEFAULT_GOAL := build