// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/recipes/fiber-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := proto.NewAddServiceClient(conn)

	// g := gin.Default()
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/add/:a/:b", func(c fiber.Ctx) error {
		a, err := strconv.ParseInt(c.Params("a"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument A",
			})
		}
		b, err := strconv.ParseInt(c.Params("b"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument B",
			})
		}
		req := &proto.Request{A: a, B: b}
		if res, err := client.Add(context.Background(), req); err == nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"result": fmt.Sprint(res.Result),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	})

	app.Get("/mult/:a/:b", func(c fiber.Ctx) error {
		a, err := strconv.ParseInt(c.Params("a"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument A",
			})
		}
		b, err := strconv.ParseInt(c.Params("b"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument B",
			})
		}
		req := &proto.Request{A: a, B: b}
		if res, err := client.Multiply(context.Background(), req); err == nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"result": fmt.Sprint(res.Result),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	})
	log.Fatal(app.Listen(":3000"))
}
