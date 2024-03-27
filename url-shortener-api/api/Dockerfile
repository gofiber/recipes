# multistage docker build. This redices the size of the final docker image.
# stage 1 to build the app
FROM golang:alpine@sha256:0466223b8544fb7d4ff04748acc4d75a608234bf4e79563bff208d2060c0dd79 as builder

RUN mkdir /build 

ADD . /build/

WORKDIR /build 

RUN go build -o main .

# stage 2 deploys the app built in stage 1
FROM alpine@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]