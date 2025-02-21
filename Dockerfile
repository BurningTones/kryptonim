FROM golang:1.22-alpine

WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main ./cmd

# Expose the application port
EXPOSE 8080

# Copy the .env file
COPY .env .env

# Set environment variables from the .env file
ENV $(cat .env | grep -v ^# | xargs)

# Run the application
CMD ["./main"]