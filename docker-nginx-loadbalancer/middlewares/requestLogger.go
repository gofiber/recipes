package middlewares

import (
	"github.com/gofiber/fiber"
	"log"
)

/*RequestLogger : Check the request*/
func RequestLogger(c *fiber.Ctx) {
	method := string(c.Fasthttp.Request.Header.Method())
	url:= string(c.Fasthttp.Request.Header.RequestURI())
	log.Println("[ "+method+" ]"+" => "+url)
	c.Next()
}

