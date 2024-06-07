FROM golang:1.22 AS builder
WORKDIR /app

COPY ./cmd/order-service/ ./
RUN go mod download
RUN go build -o order-service ./app/main.go


FROM debian:bullseye-slim
WORKDIR /app

COPY --from=builder /app/order-service ./
COPY ./cmd/order-service/configs /app/configs

CMD ["./order-service"]