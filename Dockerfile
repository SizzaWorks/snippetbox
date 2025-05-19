# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o snippetbox ./cmd/web

# Runtime stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/snippetbox .
COPY --from=builder /app/ui ./ui
COPY --from=builder /app/wizexercise.txt .
COPY --from=builder /app/version.txt .

EXPOSE 4000

ENTRYPOINT ["./snippetbox"]
