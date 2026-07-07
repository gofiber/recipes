FROM golang:1.25-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /bin/worker ./worker

FROM alpine:latest
COPY --from=build /bin/worker /bin/worker
ENTRYPOINT ["/bin/worker"]
