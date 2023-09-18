FROM golang@sha256:cffaba795c36f07e372c7191b35ceaae114d74c31c3763d442982e3a4df3b39e as build-go
WORKDIR /cloud-run-example
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/cloud-run-example .

FROM alpine:latest@sha256:7144f7bab3d4c2648d7e59409f15ec52a18006a128c733fcff20d3a4a54ba44a
RUN addgroup -S cloud-run-example && adduser -S cloud-run-example -G cloud-run-example
USER cloud-run-example
WORKDIR /home/cloud-run-example
COPY --from=build-go /bin/cloud-run-example ./
EXPOSE 3000
ENTRYPOINT ["./cloud-run-example"]
