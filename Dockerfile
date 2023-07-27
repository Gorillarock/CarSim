## Build Command: docker build -t carapp .
## Run Command: docker run -it --rm carapp

# Step 1: Build stage
FROM golang:alpine AS builder

WORKDIR /app

# Copy the entire source code from the current directory to /app
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main

# Step 2: Create the final executable image
FROM alpine:latest

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]