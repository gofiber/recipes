FROM golang@sha256:0b55ab82ac2a54a6f8f85ec8b943b9e470c39e32c109b766bbc1b801f3fa8d3b AS builder
WORKDIR /k8s-example
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/k8s-example .

FROM alpine:latest@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b
RUN apk --no-cache add ca-certificates
RUN addgroup -S k8s-example && adduser -S k8s-example -G k8s-example
USER k8s-example
WORKDIR /home/k8s-example
COPY --from=builder /bin/k8s-example ./
EXPOSE 3000
ENTRYPOINT ["./k8s-example"]
