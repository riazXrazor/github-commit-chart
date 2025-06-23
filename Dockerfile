 # Use a Go base image
FROM golang:1.13-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app .

# Use a minimal base image for the final stage
FROM alpine:latest
COPY .env ./
COPY ./client/public ./client/public
# Copy the built binary from the builder stage
COPY --from=builder /bin/app /bin/app

# Expose the port your application listens on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/bin/app"]