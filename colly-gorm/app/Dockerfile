FROM golang:1.23

RUN apt update && apt upgrade -y &&\
    apt install -y git

WORKDIR /go/src/app

COPY . ./

RUN go mod tidy && go mod verify

ENTRYPOINT [ "go", "run", "./cmd/api" ]