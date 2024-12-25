FROM golang:1.23.4@sha256:fffeec553ab94798ca27bda80814acd3f776492d7deafb1a82dd08830b76ec7c

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
