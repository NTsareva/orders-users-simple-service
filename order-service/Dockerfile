FROM golang:1.22 as builder

WORKDIR /app

RUN apt-get update && apt-get install -y git

COPY . .

RUN go mod download
RUN go build -o order-service ./cmd/main.go

FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/order-service /app/order-service

CMD ["/app/order-service"]