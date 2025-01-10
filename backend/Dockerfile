# Build stage
FROM golang:1.23.3 AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire backend source code
COPY backend/ ./

# Set working directory to cmd folder
WORKDIR /app/cmd

# Build the Go binary
RUN go build -o /app/main .

# Runtime stage
FROM debian:bookworm-slim

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Install tools and libraries for runtime
RUN apt-get update && apt-get install -y libc6 curl && rm -rf /var/lib/apt/lists/*

# Expose the HTTP port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
