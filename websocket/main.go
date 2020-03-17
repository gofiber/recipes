// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
    "log"
    "github.com/gofiber/fiber"
	"github.com/gofiber/websocket"
)

func main() {
    app := fiber.New()
    
    // Optional middleware
    app.Use("/ws", func(c *fiber.Ctx) {
        if c.Get("host") == "localhost:3000" {
			c.Locals("Host", "Localhost:3000")
			c.Next()
			return
		}
		c.Status(403).Send("Request origin not allowed")
    })
    
    // Upgraded websocket request
    app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Host")) // "Localhost:3000"
        for {
            mt, msg, err := c.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                break
            }
            log.Printf("recv: %s", msg)
            err = c.WriteMessage(mt, msg)
            if err != nil {
                log.Println("write:", err)
                break
            }
        }
    })
    
    // ws://localhost:3000/ws
    app.Listen(3000)
}
