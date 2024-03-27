FROM golang@sha256:0b55ab82ac2a54a6f8f85ec8b943b9e470c39e32c109b766bbc1b801f3fa8d3b as build-go
WORKDIR /cloud-run-example
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/cloud-run-example .

FROM alpine:latest@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b
RUN addgroup -S cloud-run-example && adduser -S cloud-run-example -G cloud-run-example
USER cloud-run-example
WORKDIR /home/cloud-run-example
COPY --from=build-go /bin/cloud-run-example ./
EXPOSE 3000
ENTRYPOINT ["./cloud-run-example"]
