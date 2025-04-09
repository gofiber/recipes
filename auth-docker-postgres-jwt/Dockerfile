FROM golang:1.24@sha256:c0b66cfec8562c8d8452c63a63ede2014526224fa1cd5ecf5acda73f8b28263b

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
