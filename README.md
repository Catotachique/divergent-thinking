# divergent-thinking

1. Microservices Architecture Overview

In a microservices architecture, each service is:

    Independent: Each service can be developed, deployed, and scaled independently.

    Single Responsibility: Each service is designed around a specific business function or domain (e.g., user service, order service, inventory service).

    Communicates over HTTP or Messaging: Services typically communicate using REST APIs or through message queues (e.g., RabbitMQ, Kafka).


### Folder structure
/divergent-thinking
  /user-service
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

  /order-service
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


## API
https://jsonplaceholder.typicode.com/