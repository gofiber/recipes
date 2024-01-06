FROM golang@sha256:672a2286da3ee7a854c3e0a56e0838918d0dbb1c18652992930293312de898a6 as build-go
WORKDIR /cloud-run-example
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/cloud-run-example .

FROM alpine:latest@sha256:51b67269f354137895d43f3b3d810bfacd3945438e94dc5ac55fdac340352f48
RUN addgroup -S cloud-run-example && adduser -S cloud-run-example -G cloud-run-example
USER cloud-run-example
WORKDIR /home/cloud-run-example
COPY --from=build-go /bin/cloud-run-example ./
EXPOSE 3000
ENTRYPOINT ["./cloud-run-example"]
