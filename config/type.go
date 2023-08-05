package config

import (
	transactionHandler "github.com/Redchlorophyll/expense-tracker/domain/transaction/httpservice"
	transactionService "github.com/Redchlorophyll/expense-tracker/domain/transaction/services"
	userHandler "github.com/Redchlorophyll/expense-tracker/domain/user/httpservice"
	userService "github.com/Redchlorophyll/expense-tracker/domain/user/services"
)

type HTTPService struct {
	UserHandler *userHandler.UserServiceHandler

	TransactionHandler *transactionHandler.TransactionHandler
}

type Service struct {
	User userService.UserServiceProvider

	Transaction transactionService.TransactionServiceProvider
}
