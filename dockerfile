# Step 1: Use official Go image to build the binary
FROM golang:1.22 AS builder

# Set working directory inside container
WORKDIR /app

# Copy source code
COPY test.go .

# Build the Go binary
RUN go build -o test-app test.go

# Step 2: Use a lightweight image for running
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/test-app .

# Command to run the app
CMD ["./test-app"]
