# ----------------------
# 1. Build Stage
# ----------------------
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o nextmed-api ./cmd/server

# ----------------------
# 2. Run Stage
# ----------------------
FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates

# Copy built binary
COPY --from=builder /app/nextmed-api .

# Copy .env (so godotenv can load it inside Docker)
COPY .env .env

EXPOSE 7777

ENTRYPOINT ["./nextmed-api"]
