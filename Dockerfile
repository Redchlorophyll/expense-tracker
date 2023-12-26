# Build stage
FROM golang:1.17.13-alpine as builder

# Set the working directory in the container
WORKDIR /app

# Copy the application files into the working directory
COPY . /app

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o main .

# Runtime stage
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy only the built binary from the builder image
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
CMD ["./main"]
