FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy just the module files first for better caching
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the rest of the application
COPY backend/ ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/server .
COPY backend/templates ./templates/
COPY frontend/static ./frontend/static/
EXPOSE 8080
CMD ["./server"]