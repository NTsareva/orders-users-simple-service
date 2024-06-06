FROM golang:1.22 as builder

WORKDIR /app

RUN apk add --no-cache git

COPY cmd/user-service/go.mod cmd/user-service/go.sum ./

RUN go mod download

COPY cmd/user-service/ .

RUN go build -o user-service ./app/main.go

FROM debian:latest

WORKDIR /app

RUN apk add glibc-2.34-r0.apk

RUN sleep 10; echo "sleeping"

COPY --from=builder /app/user-service /app/user-service
COPY cmd/user-service/configs /app/configs

CMD ["/app/user-service"]