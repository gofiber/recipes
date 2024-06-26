# multistage docker build. This redices the size of the final docker image.
# stage 1 to build the app
FROM golang:alpine@sha256:ace6cc3fe58d0c7b12303c57afe6d6724851152df55e08057b43990b927ad5e8 as builder

RUN mkdir /build 

ADD . /build/

WORKDIR /build 

RUN go build -o main .

# stage 2 deploys the app built in stage 1
FROM alpine@sha256:b89d9c93e9ed3597455c90a0b88a8bbb5cb7188438f70953fede212a0c4394e0

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]