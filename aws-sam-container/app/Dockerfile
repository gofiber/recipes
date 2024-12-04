FROM golang:1.23.4 as build-env

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /go/bin/app

FROM gcr.io/distroless/static-debian11:latest

COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.7.0 /lambda-adapter /opt/extensions/lambda-adapter
COPY --from=build-env /go/bin/app /

ENV PORT=3000
EXPOSE 3000

CMD ["/app"]