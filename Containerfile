# Stage 1: Build
FROM golang:1.23 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o service ./cmd/maps-service/service.go

# Stage 2: Run
FROM alpine

# Set up a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Set the working directory in the container
WORKDIR /home/appuser

# Copy the built application from the builder stage
COPY --from=builder /app/service .

# Expose application port (adjust if needed)
EXPOSE 8081

# Command to run the application
CMD ["./service"]
