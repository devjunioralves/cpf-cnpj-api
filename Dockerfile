FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o cmd/main ./cmd

FROM golang:1.21

WORKDIR /root/

COPY --from=builder /app/cmd/main .
COPY --from=builder /app/.env .env

EXPOSE 8080

CMD ["./main"]