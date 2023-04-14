#!/usr/bin/make -f

test: fmt
	go test -cover -timeout=1s -race ./...

fmt:
	go fmt ./... && go mod tidy

install: test
	go install github.com/mdwhatcott/valign/cmd/...
