FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./bot/cmd/main.go
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/config.yml .
COPY --from=builder /app/main .
CMD ["./main"]