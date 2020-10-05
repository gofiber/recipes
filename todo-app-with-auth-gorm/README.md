# Prerequisite

1. Make sure you have the following installed outside the current project directory and available in your `GOPATH`
    - golang
    - [air](https://github.com/cosmtrek/air) for hot reloading
    - [godotenv](https://github.com/joho/godotenv) for loading `.env` file

# Installation

1. Clone this repo
2. Run `go get`

# Running

1. Type `air` in the command line

# Environment Variables

```shell
# PORT returns the server listening port
# default: 5000
PORT=

# DB returns the name of the sqlite database
# default: gotodo.db
DB=

# TOKENKEY returns the jwt token secret
TOKENKEY=

# TOKENEXP returns the jwt token expiration duration.
# Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
# default: 10h
TOKENEXP=
```
