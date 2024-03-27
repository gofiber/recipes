FROM golang:1.20@sha256:8f9af7094d0cb27cc783c697ac5ba25efdc4da35f8526db21f7aebb0b0b4f18a as build-env

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /go/bin/app

FROM gcr.io/distroless/static-debian11:latest@sha256:046b92c933032a8ca99a66f4c79a68ac029d9a4ababd1a806a82140b3b899fd3

COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.7.0@sha256:c82920b82cd7fc20febf8bd608e907610bb32015edaf0fd6c065138d17125c12 /lambda-adapter /opt/extensions/lambda-adapter
COPY --from=build-env /go/bin/app /

ENV PORT=3000
EXPOSE 3000

CMD ["/app"]