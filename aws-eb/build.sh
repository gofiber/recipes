#!/bin/bash -xe
# See http://tldp.org/LDP/abs/html/options.html
# -x -> Print each command to stdout before executing it, expand commands
# -e -> Abort script at first error, when a command exits with non-zero status
#   (except in until or while loops, if-tests, list constructs)

# Get dependencies
go get -u github.com/gofiber/fiber/v2

# Build the binary
go build -o application application.go

# Modify permissons to make the binary executable.
chmod +x application
