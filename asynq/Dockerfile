FROM golang:1.25-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /bin/api .

FROM alpine:latest
COPY --from=build /bin/api /bin/api
EXPOSE 3000
ENTRYPOINT ["/bin/api"]
