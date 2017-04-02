FROM golang:1.8.0-alpine
MAINTAINER pocka <pockawoooh@gmail.com>

WORKDIR /work

RUN apk --update add git make gcc && \
	go get -u github.com/go-swagger/go-swagger/cmd/swagger && \
	go get github.com/go-openapi/runtime && \
	go get github.com/go-openapi/swag && \
	go get github.com/go-openapi/loads && \
	go get github.com/jessevdk/go-flags && \
	go get github.com/tylerb/graceful && \
	go get github.com/gorilla/context && \
	go get github.com/docker/go-units && \
	go get github.com/go-openapi/validate && \
	swagger version
