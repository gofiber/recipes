FROM golang:1.20@sha256:8f9af7094d0cb27cc783c697ac5ba25efdc4da35f8526db21f7aebb0b0b4f18a

RUN apt update && apt upgrade -y &&\
    apt install -y git

WORKDIR /go/src/app

COPY . ./

RUN go mod tidy && go mod verify

ENTRYPOINT [ "go", "run", "./cmd/api" ]