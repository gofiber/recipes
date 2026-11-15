# TARGETOS and TARGETARCH are the platform the image is built for. BuildKit sets them, and
# `docker build --platform` chooses them. A builder that does not set them leaves them empty,
# and then Go builds for the platform it runs on, which is linux/amd64 on Cloud Build.
# Unlike the other recipes this file does not say `FROM --platform=$BUILDPLATFORM`: builders
# without BuildKit, such as the legacy Docker builder, cannot parse it, and Cloud Run only needs
# amd64 images, which Cloud Build builds natively.
FROM golang:1.27-alpine AS build-go
ARG TARGETOS
ARG TARGETARCH
WORKDIR /cloud-run-example
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/cloud-run-example .

FROM alpine:latest
RUN addgroup -S cloud-run-example && adduser -S cloud-run-example -G cloud-run-example
USER cloud-run-example
WORKDIR /home/cloud-run-example
COPY --from=build-go /bin/cloud-run-example ./
EXPOSE 3000
ENTRYPOINT ["./cloud-run-example"]
