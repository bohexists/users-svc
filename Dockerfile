# Use the official Go image for building the application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to install dependencies
COPY go.mod go.sum ./

# Install the dependencies
RUN go mod download

# Copy the rest of the project files into the container
COPY . .

# Build the Go application
RUN go build -o users-svc cmd/main.go

# Create a minimal image for running the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder phase
COPY --from=builder /app/users-svc .

# Copy the views and config directories to the container
COPY --from=builder /app/views ./views
COPY --from=builder /app/config ./config

# Expose the port that the application will listen on
EXPOSE 8080

# Run the application
CMD ["./users-svc"]