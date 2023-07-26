# Use the official Go image from the Docker Hub
FROM golang:1.21-rc-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go Mod and Sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
