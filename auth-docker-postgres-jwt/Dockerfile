FROM golang:1.24rc3@sha256:69ffbde212822c1bb785db9de3e27c4dc0077ade807f460b78ce49e38ecffeb2

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
