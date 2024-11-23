# Stage 1: Build
FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app


# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build the application
COPY . .
RUN go build -o go_final

# Stage 2: Run
FROM alpine:3.20

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/go_final .

# Expose the port and run the application
EXPOSE 50051
CMD ["./go_final"]
