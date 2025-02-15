# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

COPY migrations/ ./migrations
COPY goose/ ./goose

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./app/main/migrations
COPY --from=builder /app/goose/goose /usr/local/bin/goose

# Set explicit permissions for both executables
RUN chmod +x /app/main && \
    chmod +x /usr/local/bin/goose

EXPOSE 8080

# Run the application
CMD ["sh", "-c", "goose -dir app/main/migrations postgres ${DB_URL} up && ./main"]