# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17-buster AS build
ADD . /go/src/go-inception-payment-scb
WORKDIR /go/src/go-inception-payment-scb
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go build -o /go-inception-payment-scb

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-inception-payment-scb /go-inception-payment-scb

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-inception-payment-scb"]