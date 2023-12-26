FROM golang:1.17

# Set the working directory in the container
WORKDIR /app

# Copy the application files into the working directory
COPY . /app

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
CMD ["./main"]
