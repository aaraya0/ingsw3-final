# Use the official Golang image as the base image
FROM golang:1.18 AS builder

# Set the working directory to the app directory
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the entire application to the working directory
COPY . .

# Build the application
RUN go build -o /app/backend ./backend/main.go

# Use the official Node.js image as the base image for the frontend
FROM node:14 AS frontend-builder

# Set the working directory to the frontend directory
WORKDIR /app/frontend

# Copy the frontend code to the working directory
COPY frontend .

# Install frontend dependencies
RUN npm install

# Build the frontend
RUN npm run build

# Use a minimal base image for the final image
FROM alpine:latest

# Set the working directory to the app directory
WORKDIR /app

# Copy the compiled backend binary
COPY --from=builder /app/backend .

# Copy the compiled frontend files
COPY --from=frontend-builder /app/frontend/build ./frontend/build

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./backend"]
