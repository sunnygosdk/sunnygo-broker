FROM golang:1.24-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o brokerApp ./cmd/api

RUN chmod +x /app/brokerApp

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]
