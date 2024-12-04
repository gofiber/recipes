FROM golang:1.23.4@sha256:81e1da9b9604cdee2bc226e90e15f9d51ae6a8cd2271dc341c42ca0926c3b83f

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
