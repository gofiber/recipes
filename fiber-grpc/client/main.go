// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/amalshaji/fiber-grpc/proto"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewAddServiceClient(conn)

	// g := gin.Default()
	app := fiber.New()

	app.Use(middleware.Logger())

	app.Get("/add/:a/:b", func(c *fiber.Ctx) {
		a, err := strconv.ParseUint(c.Params("a"), 10, 64)
		if err != nil {
			_ = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument A",
			})
		}
		b, err := strconv.ParseUint(c.Params("b"), 10, 64)
		if err != nil {
			_ = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument B",
			})
			return
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		if res, err := client.Add(context.Background(), req); err == nil {
			_ = c.Status(fiber.StatusOK).JSON(fiber.Map{
				"result": fmt.Sprint(res.Result),
			})
		} else {
			_ = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
			return
		}
	})

	app.Get("/mult/:a/:b", func(c *fiber.Ctx) {
		a, err := strconv.ParseUint(c.Params("a"), 10, 64)
		if err != nil {
			_ = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument A",
			})
		}
		b, err := strconv.ParseUint(c.Params("b"), 10, 64)
		if err != nil {
			_ = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument B",
			})
			return
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		if res, err := client.Multiply(context.Background(), req); err == nil {
			_ = c.Status(fiber.StatusOK).JSON(fiber.Map{
				"result": fmt.Sprint(res.Result),
			})
		} else {
			_ = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
			return
		}
	})
	log.Fatal(app.Listen(3000))
}
