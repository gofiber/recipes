FROM golang:1.23.3@sha256:3694e364a0357a6e885569d3bdd097f7b36839bc4fb0df2af33d411b37e27e59

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
