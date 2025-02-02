# Use multi-stage build for smaller final image
FROM golang:1.18-alpine AS build

# Set the working directory
WORKDIR /app

COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /HNG

# Run the tests in the container
FROM build AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

# Set the working directory
WORKDIR /

# Copy the binary from the build stage
COPY --from=build /HNG /HNG

# Expose port
EXPOSE 8080

USER nonroot:nonroot

# Run the binary
ENTRYPOINT [ "/HNG", "port", "8080"]