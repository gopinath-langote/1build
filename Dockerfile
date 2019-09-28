FROM golang:1.12 AS builder

RUN mkdir -p /app
WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go install ./...

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache \
        libc6-compat

COPY --from=builder /go/bin/1build /usr/bin/1build
