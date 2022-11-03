# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ADD . /app

WORKDIR /app/cmd

RUN apk add --no-cache bash
RUN CGO_ENABLED=0 go build -o /simplest-possible-thing

WORKDIR /app

CMD [ "/simplest-possible-thing" ]