FROM golang:1.8.0-alpine
MAINTAINER pocka <pockawoooh@gmail.com>

ENTRYPOINT /bin/sh

WORKDIR /work

RUN apk --update add git make && \
	go get -u github.com/go-swagger/go-swagger/cmd/swagger && \
	swagger version
