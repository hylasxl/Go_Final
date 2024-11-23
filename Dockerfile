# Use the official Golang image
FROM golang:1.23.2-alpine3.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o grpc-server .

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/grpc-server .

# Copy the Envoy configuration
COPY envoy.yaml .

# Expose the port that the service will run on
EXPOSE 50051
# Command to run the executable
CMD ["./grpc-server"]