
# Simple graphql example for [Fiber](https://github.com/gofiber/fiber) with [Fasthttp gqlgen](https://github.com/arsmn/gqlgen)

## Usage
- init your go module
`go mod init github.com/[username]/gqlgen-todos`
- install gqlgen
`go get github.com/99designs/gqlgen`
- replace with arsmn/gqlgen by adding this line in go.mod file <br/> 
`replace github.com/99designs/gqlgen v0.13.0 => github.com/arsmn/gqlgen v0.13.2`
- build the server
`go run github.com/99designs/gqlgen init`
- implement the resolvers
- run the server
`go run server.go`
- browse `http://localhost:8080/` for playground
- see more [examples](https://github.com/arsmn/gqlgen/tree/master/example) and [documentation](https://gqlgen.com)