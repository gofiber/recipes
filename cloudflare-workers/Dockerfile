# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS build

# Set destination for COPY
WORKDIR /app

# Download any Go modules
COPY container_src/go.mod container_src/go.sum ./
RUN go mod download

# Copy container source code
COPY container_src/*.go ./

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /server

# Use distroless for minimal attack surface
FROM gcr.io/distroless/static-debian12:latest

# Copy the binary
COPY --from=build /server /server

# Expose port
EXPOSE 8080

# Run
ENTRYPOINT ["/server"]