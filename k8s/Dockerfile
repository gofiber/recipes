FROM golang@sha256:672a2286da3ee7a854c3e0a56e0838918d0dbb1c18652992930293312de898a6 AS builder
WORKDIR /k8s-example
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/k8s-example .

FROM alpine:latest@sha256:51b67269f354137895d43f3b3d810bfacd3945438e94dc5ac55fdac340352f48
RUN apk --no-cache add ca-certificates
RUN addgroup -S k8s-example && adduser -S k8s-example -G k8s-example
USER k8s-example
WORKDIR /home/k8s-example
COPY --from=builder /bin/k8s-example ./
EXPOSE 3000
ENTRYPOINT ["./k8s-example"]
