package httpservice

import (
	"strconv"

	transactionDatatype "github.com/Redchlorophyll/expense-tracker/internal/domain/transaction/datatype"
	"github.com/gofiber/fiber/v2"
)

func NewHandler(cfg TransactionHandlerConfig) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: cfg.TransactionService,
	}
}

func (transaction *TransactionHandler) CreateTransactionsHandler(fiberCtx *fiber.Ctx) error {
	userId, err := strconv.ParseInt(fiberCtx.Cookies("token"), 10, 64)

	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 403,
			Message:    "Forbidden access!",
		})
	}

	var request transactionDatatype.PostTransactionRequest

	err = fiberCtx.BodyParser(&request)

	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 400,
			Message:    err.Error(),
		})
	}

	result, err := transaction.TransactionService.CreateTransactions(fiberCtx.Context(), request, int(userId))

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(result)
}

func (transaction *TransactionHandler) GetTransactionsHandler(fiberCtx *fiber.Ctx) error {
	userId, err := strconv.ParseInt(fiberCtx.Cookies("token"), 10, 64)

	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 403,
			Message:    "Forbidden access!",
		})
	}

	result, err := transaction.TransactionService.GetTransactions(fiberCtx.Context(), int(userId))

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(result)
}

func (transaction *TransactionHandler) DeleteTransactionsHandler(fiberCtx *fiber.Ctx) error {
	userId, err := strconv.ParseInt(fiberCtx.Cookies("token"), 10, 64)

	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 403,
			Message:    "Forbidden access!",
		})
	}

	var request transactionDatatype.DeleteTransactionRequest

	err = fiberCtx.BodyParser(&request)

	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 400,
			Message:    err.Error(),
		})
	}

	result, err := transaction.TransactionService.DeleteTransactions(fiberCtx.Context(), request.Transactions, int(userId))

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(result)
}

func (transaction *TransactionHandler) GetTransactionSummaryHandler(fiberCtx *fiber.Ctx) error {
	userId, err := strconv.ParseInt(fiberCtx.Cookies("token"), 10, 64)

	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 403,
			Message:    "Forbidden access!",
		})
	}

	result, err := transaction.TransactionService.GetTransactionSummary(fiberCtx.Context(), int(userId))

	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(transactionDatatype.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(result)
}
