
#build stage
FROM golang:alpine@sha256:4db4aac30880b978cae5445dd4a706215249ad4f43d28bd7cdf7906e9be8dd6b AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v . && go build -ldflags="-s -w" main.go

#final stage
FROM alpine:latest@sha256:51b67269f354137895d43f3b3d810bfacd3945438e94dc5ac55fdac340352f48
LABEL maintainer=numtostr version=0.0.1
COPY --from=builder /go/src/app/main /main
EXPOSE 4000
ENTRYPOINT /main
