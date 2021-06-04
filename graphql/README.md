
# Simple graphql example for [Fiber](https://github.com/gofiber/fiber) with [fastgql](https://github.com/arsmn/fastgql)

## Usage
- init your go module
`go mod init github.com/[username]/gqlgen-todos`
- install gqlgen
`go get github.com/arsmn/fastgql`
- build the server
`go run github.com/arsmn/fastgql init`
- implement the resolvers
- run the server
`go run server.go`
- browse `http://localhost:8080/` for playground
- see more [examples](https://github.com/arsmn/fastgql/tree/master/example) and [documentation](https://gqlgen.com)