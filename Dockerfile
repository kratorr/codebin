FROM golang:alpine AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/codebin ./cmd/main.go

FROM alpine

WORKDIR /app

COPY templates/ /app/templates
COPY --from=builder /app/codebin /app/codebin-test

ENTRYPOINT ["./codebin-test"]