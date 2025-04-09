# divergent-thinking

1. Microservices Architecture Overview

In a microservices architecture, each service is:

    Independent: Each service can be developed, deployed, and scaled independently.

    Single Responsibility: Each service is designed around a specific business function or domain (e.g., user service, order service, inventory service).

    Communicates over HTTP or Messaging: Services typically communicate using REST APIs or through message queues (e.g., RabbitMQ, Kafka).

### Commands
go mod init forex_service
go build -o forex_service

### Folder structure
/divergent-thinking
  /forex-service
    /cmd
      main.go                # Entry point for the service
    /internal
      /handler
        user_handler.go      # HTTP request handlers
      /model
        user.go              # User model
      /service
        user_service.go      # Business logic
    /api
      user.proto             # gRPC definitions (if you're using gRPC)
    Dockerfile               # Dockerfile to build the container for this service
    go.mod                   # Go module dependencies
    go.sum                   # Go module checksum

  /bet-service
    /cmd
      main.go                # Entry point for the service
    /internal
      /handler
        order_handler.go     # HTTP request handlers
      /model
        order.go             # Order model
      /service
        order_service.go     # Business logic
    /api
      order.proto            # gRPC definitions (if you're using gRPC)
    Dockerfile               # Dockerfile for the order-service container
    go.mod                   # Go module dependencies
    go.sum                   # Go module checksum

  /api-gateway
    /cmd
      main.go                # API gateway entry point (proxy for microservices)
    /handler
      gateway_handler.go     # Gateway logic to forward requests to services
    Dockerfile               # Dockerfile for the API gateway container
    go.mod                   # Go module dependencies
    go.sum                   # Go module checksum

  /docker-compose.yml        # Docker Compose file to manage all the containers

#### Explanation of the Folder Structure

/user-service, /order-service: These are the directories for each microservice (e.g., user, order). Each service has:

/cmd: This folder contains the entry point of the application (the main.go file).

/internal: This folder holds the service logic. It's broken down into subfolders like:

/handler: Contains the HTTP request handlers (e.g., user_handler.go).

/model: Defines the data models (e.g., user.go, order.go).

/service: Contains the core business logic (e.g., user_service.go or order_service.go).

Dockerfile: A file to build the Docker image for that service.

go.mod and go.sum: Go modules files to manage dependencies.

/api-gateway: This is the API Gateway service. It might have a similar structure as the other services, but its primary purpose is to route incoming requests to the appropriate microservices.

/docker-compose.yml: This is the Docker Compose file that will manage all the microservices and their containers, including dependencies like databases, messaging systems, or any other external services.

## API
https://jsonplaceholder.typicode.com/