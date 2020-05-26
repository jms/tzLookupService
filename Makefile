#!make
GO_CMD = $(shell which go)
SHELL = /usr/bin/bash

build:
	GOOS=linux GOARCH=amd64 $(GO_CMD) build -o bin/tzLookupService-linux-amd64 main.go

run:
	$(GO_CMD) run main.go

build-all:
	GOOS=freebsd GOARCH=amd64 $(GO_CMD) build -o bin/tzLookupService-freebsd-amd64 main.go
	GOOS=darwin GOARCH=amd64 $(GO_CMD) build -o bin/tzLookupService-darwin-amd64 main.go
	GOOS=linux GOARCH=amd64 $(GO_CMD) build -o bin/tzLookupService-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 $(GO_CMD) build -o bin/tzLookupService-windows-amd64 main.go

clean:
	rm -fv bin/*
	rm -frv dist
	rm -v timezone.snap.db

test:
	$(GO_CMD) test

create-tz-db:
	# data https://github.com/evansiroky/timezone-boundary-builder/releases
	curl -LO -C - https://github.com/evansiroky/timezone-boundary-builder/releases/download/2020a/timezones.geojson.zip
	unzip timezones.geojson.zip
	go run cmd/timezone.go -json "dist/combined.json" -db=timezone -type=boltdb