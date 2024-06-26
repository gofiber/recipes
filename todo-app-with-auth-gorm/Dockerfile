
#build stage
FROM golang:alpine@sha256:ace6cc3fe58d0c7b12303c57afe6d6724851152df55e08057b43990b927ad5e8 AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v . && go build -ldflags="-s -w" main.go

#final stage
FROM alpine:latest@sha256:b89d9c93e9ed3597455c90a0b88a8bbb5cb7188438f70953fede212a0c4394e0
LABEL maintainer=numtostr version=0.0.1
COPY --from=builder /go/src/app/main /main
EXPOSE 4000
ENTRYPOINT /main
