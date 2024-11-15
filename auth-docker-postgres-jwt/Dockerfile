FROM golang:1.23.3@sha256:73f06be4578c9987ce560087e2e2ea6485fb605e3910542cadd8fa09fc5f3e31

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
