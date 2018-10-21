FROM golang:1.11-stretch

RUN mkdir -p /go/src/github.com/lucasrenan/lucas-go
ENV GOPATH /go

COPY . /go/src/github.com/lucasrenan/lucas-go
WORKDIR /go/src/github.com/lucasrenan/lucas-go
