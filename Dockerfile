# syntax=docker/dockerfile:1.4
FROM golang:1.23.8-alpine AS builder

WORKDIR /posts

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY .. .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o bin/main ./cmd/posts/main.go

CMD ["/posts/bin/main"]
