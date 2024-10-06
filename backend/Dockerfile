# Backend Dockerfile

# Use the official Golang image
FROM golang:1.23 as builder

# Install Air from the new module path
RUN go install github.com/air-verse/air@latest

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the backend code into the container
COPY . .

# Expose the port your application listens on (e.g., 8080)
EXPOSE 8080

# Set the working directory to the backend directory
WORKDIR /app/backend

# Command to run the application with 'air' for live reloading
CMD ["air"]
