FROM golang:1.22 as builder

WORKDIR /app

# Создайте каталог для user-service
RUN mkdir -p /app/user-service

COPY . .

COPY ../.. /app/user-service

RUN apt-get update && apt-get install -y git

RUN go mod tidy
RUN go build -o order-service ./cmd/main.go

FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/order-service /app/order-service

CMD ["/app/order-service"]