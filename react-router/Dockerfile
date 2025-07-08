# First stage: Get Golang image from DockerHub.
FROM golang:1.24 AS backend-builder

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
FROM node:24@sha256:256a2e7037e745228f7630d578e6c1d327ab4c0a8e401c63d0d4d9dfb3c13465 AS frontend-builder

# Set our working directory for this stage.
WORKDIR /frontendcompile

# Copy lockfiles and dependencies.
COPY ./web/package.json ./web/yarn.lock ./

# Install our dependencies.
RUN yarn

# Copy our installed 'node_modules' and everything else.
COPY ./web .

# Build our application.
RUN yarn build

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
