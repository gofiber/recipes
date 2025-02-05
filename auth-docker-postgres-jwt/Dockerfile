FROM golang:1.24rc2@sha256:a5ddeaeeec69e764a7a86c6063ceabb4ab0261a448a173ba3ad0dde7b93524d6

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
