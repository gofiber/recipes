FROM golang:1.23.0@sha256:613a108a4a4b1dfb6923305db791a19d088f77632317cfc3446825c54fb862cd

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
