# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ENV GO111MODULE=on
ENV GOFLAGS=-mod=readonly

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY templates/ ./templates/
RUN go mod download

COPY *.go ./

RUN go build -o /website-go-nudget-info

EXPOSE 8080

CMD [ "/website-go-nudget-info" ]
