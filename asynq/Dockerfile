FROM golang:1.26-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /bin/api .

FROM alpine:latest
RUN addgroup -S app && adduser -S app -G app
USER app
COPY --from=build /bin/api /bin/api
EXPOSE 3000
ENTRYPOINT ["/bin/api"]
