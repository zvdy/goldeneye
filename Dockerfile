# Use the official Golang image as the base image
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o goldeneye ./cmd/goldeneye/main.go

# Use a minimal base image to reduce the size of the final image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/goldeneye .

# Set the entrypoint to the goldeneye executable
ENTRYPOINT ["./goldeneye"]

# Default command arguments (can be overridden)
CMD ["-h"]