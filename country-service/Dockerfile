FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /app/main .
COPY .env .env

EXPOSE 1323

CMD ["./main"]
