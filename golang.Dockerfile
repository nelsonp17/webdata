FROM golang:alpine

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

# Install necessary packages
RUN apk update && apk add gcc musl-dev libwebp-dev

# Copy the source code
COPY . .

# Build the application
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/app .

# Use an unprivileged user
USER appuser

# Command to run the application
CMD ["/go/bin/app"]