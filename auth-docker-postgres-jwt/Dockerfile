FROM golang:1.20@sha256:6e1a67ec11cc4eaa7a9699bc6e5385b00001669531284005ebd2ada4129138d9

# Enviroment variable
WORKDIR /usr/src/some-api

RUN go install github.com/cosmtrek/air@latest

#Copying files to work directory
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

# Run and expose the server on port 3000
EXPOSE 3000

# CMD [ "nodemon", "build/app.js" ]
