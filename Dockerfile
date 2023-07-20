# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.20.2-buster as builder

WORKDIR /

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -o main 

# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "./main" ]