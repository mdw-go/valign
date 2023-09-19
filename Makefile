#!/usr/bin/make -f

VERSION := $(shell git describe)

test: fmt
	go test -cover -timeout=1s -race ./...

fmt:
	go fmt ./... && go mod tidy

install: test
	go install -ldflags="-X 'main.Version=$(VERSION)'" github.com/mdwhatcott/valign/cmd/...
