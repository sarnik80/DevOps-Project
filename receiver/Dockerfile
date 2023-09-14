FROM golang:1.18-alpine AS builder

LABEL authors="sra"

RUN mkdir app

WORKDIR /app

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY ./receiver/go.mod ./receiver/go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY ./receiver .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o result ./

FROM alpine:3.18.2

RUN mkdir final

WORKDIR /final

COPY --from=builder /app/result .


CMD [ "./result" ]