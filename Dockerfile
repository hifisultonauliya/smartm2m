# Start from the official Golang image
# This image includes the Go tools necessary to compile and build Go applications
ARG GO_VERSION=1.22.1

# Use the official Golang image for the specified version as a base image
FROM golang:${GO_VERSION}-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Download and install any required dependencies
RUN go mod tidy

# Build the application
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Copy the pre-built binary file from the previous stage
COPY --from=builder /build/main /app/

# Set the working directory in the container
WORKDIR /app

# Command to run the executable
CMD ["./main"]
