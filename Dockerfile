
# Stage 1: Build the Go binary
FROM golang:1.23.2-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -o math-cli main.go

# Stage 2: Create a minimal image with the binary
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/math-cli .

# Set executable permissions
RUN chmod +x math-cli

# Define the default command to run the binary
ENTRYPOINT ["./math-cli"]
