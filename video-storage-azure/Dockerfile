FROM golang:1.16-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    git \
    && update-ca-certificates

FROM base AS builder
WORKDIR /app

COPY go.* /app
RUN go mod download \
    && go mod verify

COPY . /app

RUN go build -o video-storage -a .

FROM alpine:3 as prod

WORKDIR /srv

COPY .env .env
COPY --from=builder /app/video-storage /srv/video-storage

ENTRYPOINT ["/srv/video-storage"]
