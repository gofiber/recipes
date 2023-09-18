# multistage docker build. This redices the size of the final docker image.
# stage 1 to build the app
FROM golang:alpine@sha256:96634e55b363cb93d39f78fb18aa64abc7f96d372c176660d7b8b6118939d97b as builder

RUN mkdir /build 

ADD . /build/

WORKDIR /build 

RUN go build -o main .

# stage 2 deploys the app built in stage 1
FROM alpine@sha256:7144f7bab3d4c2648d7e59409f15ec52a18006a128c733fcff20d3a4a54ba44a

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]