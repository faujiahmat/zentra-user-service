# ====================
# Stage 1: Builder
# ====================
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary statis
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# ====================
# Stage 2: Runtime (Alpine)
# ====================
FROM alpine:3.20

# Install SSL certs & basic tools (optional)
RUN apk add --no-cache ca-certificates bash coreutils && update-ca-certificates

# Buat working directory
WORKDIR /app

# Copy binary dari builder stage
COPY --from=builder /app/main .

# Ganti ke user non-root (lebih aman)
RUN adduser -D appuser
USER appuser

# Expose REST & gRPC ports
EXPOSE 3400 4400

# Jalankan binary
CMD ["./main"]
