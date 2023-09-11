# Dockerfile

# Use the official Go Image as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code into the container
COPY . .

# Copy the db.yaml configuration file into the container
COPY resources/db.yaml resources/db.yaml

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
