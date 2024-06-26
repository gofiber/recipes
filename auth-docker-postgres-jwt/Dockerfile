FROM golang:1.21@sha256:b405b620c7b53ef64695c7da7c8396f411f381c1eb7da6713c585dd7eca1559b

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
