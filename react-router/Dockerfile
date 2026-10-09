# First stage: Get Golang image from DockerHub.
FROM golang:1.27 AS backend-builder

# Label this container.
LABEL appname="React App Fiber"
LABEL author="Nicholas Dwiarto W. <nicholasdwiarto@yahoo.com>"
LABEL description="Sample application to showcase serving React SPA with Fiber"

# Set our working directory for this stage.
WORKDIR /backendcompile

# Copy all of our files.
COPY . .

# Get and install all dependencies.
RUN CGO_ENABLED=0 go build -o react-router ./cmd/react-router/main.go

# Next stage: Build our frontend application.
FROM node:26@sha256:a723b54c35a76e947095a20a67d39585bb09c862e6b1adeb8a9f518f95e34fb0 AS frontend-builder

# Set our working directory for this stage.
WORKDIR /frontendcompile

# Copy lockfiles and dependencies.
COPY ./web/package.json ./web/package-lock.json ./

# Install our dependencies.
RUN npm ci

# Copy our installed 'node_modules' and everything else.
COPY ./web .

# Build our application.
RUN npm run build

# Last stage: discard everything except our executables.
FROM alpine:latest AS prod

# Set our next working directory.
WORKDIR /build

# Create directory for our React application to live.
RUN mkdir -p /web/build

# Copy our executable and our built React application.
COPY --from=backend-builder /backendcompile/react-router .
COPY --from=frontend-builder /frontendcompile/build ./web/build

# Declare entrypoints and activation commands.
EXPOSE 8080
ENTRYPOINT ["./react-router"]
