# syntax=docker/dockerfile:1

FROM golang:1.21.5-alpine

ADD . /app

WORKDIR /app

RUN apk add --no-cache bash
RUN CGO_ENABLED=0 go build -o /event-sourcing-go

WORKDIR /app

CMD [ "/event-sourcing-go" ]