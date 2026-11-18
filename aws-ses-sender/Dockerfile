FROM golang:1.27 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.24

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# The image must not contain .env (it holds AWS keys). Pass the variables at run time,
# for example: docker run --env-file .env -p 3000:3000 aws-ses-sender
COPY --from=builder /usr/src/app/app .

CMD ["./app"]
