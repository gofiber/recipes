package routes

import (
	"github.com/gofiber/fiber"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/controllers"
	apiControllers "github.com/itsursujit/fiber-boilerplate/controllers/api"
	"github.com/itsursujit/fiber-boilerplate/middlewares"
)

func ApiRoutes() {
	api := App.Group("/api")

	ApiVersions(api)
}

func ApiVersions(api fiber.Router) {
	v1Routes(api)
	v2Routes(api)
}

func v1AuthRoutes(api fiber.Router) {
	api.Post("/oauth/token", apiControllers.OAuthToken)
}

func v1Routes(api fiber.Router) {
	v1 := api.Group("/v1")
	v1AuthRoutes(v1)

	v1.Use(middlewares.Authenticate(middlewares.AuthConfig{
		SigningKey: []byte(config.AuthConfig.Api_Jwt_Secret),
	}))
	v1.Get("/users", controllers.Index)

}

func v2Routes(api fiber.Router) {

}

func UserApiRoutes() {

}
