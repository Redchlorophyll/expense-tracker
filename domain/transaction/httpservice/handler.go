package httpservice

import (
	transactionDatatype "github.com/Redchlorophyll/expense-tracker/domain/transaction/datatype"
	"github.com/gofiber/fiber/v2"
)

func NewHandler(cfg TransactionHandlerConfig) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: cfg.TransactionService,
	}
}

func (transaction *TransactionHandler) CreateTransactionsHandler(fiberCtx *fiber.Ctx) error {
	var request transactionDatatype.PostTransactionRequest

	err := fiberCtx.BodyParser(&request)

	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 400,
			Message:    err.Error(),
		})
	}

	result, err := transaction.TransactionService.CreateTransactions(fiberCtx.Context())

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(result)
}

func (transaction *TransactionHandler) GetTransactionsHandler(fiberCtx *fiber.Ctx) error {
	result, err := transaction.TransactionService.GetTransactions(fiberCtx.Context())

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(result)
}

func (transaction *TransactionHandler) DeleteTransactionsHandler(fiberCtx *fiber.Ctx) error {
	var request transactionDatatype.DeleteTransactionRequest

	err := fiberCtx.BodyParser(&request)

	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 400,
			Message:    err.Error(),
		})
	}

	result, err := transaction.TransactionService.DeleteTransactions(fiberCtx.Context())

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(result)
}

func (transaction *TransactionHandler) GetTransactionSummaryHandler(fiberCtx *fiber.Ctx) error {
	result, err := transaction.TransactionService.GetTransactionSummary(fiberCtx.Context())

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(result)
}
