FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o user-service ./app/main.go

FROM ubuntu:latest

COPY --from=builder /app/user-service /app/user-service

CMD ["/app/user-service"]