FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./bot/cmd/main.go
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/config.yml .
RUN sed -i 's|http://localhost:8080|http://api:8080|' config.yml
COPY --from=builder /app/.env .
COPY --from=builder /app/main .
CMD ["./main"]