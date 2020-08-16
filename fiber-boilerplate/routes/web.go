package routes

import (
	"github.com/gofiber/fiber"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/auth"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/controllers"
	"github.com/itsursujit/fiber-boilerplate/middlewares"
)

func WebRoutes() {
	web := App.Group("")
	web.Use(auth.AuthCookie)
	LandingRoutes(web)
	UserRoutes(web)
}

func LandingRoutes(app fiber.Router) {
	app.Use(middlewares.Authenticate(middlewares.AuthConfig{
		SigningKey:  []byte(config.AuthConfig.App_Jwt_Secret),
		TokenLookup: "cookie:fiber-boilerplate-Token",
		ErrorHandler: func(ctx *fiber.Ctx, err error) {
			auth.Logout(ctx)
			ctx.Next()
		},
	}))
	app.Get("/", controllers.Landing)
}

func UserRoutes(app fiber.Router) {
	account := app.Group("/account")
	account.Use(middlewares.Authenticate(middlewares.AuthConfig{
		SigningKey:  []byte(config.AuthConfig.App_Jwt_Secret),
		TokenLookup: "cookie:fiber-boilerplate-Token",
		ErrorHandler: func(ctx *fiber.Ctx, err error) {
			auth.Logout(ctx)
			ctx.Redirect("/login")
			return
		},
	}))
	// account.Get("/users", controllers.Index)
	account.Get("/file-manager", controllers.FileIndex)
	account.Get("/file-manager/view", controllers.ViewFile)
	account.Post("/file-manager/upload", controllers.Upload)
	account.Post("/paypal/do/order", controllers.PlaceOrderFromPaypal)
	account.Post("/paypal/do/order/validate/:amount", controllers.ValidateOrderFromPaypal)
	account.Get("/paypal/order/success/:id", controllers.PostOrderSuccessResponseFromPaypal)
	account.Post("/paypal/order/cancel/:id", controllers.PostOrderCancelResponseFromPaypal)
	account.Get("/paypal/order/:id", controllers.GetOrderDetailFromPaypal)
}
