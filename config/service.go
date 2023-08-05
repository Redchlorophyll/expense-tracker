package config

import (
	transactionService "github.com/Redchlorophyll/expense-tracker/domain/transaction/services"
	userService "github.com/Redchlorophyll/expense-tracker/domain/user/services"
)

func GetService() *Service {
	return &Service{
		User:        userService.NewUserService(userService.UserServiceConfig{}),
		Transaction: transactionService.NewTransactionService(transactionService.TransactionServiceConfig{}),
	}
}
