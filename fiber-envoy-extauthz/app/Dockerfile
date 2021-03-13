FROM golang:alpine AS builder
WORKDIR /bin
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

FROM alpine:latest AS final
WORKDIR /
COPY --from=builder /bin/app ./
EXPOSE 3000
ENTRYPOINT ["./app"]
