
#build stage
FROM golang:alpine@sha256:96634e55b363cb93d39f78fb18aa64abc7f96d372c176660d7b8b6118939d97b AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v . && go build -ldflags="-s -w" main.go

#final stage
FROM alpine:latest@sha256:7144f7bab3d4c2648d7e59409f15ec52a18006a128c733fcff20d3a4a54ba44a
LABEL maintainer=numtostr version=0.0.1
COPY --from=builder /go/src/app/main /main
EXPOSE 4000
ENTRYPOINT /main
