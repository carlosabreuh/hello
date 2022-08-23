# Build full image
FROM golang:1.16-alpine as build

ARG VERSION="unset"

WORKDIR /go/src/github.com/nytimes/dv-interview-exercise/hello
ADD . /go/src/github.com/nytimes/dv-interview-exercise

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags="-X main.version=${VERSION}" -o /go/bin/hello

ENTRYPOINT ["/go/bin/hello"]
