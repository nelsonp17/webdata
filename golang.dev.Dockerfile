FROM golang:alpine

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod .
COPY go.sum .

# Instalar air
RUN go install github.com/air-verse/air@latest

# Download dependencies
RUN go mod download

# Install necessary packages
RUN apk update && apk add gcc musl-dev libwebp-dev

# Copy the source code
COPY . .

# Build the application
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/app .

CMD ["air"]