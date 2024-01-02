# Use official golang as base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
ADD . /app

# Build the Golang application
RUN go build -o main .

# Make port 8080 available to the world outside this container
# EXPOSE 8080

# Run the executable
CMD ["./main"]
