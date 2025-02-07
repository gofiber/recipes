FROM golang:1.24rc3@sha256:f609f42515142764b42c25aac3ee2eeeae5466749caffc9ee8762c17c95662dc

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
