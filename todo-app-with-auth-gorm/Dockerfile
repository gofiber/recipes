
#build stage
FROM golang:alpine@sha256:0466223b8544fb7d4ff04748acc4d75a608234bf4e79563bff208d2060c0dd79 AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v . && go build -ldflags="-s -w" main.go

#final stage
FROM alpine:latest@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b
LABEL maintainer=numtostr version=0.0.1
COPY --from=builder /go/src/app/main /main
EXPOSE 4000
ENTRYPOINT /main
