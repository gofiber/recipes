FROM golang as build-go
WORKDIR /cloud-run-example
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/cloud-run-example .

FROM alpine:latest
RUN addgroup -S cloud-run-example && adduser -S cloud-run-example -G cloud-run-example
USER cloud-run-example
WORKDIR /home/cloud-run-example
COPY --from=build-go /bin/cloud-run-example ./
EXPOSE 3000
ENTRYPOINT ["./cloud-run-example"]
