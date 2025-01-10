FROM golang:1.23.3 AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire backend source code
COPY backend/ .

# Set working directory to cmd folder
WORKDIR /app/cmd

# Build the Go binary
RUN go build -o /app/main .

# Use a lightweight runtime image
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/main .
RUN apt-get update && apt-get install -y libc6 && rm -rf /var/lib/apt/lists/*

EXPOSE 8080
CMD ["./main"]
