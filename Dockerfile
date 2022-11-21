# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN go install github.com/zzfima/Golang-Nuget-info
RUN go install github.com/gorilla/mux

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /website-go-nudget-info

EXPOSE 8080

CMD [ "/website-go-nudget-info" ]
