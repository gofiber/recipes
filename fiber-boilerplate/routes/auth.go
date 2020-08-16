package routes

import (
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/controllers"
	"github.com/itsursujit/fiber-boilerplate/middlewares"
)

func AuthRoutes() {
	App.Get("/login",
		middlewares.RedirectToHomePageOnLogin,
		controllers.LoginGet,
	)
	App.Post("/do/login",
		middlewares.ValidateLoginPost,
		controllers.LoginPost,
	)
	App.Post("/do/logout", controllers.LogoutPost)

	App.Get("/register", middlewares.RedirectToHomePageOnLogin, controllers.RegisterGet)
	App.Post("/do/register",
		middlewares.RedirectToHomePageOnLogin,
		middlewares.ValidateRegisterPost,
		controllers.RegisterPost,
	)

	App.Get("/reset-password",
		middlewares.ValidatePasswordReset,
		controllers.PasswordReset,
	)
	App.Post("/do/reset-password",
		controllers.RequestPasswordResetPost,
	)
	App.Get("/request-password-reset", middlewares.RedirectToHomePageOnLogin, controllers.RequestPasswordReset)
	App.Post("/do/password-reset/:token",
		middlewares.RedirectToHomePageOnLogin,
		middlewares.ValidatePasswordResetPost,
		middlewares.ValidateRegisterPost,
		controllers.PasswordResetPost)
	App.Get("/resend/confirm", controllers.ResendConfirmEmail)
	App.Get("/do/verify-email",
		middlewares.ValidateConfirmToken,
		controllers.VerifyRegisteredEmail,
	)

}
