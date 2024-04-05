FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Go project into the container
COPY ./hello /go/src/app

# Installs Go dependencies
RUN go mod download
RUN go mod vendor

# Build the Go project
RUN go build -o hello .

# Run the built executable
CMD ["./hello"]