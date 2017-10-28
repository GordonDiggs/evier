.PHONY: build test

build:
	go build -v

test:
	go test $(shell go list ./... | grep -v /vendor/)
