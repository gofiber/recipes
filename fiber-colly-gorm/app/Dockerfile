FROM golang:1.20

RUN apt update && apt upgrade -y &&\
    apt install -y git\
    make openssh-client

WORKDIR /go/src/app

COPY . ./

RUN go mod init fiber-colly-gorm

RUN GOPROXY=https://goproxy.cn go mod tidy &&\
    GOPROXY=https://goproxy.cn go mod verify

ENTRYPOINT [ "go", "run", "./cmd/api" ]