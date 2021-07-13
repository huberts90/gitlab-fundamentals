FROM golang:1.15.8-alpine as dev
COPY . /usr/src/app
ENV GO111MODULE=on
ENV CGO_ENABLED=0
WORKDIR /usr/src/app

FROM dev as build
RUN mkdir -p release
RUN go build  -o release/main cmd/gitlab-fundamentals/cmd/main.go

FROM build as test
RUN go test ./internal/...
