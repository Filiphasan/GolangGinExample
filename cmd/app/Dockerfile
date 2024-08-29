# Build aşaması
FROM golang:1.20-alpine AS builder

# Set working directory to /app
WORKDIR /app

# Download Go modules
COPY ../../go.mod ../../go.sum ./
RUN go mod download

# Copy source code to working directory
COPY ../../ .

# Build app
WORKDIR /app/cmd/app
RUN go build -o /app

# Çalıştırma aşaması
FROM alpine:latest

# Download ca-certificates
RUN apk --no-cache add ca-certificates

# Copy build artifact from builder stage
COPY --from=builder /app /usr/local/bin/app

# Uygulamayı çalıştır
CMD ["app"]
