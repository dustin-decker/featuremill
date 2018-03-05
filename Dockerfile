# for CI testing
ARG GO_VERSION=1.9
ARG PROJECT_PATH=/go/src/github.com/dustin-decker/featuremill
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /go/src/github.com/dustin-decker/featuremill
RUN apk --no-cache add git
RUN go get -u github.com/golang/dep/cmd/dep
COPY ./ ${PROJECT_PATH}
RUN export PATH=$PATH:`go env GOHOSTOS`-`go env GOHOSTARCH` \
    && dep ensure \
    && go test $(go list ./... | grep -v /vendor/)