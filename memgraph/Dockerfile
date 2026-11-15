FROM golang:1.27 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the app for the platform the image is built on.
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -ldflags="-s -w" -o app .

FROM scratch

COPY --from=builder ["/build/app", "/"]

EXPOSE 3000

ENTRYPOINT ["/app"]
