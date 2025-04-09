# Stage 1: Build the Go application
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod ./

# Download the Go modules
RUN go mod tidy

# Copy the source code
COPY . .

# Set the working directory to the "cmd" folder where the main.go is located
WORKDIR /app/cmd

# Build the application binary
RUN go build -o /app/forex_service .


# Stage 2: Use a lightweight base image to run the app
FROM ubuntu:latest

WORKDIR /root/

RUN apt update && apt install -y ca-certificates

# Copy the compiled binary from the builder
COPY --from=builder /app/forex_service .

# Expose the service port
EXPOSE 8082

# Run the binary
CMD ["./forex_service"]
