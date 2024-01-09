package main

import (
	internal "github.com/Redchlorophyll/expense-tracker/internal/config"
	ExpenseTrackerConfig "github.com/Redchlorophyll/expense-tracker/internal/config/expense_tracker"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	serv := internal.GetService()
	httpService := ExpenseTrackerConfig.InitializeService(serv)

	httpService.UserHandler.SetRoute(app)
	httpService.TransactionHandler.SetRoute(app)

	app.Listen(":8080")
}
