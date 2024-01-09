package httpservice

import (
	"github.com/gofiber/fiber/v2"
)

func (userHandler *UserServiceHandler) SetRoute(app *fiber.App) {
	group := app.Group("/api/v1/user")

	group.Post("/signup", userHandler.CreateUserHandler)
	group.Get("/login", userHandler.LoginHandler)
}
