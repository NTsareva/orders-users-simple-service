FROM golang:1.22 AS builder

ARG USE_REPLACEMENT=true

WORKDIR /app

COPY ./cmd/order-service/go.mod ./cmd/order-service/go.sum ./


RUN go mod download

COPY ./cmd/order-service/ ./
RUN go build -o order-service ./app/main.go

FROM debian:bullseye-slim
WORKDIR /app
RUN apt-get update && apt-get install -y libc6

COPY --from=builder /app/order-service ./
COPY ./cmd/order-service/configs /app/configs

CMD ["./order-service"]