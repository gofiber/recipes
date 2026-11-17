FROM --platform=$BUILDPLATFORM golang:alpine AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /bin
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o app main.go

FROM alpine:latest AS final
WORKDIR /
COPY --from=builder /bin/app ./
EXPOSE 1337
ENTRYPOINT ["./app"]
