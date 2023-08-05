package main

import (
	mainConfig "github.com/Redchlorophyll/expense-tracker/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	serv := mainConfig.GetService()
	httpService := mainConfig.InitializeService(serv)

	httpService.UserHandler.SetRoute(app)
	httpService.TransactionHandler.SetRoute(app)

	app.Listen(":8080")
}
