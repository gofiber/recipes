FROM golang:1.23.0@sha256:1a6db32ea47a4910759d5bcbabeb8a8b42658e311bd8348ea4763735447c636c

# Enviroment variable
WORKDIR /usr/src/some-api

RUN go install github.com/air-verse/air@latest

#Copying files to work directory
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

# Run and expose the server on port 3000
EXPOSE 3000

# CMD [ "nodemon", "build/app.js" ]
