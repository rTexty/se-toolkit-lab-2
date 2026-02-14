# Stage 1: Build the Go application
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy the Go module files if they exist (go.mod and go.sum)
COPY go.mod go.sum* ./
# Download dependencies (if any)
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# CGO_ENABLED=0 ensures a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a minimal image for running the application
FROM alpine:3.19

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
