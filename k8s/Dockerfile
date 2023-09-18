FROM golang@sha256:cffaba795c36f07e372c7191b35ceaae114d74c31c3763d442982e3a4df3b39e AS builder
WORKDIR /k8s-example
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/k8s-example .

FROM alpine:latest@sha256:7144f7bab3d4c2648d7e59409f15ec52a18006a128c733fcff20d3a4a54ba44a
RUN apk --no-cache add ca-certificates
RUN addgroup -S k8s-example && adduser -S k8s-example -G k8s-example
USER k8s-example
WORKDIR /home/k8s-example
COPY --from=builder /bin/k8s-example ./
EXPOSE 3000
ENTRYPOINT ["./k8s-example"]
