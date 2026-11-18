FROM --platform=$BUILDPLATFORM golang:1.27 AS builder

# Set by BuildKit to the platform the image is built for.
ARG TARGETOS
ARG TARGETARCH

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container. The worker/ directory is its own Go module and is
# built by its own Dockerfile, so it is not part of this build.
COPY . .

# Build the API server for the platform the image is built for.
ENV CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN go build -ldflags="-s -w" -o api .

FROM scratch

# Copy the binary from /build to the root folder of the scratch container.
COPY --from=builder ["/build/api", "/"]

EXPOSE 3000

# Command to run when starting the container.
ENTRYPOINT ["/api"]
