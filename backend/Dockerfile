# Use an official Go runtime as a parent image
FROM golang:1.20-alpine

# Set the Working Directory
WORKDIR /app

# Copy the go modules
COPY go.mod go.sum ./

# Download modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./app"]