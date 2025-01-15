FROM golang:1.23.4@sha256:585103a29aa6d4c98bbb45d2446e1fdf41441698bbdf707d1801f5708e479f04

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
