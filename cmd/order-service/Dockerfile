FROM golang:1.22-alpine as builder

WORKDIR /app

RUN apk add --no-cache git

COPY cmd/order-service/go.mod cmd/order-service/go.sum ./

RUN go mod download

COPY cmd/order-service/ .

RUN go build -o order-service ./app/main.go

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates \
    && wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub \
    && wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.34-r0/glibc-2.34-r0.apk \
    && apk add glibc-2.34-r0.apk \
    && rm -rf glibc-2.34-r0.apk

COPY --from=builder /app/order-service /app/order-service
COPY cmd/order-service/configs /app/configs

CMD ["/app/order-service"]