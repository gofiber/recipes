FROM golang@sha256:a66eda637829ce891e9cf61ff1ee0edf544e1f6c5b0e666c7310dce231a66f28 AS builder
WORKDIR /k8s-example
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/k8s-example .

FROM alpine:latest@sha256:b89d9c93e9ed3597455c90a0b88a8bbb5cb7188438f70953fede212a0c4394e0
RUN apk --no-cache add ca-certificates
RUN addgroup -S k8s-example && adduser -S k8s-example -G k8s-example
USER k8s-example
WORKDIR /home/k8s-example
COPY --from=builder /bin/k8s-example ./
EXPOSE 3000
ENTRYPOINT ["./k8s-example"]
