package httpservice

import "github.com/Redchlorophyll/expense-tracker/domain/transaction/services"

type TransactionHandler struct {
	TransactionService services.TransactionServiceProvider
}

type TransactionHandlerConfig struct {
	TransactionService services.TransactionServiceProvider
}
