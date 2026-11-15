
FROM --platform=$BUILDPLATFORM golang:1.27 AS builder

# Set by BuildKit to the platform the image is built for.
ARG TARGETOS
ARG TARGETARCH

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Build the worker for the platform the image is built for.
ENV CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN go build -ldflags="-s -w" -o worker .

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/worker", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/worker"]