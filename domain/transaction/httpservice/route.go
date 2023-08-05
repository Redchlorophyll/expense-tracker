package httpservice

import (
	"github.com/gofiber/fiber/v2"
)

func (transactionHandler *TransactionHandler) SetRoute(app *fiber.App) {
	group := app.Group("/api/v1/transaction")

	group.Post("", transactionHandler.CreateTransactionsHandler)
	group.Get("", transactionHandler.GetTransactionsHandler)
	group.Delete("", transactionHandler.DeleteTransactionsHandler)
	group.Get("/summary", transactionHandler.GetTransactionSummaryHandler)
}
