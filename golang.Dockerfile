# Use the official Golang image as the base image
FROM golang:alpine

# Install necessary packages including Chromium
RUN apk update && apk add --no-cache \
    chromium \
    chromium-chromedriver \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont

# Create appuser
RUN adduser -D -g '' appuser

# Create app directory
RUN mkdir /app
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Set environment variable for Chromium executable path
ENV CHROMIUM_EXECUTABLE_PATH /usr/bin/chromium-browser

# Build the application
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/app .

# Use an unprivileged user
USER appuser

# Command to run the application
CMD ["/go/bin/app"]







## Use a multi-stage build to reduce the final image size.
#
## Stage 1: Build the Go application
#FROM golang:alpine AS builder
#
#WORKDIR /app
#
## Copy go.mod and go.sum if they exist
#COPY go.mod go.sum ./
#RUN go mod download
#
#COPY . .
#
## Build the Go application
#RUN go build -o main .
#
## Stage 2: Create the final image with Ubuntu and Chrome
#FROM ubuntu:latest
#
## Install dependencies for Chrome
#RUN apt-get update && \
#    apt-get install -y wget gnupg lsb-release software-properties-common
#
## Add Google Chrome repository
#RUN curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
#RUN apt-get install -y ./google-chrome-stable_current_amd64.deb
#RUN rm google-chrome-stable_current_amd64.deb
#
## Copy the built binary from the builder stage
#COPY --from=builder /app/main /app/main
#
## Set the entrypoint to run your Go application (or any other command)
#ENTRYPOINT ["/app/main"]
