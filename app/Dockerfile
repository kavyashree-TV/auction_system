
FROM golang:1.11.13 AS builder

COPY . /auction_system/

WORKDIR /auction_system/app

RUN go build

EXPOSE 8003

CMD go run /auction_system/main.go





