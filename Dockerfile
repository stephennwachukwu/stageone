# Use multi-stage build for smaller final image
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

# Build the application
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o api

# Final stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/api .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./api"]