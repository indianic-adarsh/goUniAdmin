# Step 1: Use Go 1.24 image for building the Go application
FROM golang:1.24 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the Go application source code and .env file into the container
COPY . .
COPY .env .  

# Build the Go app
RUN go build -o api ./cmd/api

# Step 2: Use a base image with a newer version of glibc (Debian Bookworm)
FROM debian:bookworm

# Install necessary libraries (glibc is included by default, but update package lists)
RUN apt-get update && apt-get install -y libc6

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary and .env file from the builder stage
COPY --from=builder /app/api .
COPY --from=builder /app/.env .  

# Expose the port the app will run on (assuming it's port 8080)
EXPOSE 8080

# Command to run the Go app
CMD ["./api"]