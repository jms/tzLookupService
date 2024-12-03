# syntax=docker/dockerfile:1
FROM golang:1.23-alpine as build-stage
WORKDIR /app
COPY go.mod go.sum ./
COPY *.go ./

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -v -o tzLookupService main.go

FROM redhat/ubi9-minimal

ENV GIN_MODE=release

WORKDIR /app

COPY --from=build-stage /app/tzLookupService ./
COPY timezone.snap.db ./

EXPOSE 8080

CMD ["/app/tzLookupService"]
