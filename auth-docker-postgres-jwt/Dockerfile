FROM golang:1.24@sha256:b6652731d1d5622f85509f72942ce2a344e3bf6dd6793b2e462cc5fb3126b566

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
