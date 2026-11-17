FROM --platform=$BUILDPLATFORM golang:1.27 AS builder

# Set by BuildKit to the platform the image is built for.
ARG TARGETOS
ARG TARGETARCH

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the app for the platform the image is built for.
ENV CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN go build -ldflags="-s -w" -o app .

FROM scratch

COPY --from=builder ["/build/app", "/"]

EXPOSE 3000

ENTRYPOINT ["/app"]
