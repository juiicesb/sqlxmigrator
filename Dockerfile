FROM golang:1.18.1-alpine3.15

WORKDIR /sqlxmigrator
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 
