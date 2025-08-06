# Build stage for SvelteKit frontend
FROM node:22-alpine AS frontend-builder

WORKDIR /app

# Copy frontend files
COPY frontend/ ./

# Install dependencies and build the frontend
RUN npm install
# Run svelte-kit sync before building (important for SvelteKit projects)
RUN npm run prepare
RUN npm run build

# Build stage for Go backend
FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

# Copy go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Copy frontend build from previous stage
COPY --from=frontend-builder /app/dist ./frontend/dist

# Install SQLite dependencies
RUN apk add --no-cache gcc musl-dev

# Enable CGO and build the Go application
ENV CGO_ENABLED=1
RUN go build -o carsearch ./cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the compiled application from the build stage
COPY --from=backend-builder /app/carsearch .

# Copy the frontend build directory
COPY --from=backend-builder /app/frontend/dist ./frontend/dist

# Set execution permissions
RUN chmod +x ./carsearch

# Expose the port the app runs on
EXPOSE 8000

# Command to run the executable
CMD ["./carsearch"]
