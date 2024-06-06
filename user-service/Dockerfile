FROM golang:1.22 as builder

WORKDIR /app

# Установите git с помощью apt-get
RUN apt-get update && apt-get install -y git

COPY . .

RUN go mod download
RUN go build -o user-service ./cmd/main.go

FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/user-service /app/user-service

CMD ["/app/user-service"]