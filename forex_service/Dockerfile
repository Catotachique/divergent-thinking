# Stage 1: Build the Go application
FROM golang:1.18-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application binary
RUN go build -o forex_service ./cmd

# Stage 2: Create a minimal image to run the service
FROM ubuntu:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/forex_service /app/

# Expose the port the service will run on
EXPOSE 8082

# Run the service
CMD ["./forex_service"]
