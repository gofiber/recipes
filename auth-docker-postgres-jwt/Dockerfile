FROM golang:1.24@sha256:52ff1b35ff8de185bf9fd26c70077190cd0bed1e9f16a2d498ce907e5c421268

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
