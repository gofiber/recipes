FROM golang:1.20@sha256:685846c2eb68f6ae4d784d56125ea9894981b3408340aa7d89362fef83e1aebb as build-env

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /go/bin/app

FROM gcr.io/distroless/static-debian11:latest@sha256:9be3fcc6abeaf985b5ecce59451acbcbb15e7be39472320c538d0d55a0834edc

COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.7.0@sha256:c82920b82cd7fc20febf8bd608e907610bb32015edaf0fd6c065138d17125c12 /lambda-adapter /opt/extensions/lambda-adapter
COPY --from=build-env /go/bin/app /

ENV PORT=3000
EXPOSE 3000

CMD ["/app"]