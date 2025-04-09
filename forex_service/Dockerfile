# Stage 1: Build the Go application
FROM golang:1.18-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod tidy

# Copy the source code
COPY . .

# Set the working directory to the "cmd" folder where the main.go is located
WORKDIR /app/cmd

# Build the application binary
RUN go build -o forex_service .

# Expose the port the service will run on
EXPOSE 8082

# Run the service
CMD ["./forex_service"]
