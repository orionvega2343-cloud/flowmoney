FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./api/cmd/main.go
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/config.yml .
RUN sed -i 's/host: localhost/host: db/' config.yml
COPY --from=builder /app/.env .
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
