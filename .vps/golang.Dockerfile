# Use the official Golang image as the base image
FROM golang:alpine

# Install necessary packages including Chromium
RUN apk update && apk add --no-cache \
    chromium \
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