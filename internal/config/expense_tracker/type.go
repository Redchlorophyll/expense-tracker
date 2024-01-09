package expense_tracker

import (
	transactionHandler "github.com/Redchlorophyll/expense-tracker/internal/domain/transaction/httpservice"
	userHandler "github.com/Redchlorophyll/expense-tracker/internal/domain/user/httpservice"
)

type HTTPService struct {
	UserHandler *userHandler.UserServiceHandler

	TransactionHandler *transactionHandler.TransactionHandler
}
