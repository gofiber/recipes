# syntax=docker/dockerfile:1

ARG GO_VERSION=1.25.0
FROM golang:${GO_VERSION}-alpine AS build

# Set destination for COPY
WORKDIR /app

# Download any Go modules
COPY container_src/go.mod container_src/go.sum ./
RUN go mod download

# Copy container source code
COPY container_src/ ./

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /server

# Use distroless for minimal attack surface
FROM gcr.io/distroless/static-debian12@sha256:cd64bec9cec257044ce3a8dd3620cf83b387920100332f2b041f19c4d2febf93

# Copy the binary
COPY --from=build /server /server

# Expose port
EXPOSE 8080

# Run
ENTRYPOINT ["/server"]
