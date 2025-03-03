FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/main ./cmd/main.go

RUN go build -o bin/worker ./cmd/consumer/worker.go

FROM golang:1.23

WORKDIR /root/

COPY --from=builder /app/bin/main .
COPY --from=builder /app/bin/worker .
COPY --from=builder /app/.env .env
COPY --from=builder /app/start.sh /start.sh

RUN chmod +x /start.sh

EXPOSE 8080

CMD ["/start.sh"]
