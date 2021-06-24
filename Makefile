#!make
include .env.dev
export

# Start
.PHONY: build
build:
	rm -rf build && mkdir build && go build -o build -v ./cmd/server

.PHONY: run
run:
	./build/server

# Test
.PHONY: test
test:
	go test -cover -v -race -timeout 30s ./...


.DEFAULT_GOAL := build