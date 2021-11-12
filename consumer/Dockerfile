# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY member/*.go ./

RUN go build -o /consumer

CMD [ "/consumer" ]