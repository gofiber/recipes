FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY resources .
COPY go.mod .
COPY go.sum .
RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
RUN mkdir -p ./storage
RUN mkdir -p ./.cache
RUN mkdir -p ./.cache/go-build
RUN chown appuser ./.cache
RUN chown appuser ./.cache/go-build
RUN chown appuser ./storage
RUN mkdir -p ./uploads
RUN chown appuser ./uploads
USER appuser
EXPOSE 1421
ENTRYPOINT ["./setupDockerHost.sh"]