FROM golang:1.24@sha256:ceb568d0de81fbef40ce4fee77eab524a0a0a8536065c51866ad8c59b7a912cf

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
