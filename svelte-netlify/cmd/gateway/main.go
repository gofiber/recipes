package main

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3/middleware/static"

	"github.com/gofiber/recipes/svelte-netlify/adapter"
	"github.com/gofiber/recipes/svelte-netlify/handler"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v3"
)

var fiberLambda *adapter.FiberLambda

func init() {
	app := fiber.New()
	app.Get("/*", static.New("./public"))
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendFile("index")
	})
	app.Get("/api/:ip", handler.CacheRequest(10*time.Minute), handler.GeoLocation)

	fiberLambda = adapter.New(app)
}

// Handler proxies app requests to AWS Lambda.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
