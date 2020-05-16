#!make
GO_CMD = $(shell which go)
SHELL = /usr/bin/bash

build:
	GOOS=linux GOARCH=amd64 $(GO_CMD) build -o bin/tzLookupService-linux-amd64 main.go
	cp timezone.snap.db bin/

run:
	$(GO_CMD) run main.go

build-all:
	GOOS=freebsd GOARCH=amd64 $(GO_CMD) build -o bin/tzLookupService-freebsd-amd64 main.go
	GOOS=darwin GOARCH=amd64 $(GO_CMD) build -o bin/tzLookupService-darwin-amd64 main.go
	GOOS=linux GOARCH=amd64 $(GO_CMD) build -o bin/tzLookupService-linux-amd64 main.go
	OOS=windows GOARCH=amd64 $(GO_CMD) build -o bin/tzLookupService-windows-amd64 main.go
	cp timezone.snap.db bin/

clean:
	rm -fv bin/*

test:
	$(GO_CMD) test

