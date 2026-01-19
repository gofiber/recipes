package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberAdapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

var fiberLambda *fiberAdapter.FiberLambda

// init the Fiber Server
func init() {
	log.Printf("Fiber cold start")
	app := fiber.New()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World"})
	})

	fiberLambda = fiberAdapter.New(app)
}

// Handler will deal with Fiber working with Lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}
