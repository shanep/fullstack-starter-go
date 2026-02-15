# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to leverage Docker's build cache
COPY go.mod .

# Download Go modules and dependencies
# Uncomment the following line if you want to download dependencies during the build process. This can speed up subsequent builds if your dependencies don't change often.
#RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary.
# Update this command below as needed
RUN go build -o /main .


# Stage 2: Create a minimal runtime image
# Using 'scratch' for the smallest possible image or 'alpine:latest' for a minimal Linux environment if needed
FROM scratch

# Copy only the built binary from the 'builder' stage
COPY --from=builder /main /main

# Expose the port your web server listens on (e.g., 3000)
EXPOSE 3000

# Command to run the application when the container starts
CMD ["/main"]